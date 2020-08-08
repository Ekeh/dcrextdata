// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package commstats

import (
	"context"
	"net/http"
	"time"

	"github.com/planetdecred/dcrextdata/app/config"
)

type CommStat struct {
	Date               time.Time         `json:"date"`
	RedditStats        map[string]Reddit `json:"reddit_stats"`
	TwitterFollowers   int               `json:"twitter_followers"`
	YoutubeSubscribers int               `json:"youtube_subscribers"`
	GithubStars        int               `json:"github_stars"`
	GithubFolks        int               `json:"github_folks"`
}

type RedditResponse struct {
	Kind string `json:"kind"`
	Data Reddit `json:"data"`
}

type Reddit struct {
	Date           time.Time `json:"date"`
	Subscribers    int       `json:"subscribers"`
	AccountsActive int       `json:"active_user_count"`
	Subreddit      string    `json:"subreddit"`
}

type Github struct {
	Date       time.Time `json:"date"`
	Stars      int       `json:"stars"`
	Folks      int       `json:"folks"`
	Repository string    `json:"repository"`
}

type Youtube struct {
	Date        time.Time `json:"date"`
	Subscribers int       `json:"subscribers"`
	Channel     string    `json:"channel"`
	ViewCount   int       `json:"view_count"`
}

type Twitter struct {
	Date      time.Time `json:"date"`
	Followers int       `json:"followers"`
	Handle    string    `json:"handle"`
}

type ChartData struct {
	Date   time.Time `json:"date"`
	Record int64     `json:"record"`
}

type DataStore interface {
	StoreRedditStat(context.Context, Reddit) error
	LastCommStatEntry() (time time.Time)
	StoreTwitterStat(ctx context.Context, twitter Twitter) error
	StoreYoutubeStat(ctx context.Context, youtube Youtube) error
	StoreGithubStat(ctx context.Context, github Github) error
	LastEntry(ctx context.Context, tableName string, receiver interface{}) error
	StoreGoogleStatsInterestOverTime(ctx context.Context, interestOverTime []GoogleInterestOverTime) error
	StoreGoogleStatsInterestByLocation(ctx context.Context, interestByLocation []GoogleInterestByLocation) error
}

type Collector struct {
	client    http.Client
	dataStore DataStore
	options   *config.CommunityStatOptions
}

// <Google trends start here>
// Timeline - it's representation of interest to trend trough period timeline. Mostly used for charts
type GoogleInterestOverTime struct {
	Id                int       `json:"time" bson:"time"`
	Geo               string    `json:"geo" bson:"geo"`
	FormattedTime     string    `json:"formatted_time" bson:"formatted_time"`
	FormattedAxisTime string    `json:"formatted_axis_time" bson:"formatted_axis_time"`
	Value             int       `json:"value" bson:"value"`
	Keyword           string    `json:"keyword" bson:"keyword"`
	Date              time.Time `json:"date"`
}

type GoogleInterestByLocation struct {
	Id            int       `json:"time" bson:"time"`
	Geo           string    `json:"geo" bson:"geo"`
	GeoCode       string    `json:"geoCode" bson:"geo_code"`
	GeoName       string    `json:"geoName" bson:"geo_name"`
	Value         int       `json:"value" bson:"value"`
	MaxValueIndex int       `json:"maxValueIndex" bson:"max_value_index"`
	Keyword       string    `json:"keyword" bson:"keyword"`
	Date          time.Time `json:"date"`
}

type Timeline struct {
	Time              string   `json:"time" bson:"time"`
	FormattedTime     string   `json:"formattedTime" bson:"formatted_time"`
	FormattedAxisTime string   `json:"formattedAxisTime" bson:"formatted_axis_time"`
	Value             []int    `json:"value" bson:"value"`
	HasData           []bool   `json:"hasData" bson:"has_data"`
	FormattedValue    []string `json:"formattedValue" bson:"formatted_value"`
}

// GeoMap - it's representation of interest by location. Mostly used for maps
type GeoMap struct {
	GeoCode        string   `json:"geoCode" bson:"geo_code"`
	GeoName        string   `json:"geoName" bson:"geo_name"`
	Value          []int    `json:"value" bson:"value"`
	FormattedValue []string `json:"formattedValue" bson:"formatted_value"`
	MaxValueIndex  int      `json:"maxValueIndex" bson:"max_value_index"`
	HasData        []bool   `json:"hasData" bson:"has_data"`
}

// ExploreRequest it's an input which can contain multiple items (keywords) to discover
// category can be found in ExploreCategories output
type ExploreRequest struct {
	ComparisonItems []*ComparisonItem `json:"comparisonItem" bson:"comparison_items"`
	Category        int               `json:"category" bson:"category"`
	Property        string            `json:"property" bson:"property"`
}

// ComparisonItem it's concrete search keyword
// with Geo (can be found with ExploreLocations method) locality and Time period
type ComparisonItem struct {
	Keyword string `json:"keyword" bson:"keyword"`
	Geo     string `json:"geo,omitempty" bson:"geo"`
	Time    string `json:"time" bson:"time"`
}

// ExploreWidget - output of Explore method, required for InterestOverTime, InterestByLocation and Related methods.
// Globally it's a structure related to Google Trends UI and contains mostly system info
type ExploreWidget struct {
	Token   string          `json:"token" bson:"token"`
	Type    string          `json:"type" bson:"type"`
	Title   string          `json:"title" bson:"title"`
	ID      string          `json:"id" bson:"id"`
	Request *WidgetResponse `json:"request" bson:"request"`
}

// WidgetResponse - system info for every available trends search mode
type WidgetResponse struct {
	Geo                interface{}             `json:"geo,omitempty" bson:"geo"`
	Time               string                  `json:"time,omitempty" bson:"time"`
	Resolution         string                  `json:"resolution,omitempty" bson:"resolution"`
	Locale             string                  `json:"locale,omitempty" bson:"locale"`
	Restriction        WidgetComparisonItem    `json:"restriction" bson:"restriction"`
	CompItem           []*WidgetComparisonItem `json:"comparisonItem" bson:"comparison_item"`
	RequestOpt         RequestOptions          `json:"requestOptions" bson:"request_option"`
	KeywordType        string                  `json:"keywordType" bson:"keyword_type"`
	Metric             []string                `json:"metric" bson:"metric"`
	Language           string                  `json:"language" bson:"language"`
	TrendinessSettings map[string]string       `json:"trendinessSettings" bson:"trendiness_settings"`
	DataMode           string                  `json:"dataMode,omitempty" bson:"data_mode"`
	UserCountryCode    string                  `json:"userCountryCode,omitempty" bson:"user_country_code"`
}

// WidgetComparisonItem - system info for comparison item part of WidgetResponse
type WidgetComparisonItem struct {
	Geo                            map[string]string   `json:"geo,omitempty" bson:"geo"`
	Time                           string              `json:"time,omitempty" bson:"time"`
	ComplexKeywordsRestriction     KeywordsRestriction `json:"complexKeywordsRestriction,omitempty" bson:"complex_keywords_restriction"`
	OriginalTimeRangeForExploreURL string              `json:"originalTimeRangeForExploreUrl,omitempty" bson:"original_time_range_for_explore_url"`
}

// KeywordsRestriction - system info for keywords limitations, not used. part of WidgetResponse
type KeywordsRestriction struct {
	Keyword []*KeywordRestriction `json:"keyword" bson:"keyword"`
}

// RequestOptions - part of WidgetResponse
type RequestOptions struct {
	Property string `json:"property" bson:"property"`
	Backend  string `json:"backend" bson:"backend"`
	Category int    `json:"category" bson:"category"`
}

// KeywordRestriction - specific keyword limitation. Part of KeywordsRestriction
type KeywordRestriction struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type exploreOut struct {
	Widgets []*ExploreWidget `json:"widgets" bson:"widgets"`
}

type multilineOut struct {
	Default multiline `json:"default" bson:"default"`
}

type multiline struct {
	TimelineData []*Timeline `json:"timelineData" bson:"timeline_data"`
}

type geoOut struct {
	Default geo `json:"default" bson:"default"`
}

type geo struct {
	GeoMapData []*GeoMap `json:"geoMapData" bson:"geomap_data"`
}

// </Google trends end here>
