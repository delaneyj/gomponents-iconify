package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const (
	IconifyVersion             = "1.0.2"
	activityPath               = `<g id="feActivity0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feActivity1" fill="currentColor"><path id="feActivity2" d="M11 19a1 1 0 0 1-2 0V8a1 1 0 1 1 2 0v11Zm-4 0a1 1 0 0 1-2 0v-7a1 1 0 0 1 2 0v7Zm6 0v-9a1 1 0 0 1 2 0v9a1 1 0 0 1-2 0Zm4-14a1 1 0 0 1 2 0v14a1 1 0 0 1-2 0V5Z"/></g></g>`
	addCartPath                = `<g id="feAddCart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAddCart1" fill="currentColor"><path id="feAddCart2" d="M8 16h8a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V4H5a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1v9h8.775L18 7V5.001c1.145 0 2 .894 2 1.999c0 .146-.017.291-.05.434l-1.151 5c-.21.915-1.052 1.566-2.024 1.566H8.073L8 13.999V16Zm-.5 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm9 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM14 5h1a1 1 0 0 1 0 2h-1v1a1 1 0 0 1-2 0V7h-1a1 1 0 0 1 0-2h1V4a1 1 0 0 1 2 0v1Z"/></g></g>`
	alignBottomPath            = `<g id="feAlignBottom0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignBottom1" fill="currentColor"><path id="feAlignBottom2" d="M4.8 19h14.4c.442 0 .8.448.8 1s-.358 1-.8 1H4.8c-.442 0-.8-.448-.8-1s.358-1 .8-1Zm6.2-4a2 2 0 1 1-4 0V5a2 2 0 1 1 4 0v10Zm6 0a2 2 0 1 1-4 0V9a2 2 0 1 1 4 0v6Z"/></g></g>`
	alignCenterPath            = `<g id="feAlignCenter0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignCenter1" fill="currentColor"><path id="feAlignCenter2" d="M11 13v-2H6.286C5.023 11 4 10.105 4 9s1.023-2 2.286-2H11V3a1 1 0 0 1 2 0v4h4.714C18.977 7 20 7.895 20 9s-1.023 2-2.286 2H13v2h3a2 2 0 1 1 0 4h-3v4a1 1 0 0 1-2 0v-4H8a2 2 0 1 1 0-4h3Z"/></g></g>`
	alignLeftPath              = `<g id="feAlignLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignLeft1" fill="currentColor"><path id="feAlignLeft2" d="M7 21a1 1 0 0 1-2 0V3a1 1 0 1 1 2 0v18Zm4-12a2 2 0 1 1 0-4h6a2 2 0 1 1 0 4h-6Zm0 3h4a2 2 0 1 1 0 4h-4a2 2 0 1 1 0-4Z"/></g></g>`
	alignRightPath             = `<g id="feAlignRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignRight1" fill="currentColor"><path id="feAlignRight2" d="M17 21a1 1 0 0 1-2 0V3a1 1 0 0 1 2 0v18ZM11 5a2 2 0 1 1 0 4H5a2 2 0 1 1 0-4h6Zm0 7a2 2 0 1 1 0 4H7a2 2 0 1 1 0-4h4Z"/></g></g>`
	alignTopPath               = `<g id="feAlignTop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignTop1" fill="currentColor"><path id="feAlignTop2" d="M4.8 3h14.4c.442 0 .8.448.8 1s-.358 1-.8 1H4.8C4.358 5 4 4.552 4 4s.358-1 .8-1ZM7 9a2 2 0 1 1 4 0v10a2 2 0 1 1-4 0V9Zm6 0a2 2 0 1 1 4 0v6a2 2 0 1 1-4 0V9Z"/></g></g>`
	alignVerticallyPath        = `<g id="feAlignVertically0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignVertically1" fill="currentColor"><path id="feAlignVertically2" d="M17 13v2a2 2 0 1 1-4 0v-2h-2v4.714C11 18.977 10.105 20 9 20s-2-1.023-2-2.286V13H3a1 1 0 0 1 0-2h4V6.286C7 5.023 7.895 4 9 4s2 1.023 2 2.286V11h2V9a2 2 0 1 1 4 0v2h4a1 1 0 0 1 0 2h-4Z"/></g></g>`
	angryPath                  = `<g id="feAngry0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAngry1" fill="currentColor"><path id="feAngry2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm4-3h-1.339s-.417-2.667-2.661-2.667S9.333 17 9.333 17H8a4 4 0 1 1 8 0Zm-5.494-8.01a1.5 1.5 0 1 1-1.52-.416L7.33 7.97a.5.5 0 0 1 .342-.94l2.82 1.026a.5.5 0 0 1 .015.934Zm2.897 0a.499.499 0 0 1 .016-.934l2.82-1.026a.5.5 0 1 1 .342.94l-1.659.603a1.5 1.5 0 1 1-1.519.417Z"/></g></g>`
	appMenuPath                = `<g id="feAppMenu0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAppMenu1" fill="currentColor"><path id="feAppMenu2" d="M16 16h4v4h-4v-4Zm-6 0h4v4h-4v-4Zm-6 0h4v4H4v-4Zm12-6h4v4h-4v-4Zm-6 0h4v4h-4v-4Zm-6 0h4v4H4v-4Zm12-6h4v4h-4V4Zm-6 0h4v4h-4V4ZM4 4h4v4H4V4Z"/></g></g>`
	apronPath                  = `<g id="feApron0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feApron1" fill="currentColor" fill-rule="nonzero"><path id="feApron2" d="m9.694 8l-1 6H6v6h12v-6h-2.694l-1-6H9.694ZM10 3v3h4V3a1 1 0 0 1 2 0v3l1 6h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h1l1-6V3a1 1 0 1 1 2 0ZM8 16h8v2H8v-2Z"/></g></g>`
	arrowDownPath              = `<g id="feArrowDown0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowDown1" fill="currentColor"><path id="feArrowDown2" d="m6 7l6 6l6-6l2 2l-8 8l-8-8z"/></g></g>`
	arrowLeftPath              = `<g id="feArrowLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowLeft1" fill="currentColor"><path id="feArrowLeft2" d="m15 4l2 2l-6 6l6 6l-2 2l-8-8z"/></g></g>`
	arrowRightPath             = `<g id="feArrowRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowRight1" fill="currentColor"><path id="feArrowRight2" d="m9.005 4l8 8l-8 8L7 18l6.005-6L7 6z"/></g></g>`
	arrowUpPath                = `<g id="feArrowUp0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowUp1" fill="currentColor"><path id="feArrowUp2" d="m4 15l8-8l8 8l-2 2l-6-6l-6 6z"/></g></g>`
	artboardPath               = `<g id="feArtboard0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArtboard1" fill="currentColor"><path id="feArtboard2" d="M6 16V8H4a1 1 0 1 1 0-2h2V4a1 1 0 1 1 2 0v2h8V4a1 1 0 0 1 2 0v2h2a1 1 0 0 1 0 2h-2v8h2a1 1 0 0 1 0 2h-2v2a1 1 0 0 1-2 0v-2H8v2a1 1 0 0 1-2 0v-2H4a1 1 0 0 1 0-2h2Zm2 0h8V8H8v8Z"/></g></g>`
	audioPlayerPath            = `<g id="feAudioPlayer0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAudioPlayer1" fill="currentColor"><path id="feAudioPlayer2" d="M18 20a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v16Zm-6 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8 4v6h8V4H8Zm4 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	backwardPath               = `<g id="feBackward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBackward1" fill="currentColor"><path id="feBackward2" d="M12 12c0-.107.023-.216.072-.316a.617.617 0 0 1 .221-.253l6.86-4.35c.276-.174.623-.06.775.254a.725.725 0 0 1 .072.316v8.698c0 .36-.255.651-.57.651a.516.516 0 0 1-.277-.082l-6.86-4.349A.671.671 0 0 1 12 12v4.35c0 .36-.255.651-.57.651a.516.516 0 0 1-.277-.082l-6.86-4.349c-.275-.174-.374-.57-.221-.885a.617.617 0 0 1 .221-.253l6.86-4.35c.276-.174.623-.06.775.254a.725.725 0 0 1 .072.316V12Z"/></g></g>`
	barPath                    = `<g id="feBar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBar1" fill="currentColor"><path id="feBar2" d="M3 16h18v2H3v-2Zm0-5h18v2H3v-2Zm0-5h18v2H3V6Z"/></g></g>`
	barChartPath               = `<g id="feBarChart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBarChart1" fill="currentColor"><path id="feBarChart2" d="M5 19h15a1 1 0 0 1 0 2H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 2 0v7Zm5-4a1 1 0 0 1-2 0V6a1 1 0 1 1 2 0v9Zm2 0V8a1 1 0 0 1 2 0v7a1 1 0 0 1-2 0Zm4-11a1 1 0 0 1 2 0v11a1 1 0 0 1-2 0V4Z"/></g></g>`
	beerPath                   = `<g id="feBeer0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBeer1" fill="currentColor"><path id="feBeer2" d="M17 16v3c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2h2a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2ZM5 5v3h10V5H5Zm0 5v9h10v-9H5Zm4 6a1 1 0 0 1-2 0v-3a1 1 0 0 1 2 0v3Zm4 0a1 1 0 0 1-2 0v-3a1 1 0 0 1 2 0v3Zm4-7h2v5h-2V9Z"/></g></g>`
	beginnerPath               = `<g id="feBeginner0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBeginner1" fill="currentColor" fill-rule="nonzero"><path id="feBeginner2" d="M12 7.529L5 5.09v10.372l7 3.251V7.529ZM5.632 3.108L12 5.326l6.368-2.218c1.047-.365 2.18.227 2.53 1.322c.067.213.102.436.102.66v10.372c0 .826-.465 1.574-1.188 1.91L12 21l-7.812-3.628C3.465 17.036 3 16.288 3 15.462V5.09C3 3.936 3.895 3 5 3c.215 0 .429.037.632.108Z"/></g></g>`
	bellPath                   = `<g id="feBell0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBell1" fill="currentColor"><path id="feBell2" d="M15 18v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6.002 6.002 0 0 1 5-5.917V3a1 1 0 0 1 2 0v1.083c2.838.476 5 2.944 5 5.917v6h1c.55 0 1 .45 1 1s-.45 1-1 1h-4Zm-7-2h8v-6a4 4 0 1 0-8 0v6Zm4 4a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1Z"/></g></g>`
	birthdayCakePath           = `<g id="feBirthdayCake0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBirthdayCake1" fill="currentColor" fill-rule="nonzero"><path id="feBirthdayCake2" d="M17 10a4 4 0 0 1 4 4v5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5a4 4 0 0 1 4-4V7h2v3h2V7h2v3h2V7h2v3ZM7 12a2 2 0 0 0-2 2v1h14v-1a2 2 0 0 0-2-2H7Zm-2 5v2h14v-2H5ZM7 4a1 1 0 0 0 1-1a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V4Zm4 0a1 1 0 0 0 1-1a1 1 0 0 1 1 1v1a1 1 0 0 1-2 0V4Zm4 0a1 1 0 0 0 1-1a1 1 0 0 1 1 1v1a1 1 0 0 1-2 0V4Z"/></g></g>`
	boldPath                   = `<g id="feBold0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBold1" fill="currentColor"><path id="feBold2" d="M5 20V4c0-1.1.9-2 2-2h5c3 0 5.966 1.4 6 4.919c0 1.838-.931 3.73-2.532 4.324v.135c2.033.514 3.532 2.027 3.532 4.73c0 1.022-.203 1.905-.573 2.653C17.337 20.96 15.09 22 12 22H7c-1.1 0-2-.9-2-2Zm3-1h4c2.566 0 4-1.2 4-3s-1.408-3-4-3H8v6Zm0-9h3c2.388 0 4-.85 4-2.55C15 5.75 13.5 5 10.996 5H8v5Z"/></g></g>`
	boltPath                   = `<g id="feBolt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBolt1" fill="currentColor"><path id="feBolt2" d="M18 2h-8L6 13h4l-2 9l9-13h-4.995z"/></g></g>`
	bookPath                   = `<g id="feBook0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBook1" fill="currentColor" fill-rule="nonzero"><path id="feBook2" d="m13 16.006l7-.047V5.992l-5.17.007l-1.814 1.814L13 16.006Zm-2-8.193L9.179 6.038L4 6.003v9.956l7 .047V7.813Zm-1-3.77L12 6l2-2l5.997-.008A2 2 0 0 1 22 5.989v9.97a2 2 0 0 1-1.986 2L14 18l-1.996 2L10 18l-6.014-.041a2 2 0 0 1-1.986-2V6.003a2 2 0 0 1 2-2l6 .04Z"/></g></g>`
	bookmarkPath               = `<g id="feBookmark0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBookmark1" fill="currentColor" fill-rule="nonzero"><path id="feBookmark2" d="M18 4H6v14.764l6-3l6 3V4ZM6 2h12a2 2 0 0 1 2 2v18l-8-4l-8 4V4a2 2 0 0 1 2-2Zm8 4h2v6h-2V6Z"/></g></g>`
	breadPath                  = `<g id="feBread0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBread1" fill="currentColor" fill-rule="nonzero"><path id="feBread2" d="M17 19v-8.689l.999-.577A1.998 1.998 0 0 0 19 8V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v1c0 .723.386 1.377 1.001 1.734L7 10.31V19h10Zm2 0a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7.535A3.998 3.998 0 0 1 3 8V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v1c0 1.48-.804 2.773-2 3.465V19Z"/></g></g>`
	browserPath                = `<g id="feBrowser0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBrowser1" fill="currentColor" fill-rule="nonzero"><path id="feBrowser2" d="M4 10v8h16v-8H4Zm0-4v2h2V6H4Zm4 0v2h12V6H8ZM4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/></g></g>`
	brushPath                  = `<g id="feBrush0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBrush1" fill="currentColor"><path id="feBrush2" d="M5.261 16.011A2 2 0 0 0 7.99 18.74A2.5 2.5 0 0 1 5.5 21H3v-2.5a2.5 2.5 0 0 1 2.261-2.489ZM19 3c1.1 0 2 1.006 2 2L8 18l-2-2L19 3Z"/></g></g>`
	bugPath                    = `<g id="feBug0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBug1" fill="currentColor"><path id="feBug2" d="M6 15H4a1 1 0 0 1 0-2h2v-1.337a1 1 0 0 1-1-.249L3.586 10A1 1 0 0 1 5 8.586l1.136 1.136a6.007 6.007 0 0 1 2.665-3.8L7.586 4.708A1 1 0 0 1 9 3.293l1.823 1.822a6.029 6.029 0 0 1 2.354 0L15 3.293a1 1 0 0 1 1.414 1.414L15.2 5.923a6.007 6.007 0 0 1 2.665 3.8L19 8.585A1 1 0 0 1 20.414 10L19 11.414a1 1 0 0 1-1 .25V13h2a1 1 0 0 1 0 2h-2c0 .483-.057.953-.165 1.404a1 1 0 0 1 1.165.182L20.414 18A1 1 0 0 1 19 19.414L17.586 18a.997.997 0 0 1-.216-.321A6 6 0 0 1 12 21a6 6 0 0 1-5.37-3.321a.997.997 0 0 1-.216.321L5 19.414A1 1 0 1 1 3.586 18L5 16.586a1 1 0 0 1 1.165-.182A6.016 6.016 0 0 1 6 15Zm9.874-5H8.126a4.002 4.002 0 0 1 7.748 0Zm0 6a4.002 4.002 0 0 1-7.748 0h7.748Z"/></g></g>`
	buildingPath               = `<g id="feBuilding0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBuilding1" fill="currentColor" fill-rule="nonzero"><path id="feBuilding2" d="M6 2h12a2 2 0 0 1 2 2v18H4V4a2 2 0 0 1 2-2Zm0 18h4v-4h4v4h4V4H6v16Zm7-14h3v3h-3V6Zm-5 5h3v3H8v-3Zm5 0h3v3h-3v-3ZM8 6h3v3H8V6Z"/></g></g>`
	busPath                    = `<g id="feBus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBus1" fill="currentColor" fill-rule="nonzero"><path id="feBus2" d="M16 19H8a2 2 0 1 1-4 0V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14a2 2 0 1 1-4 0Zm1-4a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM7 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM6 5v6h12V5H6Z"/></g></g>`
	cagePath                   = `<g id="feCage0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCage1" fill="currentColor"><path id="feCage2" d="M14 4.341c2.33.824 4 3.047 4 5.659v9h.5a1.5 1.5 0 0 1 0 3h-13a1.5 1.5 0 0 1 0-3H6v-9a6.002 6.002 0 0 1 4-5.659V4a2 2 0 1 1 4 0v.341ZM16 11v3h-3v-3h3Zm0-2h-3V6a4.183 4.183 0 0 1 3 3Zm-8 2h3v3H8v-3Zm0-2a4.183 4.183 0 0 1 3-3v3H8Zm8 7v3h-3v-3h3Zm-8 0h3v3H8v-3Z"/></g></g>`
	calendarPath               = `<g id="feCalendar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCalendar1" fill="currentColor"><path id="feCalendar2" d="M8 4h8V2h2v2h1a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1V2h2v2ZM5 8v12h14V8H5Zm2 3h2v2H7v-2Zm4 0h2v2h-2v-2Zm4 0h2v2h-2v-2Zm0 4h2v2h-2v-2Zm-4 0h2v2h-2v-2Zm-4 0h2v2H7v-2Z"/></g></g>`
	cameraPath                 = `<g id="feCamera0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCamera1" fill="currentColor"><path id="feCamera2" d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5ZM17 5v2h2V5h-2Zm-5 12a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`
	carPath                    = `<g id="feCar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCar1" fill="currentColor" fill-rule="nonzero"><path id="feCar2" d="M3 18v-5a2 2 0 0 1 2-2V8a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v3a2 2 0 0 1 2 2v5a2 2 0 1 1-4 0H7a2 2 0 1 1-4 0ZM9 6a2 2 0 0 0-2 2v3h10V8a2 2 0 0 0-2-2H9Zm-3 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm12 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	cartPath                   = `<g id="feCart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCart1" fill="currentColor"><path id="feCart2" d="M8 16h8a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V4H5a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1v2.001L8.073 5h9.854C19.072 5 20 5.895 20 7c0 .146-.017.291-.05.434l-1.151 5c-.21.915-1.052 1.566-2.024 1.566H8.073L8 13.999V16Zm-.5 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm9 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM8 7v5h8.831L18 7H8Z"/></g></g>`
	checkPath                  = `<g id="feCheck0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheck1" fill="currentColor"><path id="feCheck2" d="m6 10l-2 2l6 6L20 8l-2-2l-8 8z"/></g></g>`
	checkCirclePath            = `<g id="feCheckCircle0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheckCircle1" fill="currentColor" fill-rule="nonzero"><path id="feCheckCircle2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10ZM8 10l-2 2l5 5l7-7l-2-2l-5 5l-3-3Z"/></g></g>`
	checkCircleOPath           = `<g id="feCheckCircleO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheckCircleO1" fill="currentColor" fill-rule="nonzero"><path id="feCheckCircleO2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10ZM8 10l-2 2l5 5l7-7l-2-2l-5 5l-3-3Z"/></g></g>`
	checkVerifiedPath          = `<g id="feCheckVerified0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheckVerified1" fill="currentColor"><path id="feCheckVerified2" d="M4.252 14H4a2 2 0 1 1 0-4h.252c.189-.734.48-1.427.856-2.064l-.18-.179a2 2 0 1 1 2.83-2.828l.178.179A7.952 7.952 0 0 1 10 4.252V4a2 2 0 1 1 4 0v.252c.734.189 1.427.48 2.064.856l.179-.18a2 2 0 1 1 2.828 2.83l-.179.178c.377.637.667 1.33.856 2.064H20a2 2 0 1 1 0 4h-.252a7.952 7.952 0 0 1-.856 2.064l.18.179a2 2 0 1 1-2.83 2.828l-.178-.179a7.952 7.952 0 0 1-2.064.856V20a2 2 0 1 1-4 0v-.252a7.952 7.952 0 0 1-2.064-.856l-.179.18a2 2 0 1 1-2.828-2.83l.179-.178A7.952 7.952 0 0 1 4.252 14ZM9 10l-2 2l4 4l6-6l-2-2l-4 4l-2-2Z"/></g></g>`
	clockPath                  = `<g id="feClock0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feClock1" fill="currentColor"><path id="feClock2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm8-10a8 8 0 1 0-16 0a8 8 0 0 0 16 0Zm-4-1a1 1 0 0 1 0 2h-3c-1.1 0-2-.9-2-2V7a1 1 0 0 1 2 0v4h3Z"/></g></g>`
	closePath                  = `<g id="feClose0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feClose1" fill="currentColor"><path id="feClose2" d="M10.657 12.071L5 6.414L6.414 5l5.657 5.657L17.728 5l1.414 1.414l-5.657 5.657l5.657 5.657l-1.414 1.414l-5.657-5.657l-5.657 5.657L5 17.728z"/></g></g>`
	cloudPath                  = `<g id="feCloud0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCloud1" fill="currentColor" fill-rule="nonzero"><path id="feCloud2" d="m12.713 9.72l-1.073 1.542l-1.606-.975A2 2 0 0 0 7 12v2H5a1 1 0 0 0 0 2h11a4 4 0 1 0-3.287-6.28ZM16 6a6 6 0 1 1 0 12H5a3 3 0 0 1 0-6a4 4 0 0 1 6.071-3.423A5.993 5.993 0 0 1 16 6Z"/></g></g>`
	cocktailPath               = `<g id="feCocktail0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCocktail1" fill="currentColor" fill-rule="nonzero"><path id="feCocktail2" d="M13 19h3a1 1 0 0 1 0 2H8a1 1 0 0 1 0-2h3v-6.75L4.8 7.6A2 2 0 0 1 4 6V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1a2 2 0 0 1-.8 1.6L13 12.25V19ZM6 5v1l6 4l6-4V5H6Z"/></g></g>`
	codePath                   = `<g id="feCode0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCode1" fill="currentColor"><path id="feCode2" d="m21.004 11.657l-5.657 5.657l-1.414-1.415l4.242-4.242l-4.242-4.243L15.347 6l5.657 5.657Zm-15.176 0l4.243 4.242l-1.414 1.415L3 11.657L8.657 6l1.414 1.414l-4.243 4.243Z"/></g></g>`
	codepenPath                = `<g id="feCodepen0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCodepen1" fill="currentColor"><path id="feCodepen2" d="M22 15.047a.846.846 0 0 1-.008.112l-.006.037l-.016.072c-.003.015-.008.028-.013.042l-.022.063l-.02.042c-.008.02-.018.038-.028.057l-.025.04a.769.769 0 0 1-.108.135l-.035.034a.812.812 0 0 1-.049.04l-.038.03c-.005.003-.01.008-.015.01l-9.14 6.095a.858.858 0 0 1-.954 0l-9.14-6.094l-.014-.01a.807.807 0 0 1-.088-.071c-.012-.01-.023-.022-.034-.034c-.015-.015-.03-.03-.043-.046a.707.707 0 0 1-.066-.09a1.038 1.038 0 0 1-.054-.096l-.019-.042a.868.868 0 0 1-.022-.063c-.005-.014-.01-.027-.013-.042c-.007-.023-.01-.048-.015-.072l-.007-.037A.847.847 0 0 1 2 15.047V8.953c0-.038.003-.075.008-.112l.007-.037c.004-.024.008-.049.015-.072a.774.774 0 0 1 .145-.295a.978.978 0 0 1 .029-.038l.043-.046l.034-.034a.834.834 0 0 1 .088-.07c.005-.003.009-.008.014-.01l9.14-6.095a.86.86 0 0 1 .954 0l9.14 6.094c.005.003.01.008.015.01l.038.03a.839.839 0 0 1 .05.041l.034.034a.735.735 0 0 1 .108.136l.025.04l.029.056l.019.042l.022.063c.005.014.01.027.013.042c.007.023.011.048.016.072l.006.037a.834.834 0 0 1 .008.112v6.094ZM3.719 10.562v2.876L5.869 12l-2.15-1.438Zm7.422-2.088V4.465l-6.734 4.49l3.008 2.011l3.726-2.492Zm8.452.48L12.86 4.465v4.009l3.726 2.492l3.007-2.012ZM4.407 15.046l6.734 4.489v-4.009l-3.726-2.492l-3.008 2.012Zm8.453.48v4.009l6.733-4.49l-3.007-2.01l-3.726 2.491ZM12 9.966L8.96 12L12 14.033L15.04 12L12 9.967Zm8.281 3.472v-2.876L18.131 12l2.15 1.438Z"/></g></g>`
	coffeePath                 = `<g id="feCoffee0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCoffee1" fill="currentColor"><path id="feCoffee2" d="M4 17h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2ZM17 7h1a2 2 0 1 1 0 4h-1c-.033 0-.067 0-.1-.002A5.002 5.002 0 0 1 12 15h-2a5 5 0 0 1-5-5V7a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2ZM7 7v3a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3V7H7Z"/></g></g>`
	columnsPath                = `<g id="feColumns0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feColumns1" fill="currentColor" fill-rule="nonzero"><path id="feColumns2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm12 2v12h4V6h-4ZM4 6v12h4V6H4Zm6 0v12h4V6h-4Z"/></g></g>`
	commentPath                = `<g id="feComment0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feComment1" fill="currentColor"><path id="feComment2" d="M7.268 18.732L5 21v-4.157c-1.25-1.46-2-3.319-2-5.343C3 6.806 7.03 3 12 3s9 3.806 9 8.5s-4.03 8.5-9 8.5a9.352 9.352 0 0 1-4.732-1.268Z"/></g></g>`
	commentOPath               = `<g id="feCommentO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCommentO1" fill="currentColor" fill-rule="nonzero"><path id="feCommentO2" d="M5 21v-4.157c-1.25-1.46-2-3.319-2-5.343C3 6.806 7.03 3 12 3s9 3.806 9 8.5s-4.03 8.5-9 8.5a9.352 9.352 0 0 1-4.732-1.268L5 21Zm7-3c3.866 0 7-2.91 7-6.5S15.866 5 12 5s-7 2.91-7 6.5S8.134 18 12 18Z"/></g></g>`
	commentingPath             = `<g id="feCommenting0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCommenting1" fill="currentColor" fill-rule="nonzero"><path id="feCommenting2" d="M5 21v-4.157c-1.25-1.46-2-3.319-2-5.343C3 6.806 7.03 3 12 3s9 3.806 9 8.5s-4.03 8.5-9 8.5a9.352 9.352 0 0 1-4.732-1.268L5 21Zm7-3c3.866 0 7-2.91 7-6.5S15.866 5 12 5s-7 2.91-7 6.5S8.134 18 12 18Zm-2.5-5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm5 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/></g></g>`
	commentsPath               = `<g id="feComments0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feComments1" fill="currentColor"><path id="feComments2" d="M4.368 14.632L3 16v-2.8A5.64 5.64 0 0 1 2 10c0-3.314 2.91-6 6.5-6c3.254 0 5.95 2.207 6.425 5.088A6.57 6.57 0 0 1 16 9c3.314 0 6 2.462 6 5.5c0 1.125-.368 2.17-1 3.041V20l-1.225-1.225A6.32 6.32 0 0 1 16 20c-2.825 0-5.194-1.79-5.831-4.2c-.533.13-1.092.2-1.669.2a6.81 6.81 0 0 1-4.132-1.368ZM8.5 14c2.52 0 4.5-1.828 4.5-4c0-2.172-1.98-4-4.5-4S4 7.828 4 10c0 2.172 1.98 4 4.5 4Zm3.546 1.03C12.336 16.687 13.972 18 16 18c2.24 0 4-1.6 4-3.5S18.24 11 16 11c-.389 0-.763.048-1.117.138c-.338 1.626-1.387 3.018-2.837 3.891Z"/></g></g>`
	compressPath               = `<g id="feCompress0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCompress1" fill="currentColor"><path id="feCompress2" d="M18 12h-6V6l2 2l3-3l2 2l-3 3l2 2ZM6 12h6v6l-2-2l-3 3l-2-2l3-3l-2-2Z"/></g></g>`
	creditCardPath             = `<g id="feCreditCard0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCreditCard1" fill="currentColor" fill-rule="nonzero"><path id="feCreditCard2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm0 2v3h16V6H4Zm0 5v7h16v-7H4Zm2 3v2h4v-2H6Zm6 0v2h4v-2h-4Z"/></g></g>`
	cropPath                   = `<g id="feCrop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCrop1" fill="currentColor"><path id="feCrop2" d="M16 20v-2H8c-1.1 0-2-.9-2-2V8H4a1 1 0 1 1 0-2h2V4a1 1 0 1 1 2 0v2h9l2-2l1 1l-2 2v9h2a1 1 0 0 1 0 2h-2v2a1 1 0 0 1-2 0Zm0-4V9l-7 7h7ZM8 8v7l7-7H8Z"/></g></g>`
	cryPath                    = `<g id="feCry0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCry1" fill="currentColor"><path id="feCry2" d="M9.5 12a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5.5 4a3 3 0 0 0-6 0h1s.317-2 2-2s1.996 2 1.996 2H15Zm-.5-4a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4-7a1 1 0 0 0 1-1c0-.368-.333-1.035-1-2c-.667.965-1 1.632-1 2a1 1 0 0 0 1 1Z"/></g></g>`
	cutleryPath                = `<g id="feCutlery0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCutlery1" fill="currentColor"><path id="feCutlery2" d="M9 13v8a1 1 0 0 1-2 0v-8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2v6h2V2a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2Zm9-9v17a1 1 0 0 1-2 0v-6h-1a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2Z"/></g></g>`
	deleteLinkPath             = `<g id="feDeleteLink0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDeleteLink1" fill="currentColor" fill-rule="nonzero"><path id="feDeleteLink2" d="M11 9H7a3 3 0 0 0 0 6h4v2H7A5 5 0 0 1 7 7h4v2Zm2 6h4a3 3 0 0 0 0-6h-4V7h4a5 5 0 0 1 0 10h-4v-2Zm0-4h2a1 1 0 0 1 0 2h-2v-2Zm-4 0h2v2H9a1 1 0 0 1 0-2Z"/></g></g>`
	desktopPath                = `<g id="feDesktop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDesktop1" fill="currentColor"><path id="feDesktop2" d="M13 18h2a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2h2v-2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-7v2ZM4 6v8h16V6H4Z"/></g></g>`
	diamondPath                = `<g id="feDiamond0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDiamond1" fill="currentColor" fill-rule="nonzero"><path id="feDiamond2" d="m12 17.876l6.626-7.952L16.164 5H7.836L5.374 9.924L12 17.876ZM6.6 3h10.8l3.6 7.2L12 21L3 10.2L6.6 3Z"/></g></g>`
	differencePath             = `<g id="feDifference0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDifference1" fill="currentColor"><path id="feDifference2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-4-6h4v4h4V5h-8v4Zm-6 2v8h8v-4H9v-4H5Z"/></g></g>`
	disabledPath               = `<g id="feDisabled0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDisabled1" fill="currentColor"><path id="feDisabled2" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm-4.906-3.68L18.32 7.094A8 8 0 0 1 7.094 18.32ZM5.68 16.906A8 8 0 0 1 16.906 5.68L5.68 16.906Z"/></g></g>`
	disappointedPath           = `<g id="feDisappointed0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDisappointed1" fill="currentColor"><path id="feDisappointed2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm4-3h-1.339s-.417-2.667-2.661-2.667S9.333 17 9.333 17H8a4 4 0 1 1 8 0Zm-2.234-7.808L13 9V8l4 1v1l-1.237-.31a1.5 1.5 0 1 1-1.997-.5Zm-5.53.499L7 10V9l4-1v1l-.766.192a1.5 1.5 0 1 1-1.997.499Z"/></g></g>`
	distributeHorizontallyPath = `<g id="feDistributeHorizontally0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDistributeHorizontally1" fill="currentColor"><path id="feDistributeHorizontally2" d="M7 21a1 1 0 0 1-2 0V3a1 1 0 1 1 2 0v18Zm12 0a1 1 0 0 1-2 0V3a1 1 0 0 1 2 0v18Zm-5-4a2 2 0 1 1-4 0V7a2 2 0 1 1 4 0v10Z"/></g></g>`
	distributeVerticallyPath   = `<g id="feDistributeVertically0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDistributeVertically1" fill="currentColor"><path id="feDistributeVertically2" d="M7 10h10a2 2 0 1 1 0 4H7a2 2 0 1 1 0-4Zm-4 7h18a1 1 0 0 1 0 2H3a1 1 0 0 1 0-2ZM3 5h18a1 1 0 0 1 0 2H3a1 1 0 1 1 0-2Z"/></g></g>`
	documentPath               = `<g id="feDocument0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDocument1" fill="currentColor" fill-rule="nonzero"><path id="feDocument2" d="M15 4H6v16h12V7h-3V4ZM6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm2 9h8v2H8v-2Zm0 4h8v2H8v-2Z"/></g></g>`
	donutPath                  = `<g id="feDonut0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDonut1" fill="currentColor"><path id="feDonut2" d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 6C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-18h2v2h-2V4Zm6 5h2v2h-2V9ZM5 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm2 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-4 2h2v2h-2v-2Zm-4 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-5-4h2v2H6v-2ZM5 8h2v2H5V8Z"/></g></g>`
	downloadPath               = `<g id="feDownload0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDownload1" fill="currentColor"><path id="feDownload2" d="M5 19h14a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Zm8-5.825l3.243-3.242l1.414 1.414L12 17.004l-5.657-5.657l1.414-1.414L11 13.175V2h2v11.175Z"/></g></g>`
	dropDownPath               = `<g id="feDropDown0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropDown1" fill="currentColor"><path id="feDropDown2" d="m5 8l7 8l7-8z"/></g></g>`
	dropLeftPath               = `<g id="feDropLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropLeft1" fill="currentColor"><path id="feDropLeft2" d="m16 5l-8 7l8 7z"/></g></g>`
	dropRightPath              = `<g id="feDropRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropRight1" fill="currentColor"><path id="feDropRight2" d="M8 5v14l8-7z"/></g></g>`
	dropUpPath                 = `<g id="feDropUp0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropUp1" fill="currentColor"><path id="feDropUp2" d="m12 8l7 8H5z"/></g></g>`
	editPath                   = `<g id="feEdit0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEdit1" fill="currentColor"><path id="feEdit2" d="M5 20h14a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Zm-1-5L14 5l3 3L7 18H4v-3ZM15 4l2-2l3 3l-2.001 2.001L15 4Z"/></g></g>`
	ejectPath                  = `<g id="feEject0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEject1" fill="currentColor"><path id="feEject2" d="M7 16h10a1 1 0 0 1 0 2H7a1 1 0 0 1 0-2Zm5.973-9.565l4.83 6.048c.359.448.214 1.054-.324 1.353a1.34 1.34 0 0 1-.648.164H7.169C6.523 14 6 13.563 6 13.024c0-.193.068-.381.196-.541l4.831-6.048c.358-.449 1.084-.57 1.622-.271c.128.071.238.163.324.27Z"/></g></g>`
	elipsisHPath               = `<g id="feElipsisH0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feElipsisH1" fill="currentColor"><path id="feElipsisH2" d="M18 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM6 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm6 0a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g></g>`
	elipsisVPath               = `<g id="feElipsisV0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feElipsisV1" fill="currentColor"><path id="feElipsisV2" d="M12 20a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g></g>`
	equalizerPath              = `<g id="feEqualizer0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEqualizer1" fill="currentColor"><path id="feEqualizer2" d="M13.17 7a3.001 3.001 0 0 1 5.66 0H21a1 1 0 0 1 0 2h-2.17a3.001 3.001 0 0 1-5.66 0H3a1 1 0 1 1 0-2h10.17Zm-8 8a3.001 3.001 0 0 1 5.66 0H21a1 1 0 0 1 0 2H10.83a3.001 3.001 0 0 1-5.66 0H3a1 1 0 0 1 0-2h2.17ZM16 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-8 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	eraserPath                 = `<g id="feEraser0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEraser1" fill="currentColor"><path id="feEraser2" d="M18 4v16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2ZM8 4v4h8V4H8Z"/></g></g>`
	expandPath                 = `<g id="feExpand0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feExpand1" fill="currentColor"><path id="feExpand2" d="M14 4h6v6l-2-2l-4 4l-2-2l4-4l-2-2Zm-4 16H4v-6l2 2l4-4l2 2l-4 4l2 2Z"/></g></g>`
	exportPath                 = `<g id="feExport0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feExport1" fill="currentColor"><path id="feExport2" d="M13 5.828V17h-2V5.828L7.757 9.071L6.343 7.657L12 2l5.657 5.657l-1.414 1.414L13 5.828ZM4 16h2v4h12v-4h2v4c0 1.1-.9 2-2 2H6c-1.1 0-2-.963-2-2v-4Z"/></g></g>`
	eyePath                    = `<g id="feEye0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEye1" fill="currentColor"><path id="feEye2" d="M2 12c.945-4.564 5.063-8 10-8s9.055 3.436 10 8c-.945 4.564-5.063 8-10 8s-9.055-3.436-10-8Zm10 5a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`
	facebookPath               = `<g id="feFacebook0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFacebook1" fill="currentColor"><path id="feFacebook2" d="M15.725 22v-7.745h2.6l.389-3.018h-2.99V9.31c0-.874.243-1.47 1.497-1.47h1.598v-2.7a21.391 21.391 0 0 0-2.33-.12c-2.304 0-3.881 1.407-3.881 3.99v2.227H10v3.018h2.607V22H3.104C2.494 22 2 21.506 2 20.896V3.104C2 2.494 2.494 2 3.104 2h17.792C21.506 2 22 2.494 22 3.104v17.792c0 .61-.494 1.104-1.104 1.104h-5.171Z"/></g></g>`
	fastBackwardPath           = `<g id="feFastBackward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFastBackward1" fill="currentColor"><path id="feFastBackward2" d="M5 12c0-.107.023-.216.072-.316a.617.617 0 0 1 .221-.253l6.86-4.35c.276-.174.623-.06.775.254a.725.725 0 0 1 .072.316V12c0-.106.023-.215.072-.315a.617.617 0 0 1 .221-.253l6.86-4.35c.276-.174.623-.06.775.254a.725.725 0 0 1 .072.316v8.698c0 .36-.255.651-.57.651a.516.516 0 0 1-.277-.082l-6.86-4.349A.671.671 0 0 1 13 12v4.35c0 .36-.255.651-.57.651a.516.516 0 0 1-.277-.082l-6.86-4.349A.671.671 0 0 1 5 12v4c0 .552-.5 1-1 1s-1-.448-1-1V8a1 1 0 1 1 2 0v4Z"/></g></g>`
	fastForwardPath            = `<g id="feFastForward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFastForward1" fill="currentColor" fill-rule="nonzero"><path id="feFastForward2" d="M19 12c0 .228-.106.45-.293.57l-6.86 4.348a.516.516 0 0 1-.277.082c-.315 0-.57-.291-.57-.651v-4.35c0 .23-.106.451-.293.57l-6.86 4.35A.516.516 0 0 1 3.57 17c-.315 0-.57-.291-.57-.651V7.651c0-.11.025-.22.072-.316c.152-.314.5-.428.775-.253l6.86 4.349c.093.059.17.147.221.253c.049.1.072.209.072.315V7.651c0-.11.025-.22.072-.316c.152-.314.5-.428.775-.253l6.86 4.349c.093.059.17.147.221.253c.049.1.072.209.072.315V8a1 1 0 0 1 2 0v8c0 .552-.5 1-1 1s-1-.448-1-1v-4Z"/></g></g>`
	featherPath                = `<g id="feFeather0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFeather1" fill="currentColor"><path id="feFeather2" d="M5.993 17.877c.004.04.007.081.007.123v3a1 1 0 0 1-2 0v-3C4 9.163 11.163 2 20 2c0 8.162-6.111 14.896-14.007 15.877Zm.174-2.044A14.01 14.01 0 0 0 17.833 4.167A14.01 14.01 0 0 0 6.167 15.833Z"/></g></g>`
	feedPath                   = `<g id="feFeed0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFeed1" fill="currentColor"><path id="feFeed2" d="M6 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm12.971 0C18.473 12.408 11.592 5.527 3 5.029V3c9.941 0 18 8.059 18 18h-2.029Zm-5.012 0C13.478 15.17 8.829 10.522 3 10.041V8c7.18 0 13 5.82 13 13h-2.041Z"/></g></g>`
	filePath                   = `<g id="feFile0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFile1" fill="currentColor"><path id="feFile2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4Z"/></g></g>`
	fileAudioPath              = `<g id="feFileAudio0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileAudio1" fill="currentColor"><path id="feFileAudio2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM15 14.085V11h-4v5.5a1.5 1.5 0 1 1-1-1.415V10a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5.5a1.5 1.5 0 1 1-1-1.415Z"/></g></g>`
	fileExcelPath              = `<g id="feFileExcel0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileExcel1" fill="currentColor"><path id="feFileExcel2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9 2H6v16h12V7h-3V4Zm-8 8h3l2 2l2-2h3l-3 3l3 3h-3l-2-2l-2 2H7l3-3l-3-3Z"/></g></g>`
	fileImagePath              = `<g id="feFileImage0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileImage1" fill="currentColor"><path id="feFileImage2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM15 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-7 2l3.07-3L14 16l1-1l1 1v2H8v-2Z"/></g></g>`
	fileMoviePath              = `<g id="feFileMovie0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileMovie1" fill="currentColor"><path id="feFileMovie2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM12 13a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1h3Zm1 2.5l3-1.5v3l-3-1.5Z"/></g></g>`
	filePowerpointPath         = `<g id="feFilePowerpoint0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFilePowerpoint1" fill="currentColor"><path id="feFilePowerpoint2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM10 16v2H8v-8h5a3 3 0 0 1 0 6h-3Zm0-4v2h3a1 1 0 0 0 0-2h-3Z"/></g></g>`
	fileWordPath               = `<g id="feFileWord0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileWord1" fill="currentColor"><path id="feFileWord2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9 2H6v16h12V7h-3V4Zm-8 8h2l1 3l1-3h2l1 3l1-3h2l-2 6h-2l-1-3l-1 3H9l-2-6Z"/></g></g>`
	fileZipPath                = `<g id="feFileZip0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileZip1" fill="currentColor"><path id="feFileZip2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM8 4h2v2H8V4Zm2 2h2v2h-2V6ZM8 8h2v2H8V8Zm2 2h2v2h-2v-2Zm-1 4l1-1l1 1v4H9v-4Z"/></g></g>`
	filterPath                 = `<g id="feFilter0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFilter1" fill="currentColor" fill-rule="nonzero"><path id="feFilter2" d="M18 5.97V4H6v1.97l6 4.286l6-4.285ZM13 12v8l-2 2V12L4.838 7.598A2 2 0 0 1 4 5.971V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1.97a2 2 0 0 1-.838 1.628L13 12Z"/></g></g>`
	flagPath                   = `<g id="feFlag0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFlag1" fill="currentColor"><path id="feFlag2" d="M6 12v9a1 1 0 0 1-2 0V4c0-1.1.9-2 2-2h6c1.1 0 1.999.9 1.999 2H18a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-5a2 2 0 0 1-1.997-2H6Zm0-8v6h7v2h5.002V6H12V4H6Z"/></g></g>`
	folderPath                 = `<g id="feFolder0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFolder1" fill="currentColor" fill-rule="nonzero"><path id="feFolder2" d="M11 5h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h4l2 2ZM5 7v12h14V7H5Z"/></g></g>`
	folderOpenPath             = `<g id="feFolderOpen0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFolderOpen1" fill="currentColor" fill-rule="nonzero"><path id="feFolderOpen2" d="M11 5h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h4l2 2Zm-2 6v8h10v-8H9ZM5 7v12h2V9h12V7H5Z"/></g></g>`
	forkPath                   = `<g id="feFork0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFork1" fill="currentColor"><path id="feFork2" d="M9 7.83V12h3a3 3 0 0 0 3-3V7.83a3.001 3.001 0 1 1 2 0V9a5 5 0 0 1-5 5H9v2.17a3.001 3.001 0 1 1-2 0V7.83a3.001 3.001 0 1 1 2 0ZM8 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8-14a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM8 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	forwardPath                = `<g id="feForward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feForward1" fill="currentColor" fill-rule="nonzero"><path id="feForward2" d="M12 12V7.65c0-.11.025-.22.072-.316c.152-.314.5-.428.775-.253l6.86 4.349c.093.059.17.147.221.253c.153.314.054.71-.221.885l-6.86 4.35a.516.516 0 0 1-.277.081c-.315 0-.57-.291-.57-.651v-4.35c0 .23-.106.451-.293.57l-6.86 4.35A.516.516 0 0 1 4.57 17c-.315 0-.57-.291-.57-.651V7.651c0-.11.025-.22.072-.316c.152-.314.5-.428.775-.253l6.86 4.349c.093.059.17.147.221.253c.049.1.072.209.072.315Z"/></g></g>`
	frowingPath                = `<g id="feFrowing0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFrowing1" fill="currentColor"><path id="feFrowing2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4-5a4 4 0 1 0-8 0h8Zm-1.5-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`
	fryingPanPath              = `<g id="feFryingPan0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFryingPan1" fill="currentColor" fill-rule="nonzero"><path id="feFryingPan2" d="M13 18.268a2 2 0 1 1-2 0v-2.339A7.002 7.002 0 0 1 12 2a7 7 0 0 1 1 13.93v2.338ZM12 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></g>`
	gamepadPath                = `<g id="feGamepad0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGamepad1" fill="currentColor"><path id="feGamepad2" d="M13 15h-2a5 5 0 1 1-4-8h10a5 5 0 1 1-4 8Zm-5-4v-1a1 1 0 1 0-2 0v1H5a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2H8Zm9 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	gearPath                   = `<g id="feGear0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGear1" fill="currentColor" fill-rule="nonzero"><path id="feGear2" d="M11 19.938a7.96 7.96 0 0 1-3.906-1.618l-1.458 1.458l-1.414-1.414l1.458-1.458A7.96 7.96 0 0 1 4.062 13H2v-2h2.062A7.96 7.96 0 0 1 5.68 7.094L4.222 5.636l1.414-1.414L7.094 5.68A7.96 7.96 0 0 1 11 4.062V2h2v2.062a7.96 7.96 0 0 1 3.906 1.618l1.458-1.458l1.414 1.414l-1.458 1.458A7.96 7.96 0 0 1 19.938 11H22v2h-2.062a7.96 7.96 0 0 1-1.618 3.906l1.458 1.458l-1.414 1.414l-1.458-1.458A7.96 7.96 0 0 1 13 19.938V22h-2v-2.062ZM12 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></g>`
	giftPath                   = `<g id="feGift0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGift1" fill="currentColor"><path id="feGift2" d="M19 12v8a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1.17A3 3 0 1 1 12 5a3 3 0 1 1 5.83 1H19a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2Zm-8-4H5v2h6V8Zm2 0v2h6V8h-6Zm-6 4v8h10v-8H7Zm2-6h1V5a1 1 0 1 0-1 1Zm6 0a1 1 0 1 0-1-1v1h1Z"/></g></g>`
	gitPath                    = `<g id="feGit0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGit1" fill="currentColor"><path id="feGit2" d="M21.623 11.11L12.89 2.376a1.288 1.288 0 0 0-1.821 0L9.256 4.191l2.3 2.3a1.53 1.53 0 0 1 1.937 1.95l2.217 2.217a1.532 1.532 0 1 1-.918.864l-2.068-2.068v5.441a1.533 1.533 0 1 1-1.26-.045V9.36a1.532 1.532 0 0 1-.832-2.01L8.365 5.081l-5.988 5.987a1.289 1.289 0 0 0 0 1.822l8.733 8.732a1.288 1.288 0 0 0 1.821 0l8.692-8.692a1.288 1.288 0 0 0 0-1.822"/></g></g>`
	githubPath                 = `<g id="feGithub0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGithub1" fill="currentColor"><path id="feGithub2" d="M12 2C6.475 2 2 6.475 2 12a9.994 9.994 0 0 0 6.838 9.488c.5.087.687-.213.687-.476c0-.237-.013-1.024-.013-1.862c-2.512.463-3.162-.612-3.362-1.175c-.113-.288-.6-1.175-1.025-1.413c-.35-.187-.85-.65-.013-.662c.788-.013 1.35.725 1.538 1.025c.9 1.512 2.338 1.087 2.912.825c.088-.65.35-1.087.638-1.337c-2.225-.25-4.55-1.113-4.55-4.938c0-1.088.387-1.987 1.025-2.688c-.1-.25-.45-1.275.1-2.65c0 0 .837-.262 2.75 1.026a9.28 9.28 0 0 1 2.5-.338c.85 0 1.7.112 2.5.337c1.912-1.3 2.75-1.024 2.75-1.024c.55 1.375.2 2.4.1 2.65c.637.7 1.025 1.587 1.025 2.687c0 3.838-2.337 4.688-4.562 4.938c.362.312.675.912.675 1.85c0 1.337-.013 2.412-.013 2.75c0 .262.188.574.688.474A10.016 10.016 0 0 0 22 12c0-5.525-4.475-10-10-10Z"/></g></g>`
	githubAltPath              = `<g id="feGithubAlt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGithubAlt1" fill="currentColor"><path id="feGithubAlt2" d="M20.375 8.174c.163-.4.688-1.987-.163-4.137c0 0-1.312-.413-4.3 1.625c-1.25-.35-2.587-.4-3.912-.4c-1.325 0-2.662.05-3.912.4c-2.988-2.05-4.3-1.625-4.3-1.625c-.85 2.15-.325 3.737-.163 4.137C2.612 9.262 2 10.662 2 12.362c0 6.437 4.162 7.887 9.975 7.887S22 18.799 22 12.362c0-1.7-.613-3.1-1.625-4.188ZM12 19.024c-4.125 0-7.475-.187-7.475-4.187c0-.95.475-1.85 1.275-2.588c1.338-1.225 3.625-.575 6.2-.575c2.588 0 4.85-.65 6.2.575c.813.738 1.275 1.625 1.275 2.588c0 3.987-3.35 4.187-7.475 4.187Zm-3.137-6.262c-.825 0-1.5 1-1.5 2.225s.674 2.237 1.5 2.237c.825 0 1.5-1 1.5-2.237c0-1.238-.675-2.225-1.5-2.225Zm6.274 0c-.825 0-1.5.987-1.5 2.225c0 1.237.675 2.237 1.5 2.237s1.5-1 1.5-2.237c0-1.238-.662-2.225-1.5-2.225Z"/></g></g>`
	globePath                  = `<g id="feGlobe0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGlobe1" fill="currentColor"><path id="feGlobe2" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm2 0a8 8 0 1 0 16 0a8 8 0 0 0-16 0Zm4.252-1.552c.243-.508.643-1.948.932-1.948c.288 0 .183.723.441.786c.217.053.05-.323.653-.39s1.26.302 1.26.302s.712.177 1.369 0c0 0-.347-.526.078-.698c.424-.172 1.002.464 1.023.875c.021.411-.944.687-.944.687l.944.567s.277-.878.867-.887c.525-.01 1.183.95.875 1.383c-.308.433-.506.135-.506.135s-.94.952-1.236 1.123c-.295.17-.708 0-.708 0s-.175.324 0 .473c.355.366 1.277.73 1.277.73s2.835.461 2.923 1.039c.088.578-2.256 3.5-2.625 3.5H14c-.357-.63.577-2.644.577-2.644s-.48-.532-.577-.856c-.096-.324.186-.894.186-.894l-1.279-.589s-.68 0-1.017-.267c-.337-.267-.515-1.75-.515-1.75l-1.22-.894s-1.06 1.346-1.405 1.174c-.346-.172-.74-.45-.498-.957Zm6.57-2.098c-.844 0-2.875-1.074-2.534-1.65c0-.826 2.135-.702 2.593-.702c.458 0 1.58.354 2.203.85c.623.495-.263 1.203-.572 1.502c-.31.3-.845 0-1.69 0Z"/></g></g>`
	googlePath                 = `<g id="feGoogle0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGoogle1" fill="currentColor" fill-rule="nonzero"><path id="feGoogle2" d="M11.99 13.9v-3.72h9.36c.14.63.25 1.22.25 2.05c0 5.71-3.83 9.77-9.6 9.77c-5.52 0-10-4.48-10-10S6.48 2 12 2c2.7 0 4.96.99 6.69 2.61l-2.84 2.76c-.72-.68-1.98-1.48-3.85-1.48c-3.31 0-6.01 2.75-6.01 6.12s2.7 6.12 6.01 6.12c3.83 0 5.24-2.65 5.5-4.22h-5.51v-.01Z"/></g></g>`
	googlePlusPath             = `<g id="feGooglePlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGooglePlus1" fill="currentColor"><path id="feGooglePlus2" d="M8.364 11.455h6.009c.054.318.1.636.1 1.054c0 3.636-2.437 6.218-6.11 6.218A6.359 6.359 0 0 1 2 12.364A6.359 6.359 0 0 1 8.364 6c1.718 0 3.154.627 4.263 1.664L10.9 9.327c-.473-.454-1.3-.982-2.536-.982c-2.173 0-3.946 1.8-3.946 4.019c0 2.218 1.773 4.018 3.946 4.018c2.518 0 3.463-1.81 3.609-2.746h-3.61v-2.181Zm13.636 0v1.818h-1.818v1.818h-1.818v-1.818h-1.819v-1.818h1.819V9.636h1.818v1.819H22Z"/></g></g>`
	hashPath                   = `<g id="feHash0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHash1" fill="currentColor"><path id="feHash2" d="M10 15h4V9h-4v6Zm0 2v3a1 1 0 0 1-2 0v-3H5a1 1 0 0 1 0-2h3V9H5a1 1 0 1 1 0-2h3V4a1 1 0 1 1 2 0v3h4V4a1 1 0 0 1 2 0v3h3a1 1 0 0 1 0 2h-3v6h3a1 1 0 0 1 0 2h-3v3a1 1 0 0 1-2 0v-3h-4Z"/></g></g>`
	headphonePath              = `<g id="feHeadphone0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHeadphone1" fill="currentColor"><path id="feHeadphone2" d="M19 13a2 2 0 0 1 2 2v4a2 2 0 1 1-4 0v-9a5 5 0 0 0-10 0v9a2 2 0 1 1-4 0v-4a2 2 0 0 1 2-2v-3a7 7 0 1 1 14 0v3Z"/></g></g>`
	heartPath                  = `<g id="feHeart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHeart1" fill="currentColor"><path id="feHeart2" d="M12 20c-2.205-.48-9-4.24-9-11a5 5 0 0 1 9-3a5 5 0 0 1 9 3c0 6.76-6.795 10.52-9 11Z"/></g></g>`
	heartOPath                 = `<g id="feHeartO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHeartO1" fill="currentColor"><path id="feHeartO2" d="M12 20c-2.205-.48-9-4.24-9-11a5 5 0 0 1 9-3a5 5 0 0 1 9 3c0 6.76-6.795 10.52-9 11Zm0-2c3.12-.93 7-4.805 7-9a3 3 0 0 0-3-3c-1.305 0-2.638.833-4 2.5C10.638 6.833 9.305 6 8 6a3 3 0 0 0-3 3c0 4.195 3.88 8.07 7 9Z"/></g></g>`
	homePath                   = `<g id="feHome0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHome1" fill="currentColor" fill-rule="nonzero"><path id="feHome2" d="m12 5.561l-7 5.6V19h5v-4h4v4h5v-7.358a1 1 0 0 0-.375-.781L12 5.561ZM12 3l7.874 6.3A3 3 0 0 1 21 11.641V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7.839A2 2 0 0 1 3.75 9.6L12 3Z"/></g></g>`
	importPath                 = `<g id="feImport0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feImport1" fill="currentColor"><path id="feImport2" d="m13 13.175l3.243-3.242l1.414 1.414L12 17.004l-5.657-5.657l1.414-1.414L11 13.175V2h2v11.175ZM4 16h2v4h12v-4h2v4c0 1.1-.9 2-2 2H6c-1.1 0-2-.963-2-2v-4Z"/></g></g>`
	infoPath                   = `<g id="feInfo0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feInfo1" fill="currentColor" fill-rule="nonzero"><path id="feInfo2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-1-11v6h2v-6h-2Zm0-4h2v2h-2V7Z"/></g></g>`
	insertLinkPath             = `<g id="feInsertLink0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feInsertLink1" fill="currentColor" fill-rule="nonzero"><path id="feInsertLink2" d="M11 9H7a3 3 0 0 0 0 6h4v2H7A5 5 0 0 1 7 7h4v2Zm2 6h4a3 3 0 0 0 0-6h-4V7h4a5 5 0 0 1 0 10h-4v-2Zm-4-4h6a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Z"/></g></g>`
	instagramPath              = `<g id="feInstagram0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feInstagram1" fill="currentColor"><path id="feInstagram2" d="M12 2c-2.716 0-3.056.012-4.123.06c-1.064.049-1.791.218-2.427.465a4.901 4.901 0 0 0-1.772 1.153A4.902 4.902 0 0 0 2.525 5.45c-.247.636-.416 1.363-.465 2.427C2.011 8.944 2 9.284 2 12s.011 3.056.06 4.123c.049 1.064.218 1.791.465 2.427a4.903 4.903 0 0 0 1.153 1.772a4.903 4.903 0 0 0 1.772 1.153c.636.247 1.363.416 2.427.465c1.067.048 1.407.06 4.123.06s3.056-.012 4.123-.06c1.064-.049 1.791-.218 2.427-.465a4.902 4.902 0 0 0 1.772-1.153a4.902 4.902 0 0 0 1.153-1.772c.247-.636.416-1.363.465-2.427c.048-1.067.06-1.407.06-4.123s-.012-3.056-.06-4.123c-.049-1.064-.218-1.791-.465-2.427a4.902 4.902 0 0 0-1.153-1.772a4.901 4.901 0 0 0-1.772-1.153c-.636-.247-1.363-.416-2.427-.465C15.056 2.012 14.716 2 12 2m0 1.802c2.67 0 2.986.01 4.04.058c.976.045 1.505.207 1.858.344c.466.182.8.399 1.15.748c.35.35.566.684.748 1.15c.136.353.3.882.344 1.857c.048 1.055.058 1.37.058 4.041c0 2.67-.01 2.986-.058 4.04c-.045.976-.208 1.505-.344 1.858a3.1 3.1 0 0 1-.748 1.15c-.35.35-.684.566-1.15.748c-.353.136-.882.3-1.857.344c-1.054.048-1.37.058-4.041.058c-2.67 0-2.987-.01-4.04-.058c-.976-.045-1.505-.208-1.858-.344a3.098 3.098 0 0 1-1.15-.748a3.098 3.098 0 0 1-.748-1.15c-.137-.353-.3-.882-.344-1.857c-.048-1.055-.058-1.37-.058-4.041c0-2.67.01-2.986.058-4.04c.045-.976.207-1.505.344-1.858c.182-.466.399-.8.748-1.15c.35-.35.684-.566 1.15-.748c.353-.137.882-.3 1.857-.344c1.055-.048 1.37-.058 4.041-.058m0 11.531a3.333 3.333 0 1 1 0-6.666a3.333 3.333 0 0 1 0 6.666m0-8.468a5.135 5.135 0 1 0 0 10.27a5.135 5.135 0 0 0 0-10.27m6.538-.203a1.2 1.2 0 1 1-2.4 0a1.2 1.2 0 0 1 2.4 0"/></g></g>`
	intersectPath              = `<g id="feIntersect0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feIntersect1" fill="currentColor"><path id="feIntersect2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-5-1h4v-4h-4v4Z"/></g></g>`
	italicPath                 = `<g id="feItalic0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feItalic1" fill="currentColor"><path id="feItalic2" d="m8.97 19l3.75-14H10a1 1 0 1 1 0-2h8a1 1 0 0 1 0 2h-3.208L11.04 19H14a1 1 0 0 1 0 2H6a1 1 0 0 1 0-2h2.97Z"/></g></g>`
	keyPath                    = `<g id="feKey0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feKey1" fill="currentColor" fill-rule="nonzero"><path id="feKey2" d="M11.9 11H22v4.004c0 .546-.45.996-1 .996s-1-.45-1-.996V13h-2v2c0 .55-.45 1-1 1s-1-.45-1-1v-2h-4.1A5.002 5.002 0 0 1 2 12a5 5 0 0 1 9.9-1ZM7 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`
	keyboardPath               = `<g id="feKeyboard0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feKeyboard1" fill="currentColor"><path id="feKeyboard2" d="M2 8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8Zm15 6v2h3v-2h-3Zm-3 0v2h2v-2h-2Zm-7 0v2h6v-2H7Zm-3 0v2h2v-2H4Zm14-3v2h2v-2h-2Zm-3 0v2h2v-2h-2Zm-3 0v2h2v-2h-2Zm-3 0v2h2v-2H9Zm-5 0v2h4v-2H4Zm12-3v2h4V8h-4Zm-3 0v2h2V8h-2Zm-3 0v2h2V8h-2ZM7 8v2h2V8H7ZM4 8v2h2V8H4Z"/></g></g>`
	kitchenCookerPath          = `<g id="feKitchenCooker0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feKitchenCooker1" fill="currentColor" fill-rule="nonzero"><path id="feKitchenCooker2" d="M4.268 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-.268a2 2 0 0 1-3.464 0H7.732a2 2 0 0 1-3.464 0ZM4 6v12h16V6H4Zm2 6h12v4H6v-4Zm0-4h5v2H6V8Zm4 5v1h4v-1h-4Zm4-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/></g></g>`
	laptopPath                 = `<g id="feLaptop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLaptop1" fill="currentColor" fill-rule="nonzero"><path id="feLaptop2" d="M21 20H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2ZM5 6v8h14V6H5Zm0-2h14a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/></g></g>`
	layerPath                  = `<g id="feLayer0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLayer1" fill="currentColor" fill-rule="nonzero"><path id="feLayer2" d="m6 8l6 3l6-3l-6-3l-6 3Zm6.489-4.884l7.993 4.076c.486.248.661.81.391 1.257a.97.97 0 0 1-.39.359l-7.994 4.076a1.086 1.086 0 0 1-.978 0L3.518 8.808c-.486-.248-.661-.81-.391-1.257a.97.97 0 0 1 .39-.359l7.994-4.076c.304-.155.674-.155.978 0Zm0 13.766a1.07 1.07 0 0 1-.978 0l-7.993-4.147c-.486-.252-.661-.824-.391-1.278a.976.976 0 0 1 .39-.365a.818.818 0 0 1 .755 0l7.24 3.755c.303.158.673.158.977 0l7.239-3.755a.818.818 0 0 1 .754 0c.486.252.661.824.391 1.278a.976.976 0 0 1-.39.365l-7.994 4.147Zm0 4a1.07 1.07 0 0 1-.978 0l-7.993-4.147c-.486-.252-.661-.824-.391-1.278a.976.976 0 0 1 .39-.365a.818.818 0 0 1 .755 0l7.24 3.755c.303.158.673.158.977 0l7.239-3.755a.818.818 0 0 1 .754 0c.486.252.661.824.391 1.278a.976.976 0 0 1-.39.365l-7.994 4.147Z"/></g></g>`
	layoutPath                 = `<g id="feLayout0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLayout1" fill="currentColor" fill-rule="nonzero"><path id="feLayout2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm0 2v5h4V6H4Zm0 7v5h4v-5H4Zm6-7v12h10V6H10Z"/></g></g>`
	lineChartPath              = `<g id="feLineChart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLineChart1" fill="currentColor"><path id="feLineChart2" d="M5.823 14.177a2 2 0 1 1-1-1l2.354-2.354a2 2 0 1 1 3.646 0l2.354 2.354a1.993 1.993 0 0 1 1.646 0l3.354-3.354a2 2 0 1 1 1 1l-3.354 3.354a2 2 0 1 1-3.646 0l-2.354-2.354a1.993 1.993 0 0 1-1.646 0l-2.354 2.354Z"/></g></g>`
	linkPath                   = `<g id="feLink0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLink1" fill="currentColor"><path id="feLink2" d="M11.874 11h.252c.444-1.725 2.01-3 3.874-3h2a4 4 0 1 1 0 8h-2a4.002 4.002 0 0 1-3.874-3h-.252A4.002 4.002 0 0 1 8 16H6a4 4 0 1 1 0-8h2a4.002 4.002 0 0 1 3.874 3Zm-2.124.031A2 2 0 0 0 8 10H6a2 2 0 1 0 0 4h2a2 2 0 0 0 1.75-1.031a1 1 0 0 1 0-1.938ZM14 12.97A2 2 0 0 0 15.75 14h2a2 2 0 1 0 0-4h-2A2 2 0 0 0 14 11.031a1 1 0 0 1 0 1.938Z"/></g></g>`
	linkExternalPath           = `<g id="feLinkExternal0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLinkExternal1" fill="currentColor"><path id="feLinkExternal2" d="M6 8h5v2H6v8h8v-5h2v5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2Zm10.614-2H12V4h8v8h-2V7.442l-5.336 5.336l-1.414-1.414L16.614 6Z"/></g></g>`
	listBulletPath             = `<g id="feListBullet0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feListBullet1" fill="currentColor"><path id="feListBullet2" d="M10 4h10a1 1 0 0 1 0 2H10a1 1 0 1 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2ZM5 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g></g>`
	listOrderPath              = `<g id="feListOrder0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feListOrder1" fill="currentColor"><path id="feListOrder2" d="M10 4h10a1 1 0 0 1 0 2H10a1 1 0 1 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2ZM4.5 4a.5.5 0 0 1 0-1h1c.28 0 .5.22.5.5v3a.5.5 0 0 1-1 0V4.015L4.5 4Zm0 6h2c.28 0 .5.22.5.5v.5l-1.7 2h1.2c.28 0 .5.22.5.5s-.224.5-.5.5h-2c-.28 0-.5-.22-.5-.5V13l1.7-2H4.5c-.28 0-.5-.22-.5-.5s.199-.5.5-.5Zm2 11h-2a.5.5 0 1 1 0-1H6v-.5H4.5a.5.5 0 1 1 0-1H6V18H4.5a.5.5 0 1 1 0-1h2a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5Z"/></g></g>`
	listTaskPath               = `<g id="feListTask0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feListTask1" fill="currentColor"><path id="feListTask2" d="M9 13h10a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm0 4h10a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2Zm6-8h4a1 1 0 0 1 0 2h-4a1 1 0 0 1 0-2Zm-7.257 1.914L4 7.172l1.414-1.415l2.329 2.329L12.828 3l1.415 1.414l-6.5 6.5Z"/></g></g>`
	locationPath               = `<g id="feLocation0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLocation1" fill="currentColor" fill-rule="nonzero"><path id="feLocation2" d="M12 19c.437 0 1.479-1.187 2.411-3.312C15.357 13.534 16 10.874 16 9.353C16 6.924 14.183 5 12 5S8 6.924 8 9.353c0 1.52.643 4.181 1.589 6.335C10.52 17.813 11.563 19 12 19Zm0 2c-3.314 0-6-8.138-6-11.647C6 5.844 8.686 3 12 3s6 2.844 6 6.353S15.314 21 12 21Zm0-10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	lockPath                   = `<g id="feLock0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLock1" fill="currentColor"><path id="feLock2" d="M7 10V7a5 5 0 1 1 10 0v3h1a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h1Zm-1 2v8h12v-8H6Zm3-2h6V7a3 3 0 0 0-6 0v3Zm5 4h2v4h-2v-4Z"/></g></g>`
	loginPath                  = `<g id="feLogin0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLogin1" fill="currentColor"><path id="feLogin2" d="M9.586 11L7.05 8.464L8.464 7.05l4.95 4.95l-4.95 4.95l-1.414-1.414L9.586 13H3v-2h6.586ZM11 3h8c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2h-8v-2h8V5h-8V3Z"/></g></g>`
	logoutPath                 = `<g id="feLogout0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLogout1" fill="currentColor"><path id="feLogout2" d="M3 5c0-1.1.9-2 2-2h8v2H5v14h8v2H5c-1.1 0-2-.9-2-2V5Zm14.176 6L14.64 8.464l1.414-1.414l4.95 4.95l-4.95 4.95l-1.414-1.414L17.176 13H10.59v-2h6.586Z"/></g></g>`
	loopPath                   = `<g id="feLoop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLoop1" fill="currentColor"><path id="feLoop2" d="M6.876 15.124A6.002 6.002 0 0 0 17.658 14h2.09a8.003 8.003 0 0 1-14.316 2.568L3 19v-6h6l-2.124 2.124Zm10.249-6.249A6.004 6.004 0 0 0 6.34 10H4.25a8.005 8.005 0 0 1 14.32-2.57L21 5v6h-6l2.125-2.125Z"/></g></g>`
	magicPath                  = `<g id="feMagic0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMagic1" fill="currentColor"><path id="feMagic2" d="m3 5l2-2l16 16l-2 2L3 5Zm10 0l1-2l1 2l2 1l-2 1l-1 2l-1-2l-2-1l2-1ZM5 15l1-2l1 2l2 1l-2 1l-1 2l-1-2l-2-1l2-1ZM4 9l1 1l-1 1l-1-1l1-1Zm14 1l1 1l-1 1l-1-1l1-1Z"/></g></g>`
	mailPath                   = `<g id="feMail0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMail1" fill="currentColor" fill-rule="nonzero"><path id="feMail2" d="M4 10v8h16v-8l-8 3l-8-3Zm0-4v2l8 3l8-3V6H4Zm0-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/></g></g>`
	mapPath                    = `<g id="feMap0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMap1" fill="currentColor" fill-rule="nonzero"><path id="feMap2" d="m9 17l6 1.955V7.045L9 5v12ZM3 5l6-2l6 2l6-2v16l-6 2l-6-2l-6 2V5Z"/></g></g>`
	maskPath                   = `<g id="feMask0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMask1" fill="currentColor"><path id="feMask2" d="M5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Zm14 16c0-7.732-6.268-14-14-14v14h14Z"/></g></g>`
	medalPath                  = `<g id="feMedal0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMedal1" fill="currentColor"><path id="feMedal2" d="M12 22a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm0-2a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-12a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM9 2h2v4c-1.1 0-2-.9-2-2V2Zm4 0h2v2c0 1.1-.9 2-2 2V2Z"/></g></g>`
	megaphonePath              = `<g id="feMegaphone0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMegaphone1" fill="currentColor" fill-rule="nonzero"><path id="feMegaphone2" d="m4 8l.04 2.117L19.5 6.065A1.997 1.997 0 0 1 22 8v8a1.995 1.995 0 0 1-2.5 1.934L4 13.882V16a1 1 0 0 1-2 0V8a1 1 0 1 1 2 0Zm2 3.664v.68L20 16V8L6 11.664Z"/></g></g>`
	mentionPath                = `<g id="feMention0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMention1" fill="currentColor" fill-rule="nonzero"><path id="feMention2" d="M21 12a3.5 3.5 0 0 1-5.909 2.539A4 4 0 1 1 14 8.535V8h2v4a1.5 1.5 0 0 0 3 0a7 7 0 1 0-3.392 6h3.1A9 9 0 1 1 21 12Zm-9 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	messangerPath              = `<g id="feMessanger0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMessanger1" fill="currentColor"><path id="feMessanger2" d="m12.942 14.413l-2.56-2.66L5.45 14.48l5.406-5.736l2.56 2.66l4.93-2.727l-5.405 5.736ZM11.899 2C6.432 2 2 6.144 2 11.257c0 2.908 1.434 5.503 3.678 7.2V22l3.378-1.874c.9.252 1.855.388 2.843.388c5.468 0 9.9-4.145 9.9-9.257c0-5.113-4.432-9.257-9.9-9.257Z"/></g></g>`
	minusPath                  = `<g id="feMinus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMinus1" fill="currentColor"><path id="feMinus2" d="M4 11h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2Z"/></g></g>`
	mitarashiDangoPath         = `<g id="feMitarashiDango0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMitarashiDango1" fill="currentColor" fill-rule="nonzero"><path id="feMitarashiDango2" d="M13 17.83V21a1 1 0 0 1-2 0v-3.17a3.001 3.001 0 0 1-.659-5.33A2.997 2.997 0 0 1 9 10c0-1.043.533-1.963 1.341-2.5a3 3 0 1 1 3.318 0A2.997 2.997 0 0 1 15 10a2.997 2.997 0 0 1-1.341 2.5A3.001 3.001 0 0 1 13 17.83Z"/></g></g>`
	mobilePath                 = `<g id="feMobile0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMobile1" fill="currentColor" fill-rule="nonzero"><path id="feMobile2" d="M8 2h8a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm0 2v15h8V4H8Z"/></g></g>`
	moneyPath                  = `<g id="feMoney0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMoney1" fill="currentColor" fill-rule="nonzero"><path id="feMoney2" d="M4 5h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Zm14 2H6a2 2 0 0 1-2 2v6a2 2 0 0 1 2 2h12a2 2 0 0 1 2-2V9a2 2 0 0 1-2-2ZM8 9h2v6H8V9Zm6 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/></g></g>`
	moonPath                   = `<g id="feMoon0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMoon1" fill="currentColor"><path id="feMoon2" d="M12.97 3a8.02 8.02 0 0 0-4.054 7c0 4.418 3.522 8 7.866 8c1.146 0 2.236-.25 3.218-.698C18.39 19.544 15.787 21 12.849 21C7.962 21 4 16.97 4 12s3.962-9 8.849-9h.12Z"/></g></g>`
	mousePath                  = `<g id="feMouse0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMouse1" fill="currentColor" fill-rule="nonzero"><path id="feMouse2" d="M12 4a4 4 0 0 0-4 4v8a4 4 0 1 0 8 0V8a4 4 0 0 0-4-4Zm0-2a6 6 0 0 1 6 6v8a6 6 0 1 1-12 0V8a6 6 0 0 1 6-6Zm1 8a1 1 0 0 1-2 0V7a1 1 0 0 1 2 0v3Z"/></g></g>`
	musicPath                  = `<g id="feMusic0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMusic1" fill="currentColor"><path id="feMusic2" d="M20 4v13a3 3 0 1 1-2-2.83V6h-8v13a3 3 0 1 1-2-2.83V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2Z"/></g></g>`
	noticeActivePath           = `<g id="feNoticeActive0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticeActive1" fill="currentColor"><path id="feNoticeActive2" d="M15.085 4.853a2.501 2.501 0 1 1 2.572 3.142A5.99 5.99 0 0 1 18 10v6h1c.55 0 1 .45 1 1s-.45 1-1 1h-4v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6.002 6.002 0 0 1 5-5.917V3a1 1 0 0 1 2 0v1.083a5.961 5.961 0 0 1 2.085.77ZM12 20a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1Zm-4-4h8v-6a4 4 0 1 0-8 0v6Z"/></g></g>`
	noticeOffPath              = `<g id="feNoticeOff0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticeOff1" fill="currentColor"><path id="feNoticeOff2" d="M15 18v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6 6 0 0 1 .257-1.743L5 7a1.414 1.414 0 1 1 2-2l11 11h1a1 1 0 0 1 0 2h-4Zm-3 2a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1ZM8.876 4.876A5.962 5.962 0 0 1 11 4.083V3a1 1 0 0 1 2 0v1.083c2.838.476 5 2.944 5 5.917v4L8.876 4.876Z"/></g></g>`
	noticeOnPath               = `<g id="feNoticeOn0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticeOn1" fill="currentColor"><path id="feNoticeOn2" d="M15 18v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6.002 6.002 0 0 1 5-5.917V3a1 1 0 0 1 2 0v1.083c2.838.476 5 2.944 5 5.917v6h1c.55 0 1 .45 1 1s-.45 1-1 1h-4Zm-3 2a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1Zm1-6v-2h2a1 1 0 0 0 0-2h-2V8a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0Z"/></g></g>`
	noticePushPath             = `<defs><path id="feNoticePush0" d="M17 11a4 4 0 1 1 0-8a4 4 0 0 1 0 8ZM5 5h6v2H5v12h12v-6h2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Z"/></defs><g id="feNoticePush1" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feNoticePush2"><mask id="feNoticePush3" fill="#fff"><use href="#feNoticePush0"/></mask><use id="feNoticePush4" fill="currentColor" fill-rule="nonzero" href="#feNoticePush0"/></g></g>`
	octpusPath                 = `<g id="feOctpus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feOctpus1" fill="currentColor"><path id="feOctpus2" d="M15 19.236V21a1 1 0 0 1-2 0v-2h-2v2a1 1 0 0 1-2 0v-1.764A3 3 0 0 1 4 17a1 1 0 0 1 2 0a1 1 0 0 0 2 0v-2.255a7 7 0 1 1 8 0V17a1 1 0 0 0 2 0a1 1 0 0 1 2 0a3 3 0 0 1-5 2.236Z"/></g></g>`
	openMouthPath              = `<g id="feOpenMouth0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feOpenMouth1" fill="currentColor" fill-rule="nonzero"><path id="feOpenMouth2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm2.5-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`
	palettePath                = `<g id="fePalette0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePalette1" fill="currentColor"><path id="fePalette2" d="M10.525 21.892a4 4 0 0 0-6.77-4.232A9.954 9.954 0 0 1 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10c-.501 0-.994-.037-1.475-.108ZM9 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm3 5a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm9 5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	paperPlanePath             = `<g id="fePaperPlane0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePaperPlane1" fill="currentColor" fill-rule="nonzero"><path id="fePaperPlane2" d="m13.761 12.01l-10.76-1.084L3 4.074a1.074 1.074 0 0 1 1.554-.96l15.852 7.926a1.074 1.074 0 0 1 0 1.92l-15.85 7.926a1.074 1.074 0 0 1-1.554-.96v-6.852l10.76-1.064Z"/></g></g>`
	pausePath                  = `<g id="fePause0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePause1" fill="currentColor"><path id="fePause2" d="M10 18a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v12Zm7 0a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v12Z"/></g></g>`
	pencilPath                 = `<g id="fePencil0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePencil1" fill="currentColor"><path id="fePencil2" d="M3 18L15 6l3 3L6 21H3v-3ZM16 5l2-2l3 3l-2.001 2.001L16 5Z"/></g></g>`
	phonePath                  = `<g id="fePhone0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePhone1" fill="currentColor"><path id="fePhone2" d="M4.024 9L4 8.931C3.46 7.384 3 5.27 3 4c0-.55.45-1 1-1h3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-.837A16.054 16.054 0 0 0 15 17.837V17a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v3c0 .45-.55 1-1 1c-1.725 0-3.44-.456-5-1c-5.114-1.832-9.168-5.886-10.976-11Z"/></g></g>`
	picturePath                = `<g id="fePicture0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePicture1" fill="currentColor" fill-rule="nonzero"><path id="fePicture2" d="M4 6v12h16V6H4Zm0-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm3.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM7 14l2-2l2 2l4-4l3 3v3H7v-2Z"/></g></g>`
	pictureSquarePath          = `<g id="fePictureSquare0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePictureSquare1" fill="currentColor" fill-rule="nonzero"><path id="fePictureSquare2" d="M5 5v14h14V5H5Zm0-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Zm3.5 7a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM7 14l2-2l2 2l3-3l3 3v3H7v-3Z"/></g></g>`
	pieChartPath               = `<g id="fePieChart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePieChart1" fill="currentColor"><path id="fePieChart2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-18a8 8 0 1 0 8 8h-8V4Z"/></g></g>`
	pinterestPath              = `<g id="fePinterest0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePinterest1" fill="currentColor"><path id="fePinterest2" d="M12 2C6.477 2 2 6.477 2 12c0 4.236 2.636 7.854 6.355 9.311c-.087-.791-.166-2.004.035-2.868c.182-.78 1.173-4.97 1.173-4.97s-.299-.6-.299-1.485c0-1.391.806-2.428 1.81-2.428c.852 0 1.264.64 1.264 1.408c0 .858-.546 2.14-.828 3.33c-.236.995.5 1.806 1.481 1.806c1.777 0 3.143-1.874 3.143-4.579c0-2.394-1.72-4.068-4.177-4.068c-2.844 0-4.515 2.134-4.515 4.34c0 .859.331 1.78.744 2.281a.3.3 0 0 1 .07.287c-.076.316-.245.995-.278 1.135c-.044.182-.145.221-.334.133c-1.25-.581-2.03-2.407-2.03-3.874c0-3.155 2.291-6.051 6.607-6.051c3.469 0 6.165 2.472 6.165 5.775c0 3.447-2.173 6.22-5.19 6.22c-1.012 0-1.965-.526-2.291-1.149c0 0-.502 1.91-.623 2.378c-.226.868-.835 1.958-1.243 2.622c.936.289 1.93.445 2.96.445C17.524 22 22 17.522 22 12S17.523 2 12 2"/></g></g>`
	pizzaPath                  = `<g id="fePizza0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePizza1" fill="currentColor"><path id="fePizza2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-5-5h2v2l-2-2Zm10 0l-2 2v-2h2ZM9 7v2H7l2-2Zm6 0l2 2h-2V7Zm-2.5 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-5 4a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm9 1a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`
	playPath                   = `<g id="fePlay0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePlay1" fill="currentColor" fill-rule="nonzero"><path id="fePlay2" d="M6 5.912c0-.155.037-.307.107-.443c.23-.44.75-.599 1.163-.354l10.29 6.088c.14.083.255.206.332.355c.23.44.08.995-.332 1.239L7.27 18.885a.814.814 0 0 1-.415.115C6.383 19 6 18.592 6 18.089V5.912Z"/></g></g>`
	plugPath                   = `<g id="fePlug0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePlug1" fill="currentColor"><path id="fePlug2" d="M17 11v2h3a1 1 0 0 1 0 2h-3.268A2 2 0 0 1 15 16h-3a4.002 4.002 0 0 1-3.876-3.008A1.01 1.01 0 0 1 8 13H4a1 1 0 0 1 0-2h4c.042 0 .083.003.124.008A4.002 4.002 0 0 1 12 8h3a2 2 0 0 1 1.732 1H20a1 1 0 0 1 0 2h-3Z"/></g></g>`
	plusPath                   = `<g id="fePlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePlus1" fill="currentColor"><path id="fePlus2" d="M13 13v7a1 1 0 0 1-2 0v-7H4a1 1 0 0 1 0-2h7V4a1 1 0 0 1 2 0v7h7a1 1 0 0 1 0 2h-7Z"/></g></g>`
	pocketPath                 = `<g id="fePocket0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePocket1" fill="currentColor" fill-rule="nonzero"><path id="fePocket2" d="M21.902 4.194A1.82 1.82 0 0 0 20.197 3H3.813A1.823 1.823 0 0 0 2 4.814v6.035l.069 1.2c.29 2.73 1.707 5.116 3.9 6.779c.038.03.078.059.118.089l.025.018a9.897 9.897 0 0 0 3.91 1.727a10.06 10.06 0 0 0 4.05-.014a.261.261 0 0 0 .064-.023a9.906 9.906 0 0 0 3.753-1.691l.025-.018c.04-.03.08-.058.119-.09c2.192-1.663 3.609-4.048 3.898-6.778l.069-1.2V4.814a1.792 1.792 0 0 0-.098-.62Zm-4.235 6.287l-4.704 4.513a1.372 1.372 0 0 1-1.899 0L6.36 10.48a1.371 1.371 0 1 1 1.898-1.979l3.756 3.601l3.756-3.6a1.372 1.372 0 0 1 1.898 1.978Z"/></g></g>`
	potPath                    = `<g id="fePot0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePot1" fill="currentColor"><path id="fePot2" d="M20 11v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2h-1ZM6 11v7h12v-7H6Zm5-5V5a1 1 0 0 1 2 0v1h6a1 1 0 0 1 0 2H5a1 1 0 1 1 0-2h6Z"/></g></g>`
	printPath                  = `<g id="fePrint0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePrint1" fill="currentColor"><path id="fePrint2" d="M7 11V4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7h1a2 2 0 0 1 2 2v3c0 1.1-.9 2-2 2h-1v2a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-2H6c-1.1 0-2-.9-2-2v-3a2 2 0 0 1 2-2h1Zm2-7v7h6V4H9Zm5 9v2h2v-2h-2Zm-5 5v2h6v-2H9Z"/></g></g>`
	prototypePath              = `<g id="fePrototype0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePrototype1" fill="currentColor"><path id="fePrototype2" d="M9 12H6.732a2 2 0 1 1 0-2H9V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-7Zm2 0v7h8V5h-8v5h2.268a2 2 0 1 1 0 2H11Z"/></g></g>`
	questionPath               = `<g id="feQuestion0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feQuestion1" fill="currentColor"><path id="feQuestion2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-1-4h2v2h-2v-2Zm0-1.992s2-.008 2 0C13 13.006 16 12 16 10c0-2.21-1.773-4-3.991-4A4 4 0 0 0 8 10h2c0-1.1.9-2 2-2s2 .9 2 2c0 .9-3 2.367-3 4.008Z"/></g></g>`
	quoteLeftPath              = `<g id="feQuoteLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feQuoteLeft1" fill="currentColor"><path id="feQuoteLeft2" d="M7 21a4 4 0 0 1-4-4c0-1.473 1.333-6.14 4-14h2L7 13a4 4 0 1 1 0 8Zm10 0a4 4 0 0 1-4-4c0-1.473 1.333-6.14 4-14h2l-2 10a4 4 0 1 1 0 8Z"/></g></g>`
	quoteRightPath             = `<g id="feQuoteRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feQuoteRight1" fill="currentColor"><path id="feQuoteRight2" d="M17 11a4 4 0 1 1 4-4c0 1.473-1.333 6.14-4 14h-2l2-10ZM7 11a4 4 0 1 1 4-4c0 1.473-1.333 6.14-4 14H5l2-10Z"/></g></g>`
	ragePath                   = `<g id="feRage0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRage1" fill="currentColor"><path id="feRage2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4-5a4 4 0 1 0-8 0h1.333s.423-2.667 2.667-2.667c2.244 0 2.661 2.667 2.661 2.667H16Zm-1.5-6A1.5 1.5 0 0 0 16 9.5V8c-.828 0-3 .672-3 1.5a1.5 1.5 0 0 0 1.5 1.5Zm-5 0A1.5 1.5 0 0 0 11 9.5C11 8.672 8.828 8 8 8v1.5A1.5 1.5 0 0 0 9.5 11Z"/></g></g>`
	randomPath                 = `<g id="feRandom0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRandom1" fill="currentColor"><path id="feRandom2" d="M4 17a1 1 0 0 1 0-2h2l3-3l-3-3H4a1.001 1.001 0 0 1 0-2h3l4 4l4-4h2V5l4 3.001L17 11V9h-1l-3 3l3 3h1v-2l4 3l-4 3v-2h-2l-4-4l-4 4H4Z"/></g></g>`
	removeCartPath             = `<g id="feRemoveCart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRemoveCart1" fill="currentColor"><path id="feRemoveCart2" d="M8 16h8a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V4H5a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1v9h8.775L18 7V5.001c1.145 0 2 .894 2 1.999c0 .146-.017.291-.05.434l-1.151 5c-.21.915-1.052 1.566-2.024 1.566H8.073L8 13.999V16Zm-.5 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm9 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM15 5a1 1 0 0 1 0 2h-4a1 1 0 0 1 0-2h4Z"/></g></g>`
	riceCrackerPath            = `<g id="feRiceCracker0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRiceCracker1" fill="currentColor" fill-rule="nonzero"><path id="feRiceCracker2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-5-3.755V13a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v5.245a8 8 0 1 0-10 0Z"/></g></g>`
	rocketPath                 = `<g id="feRocket0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRocket1" fill="currentColor"><path id="feRocket2" d="M14 22h-4c-.8 0-1.602-1.123-2.274-2.726L5 22l-1-1v-4l2.383-2.383C6.14 13.325 6 12.057 6 11c0-4 3-9 6-9s6 5 6 9c0 1.058-.14 2.325-.383 3.617L20 17v4l-1 1l-2.726-2.726C15.602 20.877 14.801 22 14 22Zm-2-2h-1.615a3.136 3.136 0 0 1-.179-.249c-.347-.532-.72-1.365-1.059-2.383C8.455 15.29 8 12.755 8 11c0-3.198 2.444-7 4-7s4 3.802 4 7c0 1.755-.455 4.291-1.147 6.368c-.34 1.018-.712 1.85-1.06 2.383a3.136 3.136 0 0 1-.178.249H12Zm0-8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	scalePath                  = `<g id="feScale0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feScale1" fill="currentColor"><path id="feScale2" d="M10.874 17.998A4.002 4.002 0 0 1 3 17a4.002 4.002 0 0 1 3.002-3.874A5.001 5.001 0 0 1 9.03 8.404a6 6 0 1 1 6.567 6.567a5.001 5.001 0 0 1-4.723 3.027Zm.252-9.996a5 5 0 0 1 4.872 4.872A4.002 4.002 0 0 0 15 5a4.002 4.002 0 0 0-3.874 3.002Zm-.253 7.995a3 3 0 1 0-2.87-2.87a4.007 4.007 0 0 1 2.87 2.87ZM7 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	searchPath                 = `<g id="feSearch0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSearch1" fill="currentColor"><path id="feSearch2" d="m16.325 14.899l5.38 5.38a1.008 1.008 0 0 1-1.427 1.426l-5.38-5.38a8 8 0 1 1 1.426-1.426ZM10 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g></g>`
	searchMinusPath            = `<g id="feSearchMinus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSearchMinus1" fill="currentColor"><path id="feSearchMinus2" d="m16.325 14.899l5.38 5.38a1.008 1.008 0 0 1-1.427 1.426l-5.38-5.38a8 8 0 1 1 1.426-1.426ZM10 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm-3-5a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2H7Z"/></g></g>`
	searchPlusPath             = `<g id="feSearchPlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSearchPlus1" fill="currentColor"><path id="feSearchPlus2" d="m16.325 14.899l5.38 5.38a1.008 1.008 0 0 1-1.427 1.426l-5.38-5.38a8 8 0 1 1 1.426-1.426ZM10 16a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3-5h-2v2a1 1 0 0 1-2 0v-2H7a1 1 0 0 1 0-2h2V7a1 1 0 1 1 2 0v2h2a1 1 0 0 1 0 2Z"/></g></g>`
	sharePath                  = `<g id="feShare0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feShare1" fill="currentColor"><path id="feShare2" d="M14.839 14.92a3 3 0 1 1-.8 1.599l-4.873-2.443a3 3 0 1 1 0-4.151l4.873-2.443a3 3 0 1 1 .8 1.599l-4.877 2.438a3.022 3.022 0 0 1 0 .962l4.877 2.438ZM17 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM7 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	shieldPath                 = `<g id="feShield0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feShield1" fill="currentColor" fill-rule="nonzero"><path id="feShield2" d="m6 16.764l6 3V4H6v12.764ZM6 2h12a2 2 0 0 1 2 2v14l-8 4l-8-4V4a2 2 0 0 1 2-2Z"/></g></g>`
	shoppingBagPath            = `<g id="feShoppingBag0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feShoppingBag1" fill="currentColor" fill-rule="nonzero"><path id="feShoppingBag2" d="M8 8V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3h3a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h3Zm-3 2v9h14v-9h-3.002v3H14v-3h-4v3H8v-3H5Zm5-2h4V5h-4v3Z"/></g></g>`
	sitemapPath                = `<g id="feSitemap0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSitemap1" fill="currentColor" fill-rule="nonzero"><path id="feSitemap2" d="M19 15a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2v-2H7v2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2v-2a2 2 0 0 1 2-2h4V9a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2v2h4a2 2 0 0 1 2 2v2ZM5 17v2h2v-2H5Zm12 0v2h2v-2h-2ZM11 5v2h2V5h-2Z"/></g></g>`
	smilePath                  = `<g id="feSmile0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSmile1" fill="currentColor"><path id="feSmile2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4-9a4 4 0 1 1-8 0h8Zm-4 7a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-2.5-9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`
	smileAltPath               = `<g id="feSmileAlt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSmileAlt1" fill="currentColor"><path id="feSmileAlt2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-3-9a3 3 0 0 0 6 0h1a4 4 0 1 1-8 0h1Zm.5-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/></g></g>`
	smileHeartPath             = `<g id="feSmileHeart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSmileHeart1" fill="currentColor"><path id="feSmileHeart2" d="M10 22a8 8 0 1 1 0-16a8 8 0 0 1 0 16Zm0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3-5a3 3 0 0 1-6 0h6Zm-5-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm6.625-5c-.827-.18-3.375-1.59-3.375-4.125a1.875 1.875 0 0 1 3.375-1.125A1.875 1.875 0 0 1 22 3.875C22 6.41 19.452 7.82 18.625 8Z"/></g></g>`
	smilePlusPath              = `<g id="feSmilePlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSmilePlus1" fill="currentColor"><path id="feSmilePlus2" d="M10 22a8 8 0 1 1 0-16a8 8 0 0 1 0 16Zm0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3-5a3 3 0 0 1-6 0h6Zm-5-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8-9h1a1 1 0 0 1 0 2h-1v1a1 1 0 0 1-2 0V6h-1a1 1 0 0 1 0-2h1V3a1 1 0 0 1 2 0v1Z"/></g></g>`
	speakerPath                = `<g id="feSpeaker0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSpeaker1" fill="currentColor"><path id="feSpeaker2" d="M18 20a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v16Zm-6-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	squidPath                  = `<g id="feSquid0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSquid1" fill="currentColor"><path id="feSquid2" d="M8 10L6 8l6-6l6 6l-2 2v7a1 1 0 0 0 2 0a1 1 0 0 1 2 0a3 3 0 0 1-5 2.236V21a1 1 0 0 1-2 0v-2h-2v2a1 1 0 0 1-2 0v-1.764A3 3 0 0 1 4 17a1 1 0 0 1 2 0a1 1 0 0 0 2 0v-7Z"/></g></g>`
	starPath                   = `<g id="feStar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStar1" fill="currentColor" fill-rule="nonzero"><path id="feStar2" d="M12.5 17.925L6.629 21l1.121-6.512L3 9.875l6.564-.95L12.5 3l2.936 5.925l6.564.95l-4.75 4.613L18.371 21z"/></g></g>`
	starOPath                  = `<g id="feStarO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStarO1" fill="currentColor" fill-rule="nonzero"><path id="feStarO2" d="m9.282 17.362l3.218-1.685l3.218 1.685l-.615-3.57l2.604-2.527l-3.598-.52L12.5 7.496l-1.609 3.247l-3.598.52l2.604 2.529l-.615 3.57Zm3.218.563L6.629 21l1.121-6.512L3 9.875l6.564-.95L12.5 3l2.936 5.925l6.564.95l-4.75 4.613L18.371 21L12.5 17.925Z"/></g></g>`
	stepBackwardPath           = `<g id="feStepBackward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStepBackward1" fill="currentColor" fill-rule="nonzero"><path id="feStepBackward2" d="M11 15V5H4a1 1 0 1 1 0-2h16a1 1 0 0 1 0 2h-7v10h2.24c.15 0 .297.042.421.12c.35.219.444.663.211.991l-3.24 4.57a.74.74 0 0 1-.21.199a.79.79 0 0 1-1.054-.198l-3.24-4.57A.685.685 0 0 1 8 15.714c0-.395.34-.715.76-.715H11Zm9-6a1 1 0 0 1 0 2h-5V9h5ZM8 9h1v2H4a1 1 0 0 1 0-2h4Z"/></g></g>`
	stepForwardPath            = `<g id="feStepForward0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStepForward1" fill="currentColor" fill-rule="nonzero"><path id="feStepForward2" d="M13 9v10h7a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2h7V9H8.76C8.34 9 8 8.68 8 8.285a.68.68 0 0 1 .128-.396l3.24-4.57a.79.79 0 0 1 1.054-.199a.74.74 0 0 1 .21.198l3.24 4.57a.689.689 0 0 1-.21.992a.795.795 0 0 1-.422.12H13Zm7 4a1 1 0 0 1 0 2h-5v-2h5ZM8 13h1v2H4a1 1 0 0 1 0-2h4Z"/></g></g>`
	stopPath                   = `<g id="feStop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStop1" fill="currentColor" fill-rule="nonzero"><path id="feStop2" d="M7 5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2Z"/></g></g>`
	subtractPath               = `<g id="feSubtract0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSubtract1" fill="currentColor"><path id="feSubtract2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-4-2h8V5h-8v8Z"/></g></g>`
	sunnyOPath                 = `<g id="feSunnyO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSunnyO1" fill="currentColor" fill-rule="nonzero"><path id="feSunnyO2" d="M12 18a6 6 0 1 1 0-12a6 6 0 0 1 0 12Zm0-2a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM11 2h2v3h-2V2Zm-9 9h3v2H2v-2Zm17 0h3v2h-3v-2Zm-8 8h2v3h-2v-3Zm7.621-15l1.415 1.414l-2.122 2.122L16.5 6.12L18.621 4ZM16.5 17.414L17.914 16l2.122 2.121l-1.415 1.415l-2.121-2.122ZM6.121 16l1.415 1.414l-2.122 2.122L4 18.12L6.121 16ZM4 5.414L5.414 4l2.122 2.121L6.12 7.536L4 5.414Z"/></g></g>`
	sunrisePath                = `<g id="feSunrise0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSunrise1" fill="currentColor" fill-rule="nonzero"><path id="feSunrise2" d="M21 15H3a1 1 0 0 1 0-2h1a8 8 0 1 1 16 0h1a1 1 0 0 1 0 2Zm-3 4H6a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2ZM6 13h12a6 6 0 1 0-12 0Z"/></g></g>`
	syncPath                   = `<g id="feSync0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSync1" fill="currentColor"><path id="feSync2" d="M6.876 15.124A6.002 6.002 0 0 0 17.658 14h2.09a8.003 8.003 0 0 1-14.316 2.568L3 19v-6h6l-2.124 2.124Zm10.249-6.249A6.004 6.004 0 0 0 6.34 10H4.25a8.005 8.005 0 0 1 14.32-2.57L21 5v6h-6l2.125-2.125Z"/></g></g>`
	tablePath                  = `<g id="feTable0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTable1" fill="currentColor" fill-rule="nonzero"><path id="feTable2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm9 10v4h7v-4h-7Zm-9 0v4h7v-4H4Zm9-6v4h7V8h-7ZM4 8v4h7V8H4Z"/></g></g>`
	tabletPath                 = `<g id="feTablet0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTablet1" fill="currentColor" fill-rule="nonzero"><path id="feTablet2" d="M6 4v15h12V4H6Zm0-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Z"/></g></g>`
	tagPath                    = `<g id="feTag0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTag1" fill="currentColor" fill-rule="nonzero"><path id="feTag2" d="m4.981 14.887l4.132 4.132l9.906-9.906V4.981h-4.132L4.98 14.887ZM13.486 3.58a1.98 1.98 0 0 1 1.4-.58h4.133C20.113 3 21 3.887 21 4.981v4.132c0 .526-.209 1.03-.58 1.401l-9.906 9.906a1.981 1.981 0 0 1-2.802 0L3.58 16.288a1.981 1.981 0 0 1 0-2.802l9.906-9.906ZM16 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	targetPath                 = `<g id="feTarget0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTarget1" fill="currentColor" fill-rule="nonzero"><path id="feTarget2" d="M19.938 13A8.004 8.004 0 0 1 13 19.938V22h-2v-2.062A8.004 8.004 0 0 1 4.062 13H2v-2h2.062A8.004 8.004 0 0 1 11 4.062V2h2v2.062A8.004 8.004 0 0 1 19.938 11H22v2h-2.062ZM12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`
	taxiPath                   = `<g id="feTaxi0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTaxi1" fill="currentColor" fill-rule="nonzero"><path id="feTaxi2" d="M8.314 5.059L9 3h6l.686 2.059A4.001 4.001 0 0 1 19 9v3a2 2 0 0 1 2 2v5a2 2 0 1 1-4 0H7a2 2 0 1 1-4 0v-5a2 2 0 0 1 2-2V9a4.001 4.001 0 0 1 3.314-3.941ZM9 7a2 2 0 0 0-2 2v3h10V9a2 2 0 0 0-2-2H9Zm-3 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm12 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	terminalPath               = `<g id="feTerminal0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTerminal1" fill="currentColor" fill-rule="nonzero"><path id="feTerminal2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm0 2v12h16V6H4Zm8 8h6v2h-6v-2Zm-1.015-2.429L7.45 15.107l-1.414-1.415l2.12-2.12l-2.12-2.122L7.45 8.036l3.535 3.535Z"/></g></g>`
	textAlignCenterPath        = `<g id="feTextAlignCenter0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextAlignCenter1" fill="currentColor"><path id="feTextAlignCenter2" d="M19 7H5a1 1 0 1 1 0-2h14a1 1 0 0 1 0 2Zm-2 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm2 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm-2 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/></g></g>`
	textAlignJustifyPath       = `<g id="feTextAlignJustify0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextAlignJustify1" fill="currentColor"><path id="feTextAlignJustify2" d="M19 7H5a1 1 0 1 1 0-2h14a1 1 0 0 1 0 2Zm0 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm0 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm0 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Z"/></g></g>`
	textAlignLeftPath          = `<g id="feTextAlignLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextAlignLeft1" fill="currentColor"><path id="feTextAlignLeft2" d="M19 7H5a1 1 0 1 1 0-2h14a1 1 0 0 1 0 2Zm-4 4H5a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm4 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm-4 4H5a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/></g></g>`
	textAlignRightPath         = `<g id="feTextAlignRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextAlignRight1" fill="currentColor"><path id="feTextAlignRight2" d="M19 7H5a1 1 0 1 1 0-2h14a1 1 0 0 1 0 2Zm0 4H9a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm0 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm0 4H9a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/></g></g>`
	textSizePath               = `<g id="feTextSize0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextSize1" fill="currentColor"><path id="feTextSize2" d="m13.5 16.494l2.408-7.224h2.182L21 17.995h-1.968l-.569-1.927h-2.966l-.643 1.927h-.853l.002.005h-2.707l-.782-2.65h-4.08L5.55 18H3L7 6h3l3.5 10.494ZM7 13h3L8.496 9L7 13Zm8.908 1.36h2.182l-1.094-2.909l-1.088 2.909Z"/></g></g>`
	ticketPath                 = `<g id="feTicket0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTicket1" fill="currentColor"><path id="feTicket2" d="M2 10V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2a2 2 0 1 0 0 4v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-2a2 2 0 1 0 0-4Zm3-2v8h14V8H5Zm2 2h10v4H7v-4Z"/></g></g>`
	tiledPath                  = `<g id="feTiled0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTiled1" fill="currentColor" fill-rule="nonzero"><path id="feTiled2" d="M6 4h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm7 9v5h5v-5h-5Zm-7 0v5h5v-5H6Zm7-7v5h5V6h-5ZM6 6v5h5V6H6Z"/></g></g>`
	timelinePath               = `<g id="feTimeline0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTimeline1" fill="currentColor"><path id="feTimeline2" d="M9.17 17a3.001 3.001 0 0 1 5.66 0H20l1 1l-1 1h-5.17a3.001 3.001 0 0 1-5.66 0H3l1-1l-1-1h6.17ZM12 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm2-7l-2 2l-2-2H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-3ZM7 5v5h4l1 1l1-1h4V5H7Z"/></g></g>`
	tiredPath                  = `<g id="feTired0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTired1" fill="currentColor" fill-rule="nonzero"><path id="feTired2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10ZM8.2 10L7 8.8l.8-.8L9 9.2L10.2 8l.8.8L9.8 10l1.2 1.2l-.8.8L9 10.8L7.8 12l-.8-.8L8.2 10Zm6 0L13 8.8l.8-.8L15 9.2L16.2 8l.8.8l-1.2 1.2l1.2 1.2l-.8.8l-1.2-1.2l-1.2 1.2l-.8-.8l1.2-1.2ZM12 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`
	trainPath                  = `<g id="feTrain0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTrain1" fill="currentColor"><path id="feTrain2" d="M15 18.932c-.966.068-2.002.068-3.04.068c-1.014 0-2.021 0-2.96-.063V21H6v-1l.711-1.423C5.09 18.106 4 17.11 4 15V7c0-4 3.955-4 8-4s8 0 8 4v8c0 2.092-1.095 3.09-2.717 3.566L18 20v1h-3v-2.068ZM7 16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm10 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM6 6v6h12V6H6Z"/></g></g>`
	trashPath                  = `<g id="feTrash0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTrash1" fill="currentColor" fill-rule="nonzero"><path id="feTrash2" d="M4 5h3V4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1h3a1 1 0 0 1 0 2h-1v13a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V7H4a1 1 0 1 1 0-2Zm3 2v13h10V7H7Zm2-2h6V4H9v1Zm0 4h2v9H9V9Zm4 0h2v9h-2V9Z"/></g></g>`
	trophyPath                 = `<g id="feTrophy0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTrophy1" fill="currentColor"><path id="feTrophy2" d="M6.207 9H6a2 2 0 1 1 0-4V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v1a2 2 0 1 1 0 4h-.207A5.504 5.504 0 0 1 13 12.978V17h1a2 2 0 0 1 2 2v1h1a1 1 0 0 1 0 2H7a1 1 0 0 1 0-2h1v-1a2 2 0 0 1 2-2h1v-4.022A5.504 5.504 0 0 1 6.207 9ZM8 4v3.5a3.5 3.5 0 0 0 3.5 3.5h1A3.5 3.5 0 0 0 16 7.5V4H8Z"/></g></g>`
	truckPath                  = `<g id="feTruck0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTruck1" fill="currentColor"><path id="feTruck2" d="M10 18a3 3 0 0 1-6 0a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11.001C16.105 4 17 4.895 17 5.999L20 6c2 0 2 2.896 2 4v6a2 2 0 0 1-2 2h-1a3 3 0 0 1-6 0h-3ZM4 6v9h11V6H4Zm13 2v4h3V8h-3Zm-1 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-9 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	tumblerGlassPath           = `<g id="feTumblerGlass0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTumblerGlass1" fill="currentColor" fill-rule="nonzero"><path id="feTumblerGlass2" d="M7 3h10a2 2 0 0 1 2 2v10a6 6 0 0 1-6 6h-2a6 6 0 0 1-6-6V5a2 2 0 0 1 2-2Zm0 2v9h10V5H7Z"/></g></g>`
	twitterPath                = `<g id="feTwitter0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTwitter1" fill="currentColor" fill-rule="nonzero"><path id="feTwitter2" d="M8.283 20.263c7.547 0 11.676-6.259 11.676-11.677c0-.176 0-.352-.008-.528A8.36 8.36 0 0 0 22 5.928a8.317 8.317 0 0 1-2.36.649a4.129 4.129 0 0 0 1.808-2.273a8.163 8.163 0 0 1-2.61.993A4.096 4.096 0 0 0 15.847 4a4.109 4.109 0 0 0-4.106 4.106c0 .32.04.632.104.936a11.654 11.654 0 0 1-8.46-4.29a4.115 4.115 0 0 0 1.273 5.482A4.151 4.151 0 0 1 2.8 9.722v.056a4.113 4.113 0 0 0 3.29 4.026a4.001 4.001 0 0 1-1.08.144c-.265 0-.521-.024-.77-.072a4.104 4.104 0 0 0 3.834 2.85a8.231 8.231 0 0 1-5.098 1.76c-.328 0-.656-.016-.976-.056a11.674 11.674 0 0 0 6.283 1.833"/></g></g>`
	umbrellaPath               = `<g id="feUmbrella0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUmbrella1" fill="currentColor"><path id="feUmbrella2" d="M13 20a2 2 0 1 1-2-2v-5H3a9 9 0 0 1 8-8.945V3a1 1 0 0 1 2 0v1.055A9 9 0 0 1 21 13h-8v7Zm-7.71-9h13.42a7.003 7.003 0 0 0-13.42 0Z"/></g></g>`
	underlinePath              = `<g id="feUnderline0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUnderline1" fill="currentColor"><path id="feUnderline2" d="M17 21H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm1-17v7a6 6 0 1 1-12 0V4a1 1 0 1 1 2 0v7c0 2.21 1.79 4 4 4s4-1.79 4-4V4a1 1 0 0 1 2 0Z"/></g></g>`
	unionPath                  = `<g id="feUnion0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUnion1" fill="currentColor"><path id="feUnion2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-2-2h6V5h-8v6H5v8h8v-6Z"/></g></g>`
	unlockPath                 = `<g id="feUnlock0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUnlock1" fill="currentColor"><path id="feUnlock2" d="M7 10V7a5 5 0 1 1 10 0c0 .55-.45 1-1 1s-1-.45-1-1a3 3 0 0 0-6 0v3h9a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h1Zm-1 2v8h12v-8H6Zm8 2h2v4h-2v-4Z"/></g></g>`
	uploadPath                 = `<g id="feUpload0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUpload1" fill="currentColor"><path id="feUpload2" d="M13 5.828V17h-2V5.828L7.757 9.071L6.343 7.657L12 2l5.657 5.657l-1.414 1.414L13 5.828ZM5 19h14a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z"/></g></g>`
	usbPath                    = `<g id="feUsb0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUsb1" fill="currentColor" fill-rule="nonzero"><path id="feUsb2" d="M10 4v4h6V4h-6Zm4 6v4h2v-4h-2Zm-4-8h6a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Z"/></g></g>`
	userPath                   = `<g id="feUser0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUser1" fill="currentColor" fill-rule="nonzero"><path id="feUser2" d="M12 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0 3c3.186 0 6.045.571 8 3.063V20H4v-1.937C5.955 15.57 8.814 15 12 15Z"/></g></g>`
	userMinusPath              = `<g id="feUserMinus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUserMinus1" fill="currentColor" fill-rule="nonzero"><path id="feUserMinus2" d="M12 15c3.186 0 6.045.571 8 3.063V20H4v-1.937C5.955 15.57 8.814 15 12 15Zm0-3a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm5 2a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2h-4Z"/></g></g>`
	userPlusPath               = `<g id="feUserPlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUserPlus1" fill="currentColor"><path id="feUserPlus2" d="M12 15c3.186 0 6.045.571 8 3.063V20H4v-1.937C5.955 15.57 8.814 15 12 15Zm0-3a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm8 2a1 1 0 0 1-2 0v-1h-1a1 1 0 0 1 0-2h1v-1a1 1 0 0 1 2 0v1h1a1 1 0 0 1 0 2h-1v1Z"/></g></g>`
	usersPath                  = `<g id="feUsers0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUsers1" fill="currentColor"><path id="feUsers2" d="M8 13a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm8 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-8 2a7.98 7.98 0 0 1 6 2.708V19H2v-1.292A7.98 7.98 0 0 1 8 15Zm8 4v-2.048l-.5-.567a10.057 10.057 0 0 0-1.25-1.193A8.028 8.028 0 0 1 16 15a7.98 7.98 0 0 1 6 2.708V19h-6Z"/></g></g>`
	vectorPath                 = `<g id="feVector0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feVector1" fill="currentColor" fill-rule="nonzero"><path id="feVector2" d="M8 20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2V8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2h8a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2v8a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2H8ZM8 6a2 2 0 0 1-2 2v8a2 2 0 0 1 2 2h8a2 2 0 0 1 2-2V8a2 2 0 0 1-2-2H8ZM4 4v2h2V4H4Zm14 0v2h2V4h-2Zm0 14v2h2v-2h-2ZM4 18v2h2v-2H4Z"/></g></g>`
	videoPath                  = `<g id="feVideo0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feVideo1" fill="currentColor"><path id="feVideo2" d="m15 9.649l5.646-2.137A1 1 0 0 1 22 8.448v7.109a1 1 0 0 1-1.351.936L15 14.375V16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1.649Z"/></g></g>`
	vrPath                     = `<g id="feVr0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feVr1" fill="currentColor" fill-rule="nonzero"><path id="feVr2" d="m15 17l-2-2h-2l-2 2H7a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2Zm7-3a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0v4ZM4 14a1 1 0 0 1-2 0v-4a1 1 0 1 1 2 0v4Z"/></g></g>`
	walletPath                 = `<g id="feWallet0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWallet1" fill="currentColor"><path id="feWallet2" d="M20 9c1.1 0 2 .9 2 2v2c0 1.1-.9 2-2 2v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3Zm-2 0V6H4v12h14v-3h-2c-1.1 0-2-1.1-2-2v-1.968C14 9.9 14.9 9 16 9h2Zm-2 4h2v-2h-2v2Z"/></g></g>`
	warningPath                = `<g id="feWarning0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWarning1" fill="currentColor" fill-rule="nonzero"><path id="feWarning2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-1-6h2v2h-2v-2Zm0-10h2v8h-2V6Z"/></g></g>`
	watchPath                  = `<g id="feWatch0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWatch1" fill="currentColor"><path id="feWatch2" d="M15 6.803A5.998 5.998 0 0 1 18 12c0 2.22-1.207 4.16-3 5.197V20a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2.803A5.998 5.998 0 0 1 6 12c0-2.22 1.207-4.16 3-5.197V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2.803ZM12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm2-5a1 1 0 0 1 0 2h-1c-1.1 0-2-.9-2-2v-1a1 1 0 0 1 2 0v1h1Z"/></g></g>`
	watchAltPath               = `<g id="feWatchAlt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWatchAlt1" fill="currentColor"><path id="feWatchAlt2" d="M15 6h1a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-1v2a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-2H8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2ZM8 8v8h8V8H8Z"/></g></g>`
	wineGlassPath              = `<g id="feWineGlass0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWineGlass1" fill="currentColor"><path id="feWineGlass2" d="M9 20h2v-6a4 4 0 0 1-4-4V4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v6a4 4 0 0 1-4 4v6h2a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2ZM9 4v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4H9Z"/></g></g>`
	wordpressPath              = `<g id="feWordpress0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWordpress1" fill="currentColor"><path id="feWordpress2" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 19.542c-5.261 0-9.541-4.281-9.541-9.542S6.739 2.458 12 2.458s9.542 4.28 9.542 9.542c0 5.26-4.28 9.542-9.542 9.542M3.422 12a8.579 8.579 0 0 0 4.835 7.72L4.164 8.51A8.545 8.545 0 0 0 3.422 12m14.37-.432c0-1.06-.382-1.794-.708-2.366c-.435-.707-.843-1.305-.843-2.012c0-.788.598-1.522 1.44-1.522c.039 0 .075.004.112.007A8.547 8.547 0 0 0 12 3.42a8.569 8.569 0 0 0-7.168 3.867c.202.006.392.01.553.01c.897 0 2.286-.109 2.286-.109c.462-.027.517.652.055.707c0 0-.465.055-.982.082l3.124 9.292l1.877-5.63l-1.336-3.662c-.462-.027-.9-.082-.9-.082c-.462-.027-.408-.734.055-.707c0 0 1.416.11 2.259.11c.897 0 2.286-.11 2.286-.11c.463-.027.517.652.055.707c0 0-.465.055-.982.082l3.1 9.221l.857-2.859c.37-1.187.652-2.039.652-2.773m-5.64 1.183l-2.574 7.48a8.569 8.569 0 0 0 5.272-.137a.791.791 0 0 1-.062-.119l-2.636-7.224Zm7.377-4.866c.037.274.058.567.058.882c0 .87-.163 1.849-.653 3.073l-2.62 7.576A8.574 8.574 0 0 0 20.578 12a8.537 8.537 0 0 0-1.05-4.116"/></g></g>`
	wordpressAltPath           = `<g id="feWordpressAlt0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWordpressAlt1" fill="currentColor"><path id="feWordpressAlt2" d="M19.89 7.686A8.947 8.947 0 0 1 20.99 12a8.986 8.986 0 0 1-4.47 7.77l2.746-7.939c.513-1.282.684-2.308.684-3.22c0-.331-.022-.638-.06-.925Zm-6.651.098c.541-.028 1.029-.085 1.029-.085c.484-.057.427-.77-.057-.741c0 0-1.457.114-2.397.114c-.883 0-2.368-.114-2.368-.114c-.485-.028-.541.712-.057.74c0 0 .459.058.943.086l1.401 3.838l-1.968 5.901l-3.274-9.739c.542-.028 1.03-.085 1.03-.085c.483-.057.426-.77-.058-.741c0 0-1.456.114-2.396.114c-.17 0-.368-.004-.58-.01A8.98 8.98 0 0 1 12 3.008c2.34 0 4.472.895 6.071 2.36c-.039-.002-.076-.007-.116-.007c-.884 0-1.51.77-1.51 1.596c0 .74.427 1.368.883 2.109c.342.598.741 1.368.741 2.48c0 .769-.296 1.662-.684 2.906l-.897 2.996l-3.25-9.665Zm-1.24 13.207a9.03 9.03 0 0 1-2.54-.366l2.699-7.839l2.763 7.571c.018.045.04.086.064.124a8.972 8.972 0 0 1-2.985.51ZM3.01 12c0-1.304.28-2.541.779-3.66l4.288 11.751A8.991 8.991 0 0 1 3.01 12ZM12 2C6.487 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2Z"/></g></g>`
	wrenchPath                 = `<g id="feWrench0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feWrench1" fill="currentColor"><path id="feWrench2" d="M14 11.584V20a2 2 0 1 1-4 0v-8.416a5.001 5.001 0 0 1 2.391-9.569L10 6l4 2l2.215-3.691A5.001 5.001 0 0 1 14 11.584Z"/></g></g>`
	yakiDangoPath              = `<g id="feYakiDango0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feYakiDango1" fill="currentColor" fill-rule="nonzero"><path id="feYakiDango2" d="M13 17.83V21a1 1 0 0 1-2 0v-3.17a3.001 3.001 0 0 1-.659-5.33A2.997 2.997 0 0 1 9 10c0-1.043.533-1.963 1.341-2.5a3 3 0 1 1 3.318 0A2.997 2.997 0 0 1 15 10a2.997 2.997 0 0 1-1.341 2.5A3.001 3.001 0 0 1 13 17.83ZM12 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`
	youtubePath                = `<g id="feYoutube0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feYoutube1" fill="currentColor"><path id="feYoutube2" d="M9.935 14.628v-5.62l5.403 2.82l-5.403 2.8ZM21.8 8.035s-.195-1.379-.795-1.986c-.76-.796-1.613-.8-2.004-.847C16.203 5 12.004 5 12.004 5h-.008s-4.198 0-6.997.202c-.391.047-1.243.05-2.004.847c-.6.607-.795 1.986-.795 1.986S2 9.653 2 11.272v1.517c0 1.618.2 3.237.2 3.237s.195 1.378.795 1.985c.76.797 1.76.771 2.205.855c1.6.153 6.8.2 6.8.2s4.203-.006 7.001-.208c.391-.047 1.244-.05 2.004-.847c.6-.607.795-1.985.795-1.985s.2-1.619.2-3.237v-1.517c0-1.619-.2-3.237-.2-3.237Z"/></g></g>`
)

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func IconFromName(name string) g.Node {
	switch name {
	case "activity":
		return Activity()
	case "addCart":
		return AddCart()
	case "alignBottom":
		return AlignBottom()
	case "alignCenter":
		return AlignCenter()
	case "alignLeft":
		return AlignLeft()
	case "alignRight":
		return AlignRight()
	case "alignTop":
		return AlignTop()
	case "alignVertically":
		return AlignVertically()
	case "angry":
		return Angry()
	case "appMenu":
		return AppMenu()
	case "apron":
		return Apron()
	case "arrowDown":
		return ArrowDown()
	case "arrowLeft":
		return ArrowLeft()
	case "arrowRight":
		return ArrowRight()
	case "arrowUp":
		return ArrowUp()
	case "artboard":
		return Artboard()
	case "audioPlayer":
		return AudioPlayer()
	case "backward":
		return Backward()
	case "bar":
		return Bar()
	case "barChart":
		return BarChart()
	case "beer":
		return Beer()
	case "beginner":
		return Beginner()
	case "bell":
		return Bell()
	case "birthdayCake":
		return BirthdayCake()
	case "bold":
		return Bold()
	case "bolt":
		return Bolt()
	case "book":
		return Book()
	case "bookmark":
		return Bookmark()
	case "bread":
		return Bread()
	case "browser":
		return Browser()
	case "brush":
		return Brush()
	case "bug":
		return Bug()
	case "building":
		return Building()
	case "bus":
		return Bus()
	case "cage":
		return Cage()
	case "calendar":
		return Calendar()
	case "camera":
		return Camera()
	case "car":
		return Car()
	case "cart":
		return Cart()
	case "check":
		return Check()
	case "checkCircle":
		return CheckCircle()
	case "checkCircleO":
		return CheckCircleO()
	case "checkVerified":
		return CheckVerified()
	case "clock":
		return Clock()
	case "close":
		return Close()
	case "cloud":
		return Cloud()
	case "cocktail":
		return Cocktail()
	case "code":
		return Code()
	case "codepen":
		return Codepen()
	case "coffee":
		return Coffee()
	case "columns":
		return Columns()
	case "comment":
		return Comment()
	case "commentO":
		return CommentO()
	case "commenting":
		return Commenting()
	case "comments":
		return Comments()
	case "compress":
		return Compress()
	case "creditCard":
		return CreditCard()
	case "crop":
		return Crop()
	case "cry":
		return Cry()
	case "cutlery":
		return Cutlery()
	case "deleteLink":
		return DeleteLink()
	case "desktop":
		return Desktop()
	case "diamond":
		return Diamond()
	case "difference":
		return Difference()
	case "disabled":
		return Disabled()
	case "disappointed":
		return Disappointed()
	case "distributeHorizontally":
		return DistributeHorizontally()
	case "distributeVertically":
		return DistributeVertically()
	case "document":
		return Document()
	case "donut":
		return Donut()
	case "download":
		return Download()
	case "dropDown":
		return DropDown()
	case "dropLeft":
		return DropLeft()
	case "dropRight":
		return DropRight()
	case "dropUp":
		return DropUp()
	case "edit":
		return Edit()
	case "eject":
		return Eject()
	case "elipsisH":
		return ElipsisH()
	case "elipsisV":
		return ElipsisV()
	case "equalizer":
		return Equalizer()
	case "eraser":
		return Eraser()
	case "expand":
		return Expand()
	case "export":
		return Export()
	case "eye":
		return Eye()
	case "facebook":
		return Facebook()
	case "fastBackward":
		return FastBackward()
	case "fastForward":
		return FastForward()
	case "feather":
		return Feather()
	case "feed":
		return Feed()
	case "file":
		return File()
	case "fileAudio":
		return FileAudio()
	case "fileExcel":
		return FileExcel()
	case "fileImage":
		return FileImage()
	case "fileMovie":
		return FileMovie()
	case "filePowerpoint":
		return FilePowerpoint()
	case "fileWord":
		return FileWord()
	case "fileZip":
		return FileZip()
	case "filter":
		return Filter()
	case "flag":
		return Flag()
	case "folder":
		return Folder()
	case "folderOpen":
		return FolderOpen()
	case "fork":
		return Fork()
	case "forward":
		return Forward()
	case "frowing":
		return Frowing()
	case "fryingPan":
		return FryingPan()
	case "gamepad":
		return Gamepad()
	case "gear":
		return Gear()
	case "gift":
		return Gift()
	case "git":
		return Git()
	case "github":
		return Github()
	case "githubAlt":
		return GithubAlt()
	case "globe":
		return Globe()
	case "google":
		return Google()
	case "googlePlus":
		return GooglePlus()
	case "hash":
		return Hash()
	case "headphone":
		return Headphone()
	case "heart":
		return Heart()
	case "heartO":
		return HeartO()
	case "home":
		return Home()
	case "import":
		return Import()
	case "info":
		return Info()
	case "insertLink":
		return InsertLink()
	case "instagram":
		return Instagram()
	case "intersect":
		return Intersect()
	case "italic":
		return Italic()
	case "key":
		return Key()
	case "keyboard":
		return Keyboard()
	case "kitchenCooker":
		return KitchenCooker()
	case "laptop":
		return Laptop()
	case "layer":
		return Layer()
	case "layout":
		return Layout()
	case "lineChart":
		return LineChart()
	case "link":
		return Link()
	case "linkExternal":
		return LinkExternal()
	case "listBullet":
		return ListBullet()
	case "listOrder":
		return ListOrder()
	case "listTask":
		return ListTask()
	case "location":
		return Location()
	case "lock":
		return Lock()
	case "login":
		return Login()
	case "logout":
		return Logout()
	case "loop":
		return Loop()
	case "magic":
		return Magic()
	case "mail":
		return Mail()
	case "map":
		return Map()
	case "mask":
		return Mask()
	case "medal":
		return Medal()
	case "megaphone":
		return Megaphone()
	case "mention":
		return Mention()
	case "messanger":
		return Messanger()
	case "minus":
		return Minus()
	case "mitarashiDango":
		return MitarashiDango()
	case "mobile":
		return Mobile()
	case "money":
		return Money()
	case "moon":
		return Moon()
	case "mouse":
		return Mouse()
	case "music":
		return Music()
	case "noticeActive":
		return NoticeActive()
	case "noticeOff":
		return NoticeOff()
	case "noticeOn":
		return NoticeOn()
	case "noticePush":
		return NoticePush()
	case "octpus":
		return Octpus()
	case "openMouth":
		return OpenMouth()
	case "palette":
		return Palette()
	case "paperPlane":
		return PaperPlane()
	case "pause":
		return Pause()
	case "pencil":
		return Pencil()
	case "phone":
		return Phone()
	case "picture":
		return Picture()
	case "pictureSquare":
		return PictureSquare()
	case "pieChart":
		return PieChart()
	case "pinterest":
		return Pinterest()
	case "pizza":
		return Pizza()
	case "play":
		return Play()
	case "plug":
		return Plug()
	case "plus":
		return Plus()
	case "pocket":
		return Pocket()
	case "pot":
		return Pot()
	case "print":
		return Print()
	case "prototype":
		return Prototype()
	case "question":
		return Question()
	case "quoteLeft":
		return QuoteLeft()
	case "quoteRight":
		return QuoteRight()
	case "rage":
		return Rage()
	case "random":
		return Random()
	case "removeCart":
		return RemoveCart()
	case "riceCracker":
		return RiceCracker()
	case "rocket":
		return Rocket()
	case "scale":
		return Scale()
	case "search":
		return Search()
	case "searchMinus":
		return SearchMinus()
	case "searchPlus":
		return SearchPlus()
	case "share":
		return Share()
	case "shield":
		return Shield()
	case "shoppingBag":
		return ShoppingBag()
	case "sitemap":
		return Sitemap()
	case "smile":
		return Smile()
	case "smileAlt":
		return SmileAlt()
	case "smileHeart":
		return SmileHeart()
	case "smilePlus":
		return SmilePlus()
	case "speaker":
		return Speaker()
	case "squid":
		return Squid()
	case "star":
		return Star()
	case "starO":
		return StarO()
	case "stepBackward":
		return StepBackward()
	case "stepForward":
		return StepForward()
	case "stop":
		return Stop()
	case "subtract":
		return Subtract()
	case "sunnyO":
		return SunnyO()
	case "sunrise":
		return Sunrise()
	case "sync":
		return Sync()
	case "table":
		return Table()
	case "tablet":
		return Tablet()
	case "tag":
		return Tag()
	case "target":
		return Target()
	case "taxi":
		return Taxi()
	case "terminal":
		return Terminal()
	case "textAlignCenter":
		return TextAlignCenter()
	case "textAlignJustify":
		return TextAlignJustify()
	case "textAlignLeft":
		return TextAlignLeft()
	case "textAlignRight":
		return TextAlignRight()
	case "textSize":
		return TextSize()
	case "ticket":
		return Ticket()
	case "tiled":
		return Tiled()
	case "timeline":
		return Timeline()
	case "tired":
		return Tired()
	case "train":
		return Train()
	case "trash":
		return Trash()
	case "trophy":
		return Trophy()
	case "truck":
		return Truck()
	case "tumblerGlass":
		return TumblerGlass()
	case "twitter":
		return Twitter()
	case "umbrella":
		return Umbrella()
	case "underline":
		return Underline()
	case "union":
		return Union()
	case "unlock":
		return Unlock()
	case "upload":
		return Upload()
	case "usb":
		return Usb()
	case "user":
		return User()
	case "userMinus":
		return UserMinus()
	case "userPlus":
		return UserPlus()
	case "users":
		return Users()
	case "vector":
		return Vector()
	case "video":
		return Video()
	case "vr":
		return Vr()
	case "wallet":
		return Wallet()
	case "warning":
		return Warning()
	case "watch":
		return Watch()
	case "watchAlt":
		return WatchAlt()
	case "wineGlass":
		return WineGlass()
	case "wordpress":
		return Wordpress()
	case "wordpressAlt":
		return WordpressAlt()
	case "wrench":
		return Wrench()
	case "yakiDango":
		return YakiDango()
	case "youtube":
		return Youtube()
	default:
		panic("Unknown icon name: " + name)
	}
}

func Activity(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(activityPath), g.Group(children))
}

func AddCart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(addCartPath), g.Group(children))
}

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignBottomPath), g.Group(children))
}

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignCenterPath), g.Group(children))
}

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignLeftPath), g.Group(children))
}

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignRightPath), g.Group(children))
}

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignTopPath), g.Group(children))
}

func AlignVertically(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticallyPath), g.Group(children))
}

func Angry(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(angryPath), g.Group(children))
}

func AppMenu(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(appMenuPath), g.Group(children))
}

func Apron(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(apronPath), g.Group(children))
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownPath), g.Group(children))
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftPath), g.Group(children))
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightPath), g.Group(children))
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpPath), g.Group(children))
}

func Artboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(artboardPath), g.Group(children))
}

func AudioPlayer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(audioPlayerPath), g.Group(children))
}

func Backward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(backwardPath), g.Group(children))
}

func Bar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barPath), g.Group(children))
}

func BarChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartPath), g.Group(children))
}

func Beer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beerPath), g.Group(children))
}

func Beginner(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beginnerPath), g.Group(children))
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellPath), g.Group(children))
}

func BirthdayCake(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(birthdayCakePath), g.Group(children))
}

func Bold(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boldPath), g.Group(children))
}

func Bolt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boltPath), g.Group(children))
}

func Book(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookPath), g.Group(children))
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookmarkPath), g.Group(children))
}

func Bread(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(breadPath), g.Group(children))
}

func Browser(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(browserPath), g.Group(children))
}

func Brush(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(brushPath), g.Group(children))
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bugPath), g.Group(children))
}

func Building(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(buildingPath), g.Group(children))
}

func Bus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(busPath), g.Group(children))
}

func Cage(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cagePath), g.Group(children))
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarPath), g.Group(children))
}

func Camera(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cameraPath), g.Group(children))
}

func Car(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(carPath), g.Group(children))
}

func Cart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cartPath), g.Group(children))
}

func Check(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkPath), g.Group(children))
}

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkCirclePath), g.Group(children))
}

func CheckCircleO(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkCircleOPath), g.Group(children))
}

func CheckVerified(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkVerifiedPath), g.Group(children))
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockPath), g.Group(children))
}

func Close(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(closePath), g.Group(children))
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudPath), g.Group(children))
}

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cocktailPath), g.Group(children))
}

func Code(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codePath), g.Group(children))
}

func Codepen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codepenPath), g.Group(children))
}

func Coffee(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(coffeePath), g.Group(children))
}

func Columns(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(columnsPath), g.Group(children))
}

func Comment(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(commentPath), g.Group(children))
}

func CommentO(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(commentOPath), g.Group(children))
}

func Commenting(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(commentingPath), g.Group(children))
}

func Comments(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(commentsPath), g.Group(children))
}

func Compress(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(compressPath), g.Group(children))
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(creditCardPath), g.Group(children))
}

func Crop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cropPath), g.Group(children))
}

func Cry(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cryPath), g.Group(children))
}

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cutleryPath), g.Group(children))
}

func DeleteLink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(deleteLinkPath), g.Group(children))
}

func Desktop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(desktopPath), g.Group(children))
}

func Diamond(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diamondPath), g.Group(children))
}

func Difference(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(differencePath), g.Group(children))
}

func Disabled(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(disabledPath), g.Group(children))
}

func Disappointed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(disappointedPath), g.Group(children))
}

func DistributeHorizontally(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(distributeHorizontallyPath), g.Group(children))
}

func DistributeVertically(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(distributeVerticallyPath), g.Group(children))
}

func Document(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(documentPath), g.Group(children))
}

func Donut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(donutPath), g.Group(children))
}

func Download(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(downloadPath), g.Group(children))
}

func DropDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropDownPath), g.Group(children))
}

func DropLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropLeftPath), g.Group(children))
}

func DropRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropRightPath), g.Group(children))
}

func DropUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropUpPath), g.Group(children))
}

func Edit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(editPath), g.Group(children))
}

func Eject(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ejectPath), g.Group(children))
}

func ElipsisH(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(elipsisHPath), g.Group(children))
}

func ElipsisV(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(elipsisVPath), g.Group(children))
}

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(equalizerPath), g.Group(children))
}

func Eraser(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eraserPath), g.Group(children))
}

func Expand(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(expandPath), g.Group(children))
}

func Export(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(exportPath), g.Group(children))
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eyePath), g.Group(children))
}

func Facebook(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(facebookPath), g.Group(children))
}

func FastBackward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fastBackwardPath), g.Group(children))
}

func FastForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fastForwardPath), g.Group(children))
}

func Feather(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(featherPath), g.Group(children))
}

func Feed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(feedPath), g.Group(children))
}

func File(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePath), g.Group(children))
}

func FileAudio(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileAudioPath), g.Group(children))
}

func FileExcel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileExcelPath), g.Group(children))
}

func FileImage(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileImagePath), g.Group(children))
}

func FileMovie(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileMoviePath), g.Group(children))
}

func FilePowerpoint(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePowerpointPath), g.Group(children))
}

func FileWord(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileWordPath), g.Group(children))
}

func FileZip(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileZipPath), g.Group(children))
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filterPath), g.Group(children))
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flagPath), g.Group(children))
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderPath), g.Group(children))
}

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderOpenPath), g.Group(children))
}

func Fork(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(forkPath), g.Group(children))
}

func Forward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(forwardPath), g.Group(children))
}

func Frowing(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(frowingPath), g.Group(children))
}

func FryingPan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fryingPanPath), g.Group(children))
}

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gamepadPath), g.Group(children))
}

func Gear(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gearPath), g.Group(children))
}

func Gift(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(giftPath), g.Group(children))
}

func Git(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitPath), g.Group(children))
}

func Github(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(githubPath), g.Group(children))
}

func GithubAlt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(githubAltPath), g.Group(children))
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(globePath), g.Group(children))
}

func Google(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(googlePath), g.Group(children))
}

func GooglePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(googlePlusPath), g.Group(children))
}

func Hash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hashPath), g.Group(children))
}

func Headphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headphonePath), g.Group(children))
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartPath), g.Group(children))
}

func HeartO(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartOPath), g.Group(children))
}

func Home(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(homePath), g.Group(children))
}

func Import(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(importPath), g.Group(children))
}

func Info(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(infoPath), g.Group(children))
}

func InsertLink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(insertLinkPath), g.Group(children))
}

func Instagram(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(instagramPath), g.Group(children))
}

func Intersect(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(intersectPath), g.Group(children))
}

func Italic(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(italicPath), g.Group(children))
}

func Key(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keyPath), g.Group(children))
}

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keyboardPath), g.Group(children))
}

func KitchenCooker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(kitchenCookerPath), g.Group(children))
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(laptopPath), g.Group(children))
}

func Layer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layerPath), g.Group(children))
}

func Layout(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutPath), g.Group(children))
}

func LineChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lineChartPath), g.Group(children))
}

func Link(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkPath), g.Group(children))
}

func LinkExternal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkExternalPath), g.Group(children))
}

func ListBullet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listBulletPath), g.Group(children))
}

func ListOrder(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listOrderPath), g.Group(children))
}

func ListTask(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listTaskPath), g.Group(children))
}

func Location(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(locationPath), g.Group(children))
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lockPath), g.Group(children))
}

func Login(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(loginPath), g.Group(children))
}

func Logout(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(logoutPath), g.Group(children))
}

func Loop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(loopPath), g.Group(children))
}

func Magic(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(magicPath), g.Group(children))
}

func Mail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailPath), g.Group(children))
}

func Map(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mapPath), g.Group(children))
}

func Mask(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(maskPath), g.Group(children))
}

func Medal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(medalPath), g.Group(children))
}

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(megaphonePath), g.Group(children))
}

func Mention(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mentionPath), g.Group(children))
}

func Messanger(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messangerPath), g.Group(children))
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minusPath), g.Group(children))
}

func MitarashiDango(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mitarashiDangoPath), g.Group(children))
}

func Mobile(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mobilePath), g.Group(children))
}

func Money(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moneyPath), g.Group(children))
}

func Moon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moonPath), g.Group(children))
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePath), g.Group(children))
}

func Music(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(musicPath), g.Group(children))
}

func NoticeActive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(noticeActivePath), g.Group(children))
}

func NoticeOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(noticeOffPath), g.Group(children))
}

func NoticeOn(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(noticeOnPath), g.Group(children))
}

func NoticePush(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(noticePushPath), g.Group(children))
}

func Octpus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(octpusPath), g.Group(children))
}

func OpenMouth(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(openMouthPath), g.Group(children))
}

func Palette(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(palettePath), g.Group(children))
}

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paperPlanePath), g.Group(children))
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pausePath), g.Group(children))
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pencilPath), g.Group(children))
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phonePath), g.Group(children))
}

func Picture(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(picturePath), g.Group(children))
}

func PictureSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pictureSquarePath), g.Group(children))
}

func PieChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pieChartPath), g.Group(children))
}

func Pinterest(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pinterestPath), g.Group(children))
}

func Pizza(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pizzaPath), g.Group(children))
}

func Play(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(playPath), g.Group(children))
}

func Plug(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plugPath), g.Group(children))
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plusPath), g.Group(children))
}

func Pocket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pocketPath), g.Group(children))
}

func Pot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(potPath), g.Group(children))
}

func Print(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(printPath), g.Group(children))
}

func Prototype(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(prototypePath), g.Group(children))
}

func Question(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(questionPath), g.Group(children))
}

func QuoteLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(quoteLeftPath), g.Group(children))
}

func QuoteRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(quoteRightPath), g.Group(children))
}

func Rage(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ragePath), g.Group(children))
}

func Random(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(randomPath), g.Group(children))
}

func RemoveCart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(removeCartPath), g.Group(children))
}

func RiceCracker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(riceCrackerPath), g.Group(children))
}

func Rocket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rocketPath), g.Group(children))
}

func Scale(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scalePath), g.Group(children))
}

func Search(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchPath), g.Group(children))
}

func SearchMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchMinusPath), g.Group(children))
}

func SearchPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchPlusPath), g.Group(children))
}

func Share(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sharePath), g.Group(children))
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldPath), g.Group(children))
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shoppingBagPath), g.Group(children))
}

func Sitemap(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sitemapPath), g.Group(children))
}

func Smile(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smilePath), g.Group(children))
}

func SmileAlt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smileAltPath), g.Group(children))
}

func SmileHeart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smileHeartPath), g.Group(children))
}

func SmilePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smilePlusPath), g.Group(children))
}

func Speaker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(speakerPath), g.Group(children))
}

func Squid(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squidPath), g.Group(children))
}

func Star(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(starPath), g.Group(children))
}

func StarO(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(starOPath), g.Group(children))
}

func StepBackward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stepBackwardPath), g.Group(children))
}

func StepForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stepForwardPath), g.Group(children))
}

func Stop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stopPath), g.Group(children))
}

func Subtract(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(subtractPath), g.Group(children))
}

func SunnyO(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunnyOPath), g.Group(children))
}

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunrisePath), g.Group(children))
}

func Sync(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(syncPath), g.Group(children))
}

func Table(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tablePath), g.Group(children))
}

func Tablet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tabletPath), g.Group(children))
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tagPath), g.Group(children))
}

func Target(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(targetPath), g.Group(children))
}

func Taxi(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(taxiPath), g.Group(children))
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(terminalPath), g.Group(children))
}

func TextAlignCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textAlignCenterPath), g.Group(children))
}

func TextAlignJustify(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textAlignJustifyPath), g.Group(children))
}

func TextAlignLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textAlignLeftPath), g.Group(children))
}

func TextAlignRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textAlignRightPath), g.Group(children))
}

func TextSize(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textSizePath), g.Group(children))
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ticketPath), g.Group(children))
}

func Tiled(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tiledPath), g.Group(children))
}

func Timeline(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(timelinePath), g.Group(children))
}

func Tired(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tiredPath), g.Group(children))
}

func Train(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trainPath), g.Group(children))
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trashPath), g.Group(children))
}

func Trophy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trophyPath), g.Group(children))
}

func Truck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(truckPath), g.Group(children))
}

func TumblerGlass(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tumblerGlassPath), g.Group(children))
}

func Twitter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(twitterPath), g.Group(children))
}

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(umbrellaPath), g.Group(children))
}

func Underline(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(underlinePath), g.Group(children))
}

func Union(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unionPath), g.Group(children))
}

func Unlock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unlockPath), g.Group(children))
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(uploadPath), g.Group(children))
}

func Usb(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(usbPath), g.Group(children))
}

func User(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userPath), g.Group(children))
}

func UserMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userMinusPath), g.Group(children))
}

func UserPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userPlusPath), g.Group(children))
}

func Users(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(usersPath), g.Group(children))
}

func Vector(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(vectorPath), g.Group(children))
}

func Video(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(videoPath), g.Group(children))
}

func Vr(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(vrPath), g.Group(children))
}

func Wallet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(walletPath), g.Group(children))
}

func Warning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(warningPath), g.Group(children))
}

func Watch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(watchPath), g.Group(children))
}

func WatchAlt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(watchAltPath), g.Group(children))
}

func WineGlass(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wineGlassPath), g.Group(children))
}

func Wordpress(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wordpressPath), g.Group(children))
}

func WordpressAlt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wordpressAltPath), g.Group(children))
}

func Wrench(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wrenchPath), g.Group(children))
}

func YakiDango(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(yakiDangoPath), g.Group(children))
}

func Youtube(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(youtubePath), g.Group(children))
}
