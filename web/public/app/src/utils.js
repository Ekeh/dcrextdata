import dompurify from 'dompurify'
import humanize from './helpers/humanize_helper'
import { map } from 'lodash-es'

const Dygraph = require('../../dist/js/dygraphs.min.js')

export const appName = 'dcrextdata'

export const hide = (el) => {
  el.classList.add('d-none')
  el.classList.add('d-hide')
}

export const hideAll = (els) => {
  els.forEach(el => {
    el.classList.add('d-none')
    el.classList.add('d-hide')
  })
}

export const show = (el) => {
  el.classList.remove('d-none')
  el.classList.remove('d-hide')
}

export const showAll = (els) => {
  els.forEach(el => {
    el.classList.remove('d-none')
    el.classList.remove('d-hide')
  })
}

export const setAllValues = (targets, value) => {
  targets.forEach(el => {
    el.innerHTML = value
  })
}

export const showLoading = (loadingTarget, elementsToHide) => {
  show(loadingTarget)
  if (!elementsToHide || !elementsToHide.length) return
  elementsToHide.forEach(element => {
    hide(element)
  })
}

export const hideLoading = (loadingTarget, elementsToShow) => {
  hide(loadingTarget)
  if (!elementsToShow || !elementsToShow.length) return
  elementsToShow.forEach(element => {
    show(element)
  })
}

export const isHidden = (el) => {
  return el.classList.contains('d-none')
}

export function legendFormatter (data) {
  let html = ''
  if (data.x == null) {
    let dashLabels = data.series.reduce((nodes, series) => {
      return `${nodes} <div class="pr-2">${series.dashHTML} ${series.labelHTML}</div>`
    }, '')
    html = `<div class="d-flex flex-wrap justify-content-center align-items-center">
              <div class="pr-3">${this.getLabels()[0]}: N/A</div>
              <div class="d-flex flex-wrap">${dashLabels}</div>
            </div>`
  } else {
    data.series.sort((a, b) => a.y > b.y ? -1 : 1)
    let extraHTML = ''

    let yVals = data.series.reduce((nodes, series) => {
      if (!series.isVisible) return nodes
      let yVal = series.yHTML
      if (series.y === undefined) {
        yVal = 'N/A'
      } else {
        yVal = series.y
      }
      // propotion missed/live has % sign
      if (series.y !== undefined && (series.label.toLowerCase() === 'proportion live (%)' || series.label.toLowerCase() === 'proportion missed (%)')) {
        yVal += '%'
      }
      if (yVal === undefined) {
        yVal = 'N/A'
      }
      return `${nodes} <div class="pr-2">${series.dashHTML} ${series.labelHTML}: ${yVal}</div>`
    }, '')

    let xHTML = data.xHTML
    if (data.dygraph.getLabels()[0] === 'Date') {
      xHTML = humanize.date(data.x, false, false)
    }

    html = `<div class="d-flex flex-wrap justify-content-center align-items-center">
                <div class="pr-3">${this.getLabels()[0]}: ${xHTML}</div>
                <div class="d-flex flex-wrap"> ${yVals}</div>
            </div>${extraHTML}`
  }

  dompurify.sanitize(html)
  return html
}

export function barChartPlotter (e) {
  const ctx = e.drawingContext
  const points = e.points
  const yBottom = e.dygraph.toDomYCoord(0)

  ctx.fillStyle = darkenColor(e.color)

  // Find the minimum separation between x-values.
  // This determines the bar width.
  let minSep = Infinity
  for (let i = 1; i < points.length; i++) {
    const sep = points[i].canvasx - points[i - 1].canvasx
    if (sep < minSep) minSep = sep
  }
  const barWidth = Math.max(Math.floor(2.0 / 3 * minSep), 5)

  // Do the actual plotting.
  for (let i = 0; i < points.length; i++) {
    const p = points[i]
    const centerx = p.canvasx

    ctx.fillRect(centerx - barWidth / 2, p.canvasy, barWidth, yBottom - p.canvasy)
    ctx.strokeRect(centerx - barWidth / 2, p.canvasy, barWidth, yBottom - p.canvasy)
  }
}

function darkenColor (colorStr) {
  // Defined in dygraph-utils.js
  var color = Dygraph.toRGB_(colorStr)
  color.r = Math.floor((255 + color.r) / 2)
  color.g = Math.floor((255 + color.g) / 2)
  color.b = Math.floor((255 + color.b) / 2)
  return 'rgb(' + color.r + ',' + color.g + ',' + color.b + ')'
}

export var options = {
  axes: { y: { axisLabelWidth: 100 } },
  axisLabelFontSize: 12,
  retainDateWindow: false,
  showRangeSelector: true,
  rangeSelectorHeight: 40,
  drawPoints: true,
  pointSize: 0.25,
  legend: 'always',
  labelsSeparateLines: true,
  highlightCircleSize: 4,
  yLabelWidth: 20,
  drawAxesAtZero: true
}

export function getRandomColor () {
  const letters = '0123456789ABCDEF'
  let color = '#'
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)]
  }
  return color
}

export function setActiveOptionBtn (opt, optTargets) {
  optTargets.forEach(li => {
    if (li.dataset.option === opt) {
      li.classList.add('active')
    } else {
      li.classList.remove('active')
    }
  })
}

export function setActiveRecordSetBtn (opt, optTargets) {
  optTargets.forEach(li => {
    if (li.dataset.option === opt) {
      li.classList.add('active')
    } else {
      li.classList.remove('active')
    }
  })
}

export function displayPillBtnOption (opt, optTargets) {
  optTargets.forEach(li => {
    if (opt === 'chart' && li.dataset.option === 'both') {
      li.classList.add('d-hide')
    } else {
      li.classList.remove('d-hide')
    }
  })
}

export function selectedOption (optTargets) {
  var key = false
  optTargets.forEach((el) => {
    if (el.classList.contains('active')) key = el.dataset.option
  })
  return key
}

export function insertQueryParam (name, value, defaultValue) {
  if (value === defaultValue) return
  const urlParams = new URLSearchParams(window.location.search)
  const oldValue = urlParams.get(name)
  if (oldValue !== null) {
    return false
  }
  urlParams.append(name, value)
  const baseUrl = window.location.href.replace(window.location.search, '')
  let q = urlParams.toString()
  if (q.length > 0) {
    q = `?${q}`
  }
  window.history.pushState(window.history.state, appName, `${baseUrl}${q}`)
  return true
}

export function updateQueryParam (name, value, defaultValue) {
  let urlParams = new URLSearchParams(window.location.search)
  if (!urlParams.has(name)) {
    return false
  }
  if (value === defaultValue) {
    urlParams.delete(name)
  } else {
    urlParams.set(name, value)
  }
  const baseUrl = window.location.href.replace(window.location.search, '')
  let q = urlParams.toString()
  if (q.length > 0) {
    q = `?${q}`
  }
  window.history.pushState(window.history.state, appName, `${baseUrl}${q}`)
  return true
}

export function insertOrUpdateQueryParam (name, value, defaultValue) {
  const urlParams = new URLSearchParams(window.location.search)
  return !urlParams.has(name) ? insertQueryParam(name, value, defaultValue) : updateQueryParam(name, value, defaultValue)
}

export function trimUrl (keepSet) {
  if (window.location.search.length === 0) return
  let urlParams = new URLSearchParams(window.location.search)
  let newParam = new URLSearchParams()
  for (let i = 0; i <= keepSet.length; i++) {
    const key = keepSet[i]
    if (!urlParams.has(key)) continue
    newParam.append(key, urlParams.get(key))
  }
  const baseUrl = window.location.href.replace(window.location.search, '')
  let q = newParam.toString()
  if (q.length > 0) {
    q = `?${q}`
  }
  window.history.replaceState(window.history.state, appName, `${baseUrl}${q}`)
}

export function removeUrlParam (name) {
  if (window.location.search.length === 0) return
  let urlParams = new URLSearchParams(window.location.search)
  if (!urlParams.has(name)) {
    return false
  }
  urlParams.delete(name)
  const baseUrl = window.location.href.replace(window.location.search, '')
  let q = urlParams.toString()
  if (q.length > 0) {
    q = `?${q}`
  }
  window.history.replaceState(window.history.state, appName, `${baseUrl}${q}`)
  return true
}

export function getParameterByName (name, url) {
  const urlParams = new URLSearchParams(window.location.search)
  return urlParams.get(name)
}

export function zipXYZData (gData, isHeightAxis, isDayBinned, yCoefficient, zCoefficient, windowS) {
  windowS = windowS || 1
  yCoefficient = yCoefficient || 1
  zCoefficient = zCoefficient || 1
  return map(gData.x, (n, i) => {
    let xAxisVal
    if (isHeightAxis && isDayBinned) {
      xAxisVal = n
    } else if (isHeightAxis) {
      xAxisVal = n * windowS
    } else {
      xAxisVal = new Date(n * 1000)
    }
    const data = [xAxisVal, gData.y[i] * yCoefficient]
    if (gData.z) {
      data.push(gData.z[i] * zCoefficient)
    }

    return data
  })
}

export function updateZoomSelector (targets, minDate, maxDate) {
  const duration = maxDate - minDate
  const days = duration / (1000 * 60 * 60 * 24)
  targets.forEach(el => {
    let showElement = false
    switch (el.dataset.option) {
      case 'day':
      case 'all':
        showElement = days >= 1
        break
      case 'week':
        showElement = days >= 7
        break
      case 'month':
        showElement = days >= 30
        break
      case 'year':
        showElement = days >= 365
        break
    }

    if (showElement) {
      show(el)
    } else {
      hide(el)
    }
  })
}

export function formatDate (date, format) {
  if (!format || format === '') {
    format = 'yyyy-MM-dd hh:mm'
  }

  let dd = date.getDate()
  let mm = date.getMonth() + 1
  let yyyy = date.getFullYear()
  let milliseconds = date.getMilliseconds()
  let seconds = date.getSeconds()
  let minutes = date.getMinutes()
  let hour = date.getHours()

  if (dd < 10) {
    dd = '0' + dd
  }

  if (mm < 10) {
    mm = '0' + mm
  }

  if (hour < 10) {
    hour = '0' + hour
  }

  if (minutes < 10) {
    minutes = '0' + minutes
  }

  if (seconds < 10) {
    seconds = '0' + seconds
  }

  let dateFormatted = format.replace('yyyy', yyyy).replace('MM', mm).replace('dd', dd)
  dateFormatted = dateFormatted.replace('hh', hour).replace('mm', minutes)
  dateFormatted = dateFormatted.replace('ss', seconds).replace('sss', milliseconds)
  return dateFormatted
}

export function getNumberOfPages (recordsCount, pageSize) {
  const rem = recordsCount % pageSize
  let pageCount = (recordsCount - rem) / pageSize
  if (rem > 0) {
    pageCount += 1
  }
  return pageCount
}

/* eslint no-undef: 0 */
export function notifySuccess (title, message) {
  toastr.success(title, message)
}

export function notifyFailure (title, message) {
  toastr.error(title, message)
}

/* eslint no-undef: 1 */
