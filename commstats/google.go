package commstats

import (
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/planetdecred/dcrextdata/app"

	"github.com/planetdecred/dcrextdata/app/config"
	"github.com/planetdecred/dcrextdata/app/helpers"
	"github.com/planetdecred/dcrextdata/postgres/models"
)

const (
	googleAPI     = "https://trends.google.com/trends/api"
	gSIntOverTime = "/widgetdata/multiline"
	gSIntOverReg  = "/widgetdata/comparedgeo"
	gSExplore     = "/explore"
	paramHl       = "hl"
	paramCat      = "cat"
	paramGeo      = "geo"
	paramReq      = "req"
	paramTZ       = "tz"
	paramToken    = "token"

	intOverTimeWidgetID = "TIMESERIES"
	intOverRegionID     = "GEO_MAP"

	locUS  = "US"
	catAll = "all"
	langEn = "EN"

	errParsing        = "failed to parse json"
	errReqDataF       = "request data: code = %d, status = %s"
	errInvalidRequest = "invalid request param"
	errCreateRequest  = "failed to create request"
	errDoRequest      = "failed to perform request"

	headerKeyAccept    = "Accept"
	headerKeyCookie    = "Cookie"
	headerKeySetCookie = "Set-Cookie"
	contentTypeJSON    = "application/json"
	compareDataMode    = "PERCENTAGES"
)

var (
	defaultParams = map[string]string{
		paramTZ:  "0",
		paramCat: "all",
		"fi":     "0",
		"fs":     "0",
		paramHl:  "EN",
		"ri":     "300",
		"rs":     "20",
	}
	cookie string
	// ErrRequestFailed - response status != 200
	ErrRequestFailed = errors.New("failed to perform http request")
	// ErrInvalidWidgetType - provided widget is invalid or is used for another method
	ErrInvalidWidgetType = errors.New("invalid widget type")

	interestOverTime     []GoogleInterestOverTime
	interestByLocation   []GoogleInterestByLocation
	googleTrendsKeywords []string = config.DefaultGoogleTrendsSearchKeywords
)

func GoogleTrendsKeywords() []string {
	return googleTrendsKeywords
}

func GoogleTrendsInterestOverTime() []GoogleInterestOverTime {
	return interestOverTime
}

func GoogleTrendsInterestByLocation() []GoogleInterestByLocation {
	return interestByLocation
}

func (c *Collector) startGoogleCollector(ctx context.Context) {
	//Enable only dcrd trends overtime
	if c.options.GoogleTrendType == config.DefaultGoogleTrendTypeBoth {
		go c.startGoogleCollectorInterestOverTime(ctx)
		go c.startGoogleCollectorInterestByLocation(ctx)
	} else if c.options.GoogleTrendType == config.DefaultGoogleTrendTypeTrendOverTime {
		go c.startGoogleCollectorInterestOverTime(ctx)
	} else if c.options.GoogleTrendType == config.DefaultGoogleTrendTypeTrendByLocation {
		go c.startGoogleCollectorInterestByLocation(ctx)
	} else {
		log.Warnf("Unknow Google trend type '%s'. Falling back to default trend '%s'", c.options.GoogleTrendType, config.DefaultGoogleTrendTypeBoth)
		go c.startGoogleCollectorInterestOverTime(ctx)
		go c.startGoogleCollectorInterestByLocation(ctx)
	}
}

func (c *Collector) startGoogleCollectorInterestOverTime(ctx context.Context) {
	var lastCollectionDate time.Time
	err := c.dataStore.LastEntry(ctx, models.TableNames.Googleinterestovertime, &lastCollectionDate)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("Cannot fetch last Google trends interest over time entry time, %s", err.Error())
		return
	}
	secondsPassed := time.Since(lastCollectionDate)
	period := time.Duration(c.options.GoogleTrendsStatInterval) * time.Minute
	if secondsPassed < period {
		timeLeft := period - secondsPassed
		log.Infof("Fetching Google trends interest over time stats every %dm, collected %s ago, will fetch in %s.", period/time.Minute,
			helpers.DurationToString(secondsPassed), helpers.DurationToString(timeLeft))
		time.Sleep(timeLeft)
	}

	registerStarter := func() {
		// continually check the state of the app until its free to run this module
		//app.MarkBusyIfFree()
	}

	registerStarter()
	c.collectAndStoreGoogleTrendsInterestOverTime(ctx)
	app.ReleaseForNewModule()

	ticker := time.NewTicker(time.Duration(c.options.GoogleTrendsStatInterval) * time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			registerStarter()
			c.collectAndStoreGoogleTrendsInterestOverTime(ctx)
			app.ReleaseForNewModule()
		}
	}
}

func (c *Collector) startGoogleCollectorInterestByLocation(ctx context.Context) {
	var lastCollectionDate time.Time
	err := c.dataStore.LastEntry(ctx, models.TableNames.Googleinterestbylocation, &lastCollectionDate)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("Cannot fetch last Google trends interest by location entry time, %s", err)
		return
	}
	secondsPassed := time.Since(lastCollectionDate)
	period := time.Duration(c.options.GoogleTrendsStatInterval) * time.Minute
	if secondsPassed < period {
		timeLeft := period - secondsPassed
		log.Infof("Fetching Google trends interest by location stats every %dm, collected %s ago, will fetch in %s.", period/time.Minute,
			helpers.DurationToString(secondsPassed), helpers.DurationToString(timeLeft))
		time.Sleep(timeLeft)
	}
	registerStarter := func() {
		// continually check the state of the app until its free to run this module
		//app.MarkBusyIfFree()
	}

	registerStarter()
	c.collectAndStoreGoogleTrendsInterestByLocation(ctx)
	app.ReleaseForNewModule()

	ticker := time.NewTicker(time.Duration(c.options.GoogleTrendsStatInterval) * time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			registerStarter()
			c.collectAndStoreGoogleTrendsInterestByLocation(ctx)
			app.ReleaseForNewModule()
		}
	}
}

func (c *Collector) collectAndStoreGoogleTrendsInterestOverTime(ctx context.Context) {
	for _, keyword := range c.options.GoogleTrendsSearchKeywords {
		log.Infof("Collecting Google trends interest over time for '%s' keyword", keyword)
		explore, err := c.explore(ctx, keyword, langEn)
		for retry := 0; err != nil; retry++ {
			if retry == retryLimit {
				log.Error(err)
				return
			}
			log.Warn(err)
			explore, err = c.explore(ctx, keyword, langEn)
		}
		if err != nil {
			log.Errorf("Failed to explore google trend widgets, %s", err.Error())
			return
		}
		overTime, err := c.fetchInterestOverTime(ctx, explore[0], langEn)
		for retry := 0; err != nil; retry++ {
			if retry == retryLimit {
				log.Error(err)
				return
			}
			log.Warn(err)
			overTime, err = c.fetchInterestOverTime(ctx, explore[0], langEn)
		}
		interestOverTime = c.formatResultInterestOverTime(overTime, keyword)
		err = c.dataStore.StoreGoogleStatsInterestOverTime(ctx, interestOverTime)
		if err != nil {
			log.Errorf("Unable to save Google trends stat interest over time for '%s', %s", keyword, err.Error())
			return
		}
		if len(interestOverTime) > 0 {
			log.Infof("New Google trends interest over time stat for '%s' collected at %s", keyword,
				interestOverTime[0].Date.Format(dateMiliTemplate))
		} else {
			log.Infof("No Google trends interest over time stat collected for '%s' keyword.", keyword)
		}
	}
}

func (c *Collector) collectAndStoreGoogleTrendsInterestByLocation(ctx context.Context) {
	for _, keyword := range c.options.GoogleTrendsSearchKeywords {
		log.Infof("Collecting Google trends interest by location for '%s' keyword", keyword)
		explore, err := c.explore(ctx, keyword, langEn)
		for retry := 0; err != nil; retry++ {
			if retry == retryLimit {
				log.Error(err)
				return
			}
			log.Warn(err)
			explore, err = c.explore(ctx, keyword, langEn)
		}
		if err != nil {
			log.Errorf("Failed to explore Google trends widgets, %s", err.Error())
			return
		}
		overReg, err := c.fetchInterestByLocation(ctx, explore[1], langEn)
		for retry := 0; err != nil; retry++ {
			if retry == retryLimit {
				log.Error(err)
				return
			}
			log.Warn(err)
			overReg, err = c.fetchInterestByLocation(ctx, explore[1], langEn)
		}
		interestByLocation = c.formatResultInterestByLocation(overReg, keyword)
		err = c.dataStore.StoreGoogleStatsInterestByLocation(ctx, interestByLocation)
		if err != nil {
			log.Error("Unable to save Google trends interest by location stat for '%s', %s", keyword, err.Error())
			return
		}
		if len(interestByLocation) > 0 {
			log.Infof("New Google trends interest by location stat for '%s' collected at %s", keyword,
				interestByLocation[0].Date.Format(dateMiliTemplate))
		} else {
			log.Infof("No Google trends interest  by location stat collected for '%s' keyword.", keyword)
		}
	}
}

// InterestOverTime as list of `Timeline` dots for chart.
func (c *Collector) fetchInterestOverTime(ctx context.Context, w *ExploreWidget, hl string) ([]*Timeline, error) {
	if w.ID != intOverTimeWidgetID {
		return nil, ErrInvalidWidgetType
	}

	u, _ := url.Parse(googleAPI + gSIntOverTime)

	p := make(url.Values)
	p.Set(paramTZ, "0")
	p.Set(paramHl, hl)
	p.Set(paramToken, w.Token)

	for i, v := range w.Request.CompItem {
		if len(v.Geo) == 0 {
			w.Request.CompItem[i].Geo[""] = ""
		}
	}

	// marshal request for query param
	mReq, err := jsoniter.MarshalToString(w.Request)
	if err != nil {
		return nil, errors.Wrapf(err, errInvalidRequest)
	}

	p.Set(paramReq, mReq)
	u.RawQuery = p.Encode()
	b, err := c.do(ctx, u)
	if err != nil {
		return nil, err
	}

	// google api returns not valid json :(
	str := strings.Replace(string(b), ")]}',", "", 1)

	out := new(multilineOut)
	if err := c.unmarshal(str, out); err != nil {
		return nil, err
	}

	return out.Default.TimelineData, nil
}

func (c *Collector) formatResultInterestOverTime(overTime []*Timeline, keyword string) []GoogleInterestOverTime {
	interests := make([]GoogleInterestOverTime, 0, len(overTime))
	for _, v := range overTime {
		overTimeTrend := GoogleInterestOverTime{
			Geo:               c.options.GoogleTrendsGeolocation,
			FormattedTime:     v.FormattedTime,
			FormattedAxisTime: v.FormattedAxisTime,
			Value:             v.Value[0],
			Date:              helpers.NowUTC(),
			Keyword:           keyword,
		}
		interests = append(interests, overTimeTrend)
	}
	return interests
}

func (c *Collector) formatResultInterestByLocation(overTime []*GeoMap, keyword string) []GoogleInterestByLocation {
	interests := make([]GoogleInterestByLocation, 0, len(overTime))
	for _, v := range overTime {
		interestByLocationTrend := GoogleInterestByLocation{
			Geo:           c.options.GoogleTrendsGeolocation,
			GeoCode:       v.GeoCode,
			GeoName:       v.GeoName,
			Value:         v.Value[0],
			MaxValueIndex: v.MaxValueIndex,
			Date:          helpers.NowUTC(),
			Keyword:       keyword,
		}
		interests = append(interests, interestByLocationTrend)
	}
	return interests
}

// InterestByLocation as list of `GeoMap`, with geo codes and interest values.
func (c *Collector) fetchInterestByLocation(ctx context.Context, w *ExploreWidget, hl string) ([]*GeoMap, error) {
	if w.ID != intOverRegionID {
		return nil, ErrInvalidWidgetType
	}

	u, _ := url.Parse(googleAPI + gSIntOverReg)

	p := make(url.Values)
	p.Set(paramTZ, "0")
	p.Set(paramHl, hl)
	p.Set(paramToken, w.Token)

	if len(w.Request.CompItem) > 1 {
		w.Request.DataMode = compareDataMode
	}

	// marshal request for query param
	mReq, err := jsoniter.MarshalToString(w.Request)
	if err != nil {
		return nil, errors.Wrapf(err, errInvalidRequest)
	}

	p.Set(paramReq, mReq)
	u.RawQuery = p.Encode()
	log.Info(u)
	b, err := c.do(ctx, u)
	if err != nil {
		return nil, err
	}

	// google api returns not valid json x:(
	str := strings.Replace(string(b), ")]}',", "", 1)
	out := new(geoOut)
	if err := c.unmarshal(str, out); err != nil {
		return nil, err
	}
	return out.Default.GeoMapData, nil
}

// Explore list of widgets with tokens. Every widget
// is related to specific method (`InterestOverTime`, `InterestOverLoc`, `RelatedSearches`, `Suggestions`)
// and contains required token and request information.
func (c *Collector) explore(ctx context.Context, keyword string, hl string) ([]*ExploreWidget, error) {
	// hook for using incorrect `time` request (backward compatibility)
	log.Infof("Exploring Google trends for '%s' geolocation.", c.options.GoogleTrendsGeolocation)
	r := &ExploreRequest{
		ComparisonItems: []*ComparisonItem{
			{
				Keyword: keyword,
				Geo:     c.options.GoogleTrendsGeolocation,
				Time:    "today 12-m", // For more info visit https://wiki.q-researchsoftware.com/wiki/Data_-_Google_Trends
			},
		},
		Category: 31, // Programming category
		Property: "",
	}
	for _, r := range r.ComparisonItems {
		r.Time = strings.ReplaceAll(r.Time, "+", " ")
	}

	u, _ := url.Parse(googleAPI + gSExplore)

	p := make(url.Values)
	p.Set(paramTZ, "0")
	p.Set(paramHl, hl)

	// marshal request for query param
	mReq, err := jsoniter.MarshalToString(r)
	if err != nil {
		return nil, errors.Wrapf(err, errInvalidRequest)
	}

	p.Set(paramReq, mReq)
	u.RawQuery = p.Encode()
	b, err := c.do(ctx, u)
	if err != nil {
		return nil, err
	}

	// google api returns not valid json :(
	str := strings.Replace(string(b), ")]}'", "", 1)

	out := new(exploreOut)
	if err := c.unmarshal(str, out); err != nil {
		return nil, err
	}

	return out.Widgets, nil
}

// Performs http request with a given url
func (c *Collector) do(ctx context.Context, u *url.URL) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, errCreateRequest)
	}

	r.Header.Add(headerKeyAccept, contentTypeJSON)

	resp, err := c.client.Do(r)
	if err != nil {
		return nil, errors.Wrap(err, errDoRequest)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		cookies := strings.Split(resp.Header.Get(headerKeySetCookie), ";")
		if len(cookies) > 0 {
			cookie = cookies[0]
			r.Header.Set(headerKeyCookie, cookies[0])
			resp, err = c.client.Do(r)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(ErrRequestFailed, errReqDataF, resp.StatusCode, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *Collector) unmarshal(str string, dest interface{}) error {
	if err := jsoniter.UnmarshalFromString(str, dest); err != nil {
		return errors.Wrap(err, errParsing)
	}

	return nil
}
