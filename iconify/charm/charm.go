package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const (
	IconifyVersion          = "0.12.1"
	anchorPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.75V14M3.25 7.75h-1.5c0 4 2.5 6.5 6.25 6.5s6.25-2.5 6.25-6.5h-1.5"/><circle cx="8" cy="3.5" r="1.75"/></g>`
	appsPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm0-8h4.5v4.5h-4.5z"/>`
	appsMinusPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm5.05-6h-5"/>`
	appsPlusPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h4.5v4.5h-4.5zm0 8h4.5v4.5h-4.5zm8 0h4.5v4.5h-4.5zm5.05-6h-5m2.5-2.5v5"/>`
	archivePath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v3.5H1.75zm5 6.5h2.5m-6.5-2.5v7.5h10.5v-7.5"/>`
	arrowDownPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.25 8.75l4.5 4.5l4.5-4.5m-4.5-6v10.5"/>`
	arrowDownLeftPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 11.75h-6.5v-6.5m7.5-1l-7.5 7.5"/>`
	arrowDownRightPath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 11.75h6.5v-6.5m-7.5-1l7.5 7.5"/>`
	arrowLeftPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.25 3.75l-4.5 4.5l4.5 4.5m6-4.5H2.75"/>`
	arrowRightPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.75 3.25l4.5 4.5l-4.5 4.5m-6-4.5h10.5"/>`
	arrowUpPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.75 7.25l4.5-4.5l4.5 4.5m-4.5 6V2.75"/>`
	arrowUpLeftPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 4.25h-6.5v6.5m7.5 1l-7.5-7.5"/>`
	arrowUpRightPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 4.25h6.5v6.5m-7.5 1l7.5-7.5"/>`
	atSignPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 8c0 3.25 4 3.25 4 0A6.25 6.25 0 1 0 8 14.25c2.25 0 3.25-1 3.25-1"/><circle cx="8" cy="8" r="2.25"/></g>`
	atomPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><ellipse cx="11.3" rx="8.28" ry="3.17" transform="rotate(45)"/><ellipse cy="11.3" rx="8.28" ry="3.17" transform="rotate(315)"/><path d="M8 8v0"/></g>`
	bellPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75c-2.468 0-4.25 1.5-4.25 3.5v3l-2 3.5h12.5l-2-3.5v-3c0-2-1.166-3.5-4.25-3.5zm-2.25 10.5c0 3 4.5 3 4.5 0"/>`
	bellSlashPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 12.25c0 3 4.5 3 4.5 0m2-4v-3c0-2-1.166-3.5-4.25-3.5m-3.75 2c-.53.585-.5.674-.5 1.5v3l-2 3.5h8.5"/><path d="M5.75 12.25c0 3 4.5 3 4.5 0m-7.5-10.5l10.5 12.5"/></g>`
	binPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 4.25v-2.5h4.5v2.5m-6.5 1v9h8.5v-9m-9.5-.5h10.5"/>`
	binaryPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 1.75h3v4.5h-3zm6.5 4.5h3m-3-4.5h1.5v4m-1.5 4h3v4.5h-3zm-6.5 4.5h3m-3-4.5h1.5v4"/>`
	blockPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="m4.25 11.75l8-8"/></g>`
	bluetoothPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.25L12.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25"/>`
	bluetoothConnectedPath  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.25L12.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25M1.75 8h1.5m9.5 0h1.5"/>`
	bluetoothSearchingPath  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.25L10.25 5l-4.5-3.25v12.5l4.5-3.25l-8.5-6.25m11.5 1.5s1 .5 1 1.75s-1 1.75-1 1.75m-2-1.75v0"/>`
	bluetoothSlashPath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 6.25L12.25 5l-4.5-3.25v2.5m4.5 6.75l-4.5 3.25v-6l-4 3m-2-8l12.5 9"/>`
	bookPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.75 11.75v2m1.5.5h-9c-.75 0-1.5-.5-1.5-1.5s.75-1.5 1.5-1.5h9v-9.5h-9c-.75 0-1.5.75-1.5 1.5v9.5"/>`
	bookOpenPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 3.75c-1.75-1-2.25-1-6.25-1v9.5c4 0 4.5 0 6.25 1c1.75-1 3.25-1 6.25-1v-9.5c-4 0-4.5 0-6.25 1zm0 .5v8.5"/>`
	bookmarkPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 1.75h8.5v12.5L8 9.75l-4.25 4.5z"/>`
	briefcasePath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75h12.5v9.5H1.75z"/><path d="M1.75 6.25s-.5 3.5 3 3.5h6.5c3.5 0 3-3.5 3-3.5m-8.5-2v-2.5h4.5v2.5"/></g>`
	bugPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="10" r="4.25"/><path d="M14.25 10.25h-1.5m-1 2.5l1.5 1.5m0-8.5l-1.5 1.5m-10 3h1.5m1 2.5l-1.5 1.5m0-8.5l1.5 1.5m1.5-1.5s-.75-3 2.25-3s2.25 3 2.25 3"/></g>`
	calendarPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v10.5H1.75zm9.5-2v1.5m-6.5-1.5v1.5m-2.5 4h11.5"/>`
	cameraPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75v8.5h12.5v-8.5h-3l-1.5-2h-3.5l-1.5 2z"/><circle cx="8" cy="8.5" r="2.25"/></g>`
	cameraVideoPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75h7.5v7.5h-7.5zm8 2.5l4.5-2.5v7.5l-4.5-2.5"/>`
	cameraVideoSlashPath    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 10.75l3 1.5v-7.5l-5 2.5v-2.5h-2.5m1.5 7.5h-6.5v-7.5h1.5m-1.5-2.5l8.5 12"/>`
	candyPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="3.25"/><path d="M7.25 11.25c0 1-.5 2.5-1.5 3c-.75 0-1.5-1-2-2c-1-.5-2-1.5-2-2c.5-1 2-1.5 3-1.5m4-4c0-1 .5-2.5 1.5-3c.75 0 1.5 1 2 2c1 .5 2 1.5 2 2c-.5 1-2 1.5-3 1.5"/></g>`
	cardsPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75H10v11.5H1.75zm8.25 1l4.25 2l-4.25 7.5"/>`
	castPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.25v-2.5h12.5v10.5h-4.5m-8-5a5 5 0 0 1 5 5m-5-2.5a2.5 2.5 0 0 1 2.5 2.5m-2.5 0v0"/>`
	certificatePath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.25 1.75h-8.5v11.5h2.5m3.5-3.5l-.5 4.5l2.25-1l2.25 1l-.5-4.5"/><circle cx="10.5" cy="7.5" r="2.75"/></g>`
	chartBarPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75v12.5h12.5m-9-3v-2.5m4 2.5v-5.5m4 5.5v-8.5"/>`
	chartLinePath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 11.25l2.5-4.5l2.5 2.5l3.5-6m-11.5-1.5v12.5h12.5"/>`
	chevronDownPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.75L8 10.25l4.25-4.5"/>`
	chevronLeftPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 3.75L5.75 8l4.5 4.25"/>`
	chevronRightPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 12.25L10.25 8l-4.5-4.25"/>`
	chevronUpPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 10.25L8 5.75l-4.25 4.5"/>`
	chevronsDownPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 3.75L8 8.25l4.25-4.5m-8.5 5L8 13.25l4.25-4.5"/>`
	chevronsLeftPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 3.75L7.75 8l4.5 4.25m-5-8.5L2.75 8l4.5 4.25"/>`
	chevronsRightPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 12.25L8.25 8l-4.5-4.25m5 8.5L13.25 8l-4.5-4.25"/>`
	chevronsUpPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12.25L8 7.75l-4.25 4.5m8.5-5L8 2.75l-4.25 4.5"/>`
	chevronsUpDownPath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 10.75L8 14.25l-3.25-3.5m6.5-5.5L8 1.75l-3.25 3.5"/>`
	chipPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75h10.5v10.5H2.75z"/><path d="M6.25 6.25h3.5v3.5h-3.5zm-4 4h-1m1-4.5h-1m13.5 4.5h-1m1-4.5h-1m-3.5 8v1m-4.5-1v1m4.5-13.5v1m-4.5-1v1"/></g>`
	circlePath              = `<circle cx="8" cy="8" r="6.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>`
	circleCrossPath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.25 5.75l-4.5 4.5m0-4.5l4.5 4.5"/><circle cx="8" cy="8" r="6.25"/></g>`
	circleMinusPath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M4.75 8h6.5"/></g>`
	circleTickPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8.75c-.5 2.5-2.385 4.854-5.03 5.38A6.25 6.25 0 0 1 3.373 3.798C5.187 1.8 8.25 1.25 10.75 2.25"/><path d="m5.75 7.75l2.5 2.5l6-6.5"/></g>`
	circleWarningPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8 10.75v0m0-6v3.5"/></g>`
	clipboardPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 1.75h4.5v3.5h-4.5z"/><path d="M5.25 2.75h-2.5v11.5h10.5V2.75h-2.5"/></g>`
	clipboardTickPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 1.75h4.5v3.5h-4.5zm4 11.05l1.5 1.5l3-2.5m-9-9h-2.5v11.5h4.5m6-5V2.8h-2.5"/>`
	clockPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8.25 4.75v3.5l-2.5 2"/></g>`
	clockAlarmPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.75 1.75l2.5 2m-10-2l-2.5 2m10.5 9.5l1 1m-9.5-1l-1 1m5.5-7.5v2.5l-1.5 1"/><circle cx="8" cy="9" r="5.25"/></g>`
	cloudPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 3.75A3.25 3.25 0 0 0 3.75 7c.002.255.033.508.094.756h-.002A2.25 2.25 0 0 0 4 12.25h7.5a2.75 2.75 0 1 0-1.252-5.2L10.25 7A3.25 3.25 0 0 0 7 3.75z"/>`
	cloverPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 2.75C4.25 4.25 6 6 8 8c2-2 3.75-3.75 3.25-5.25s-2.5-1-3.25.5c-.75-1.5-2.75-2-3.25-.5zM8 8c2 2 3.75 3.75 5.25 3.25s1-2.5-.5-3.25c1.5-.75 2-2.75.5-3.25S10 6 8 8zm0 0c-2 2-3.75 3.75-3.25 5.25s2.5 1 3.25-.5c.75 1.5 2.75 2 3.25.5S10 10 8 8zm0 0C6 6 4.25 4.25 2.75 4.75s-1 2.5.5 3.25c-1.5.75-2 2.75-.5 3.25S6 10 8 8z"/>`
	codePath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 11.25L1.75 8l3.5-3.25m5.5 6.5L14.25 8l-3.5-3.25"/>`
	coffeePath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 11.25c4.5 0 4.5-5.5 0-5.5h-9v5c0 5 8.5 5 8.5 0v-5m-1.5-4v1.5m-3-1.5v1.5m-3-1.5v1.5"/>`
	cogPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="1.75"/><path d="m6.75 1.75l-.5 1.5l-1.5 1l-2-.5l-1 2l1.5 1.5v1.5l-1.5 1.5l1 2l2-.5l1.5 1l.5 1.5h2.5l.5-1.5l1.5-1l2 .5l1-2l-1.5-1.5v-1.5l1.5-1.5l-1-2l-2 .5l-1.5-1l-.5-1.5z"/></g>`
	compassPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="m6.75 6.75l-1 4l3.5-1.5l1-4z"/></g>`
	conicalFlaskPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 1.75h6.5m-6.5 8h6.5m-5.5-7.5v4.5l-4 7.5h12.5l-4-7.5v-4.5"/>`
	containerPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 12.2l5.5 2l7-4.5v-6l-5.5-2l-7 4.5z"/><path d="M10.8 6.25v5.5m-3.5-3.5v6m-5.5-8l5.5 2l7-4.5"/></g>`
	copyPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 4.25v-2.5h-9.5v9.5h2.5m.5-6.5v9.5h9.5v-9.5z"/>`
	copyleftPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M6 6.75s.75-1 2-1s2.25 1 2.25 2.25s-1 2.25-2.25 2.25s-2-1-2-1"/></g>`
	copyrightPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M10 6.75s-.75-1-2-1s-2.25 1-2.25 2.25s1 2.25 2.25 2.25s2-1 2-1"/></g>`
	creditCardPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v9.5H1.75zm8 6.5h1.5m-9-3h11.5"/>`
	cropPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 1.75v10h10M11.8 14.2v-2.5m0-2.5v-5h-5m-2.5 0H1.8"/>`
	crossPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.25 4.75l-6.5 6.5m0-6.5l6.5 6.5"/>`
	crosshairPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.25v-3m0 11.5v-3M10.75 8h3M2.25 8h3"/><circle cx="8" cy="8" r="6.25"/></g>`
	cubePath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5zM8 14V8m5.75-3L8 8M2 5l6 3"/>`
	cursorPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 1.75l4.5 12.5l2.5-5.5l5.5-2.5zm7.5 7.5l4 4"/>`
	databasePath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-3.75 0-5.25 2-5.25 2v8.5s1.5 2 5.25 2s5.25-2 5.25-2v-8.5s-1.5-2-5.25-2z"/><path d="M2.75 8.25s1.5 2 5.25 2s5.25-2 5.25-2m-10.5-4s1.5 2 5.25 2s5.25-2 5.25-2"/></g>`
	diamondPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.25 8L8 14.75L14.75 8L8 1.25z"/>`
	diffPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 13.75h8m0-7.5h-8m4-4v8"/>`
	discPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><circle cx="8" cy="8" r="1.75"/></g>`
	downloadPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 13.25h9m-8.5-6.5l4 3.5l4-3.5m-4-5v8.5"/>`
	dropletPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 9a5.25 5.25 0 1 0 10.5 0C13.25 5.75 8 1.75 8 1.75S2.75 5.75 2.75 9z"/>`
	eraserPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.25h-9.5l-3-3l7.5-7.5l5 5l-5.5 5.5m-3.5-6.5l5 5"/>`
	extensionsPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8.75h5.5v5.5m5-12.5v4m-2-2h4m-12.5-1v11.5h11.5v-5.5h-6v-6z"/></g>`
	eyePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8s2-4.25 6.25-4.25S14.25 8 14.25 8s-2 4.25-6.25 4.25S1.75 8 1.75 8z"/><circle cx="8" cy="8" r="1.25" fill="currentColor"/></g>`
	eyeSlashPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.75 3.75c3.5.5 5.5 4.25 5.5 4.25s-.5 1.25-1.5 2.25m-2.5 1.5c-6 2-8.5-3.75-8.5-3.75s.5-1.75 3-3.25"/><path fill="currentColor" d="M8.625 9.083a1.25 1.25 0 0 1-1.649-.366a1.25 1.25 0 0 1 .22-1.675L8 8z"/><path d="m3.75 1.75l8.5 12.5"/></g>`
	faceFrownPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 5s.5-1 2.25-1s2.25 1 2.25 1"/></g>`
	faceNeutralPath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 4.5h4.5"/></g>`
	faceSmilePath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M9.75 6.25v-.5m-3.5.5v-.5m-.5 4s.5 1.5 2.25 1.5s2.25-1.5 2.25-1.5"/></g>`
	filePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 1.75h5.5l5 5v7.5H2.75z"/><path d="M7.75 2.25v5h5"/></g>`
	fileBinaryPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5M1.75 10.8h3v3.5h-3z"/><path d="M7.25 14.2h3m-3-3.5h1.5v3m-1-11.45v5h5.05"/></g>`
	fileCodePath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-2"/><path d="M7.75 2.25v5h5.05M6.75 10.8l2 1.75l-2 1.75m-3-3.5l-2 1.75l2 1.75"/></g>`
	fileSymlinkPath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-4"/><path d="M7.75 2.25v5h5m-10 7l3.5-3.5m0 3v-3h-3"/></g>`
	filesPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.25 1.75l4 4v5.5h-7.5v-9.5z"/><path d="M9.25 2.25v3.5h3.5m-2.5 6v2.5h-7.5v-9.5h2.5"/></g>`
	filterPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h12.5v1.5l-5 5.5v4l-2.5 1.5v-5.5l-5-5.5z"/>`
	flagPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 14.25v-11s2-1.5 4-1.5s2.5 1.5 4.5 1.5s4-1.5 4-1.5v7s-2 1.5-4 1.5s-2.5-1.5-4.5-1.5s-4 1.5-4 1.5"/>`
	flamePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 7.75c2 2 2.5-2.5 3.5-2s1.5 2 1.5 3.25c0 3.25-2.35 5.25-5.25 5.25s-5.25-2.5-5.25-6s3.5-7 5.5-7c0 0-2 4.5 0 6.5z"/>`
	floppyDiskPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75v10.5h10.5v-7.5l-3-3z"/><path d="M5.75 13.25v-3.5h4.5v3.5"/></g>`
	folderPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75v10.5h12.5v-8.5h-6l-1.5-2z"/>`
	folderSymlinkPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1.75 13.25l3.5-3.5m0 3v-3h-3"/><path d="M8.25 13.25h6v-8.5h-6l-1.5-2h-5v4"/></g>`
	foldersPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.75 2.25v8h9.5v-6.5h-5l-1.5-1.5z"/><path d="M4.75 5.25h-3v8h9.5v-3"/></g>`
	forwardPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 13.25c.5-6 5.5-7.5 8-7v-3.5L14.25 8l-4.5 5.25v-3.5c-2.5-.5-6.5.5-8 3.5z"/>`
	gamepadPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 3.75c-2 5-2 9 0 9.5s2.5-2 2.5-2h4.5s.5 2.5 2.5 2s2-4.5 0-9.5h-2l-1 1h-3.5l-1-1z"/>`
	gemPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 2.75h6.5l3 3.5l-6.25 7l-6.25-7z"/>`
	giftPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75h12.5v3.5H1.75z"/><path d="M10.25 4.75H8c0-2 .5-3 2.25-3c2 0 2 3 0 3zm-4.5 0H8c0-2-.5-3-2.25-3c-2 0-2 3 0 3zm2.25 9V5M2.75 8.75v5.5h10.5v-5.5"/></g>`
	gitBranchPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="11.5" cy="3.5" r="1.75"/><circle cx="4.5" cy="12.5" r="1.75"/><path d="M5.25 8.25c3 0 6 .5 6-2.5m-6.5 4.5v-4.5"/></g>`
	gitCherryPickPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="8" r="2.25"/><path d="M5 10.75v3.5m0-12.5v3.5M11.75 8h1.5m-4.5-3.25h1.5l1 3.25l-1 3.25h-1.5"/></g>`
	gitCommitPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="2.25"/><path d="M8 10.75v3.5m0-12.5v3.5"/></g>`
	gitComparePath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M3.75 5.75v5c0 1 .5 1.5 1.5 1.5h2m-.5 2l1.5-2l-1.5-2m5.5 0v-5c0-1-.5-1.5-1.5-1.5h-2m.5-2l-1.5 2l1.5 2"/></g>`
	gitForkPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="12.5" r="1.75"/><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="11.5" cy="3.5" r="1.75"/><path d="M8 8.75v1.5m-3.25-4.5c0 3.5 6.5 3.5 6.5 0"/></g>`
	gitMergePath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4.5" cy="3.5" r="1.75"/><circle cx="4.5" cy="12.5" r="1.75"/><circle cx="12.5" cy="8.5" r="1.75"/><path d="M4.75 10.25v-4.5c1 2 2 3 5.5 3"/></g>`
	gitRequestPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="m9.25 1.75l-1.5 2l1.5 2m3 4.5v-5c0-1-.5-1.5-1.5-1.5h-2m-5 2v4.5"/></g>`
	gitRequestCrossPath     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M12.25 7.25v3m-8.5-4.5v4.5m10.5-8.5l-3.5 3.5m0-3.5l3.5 3.5"/></g>`
	gitRequestDraftPath     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M7.75 2.75h.5m2.5 0h.5m1.5 2.5v-.5m0 3v.5m-9-2.5v4.5"/></g>`
	githubPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.75 14.25s-.5-2 .5-3c0 0-2 0-3.5-1.5s-1-4.5 0-5.5c-.5-1.5.5-2.5.5-2.5s1.5 0 2.5 1c1-.5 3.5-.5 4.5 0c1-1 2.5-1 2.5-1s1 1 .5 2.5c1 1 1.5 4 0 5.5s-3.5 1.5-3.5 1.5c1 1 .5 3 .5 3m-5-.5c-1.5.5-3-.5-3.5-1"/>`
	gitlabPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 14.25l-6.25-4.5l2-8l2 5.5h4.5l2-5.5l2 8z"/>`
	glassesPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="11" r="2.25"/><circle cx="12" cy="11" r="2.25"/><path d="M14.25 10.75c-1.5-6-2-6.5-3.5-7m-9 7c1.5-6 2-6.5 3.5-7m1 7c1-1 2.5-1 3.5 0"/></g>`
	globePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M2 8.25h12M8.25 14.2C11 11 11 5 8.25 1.8m-.5 12.4C5 11 5 5 7.75 1.8"/></g>`
	grabHorizontalPath      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="5.5" r=".75"/><circle cx="8" cy="5.5" r=".75"/><circle cx="13.5" cy="5.5" r=".75"/><circle cx="2.5" cy="10.5" r=".75"/><circle cx="8" cy="10.5" r=".75"/><circle cx="13.5" cy="10.5" r=".75"/></g>`
	grabVerticalPath        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5.5" cy="2.5" r=".75"/><circle cx="5.5" cy="8" r=".75"/><circle cx="5.5" cy="13.5" r=".75"/><circle cx="10.496" cy="2.5" r=".75"/><circle cx="10.496" cy="8" r=".75"/><circle cx="10.496" cy="13.5" r=".75"/></g>`
	graduateCapPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 9.25V6L8 2.75L1.75 6L8 9.25l3.25-1.5v3.5c0 1-1.5 2-3.25 2s-3.25-1-3.25-2v-3.5"/>`
	hashPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 10.25h9.5m-8.5-4.5h9.5m-2.5-4l-1.5 12.5m-2.5-12.5l-1.5 12.5"/>`
	headphonesPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.75c0-2.5 3.5-2 3.5-2v4.5s-3.5.5-3.5-2.5v-3.5c0-3 .5-6.5 6.25-6.5s6.25 3.5 6.25 6.5v3.5c0 3-3.5 2.5-3.5 2.5v-4.5s3.5-.5 3.5 2"/>`
	heartPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 9.75c3 3.5 4.75 4.5 4.75 4.5s1.75-1 4.75-4.5s1-7-1.5-7s-3.25 3-3.25 3s-.75-3-3.25-3s-4.5 3.5-1.5 7z"/>`
	helpPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M5.75 6.75c0-1 1-2 2.25-2s2.25 1.034 2.25 2c0 1.5-1.5 1.5-2.25 2m0 2.5v0"/></g>`
	hexagonPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5z"/>`
	homePath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 5.75v7.5h8.5v-7.5m-10.5 1.5L8 1.75l6.25 5.5"/>`
	hourglassPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.75 13.75c0-5-2-4-2-5.75s2-.75 2-5.75m-7.5 11.5c0-5 2-4 2-5.75s-2-.75-2-5.75m-1.5-.5h10.5m-10.5 12.5h10.5"/>`
	idPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><circle cx="8" cy="7.5" r="2.25"/><path d="M4.75 12.75c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`
	imagePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="m3.75 13.2l6.5-5.5l4 3"/><circle cx="5.25" cy="6.25" r=".5" fill="currentColor"/></g>`
	inboxPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 13.25h12.5v-5l-2.5-5.5h-7.5l-2.5 5.5z"/><path d="M2.25 8.75h3l1 1.5h3.5l1-1.5h3"/></g>`
	infinityPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 5c2.5 1 3.5 5 6 6s3.25-1.25 3.25-3S13.5 4 11 5s-3.5 5-6 6s-3.25-1.25-3.25-3S2.5 4 5 5z"/>`
	infoPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="6.25"/><path d="M8 5.25v0m0 6v-3.5"/></g>`
	keyPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 1.75a4.25 4.25 0 0 0-4.104 5.354L1.75 11.25v3h3v-1.5h1.5v-1.5h1.5L8.9 10.1c.359.098.728.148 1.1.15a4.25 4.25 0 0 0 0-8.5z"/><circle cx="10.75" cy="5.25" r=".5" fill="currentColor"/></g>`
	laptopPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h10.5v7.5H2.75zm0 7.5l-1 3h12.5l-1-3"/>`
	layoutColumnsPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.25.5v9.5"/>`
	layoutDashboardPath     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.5 4h5.5m-11.5 2.5h5.5m.25-6v9.5"/>`
	layoutGridPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12m-3.75-4.75v9.5m-4.5-9.5v9.5"/>`
	layoutListPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm3.5.5v9.5m-3-6.5h11.5m-11.5 3.5h11.5"/>`
	layoutRowsPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12"/>`
	layoutSidebarPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm4.5.25v9.5"/>`
	layoutStackHPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zM2 8h12M8 8v4.75"/>`
	layoutStackVPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 2.75h12.5v10.5H1.75zm6.25.5v9.5M8 8h6"/>`
	lightbulbPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 14.25h2.5M8 1.75c-2.75 0-4.25 2-4.25 4s2 2.5 2 4.5v1h4.5v-1c0-2 2-2.5 2-4.5s-1.5-4-4.25-4z"/>`
	lightningBoltPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.25 1.75l-6.5 7.5l4.5.5l-.5 4.5l6.5-7.5l-4.5-.5z"/>`
	linkPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 4.75c3 0 4.5 1.5 4.5 3.25s-1.5 3.25-4.5 3.25M5.75 8h4.5m-4-3.25c-3 0-4.5 1.5-4.5 3.25s1.5 3.25 4.5 3.25"/>`
	linkExternalPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 2.75h-5.5v10.5h10.5v-5.5m0-5l-5.5 5.5m3-6.5h3.5v3.5"/>`
	linkSlashPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.75 1.75l-5.5 12.5m4.5-9.5c3 0 4.5 1.5 4.5 3.25s-1.5 3.25-4.5 3.25m-3.5-6.5c-3 0-4.5 1.5-4.5 3.25s1.5 3.25 4.5 3.25"/>`
	mailPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 3.75h12.5v9.5H1.75z"/><path d="m2.25 4.25l5.75 5l5.75-5"/></g>`
	mapPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 5.25v8.5m-4.5-10.5v8.5m-4 2.5v-9.5l4-2l4.5 2l4-2v9.5l-4 2l-4.5-2z"/>`
	mapPinPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.25 7c0 3.75-5.25 7.25-5.25 7.25S2.75 10.75 2.75 7a5.25 5.25 0 0 1 10.5 0z"/><circle cx="8" cy="7" r="1.25" fill="currentColor"/></g>`
	mediaBackPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.25 13.25L4.75 8l8.5-5.25zm-11.5-9.5v8.5"/>`
	mediaEjectPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 11.25h10.5L8 2.75zm10.5 3H2.75"/>`
	mediaFastForwardPath    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 3.75v8.5l6-4.25zm-6.5 0v8.5l6-4.25z"/>`
	mediaPausePath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h3.5v10.5h-3.5zm7 0h3.5v10.5h-3.5z"/>`
	mediaPlayPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75v10.5L12.25 8z"/>`
	mediaRewindPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.75v8.5L1.75 8zm6.5 0v8.5L8.25 8z"/>`
	mediaSkipPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 13.25L11.25 8l-8.5-5.25zm11.5-9.5v8.5"/>`
	menuHamburgerPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 12.25h10.5m-10.5-4h10.5m-10.5-4h10.5"/>`
	menuKebabPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="2.5" r=".75"/><circle cx="8" cy="8" r=".75"/><circle cx="8" cy="13.5" r=".75"/></g>`
	menuMeatballPath        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="8" r=".75"/><circle cx="8" cy="8" r=".75"/><circle cx="13.5" cy="8" r=".75"/></g>`
	messagePath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 14.25V2.75h12.5v8.5h-8.5z"/>`
	messagesPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 14.25v-9h-9.5v6h6z"/><path d="m4.75 7.25l-3 3v-8.5h10v3"/></g>`
	microphonePath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-2.25 0-2.25 2-2.25 3v1.5c0 1 0 3 2.25 3s2.25-2 2.25-3v-1.5c0-1 0-3-2.25-3z"/><path d="M8 13v1.25m-5.25-6.5s0 4.5 5.25 4.508c5.25.008 5.25-4.508 5.25-4.508"/></g>`
	minusPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.25 7.75H2.75"/>`
	mobilePath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 1.75h8.5v12.5h-8.5zm4.5 10h-.5"/>`
	monitorPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75h12.5v9.5H1.75zm3 12.5h6.5M8 11.75v2.5"/>`
	monitorArrowPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.2 7.75v3.5H1.7v-9.5h6.5M4.75 14.2h6.5M8 11.7v2.5m1.75-7.95l4.5-4.5m-3.5-.5h4v4"/>`
	monitorCrossPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.2 7.75v3.5H1.7v-9.5h6.5M4.75 14.2h6.5M8 11.7v2.5m6.2-12.45l-3.5 3.5m0-3.5l3.5 3.5"/>`
	moonPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8c0 3.45 2.8 6.25 6.25 6.25c3.41-.003 6.25-3 6.25-6c-1 .5-4 1.5-6-.5s-1-5-.5-6c-3 0-6 2.84-6 6.25z"/>`
	movePath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.25 10.25l2-2.25l-2-2.25m-2-2L8 1.75l-2.25 2m-2 2L1.75 8l2 2.25m2 2l2.25 2l2.25-2M8 1.75v12M13.75 8h-12"/>`
	musicPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="12" r="2.25"/><circle cx="12" cy="11" r="2.25"/><path d="M6.25 12V2.75l8-1V11"/></g>`
	newspaperPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.2 14.2h.5c1.5 0 2.5-1 2.5-2.5v-6h-3m-9.5-4h9.5v12.5h-7c-1.5 0-2.5-1-2.5-2.5V2.26zm3.05 9.5h3.5"/><path d="M4.75 4.75h3.5v3.5h-3.5z"/></g>`
	northStarPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7.75h-12m6-6v12m-3.5-2.5l7-7m0 7l-7-7"/>`
	notesPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 1.75h10.5v12.5H2.75zm3 6h4.5m-4.5 3h2.5m-2.5-6h4.5"/>`
	notesCrossPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 14.25h-5.5V1.75h10.5v6.5m1 2.5l-3.5 3.5m-5-6.5h4.5m-4.5 3h1.5m-1.5-6h4.5m.5 6l3.5 3.5"/>`
	notesTickPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.25 14.25h-4.5V1.75h10.5v7.5m-3.5 3.5l1.5 1.5l3-2.5m-8.5-4h4.5m-4.5 3h1.5m-1.5-6h4.5"/>`
	nutPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.25l6.25 3.5v6.5L8 14.75l-6.25-3.5v-6.5z"/><circle cx="8" cy="8" r="2.25"/></g>`
	octagonPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 1.75h5.5l3.5 3.5v5.5l-3.5 3.5h-5.5l-3.5-3.5v-5.5z"/>`
	octagonWarningPath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 1.75h5.5l3.5 3.5v5.5l-3.5 3.5h-5.5l-3.5-3.5v-5.5zM8 11.25v0m0-6.5v3.5"/>`
	organisationPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 1.75h3.5v3.5h-3.5zm4 9h3.5v3.5h-3.5zm-8 0h3.5v3.5h-3.5zm5.75-5v2m-3.75 2.5v-2h7.5v2"/>`
	packagePath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v8.5h12.5v-8.5l-3.5-4h-5.5zm6.25-4v3.5m-5.75.5h11.5"/>`
	padlockPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.75h10.5v7.5H2.75zm2-.5s-1-4.5 3.25-4.5s3.25 4.5 3.25 4.5"/>`
	paperPlanePath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 1.75l12.5 6l-12.5 6.5l1.5-6.5zm2 6h3.5"/>`
	paperclipPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.25 10.25v-7s0-1.5-1.75-1.5s-1.75 1.5-1.75 1.5v8s0 3 3.25 3s3.25-3 3.25-3v-4.5"/>`
	pencilPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11.25v3h3l9.5-9.5l-3-3zm7-6.5l2.5 2.5"/>`
	peoplePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="5" cy="9" r="2.25"/><circle cx="11" cy="4" r="2.25"/><path d="M7.75 9.25c0-1 .75-3 3.25-3s3.25 2 3.25 3m-12.5 5c0-1 .75-3 3.25-3s3.25 2 3.25 3"/></g>`
	personPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="6" r="3.25"/><path d="M2.75 14.25c0-2.5 2-5 5.25-5s5.25 2.5 5.25 5"/></g>`
	phonePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5z"/>`
	phoneCallPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 0c2.5 0 4.5 2 4.5 4.5m-4.5-2c1 0 2 1 2 2"/>`
	phoneCrossPath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm11.5 1l-3.5 3.5m0-3.5l3.5 3.5"/>`
	phoneForwardPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 3h4.5m-2 2l2-2l-2-2"/>`
	phoneIncomingPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5z"/><path d="m13.25 2.75l-3.5 3.5m0-3v3h3"/></g>`
	phoneOutgoingPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 1.75c0 8.5 4 12.5 12.5 12.5v-4l-3.5-1l-1 1.5c-2 0-4.5-2.5-4.5-4.5l1.5-1l-1-3.5zm8 4.5l3.5-3.5m0 3v-3h-3"/>`
	pinPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.25 10.25l4 4m-12.5-7.5l5-5s1 2 2 3s4.5 2 4.5 2l-6.5 6.5s-1-3.5-2-4.5s-3-2-3-2z"/>`
	plantPotPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 6.75c0 1.25-.75 3-.75 3m.25-2.5s.75-2-1-3.5s-4.5-1-4.5-1s0 2 1.5 3.5s4 1 4 1zm.5-1s-.75-2 1-3.5s4.5-1 4.5-1s0 2-1.5 3.5s-4 1-4 1zm-4 3.5h6.5s.5 4.5-3.25 4.5s-3.25-4.5-3.25-4.5z"/>`
	plusPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.75 7.75h-10m5-5v10"/>`
	powerPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75v6.5m4.25-5s2 1.298 2 4.75a6.25 6.25 0 1 1-12.5 0c0-3.452 2-4.75 2-4.75"/>`
	printerPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.75 9.75h6.5v4.5h-6.5z"/><path d="M4.75 4.25v-2.5h6.5v2.5m-7 8h-2.5v-7.5h12.5v7.5h-2.5"/></g>`
	pulsePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25h2.5l2-4.5l3.5 8.5l2-4h2.5"/>`
	quotePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.25 3.75h-4.5v5.5c0 3.5 2.5 4.5 4.5 4c-1.5-1.5-1.5-2.5-1.5-4h1.5zm7 0h-4.5v5.5c0 3.5 2.5 4.5 4.5 4c-1.5-1.5-1.5-2.5-1.5-4h1.5z"/>`
	refreshPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 10.75h-3m12.5-2c0 3-2.798 5.5-6.25 5.5c-3.75 0-6.25-3.5-6.25-3.5v3.5m9.5-9h3m-12.5 2c0-3 2.798-5.5 6.25-5.5c3.75 0 6.25 3.5 6.25 3.5v-3.5"/>`
	replyPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.25c-.5-6-5.5-7.5-8-7v-3.5L1.75 8l4.5 5.25v-3.5c2.5-.5 6.5.5 8 3.5z"/>`
	robotPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 5.75h12.5v7.5H1.75z"/><path d="M10.75 8.75v1.5m-5.5-1.5v1.5m-.5-7.5l3.25 3l3.25-3"/></g>`
	rocketPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.25 9.75l-2-.5s0-1.5.5-3s4-1.5 4-1.5m-.5 7l.5 2s1.5 0 3-.5s1.5-4 1.5-4m-7 .5l2 2s5-2 6.5-4.5s1.5-5.5 1.5-5.5s-3 0-5.5 1.5s-4.5 6.5-4.5 6.5z"/><path fill="currentColor" d="m1.75 14.25l2-1l-1-1z"/><circle cx="10.25" cy="5.75" r=".5" fill="currentColor"/></g>`
	rotateAntiClockwisePath = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 5.25h-3m0 3.5c0 2.5 2.798 5.5 6.25 5.5a6.25 6.25 0 1 0 0-12.5c-3.75 0-6.25 3.5-6.25 3.5v-3.5"/>`
	rotateClockwisePath     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.25 5.25h3m0 3.5c0 2.5-2.798 5.5-6.25 5.5a6.25 6.25 0 1 1 0-12.5c3.75 0 6.25 3.5 6.25 3.5v-3.5"/>`
	scalesPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 3.75c1 1 2.5 1.5 4 0h4.5c1.5 1.5 3 1 4 0M8 1.75v12m-3.25.5h6.5"/><path d="m12.75 4.75l-2 5c.5 1 3.5 1 4 0zm-9.5 0l-2 5c.5 1 3.5 1 4 0z"/></g>`
	screenMaximisePath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 14.25h-3.5v-3.5m12.5 0v3.5h-3.5m0-12.5h3.5v3.5m-12.5 0v-3.5h3.5"/>`
	screenMinimisePath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 10.75h3.5v3.5m5.5 0v-3.5h3.5m0-5.5h-3.5v-3.5m-5.5 0v3.5h-3.5"/>`
	searchPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.25 11.25l3 3"/><circle cx="7.5" cy="7.5" r="4.75"/></g>`
	serverPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.25h12.5v10H1.75zm.5 5h11.5m-9 2.5v0m0-5v0m6.5 0h-3m3 5h-3"/>`
	sharePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="4" cy="8" r="2.25"/><circle cx="12" cy="12" r="2.25"/><circle cx="12" cy="4" r="2.25"/><path d="m6 9l4 2M6 7l4-2"/></g>`
	shieldPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5z"/>`
	shieldCrossPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm1.75 4l-3.5 3.5m0-3.5l3.5 3.5"/>`
	shieldKeyholePath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm0 5.5v3"/><circle cx="8" cy="6.5" r=".75" fill="currentColor"/></g>`
	shieldTickPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5z"/><path d="m5.75 7.75l1.5 1.5l3-3.5"/></g>`
	shieldWarningPath       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l5.25 2v5c0 2.25-2 4.5-5.25 5.5c-3.25-1-5.25-3-5.25-5.5v-5zm0 9v0m0-5.5v3"/>`
	shoppingBagPath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 4.75h10.5v9.5H2.75z"/><path d="M5.75 7.75c0 1.5 1 2.5 2.25 2.5s2.25-1 2.25-2.5m-7.5-3l1.5-3h7.5l1.5 3"/></g>`
	signInPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 2.25h-3.5v12h3.5m4-9.5l-3.5 3.5l3.5 3.5m5-3.5h-8.5"/>`
	signOutPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 2.25h-3.5v12h3.5m5.5-9.5l3.5 3.5l-3.5 3.5m-5-3.5h8.5"/>`
	signpostPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 9.25h10.5l2-2.25l-2-2.25H1.75zm5.5.5v4.5m0-12.5v2.5"/>`
	skullPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 11.25h3v3h6.5v-3h3s1-9.5-6.25-9.5s-6.25 9.5-6.25 9.5z"/><circle cx="5.25" cy="7.75" r=".5" fill="currentColor"/><circle cx="10.75" cy="7.75" r=".5" fill="currentColor"/></g>`
	snowflakePath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7.75h-12m6-6v12m-2.5-1l2.5-2.5l2.5 2.5m-7.5-7.5l2.5 2.5l-2.5 2.5m7.5-7.5l-2.5 2.5l-2.5-2.5m7.5 7.5l-2.5-2.5l2.5-2.5"/>`
	soundDownPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm9 .5s1 .5 1 1.75s-1 1.75-1 1.75"/>`
	soundMutePath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm12.5 0l-3.5 4.5m0-4.5l3.5 4.5"/>`
	soundUpPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.75v4.5h2.5l4 3V2.75l-4 3zm9 .5s1 .5 1 1.75s-1 1.75-1 1.75m1-6.5c2 1 3 2.5 3 4.75s-1 3.75-3 4.75"/>`
	speakerPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 1.75h9.5v12.5h-9.5zm5 2.5h-.5"/><circle cx="8" cy="9.5" r="2.25"/></g>`
	squarePath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 2.75h10.5v10.5H2.75z"/>`
	squareCrossPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.25 5.75l-4.5 4.5m0-4.5l4.5 4.5m-7.5-7.5h10.5v10.5H2.75z"/>`
	squareTickPath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 2.75h-7.5v10.5h10.5v-3.5"/><path d="m5.75 7.75l2.5 2.5l6-6.5"/></g>`
	stackPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 11L8 14.25L14.25 11M1.75 8L8 11.25L14.25 8M8 1.75L1.75 5L8 8.25L14.25 5z"/>`
	stackPopPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 6.75L1.75 8L8 11.25L14.25 8l-2.5-1.25M1.75 11L8 14.25L14.25 11M8 8.25v-6.5m-2.25 2l2.25-2l2.25 2"/>`
	stackPushPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 7.25L1.75 8L8 11.25L14.25 8l-1.5-.75M1.75 11L8 14.25L14.25 11"/><path d="M8 8.25v-6.5m-2.25 4.5l2.25 2l2.25-2"/></g>`
	starPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-2.25 4l-4 .5l3 3.5l-1 4.5l4.25-2l4.25 2l-1-4.5l3-3.5l-4-.5z"/>`
	stickyNotePath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.25 13.25h-6.5V2.75h10.5v6.5z"/><path d="M8.75 13.25v-4.5h4.5"/></g>`
	sunPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="8" cy="8" r="3.25"/><path d="m2.75 13.25l.5-.5m9.5 0l.5.5m-.5-10l.5-.5m-10 .5l-.5-.5M2.25 8h-1m13.5 0h-1M8 13.75v1m0-13.5v1"/></g>`
	swapHorizontalPath      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.75 8.25l-3 3l3 3m7.5-3H2.75m7.5-9.5l3 3l-3 3m-7.5-3h10.5"/>`
	swapVerticalPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.75 5.75l-3-3l-3 3m3 7.5V2.75m9.5 7.5l-3 3l-3-3m3-7.5v10.5"/>`
	swordPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 9.25l1.5 2.5l2 1.5m-4.5 0l1 1m1.5-2.5l-1.5 1.5m3-1l8.5-8.5v-2h-2l-8.5 8.5"/>`
	swordsPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.75 9.25l1.5 2.5l2 1.5m-4.5 0l1 1m1.5-2.5l-1.5 1.5m3-1l8.5-8.5v-2h-2l-8.5 8.5"/><path d="M10.25 12.25L8 10m2-2l2.25 2.25m1-1l-1.5 2.5l-2 1.5m4.5 0l-1 1m-1.5-2.5l1.5 1.5M6 8L1.75 3.75v-2h2L8 6"/></g>`
	tabletPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 1.75h10.5v12.5H2.75zm5.5 10h-.5"/>`
	tagPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.25 14.25l-5.5-5.5l7-7h5.5v5.5z"/><circle cx="11" cy="5" r=".5" fill="currentColor"/></g>`
	telescopePath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 5.75l1 2.5m3.5-4.5l1.5 3.5m-9 0l1 2.5l11.5-3.5l-2-4.5zm6 3.95v3m-3-.5L7 11.2l1.75-.5l2.5 3"/>`
	tentPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.25 14.25L8 10l2.75 4.25"/><path d="m9.75 1.75l-8 12.5h12.5l-8-12.5"/></g>`
	terminalPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="M8.75 10.25h2.5m-6.5-4.5L7.25 8l-2.5 2.25"/></g>`
	thumbDownPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 10.25c1.5 0 3 4 4.5 4v-4h4.5s-.5-7.5-3.5-7.5h-5.5zm0 0h-3.5v-7.5h3.5"/>`
	thumbUpPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 5.75c1.5 0 3-4 4.5-4v4h4.5s-.5 7.5-3.5 7.5h-5.5zm0 0h-3.5v7.5h3.5"/>`
	tickPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 8.75l3.5 3.5l7-7.5"/>`
	tickDoublePath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.75 9.75l2.5 2.5m3.5-4l2.5-2.5m-4.5 4l2.5 2.5l6-6.5"/>`
	ticketPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.75h12.5v3s-2 0-2 1.75s2 1.75 2 1.75v3H1.75v-3s2 0 2-1.75s-2-1.75-2-1.75zm7 8v1.5m0-9.5v1.5m0 2.5v1.5"/>`
	treeFirPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-4.25 5.5h2.5l-3.5 4h4v3h2.5v-3h4l-3.5-4h2.5z"/>`
	trianglePath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 2.75l-6.25 11.5h12.5z"/>`
	trophyPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 10.75h6.5v3.5h-6.5zm3.25-2v2m-3.25-9c-1.5 0-3 .5-3 2.25s1.5 2.25 3 2.25m6.5-4.5c1.5 0 3 .5 3 2.25s-1.5 2.25-3 2.25m-6.5-4.5h6.5v3.5c0 1.5-1 3-3.25 3s-3.25-1.5-3.25-3z"/>`
	umbrellaPath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 8.25s.5-6.5 6.25-6.5s6.25 6.5 6.25 6.5zm6 .5v4s0 1.5 1.5 1.5s1.5-1.5 1.5-1.5"/>`
	uploadPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 2.75h9m-8.5 6.5l4-3.5l4 3.5m-4 5v-8.5"/>`
	wifiPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/>`
	wifiFairPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/><path d="M4.25 8c2-1.75 5.5-1.75 7.5 0"/></g>`
	wifiPoorPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 4.75L8 13.25l6.25-8.5C11 2 5 2 1.75 4.75z"/><path d="M5 9c.75-1.75 5.25-1.75 6 0"/></g>`
	wifiSlashPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 3.25c-1.5 0-3.5 1.5-3.5 1.5L8 13.25l2.25-3m-1.5-7.5s2.977-.135 5.5 2l-2 2.5m-8-5.5l8 10.5"/>`
	wifiWarningPath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 4.75C11 2 5 2 1.75 4.75L8 13.25l1-1.5m3.25 2v0m0-6v3.5"/>`
	zoomInPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="7.5" cy="7.5" r="4.75"/><path d="M9.25 7.5h-3.5M7.5 5.75v3.5m3.75 2l3 3"/></g>`
	zoomOutPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="7.5" cy="7.5" r="4.75"/><path d="M9.25 7.5h-3.5m5.5 3.75l3 3"/></g>`
)

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 0 0")
)

func IconFromName(name string) g.Node {
	switch name {
	case "anchor":
		return Anchor()
	case "apps":
		return Apps()
	case "appsMinus":
		return AppsMinus()
	case "appsPlus":
		return AppsPlus()
	case "archive":
		return Archive()
	case "arrowDown":
		return ArrowDown()
	case "arrowDownLeft":
		return ArrowDownLeft()
	case "arrowDownRight":
		return ArrowDownRight()
	case "arrowLeft":
		return ArrowLeft()
	case "arrowRight":
		return ArrowRight()
	case "arrowUp":
		return ArrowUp()
	case "arrowUpLeft":
		return ArrowUpLeft()
	case "arrowUpRight":
		return ArrowUpRight()
	case "atSign":
		return AtSign()
	case "atom":
		return Atom()
	case "bell":
		return Bell()
	case "bellSlash":
		return BellSlash()
	case "bin":
		return Bin()
	case "binary":
		return Binary()
	case "block":
		return Block()
	case "bluetooth":
		return Bluetooth()
	case "bluetoothConnected":
		return BluetoothConnected()
	case "bluetoothSearching":
		return BluetoothSearching()
	case "bluetoothSlash":
		return BluetoothSlash()
	case "book":
		return Book()
	case "bookOpen":
		return BookOpen()
	case "bookmark":
		return Bookmark()
	case "briefcase":
		return Briefcase()
	case "bug":
		return Bug()
	case "calendar":
		return Calendar()
	case "camera":
		return Camera()
	case "cameraVideo":
		return CameraVideo()
	case "cameraVideoSlash":
		return CameraVideoSlash()
	case "candy":
		return Candy()
	case "cards":
		return Cards()
	case "cast":
		return Cast()
	case "certificate":
		return Certificate()
	case "chartBar":
		return ChartBar()
	case "chartLine":
		return ChartLine()
	case "chevronDown":
		return ChevronDown()
	case "chevronLeft":
		return ChevronLeft()
	case "chevronRight":
		return ChevronRight()
	case "chevronUp":
		return ChevronUp()
	case "chevronsDown":
		return ChevronsDown()
	case "chevronsLeft":
		return ChevronsLeft()
	case "chevronsRight":
		return ChevronsRight()
	case "chevronsUp":
		return ChevronsUp()
	case "chevronsUpDown":
		return ChevronsUpDown()
	case "chip":
		return Chip()
	case "circle":
		return Circle()
	case "circleCross":
		return CircleCross()
	case "circleMinus":
		return CircleMinus()
	case "circleTick":
		return CircleTick()
	case "circleWarning":
		return CircleWarning()
	case "clipboard":
		return Clipboard()
	case "clipboardTick":
		return ClipboardTick()
	case "clock":
		return Clock()
	case "clockAlarm":
		return ClockAlarm()
	case "cloud":
		return Cloud()
	case "clover":
		return Clover()
	case "code":
		return Code()
	case "coffee":
		return Coffee()
	case "cog":
		return Cog()
	case "compass":
		return Compass()
	case "conicalFlask":
		return ConicalFlask()
	case "container":
		return Container()
	case "copy":
		return Copy()
	case "copyleft":
		return Copyleft()
	case "copyright":
		return Copyright()
	case "creditCard":
		return CreditCard()
	case "crop":
		return Crop()
	case "cross":
		return Cross()
	case "crosshair":
		return Crosshair()
	case "cube":
		return Cube()
	case "cursor":
		return Cursor()
	case "database":
		return Database()
	case "diamond":
		return Diamond()
	case "diff":
		return Diff()
	case "disc":
		return Disc()
	case "download":
		return Download()
	case "droplet":
		return Droplet()
	case "eraser":
		return Eraser()
	case "extensions":
		return Extensions()
	case "eye":
		return Eye()
	case "eyeSlash":
		return EyeSlash()
	case "faceFrown":
		return FaceFrown()
	case "faceNeutral":
		return FaceNeutral()
	case "faceSmile":
		return FaceSmile()
	case "file":
		return File()
	case "fileBinary":
		return FileBinary()
	case "fileCode":
		return FileCode()
	case "fileSymlink":
		return FileSymlink()
	case "files":
		return Files()
	case "filter":
		return Filter()
	case "flag":
		return Flag()
	case "flame":
		return Flame()
	case "floppyDisk":
		return FloppyDisk()
	case "folder":
		return Folder()
	case "folderSymlink":
		return FolderSymlink()
	case "folders":
		return Folders()
	case "forward":
		return Forward()
	case "gamepad":
		return Gamepad()
	case "gem":
		return Gem()
	case "gift":
		return Gift()
	case "gitBranch":
		return GitBranch()
	case "gitCherryPick":
		return GitCherryPick()
	case "gitCommit":
		return GitCommit()
	case "gitCompare":
		return GitCompare()
	case "gitFork":
		return GitFork()
	case "gitMerge":
		return GitMerge()
	case "gitRequest":
		return GitRequest()
	case "gitRequestCross":
		return GitRequestCross()
	case "gitRequestDraft":
		return GitRequestDraft()
	case "github":
		return Github()
	case "gitlab":
		return Gitlab()
	case "glasses":
		return Glasses()
	case "globe":
		return Globe()
	case "grabHorizontal":
		return GrabHorizontal()
	case "grabVertical":
		return GrabVertical()
	case "graduateCap":
		return GraduateCap()
	case "hash":
		return Hash()
	case "headphones":
		return Headphones()
	case "heart":
		return Heart()
	case "help":
		return Help()
	case "hexagon":
		return Hexagon()
	case "home":
		return Home()
	case "hourglass":
		return Hourglass()
	case "id":
		return Id()
	case "image":
		return Image()
	case "inbox":
		return Inbox()
	case "infinity":
		return Infinity()
	case "info":
		return Info()
	case "key":
		return Key()
	case "laptop":
		return Laptop()
	case "layoutColumns":
		return LayoutColumns()
	case "layoutDashboard":
		return LayoutDashboard()
	case "layoutGrid":
		return LayoutGrid()
	case "layoutList":
		return LayoutList()
	case "layoutRows":
		return LayoutRows()
	case "layoutSidebar":
		return LayoutSidebar()
	case "layoutStackH":
		return LayoutStackH()
	case "layoutStackV":
		return LayoutStackV()
	case "lightbulb":
		return Lightbulb()
	case "lightningBolt":
		return LightningBolt()
	case "link":
		return Link()
	case "linkExternal":
		return LinkExternal()
	case "linkSlash":
		return LinkSlash()
	case "mail":
		return Mail()
	case "map":
		return Map()
	case "mapPin":
		return MapPin()
	case "mediaBack":
		return MediaBack()
	case "mediaEject":
		return MediaEject()
	case "mediaFastForward":
		return MediaFastForward()
	case "mediaPause":
		return MediaPause()
	case "mediaPlay":
		return MediaPlay()
	case "mediaRewind":
		return MediaRewind()
	case "mediaSkip":
		return MediaSkip()
	case "menuHamburger":
		return MenuHamburger()
	case "menuKebab":
		return MenuKebab()
	case "menuMeatball":
		return MenuMeatball()
	case "message":
		return Message()
	case "messages":
		return Messages()
	case "microphone":
		return Microphone()
	case "minus":
		return Minus()
	case "mobile":
		return Mobile()
	case "monitor":
		return Monitor()
	case "monitorArrow":
		return MonitorArrow()
	case "monitorCross":
		return MonitorCross()
	case "moon":
		return Moon()
	case "move":
		return Move()
	case "music":
		return Music()
	case "newspaper":
		return Newspaper()
	case "northStar":
		return NorthStar()
	case "notes":
		return Notes()
	case "notesCross":
		return NotesCross()
	case "notesTick":
		return NotesTick()
	case "nut":
		return Nut()
	case "octagon":
		return Octagon()
	case "octagonWarning":
		return OctagonWarning()
	case "organisation":
		return Organisation()
	case "package":
		return Package()
	case "padlock":
		return Padlock()
	case "paperPlane":
		return PaperPlane()
	case "paperclip":
		return Paperclip()
	case "pencil":
		return Pencil()
	case "people":
		return People()
	case "person":
		return Person()
	case "phone":
		return Phone()
	case "phoneCall":
		return PhoneCall()
	case "phoneCross":
		return PhoneCross()
	case "phoneForward":
		return PhoneForward()
	case "phoneIncoming":
		return PhoneIncoming()
	case "phoneOutgoing":
		return PhoneOutgoing()
	case "pin":
		return Pin()
	case "plantPot":
		return PlantPot()
	case "plus":
		return Plus()
	case "power":
		return Power()
	case "printer":
		return Printer()
	case "pulse":
		return Pulse()
	case "quote":
		return Quote()
	case "refresh":
		return Refresh()
	case "reply":
		return Reply()
	case "robot":
		return Robot()
	case "rocket":
		return Rocket()
	case "rotateAntiClockwise":
		return RotateAntiClockwise()
	case "rotateClockwise":
		return RotateClockwise()
	case "scales":
		return Scales()
	case "screenMaximise":
		return ScreenMaximise()
	case "screenMinimise":
		return ScreenMinimise()
	case "search":
		return Search()
	case "server":
		return Server()
	case "share":
		return Share()
	case "shield":
		return Shield()
	case "shieldCross":
		return ShieldCross()
	case "shieldKeyhole":
		return ShieldKeyhole()
	case "shieldTick":
		return ShieldTick()
	case "shieldWarning":
		return ShieldWarning()
	case "shoppingBag":
		return ShoppingBag()
	case "signIn":
		return SignIn()
	case "signOut":
		return SignOut()
	case "signpost":
		return Signpost()
	case "skull":
		return Skull()
	case "snowflake":
		return Snowflake()
	case "soundDown":
		return SoundDown()
	case "soundMute":
		return SoundMute()
	case "soundUp":
		return SoundUp()
	case "speaker":
		return Speaker()
	case "square":
		return Square()
	case "squareCross":
		return SquareCross()
	case "squareTick":
		return SquareTick()
	case "stack":
		return Stack()
	case "stackPop":
		return StackPop()
	case "stackPush":
		return StackPush()
	case "star":
		return Star()
	case "stickyNote":
		return StickyNote()
	case "sun":
		return Sun()
	case "swapHorizontal":
		return SwapHorizontal()
	case "swapVertical":
		return SwapVertical()
	case "sword":
		return Sword()
	case "swords":
		return Swords()
	case "tablet":
		return Tablet()
	case "tag":
		return Tag()
	case "telescope":
		return Telescope()
	case "tent":
		return Tent()
	case "terminal":
		return Terminal()
	case "thumbDown":
		return ThumbDown()
	case "thumbUp":
		return ThumbUp()
	case "tick":
		return Tick()
	case "tickDouble":
		return TickDouble()
	case "ticket":
		return Ticket()
	case "treeFir":
		return TreeFir()
	case "triangle":
		return Triangle()
	case "trophy":
		return Trophy()
	case "umbrella":
		return Umbrella()
	case "upload":
		return Upload()
	case "wifi":
		return Wifi()
	case "wifiFair":
		return WifiFair()
	case "wifiPoor":
		return WifiPoor()
	case "wifiSlash":
		return WifiSlash()
	case "wifiWarning":
		return WifiWarning()
	case "zoomIn":
		return ZoomIn()
	case "zoomOut":
		return ZoomOut()
	default:
		panic("Unknown icon name: " + name)
	}
}

func Anchor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(anchorPath), g.Group(children))
}

func Apps(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(appsPath), g.Group(children))
}

func AppsMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(appsMinusPath), g.Group(children))
}

func AppsPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(appsPlusPath), g.Group(children))
}

func Archive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(archivePath), g.Group(children))
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownPath), g.Group(children))
}

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownLeftPath), g.Group(children))
}

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownRightPath), g.Group(children))
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

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpLeftPath), g.Group(children))
}

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpRightPath), g.Group(children))
}

func AtSign(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(atSignPath), g.Group(children))
}

func Atom(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(atomPath), g.Group(children))
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellPath), g.Group(children))
}

func BellSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellSlashPath), g.Group(children))
}

func Bin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(binPath), g.Group(children))
}

func Binary(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(binaryPath), g.Group(children))
}

func Block(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(blockPath), g.Group(children))
}

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothPath), g.Group(children))
}

func BluetoothConnected(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothConnectedPath), g.Group(children))
}

func BluetoothSearching(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothSearchingPath), g.Group(children))
}

func BluetoothSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothSlashPath), g.Group(children))
}

func Book(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookPath), g.Group(children))
}

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookOpenPath), g.Group(children))
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookmarkPath), g.Group(children))
}

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(briefcasePath), g.Group(children))
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bugPath), g.Group(children))
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarPath), g.Group(children))
}

func Camera(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cameraPath), g.Group(children))
}

func CameraVideo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cameraVideoPath), g.Group(children))
}

func CameraVideoSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cameraVideoSlashPath), g.Group(children))
}

func Candy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(candyPath), g.Group(children))
}

func Cards(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cardsPath), g.Group(children))
}

func Cast(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(castPath), g.Group(children))
}

func Certificate(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(certificatePath), g.Group(children))
}

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chartBarPath), g.Group(children))
}

func ChartLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chartLinePath), g.Group(children))
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronDownPath), g.Group(children))
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronLeftPath), g.Group(children))
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronRightPath), g.Group(children))
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronUpPath), g.Group(children))
}

func ChevronsDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsDownPath), g.Group(children))
}

func ChevronsLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsLeftPath), g.Group(children))
}

func ChevronsRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsRightPath), g.Group(children))
}

func ChevronsUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsUpPath), g.Group(children))
}

func ChevronsUpDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsUpDownPath), g.Group(children))
}

func Chip(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chipPath), g.Group(children))
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circlePath), g.Group(children))
}

func CircleCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleCrossPath), g.Group(children))
}

func CircleMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleMinusPath), g.Group(children))
}

func CircleTick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleTickPath), g.Group(children))
}

func CircleWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleWarningPath), g.Group(children))
}

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardPath), g.Group(children))
}

func ClipboardTick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardTickPath), g.Group(children))
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockPath), g.Group(children))
}

func ClockAlarm(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockAlarmPath), g.Group(children))
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudPath), g.Group(children))
}

func Clover(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloverPath), g.Group(children))
}

func Code(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codePath), g.Group(children))
}

func Coffee(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(coffeePath), g.Group(children))
}

func Cog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cogPath), g.Group(children))
}

func Compass(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(compassPath), g.Group(children))
}

func ConicalFlask(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(conicalFlaskPath), g.Group(children))
}

func Container(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(containerPath), g.Group(children))
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyPath), g.Group(children))
}

func Copyleft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyleftPath), g.Group(children))
}

func Copyright(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyrightPath), g.Group(children))
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(creditCardPath), g.Group(children))
}

func Crop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cropPath), g.Group(children))
}

func Cross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(crossPath), g.Group(children))
}

func Crosshair(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(crosshairPath), g.Group(children))
}

func Cube(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cubePath), g.Group(children))
}

func Cursor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cursorPath), g.Group(children))
}

func Database(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(databasePath), g.Group(children))
}

func Diamond(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diamondPath), g.Group(children))
}

func Diff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diffPath), g.Group(children))
}

func Disc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(discPath), g.Group(children))
}

func Download(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(downloadPath), g.Group(children))
}

func Droplet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropletPath), g.Group(children))
}

func Eraser(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eraserPath), g.Group(children))
}

func Extensions(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(extensionsPath), g.Group(children))
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eyePath), g.Group(children))
}

func EyeSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eyeSlashPath), g.Group(children))
}

func FaceFrown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(faceFrownPath), g.Group(children))
}

func FaceNeutral(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(faceNeutralPath), g.Group(children))
}

func FaceSmile(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(faceSmilePath), g.Group(children))
}

func File(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePath), g.Group(children))
}

func FileBinary(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileBinaryPath), g.Group(children))
}

func FileCode(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCodePath), g.Group(children))
}

func FileSymlink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileSymlinkPath), g.Group(children))
}

func Files(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filesPath), g.Group(children))
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filterPath), g.Group(children))
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flagPath), g.Group(children))
}

func Flame(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flamePath), g.Group(children))
}

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(floppyDiskPath), g.Group(children))
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderPath), g.Group(children))
}

func FolderSymlink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderSymlinkPath), g.Group(children))
}

func Folders(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(foldersPath), g.Group(children))
}

func Forward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(forwardPath), g.Group(children))
}

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gamepadPath), g.Group(children))
}

func Gem(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gemPath), g.Group(children))
}

func Gift(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(giftPath), g.Group(children))
}

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitBranchPath), g.Group(children))
}

func GitCherryPick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitCherryPickPath), g.Group(children))
}

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitCommitPath), g.Group(children))
}

func GitCompare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitComparePath), g.Group(children))
}

func GitFork(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitForkPath), g.Group(children))
}

func GitMerge(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitMergePath), g.Group(children))
}

func GitRequest(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitRequestPath), g.Group(children))
}

func GitRequestCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitRequestCrossPath), g.Group(children))
}

func GitRequestDraft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitRequestDraftPath), g.Group(children))
}

func Github(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(githubPath), g.Group(children))
}

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitlabPath), g.Group(children))
}

func Glasses(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(glassesPath), g.Group(children))
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(globePath), g.Group(children))
}

func GrabHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(grabHorizontalPath), g.Group(children))
}

func GrabVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(grabVerticalPath), g.Group(children))
}

func GraduateCap(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(graduateCapPath), g.Group(children))
}

func Hash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hashPath), g.Group(children))
}

func Headphones(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headphonesPath), g.Group(children))
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartPath), g.Group(children))
}

func Help(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(helpPath), g.Group(children))
}

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hexagonPath), g.Group(children))
}

func Home(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(homePath), g.Group(children))
}

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hourglassPath), g.Group(children))
}

func Id(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(idPath), g.Group(children))
}

func Image(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(imagePath), g.Group(children))
}

func Inbox(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(inboxPath), g.Group(children))
}

func Infinity(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(infinityPath), g.Group(children))
}

func Info(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(infoPath), g.Group(children))
}

func Key(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keyPath), g.Group(children))
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(laptopPath), g.Group(children))
}

func LayoutColumns(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutColumnsPath), g.Group(children))
}

func LayoutDashboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutDashboardPath), g.Group(children))
}

func LayoutGrid(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutGridPath), g.Group(children))
}

func LayoutList(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutListPath), g.Group(children))
}

func LayoutRows(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutRowsPath), g.Group(children))
}

func LayoutSidebar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutSidebarPath), g.Group(children))
}

func LayoutStackH(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutStackHPath), g.Group(children))
}

func LayoutStackV(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutStackVPath), g.Group(children))
}

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lightbulbPath), g.Group(children))
}

func LightningBolt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lightningBoltPath), g.Group(children))
}

func Link(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkPath), g.Group(children))
}

func LinkExternal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkExternalPath), g.Group(children))
}

func LinkSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkSlashPath), g.Group(children))
}

func Mail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailPath), g.Group(children))
}

func Map(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mapPath), g.Group(children))
}

func MapPin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mapPinPath), g.Group(children))
}

func MediaBack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaBackPath), g.Group(children))
}

func MediaEject(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaEjectPath), g.Group(children))
}

func MediaFastForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaFastForwardPath), g.Group(children))
}

func MediaPause(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaPausePath), g.Group(children))
}

func MediaPlay(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaPlayPath), g.Group(children))
}

func MediaRewind(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaRewindPath), g.Group(children))
}

func MediaSkip(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mediaSkipPath), g.Group(children))
}

func MenuHamburger(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(menuHamburgerPath), g.Group(children))
}

func MenuKebab(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(menuKebabPath), g.Group(children))
}

func MenuMeatball(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(menuMeatballPath), g.Group(children))
}

func Message(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messagePath), g.Group(children))
}

func Messages(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messagesPath), g.Group(children))
}

func Microphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(microphonePath), g.Group(children))
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minusPath), g.Group(children))
}

func Mobile(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mobilePath), g.Group(children))
}

func Monitor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorPath), g.Group(children))
}

func MonitorArrow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorArrowPath), g.Group(children))
}

func MonitorCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorCrossPath), g.Group(children))
}

func Moon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moonPath), g.Group(children))
}

func Move(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(movePath), g.Group(children))
}

func Music(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(musicPath), g.Group(children))
}

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(newspaperPath), g.Group(children))
}

func NorthStar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(northStarPath), g.Group(children))
}

func Notes(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(notesPath), g.Group(children))
}

func NotesCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(notesCrossPath), g.Group(children))
}

func NotesTick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(notesTickPath), g.Group(children))
}

func Nut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(nutPath), g.Group(children))
}

func Octagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(octagonPath), g.Group(children))
}

func OctagonWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(octagonWarningPath), g.Group(children))
}

func Organisation(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(organisationPath), g.Group(children))
}

func Package(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packagePath), g.Group(children))
}

func Padlock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(padlockPath), g.Group(children))
}

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paperPlanePath), g.Group(children))
}

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paperclipPath), g.Group(children))
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pencilPath), g.Group(children))
}

func People(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(peoplePath), g.Group(children))
}

func Person(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(personPath), g.Group(children))
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phonePath), g.Group(children))
}

func PhoneCall(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneCallPath), g.Group(children))
}

func PhoneCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneCrossPath), g.Group(children))
}

func PhoneForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneForwardPath), g.Group(children))
}

func PhoneIncoming(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneIncomingPath), g.Group(children))
}

func PhoneOutgoing(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneOutgoingPath), g.Group(children))
}

func Pin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pinPath), g.Group(children))
}

func PlantPot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plantPotPath), g.Group(children))
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plusPath), g.Group(children))
}

func Power(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(powerPath), g.Group(children))
}

func Printer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(printerPath), g.Group(children))
}

func Pulse(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pulsePath), g.Group(children))
}

func Quote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(quotePath), g.Group(children))
}

func Refresh(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(refreshPath), g.Group(children))
}

func Reply(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(replyPath), g.Group(children))
}

func Robot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(robotPath), g.Group(children))
}

func Rocket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rocketPath), g.Group(children))
}

func RotateAntiClockwise(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rotateAntiClockwisePath), g.Group(children))
}

func RotateClockwise(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rotateClockwisePath), g.Group(children))
}

func Scales(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scalesPath), g.Group(children))
}

func ScreenMaximise(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(screenMaximisePath), g.Group(children))
}

func ScreenMinimise(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(screenMinimisePath), g.Group(children))
}

func Search(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchPath), g.Group(children))
}

func Server(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(serverPath), g.Group(children))
}

func Share(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sharePath), g.Group(children))
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldPath), g.Group(children))
}

func ShieldCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldCrossPath), g.Group(children))
}

func ShieldKeyhole(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldKeyholePath), g.Group(children))
}

func ShieldTick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldTickPath), g.Group(children))
}

func ShieldWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldWarningPath), g.Group(children))
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shoppingBagPath), g.Group(children))
}

func SignIn(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signInPath), g.Group(children))
}

func SignOut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signOutPath), g.Group(children))
}

func Signpost(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signpostPath), g.Group(children))
}

func Skull(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(skullPath), g.Group(children))
}

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(snowflakePath), g.Group(children))
}

func SoundDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(soundDownPath), g.Group(children))
}

func SoundMute(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(soundMutePath), g.Group(children))
}

func SoundUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(soundUpPath), g.Group(children))
}

func Speaker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(speakerPath), g.Group(children))
}

func Square(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squarePath), g.Group(children))
}

func SquareCross(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareCrossPath), g.Group(children))
}

func SquareTick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareTickPath), g.Group(children))
}

func Stack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stackPath), g.Group(children))
}

func StackPop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stackPopPath), g.Group(children))
}

func StackPush(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stackPushPath), g.Group(children))
}

func Star(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(starPath), g.Group(children))
}

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stickyNotePath), g.Group(children))
}

func Sun(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunPath), g.Group(children))
}

func SwapHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swapHorizontalPath), g.Group(children))
}

func SwapVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swapVerticalPath), g.Group(children))
}

func Sword(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swordPath), g.Group(children))
}

func Swords(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swordsPath), g.Group(children))
}

func Tablet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tabletPath), g.Group(children))
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tagPath), g.Group(children))
}

func Telescope(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(telescopePath), g.Group(children))
}

func Tent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tentPath), g.Group(children))
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(terminalPath), g.Group(children))
}

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thumbDownPath), g.Group(children))
}

func ThumbUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thumbUpPath), g.Group(children))
}

func Tick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tickPath), g.Group(children))
}

func TickDouble(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tickDoublePath), g.Group(children))
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ticketPath), g.Group(children))
}

func TreeFir(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(treeFirPath), g.Group(children))
}

func Triangle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trianglePath), g.Group(children))
}

func Trophy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trophyPath), g.Group(children))
}

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(umbrellaPath), g.Group(children))
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(uploadPath), g.Group(children))
}

func Wifi(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiPath), g.Group(children))
}

func WifiFair(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiFairPath), g.Group(children))
}

func WifiPoor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiPoorPath), g.Group(children))
}

func WifiSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiSlashPath), g.Group(children))
}

func WifiWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiWarningPath), g.Group(children))
}

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(zoomInPath), g.Group(children))
}

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(zoomOutPath), g.Group(children))
}
