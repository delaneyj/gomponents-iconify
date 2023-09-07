package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const (
	IconifyVersion                      = ""
	accessibilityPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="4" r="1"/><path d="m18 19l1-7l-6 1M5 8l3-3l5.5 3l-2.36 3.5m-6.9 3a5 5 0 0 0 6.88 6"/><path d="M13.76 17.5a5 5 0 0 0-6.88-6"/></g>`
	activityPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12h-4l-3 9L9 3l-3 9H2"/>`
	activitySquarePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M17 12h-2l-2 5l-2-10l-2 5H7"/></g>`
	airVentPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2M6 8h12m.3 9.7a2.5 2.5 0 0 1-3.16 3.83a2.53 2.53 0 0 1-1.14-2V12m-7.4 3.6A2 2 0 1 0 10 17v-5"/>`
	airplayPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1"/><path d="m12 15l5 6H7l5-6z"/></g>`
	alarmCheckPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21M9 13l2 2l4-4"/></g>`
	alarmClockPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M12 9v4l2 2M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21"/></g>`
	alarmClockOffPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.87 6.87a8 8 0 1 0 11.26 11.26m1.77-3.88a8 8 0 0 0-9.15-9.15M22 6l-3-3M6.26 18.67L4 21M2 2l20 20M4 4L2 6"/>`
	alarmMinusPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21M9 13h6"/></g>`
	alarmPlusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21m-8-11v6m-3-3h6"/></g>`
	albumPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M11 3v8l3-3l3 3V3"/></g>`
	alertCirclePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 8v4m0 4h.01"/></g>`
	alertOctagonPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2zM12 8v4m0 4h.01"/>`
	alertTrianglePath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.73 18l-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3ZM12 9v4m0 4h.01"/>`
	alignCenterPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m14 6H7m12 6H5"/>`
	alignCenterHorizontalPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h20m-12 4v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4m6-8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4m16 8v1a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1m0-8V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1"/>`
	alignCenterVerticalPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20M8 10H4a2 2 0 0 1-2-2V6c0-1.1.9-2 2-2h4m8 6h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4M8 20H7a2 2 0 0 1-2-2v-2c0-1.1.9-2 2-2h1m8 0h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1"/>`
	alignEndHorizontalPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="16" x="4" y="2" rx="2"/><rect width="6" height="9" x="14" y="9" rx="2"/><path d="M22 22H2"/></g>`
	alignEndVerticalPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="6" x="2" y="4" rx="2"/><rect width="9" height="6" x="9" y="14" rx="2"/><path d="M22 22V2"/></g>`
	alignHorizontalDistributeCenterPath = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M17 22v-5m0-10V2M7 22v-3M7 5V2"/></g>`
	alignHorizontalDistributeEndPath    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M10 2v20M20 2v20"/></g>`
	alignHorizontalDistributeStartPath  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M4 2v20M14 2v20"/></g>`
	alignHorizontalJustifyCenterPath    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="2" y="5" rx="2"/><rect width="6" height="10" x="16" y="7" rx="2"/><path d="M12 2v20"/></g>`
	alignHorizontalJustifyEndPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="2" y="5" rx="2"/><rect width="6" height="10" x="12" y="7" rx="2"/><path d="M22 2v20"/></g>`
	alignHorizontalJustifyStartPath     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="6" y="5" rx="2"/><rect width="6" height="10" x="16" y="7" rx="2"/><path d="M2 2v20"/></g>`
	alignHorizontalSpaceAroundPath      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="10" x="9" y="7" rx="2"/><path d="M4 22V2m16 20V2"/></g>`
	alignHorizontalSpaceBetweenPath     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="3" y="5" rx="2"/><rect width="6" height="10" x="15" y="7" rx="2"/><path d="M3 2v20M21 2v20"/></g>`
	alignJustifyPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18M3 12h18M3 18h18"/>`
	alignLeftPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m12 6H3m14 6H3"/>`
	alignRightPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m18 6H9m12 6H7"/>`
	alignStartHorizontalPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="16" x="4" y="6" rx="2"/><rect width="6" height="9" x="14" y="6" rx="2"/><path d="M22 2H2"/></g>`
	alignStartVerticalPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="9" height="6" x="6" y="14" rx="2"/><rect width="16" height="6" x="6" y="4" rx="2"/><path d="M2 2v20"/></g>`
	alignVerticalDistributeCenterPath   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M22 7h-5M7 7H1m21 10h-3M5 17H2"/></g>`
	alignVerticalDistributeEndPath      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 20h20M2 10h20"/></g>`
	alignVerticalDistributeStartPath    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 14h20M2 4h20"/></g>`
	alignVerticalJustifyCenterPath      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="16" rx="2"/><rect width="10" height="6" x="7" y="2" rx="2"/><path d="M2 12h20"/></g>`
	alignVerticalJustifyEndPath         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="12" rx="2"/><rect width="10" height="6" x="7" y="2" rx="2"/><path d="M2 22h20"/></g>`
	alignVerticalJustifyStartPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="16" rx="2"/><rect width="10" height="6" x="7" y="6" rx="2"/><path d="M2 2h20"/></g>`
	alignVerticalSpaceAroundPath        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="10" height="6" x="7" y="9" rx="2"/><path d="M22 20H2M22 4H2"/></g>`
	alignVerticalSpaceBetweenPath       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="15" rx="2"/><rect width="10" height="6" x="7" y="3" rx="2"/><path d="M2 21h20M2 3h20"/></g>`
	ampersandPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 12c0 4.4-3.6 8-8 8A4.5 4.5 0 0 1 5 15.5c0-6 8-4 8-8.5a3 3 0 1 0-6 0c0 3 2.5 8.5 12 13m-3-8h3"/>`
	ampersandsPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6c0 1.7 1.3 3 3 3c2.8 0 5-2.2 5-5m12 5c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6c0 1.7 1.3 3 3 3c2.8 0 5-2.2 5-5"/>`
	anchorPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="3"/><path d="M12 22V8m-7 4H2a10 10 0 0 0 20 0h-3"/></g>`
	angryPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 16s-1.5-2-4-2s-4 2-4 2m-.5-8L10 9m4 0l2.5-1M9 10h0m6 0h0"/></g>`
	annoyedPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 15h8M8 9h2m4 0h2"/></g>`
	antennaPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12L7 2m0 10l5-10m0 10l5-10m0 10l5-10M4.5 7h15M12 16v6"/>`
	aperturePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m14.31 8l5.74 9.94M9.69 8h11.48M7.38 12l5.74-9.94M9.69 16L3.95 6.06M14.31 16H2.83m13.79-4l-5.74 9.94"/></g>`
	appWindowPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M10 4v4M2 8h20M6 4v4"/></g>`
	applePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 20.94c1.5 0 2.75 1.06 4 1.06c3 0 6-8 6-12.22A4.91 4.91 0 0 0 17 5c-2.22 0-4 1.44-5 2c-1-.56-2.78-2-5-2a4.9 4.9 0 0 0-5 4.78C2 14 5 22 8 22c1.25 0 2.5-1.06 4-1.06Z"/><path d="M10 2c1 .5 2 2 2 5"/></g>`
	archivePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8m-10 4h4"/></g>`
	archiveRestorePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h2M20 8v11a2 2 0 0 1-2 2h-2m-7-6l3-3l3 3m-3-3v9"/></g>`
	archiveXPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8M9.5 17l5-5m-5 0l5 5"/></g>`
	areaChartPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><path d="M7 12v5h12V8l-5 5l-4-4Z"/></g>`
	armchairPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 9V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3"/><path d="M3 11v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v2H7v-2a2 2 0 0 0-4 0Zm2 7v2m14-2v2"/></g>`
	arrowBigDownPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6v6h4l-7 7l-7-7h4V6h6z"/>`
	arrowBigDownDashPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5H9m6 4v3h4l-7 7l-7-7h4V9h6z"/>`
	arrowBigLeftPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15h-6v4l-7-7l7-7v4h6v6z"/>`
	arrowBigLeftDashPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 15V9m-4 6h-3v4l-7-7l7-7v4h3v6z"/>`
	arrowBigRightPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9h6V5l7 7l-7 7v-4H6V9z"/>`
	arrowBigRightDashPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9v6m4-6h3V5l7 7l-7 7v-4H9V9z"/>`
	arrowBigUpPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 18v-6H5l7-7l7 7h-4v6H9z"/>`
	arrowBigUpDashPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19h6m-6-4v-3H5l7-7l7 7h-4v3H9z"/>`
	arrowDownPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m7-7l-7 7l-7-7"/>`
	arrowDownAZPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m13 4h-5m0 2V6.5a2.5 2.5 0 0 1 5 0V10m-5 4h5l-5 6h5"/>`
	arrowDownCirclePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 8v8m-4-4l4 4l4-4"/></g>`
	arrowDownFromLinePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 3H5m7 18V7m-6 8l6 6l6-6"/>`
	arrowDownLeftPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7L7 17m10 0H7V7"/>`
	arrowDownLeftFromCirclePath         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12a10 10 0 1 1 10 10M2 22l10-10M8 22H2v-6"/>`
	arrowDownLeftSquarePath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 8l-8 8m8 0H8V8"/></g>`
	arrowDownNarrowWidePath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m4 0h4m-4 4h7m-7 4h10"/>`
	arrowDownOneZeroPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 16l4 4l4-4m-4 4V4m10 6V4h-2m0 6h4"/><rect width="4" height="6" x="15" y="14" ry="2"/></g>`
	arrowDownRightPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10m0-10v10H7"/>`
	arrowDownRightFromCirclePath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22a10 10 0 1 1 10-10m0 10L12 12m10 4v6h-6"/>`
	arrowDownRightSquarePath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m8 8l8 8m0-8v8H8"/></g>`
	arrowDownSquarePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 8v8m-4-4l4 4l4-4"/></g>`
	arrowDownToDotPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2v14m7-7l-7 7l-7-7"/><circle cx="12" cy="21" r="1"/></g>`
	arrowDownToLinePath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17V3m-6 8l6 6l6-6m1 10H5"/>`
	arrowDownUpPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m14 4l-4-4l-4 4m4-4v16"/>`
	arrowDownWideNarrowPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m4 0h10M11 8h7m-7 4h4"/>`
	arrowDownZAPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4M7 4v16m8-16h5l-5 6h5m-5 10v-3.5a2.5 2.5 0 0 1 5 0V20m0-2h-5"/>`
	arrowDownZeroOnePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 16l4 4l4-4m-4 4V4"/><rect width="4" height="6" x="15" y="4" ry="2"/><path d="M17 20v-6h-2m0 6h4"/></g>`
	arrowLeftPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l-7-7l7-7m7 7H5"/>`
	arrowLeftCirclePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 12H8m4-4l-4 4l4 4"/></g>`
	arrowLeftFromLinePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 6l-6 6l6 6m-6-6h14m4 7V5"/>`
	arrowLeftRightPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3L4 7l4 4M4 7h16m-4 14l4-4l-4-4m4 4H4"/>`
	arrowLeftSquarePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m12 8l-4 4l4 4m4-4H8"/></g>`
	arrowLeftToLinePath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19V5m10 1l-6 6l6 6m-6-6h14"/>`
	arrowRightPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7-7l7 7l-7 7"/>`
	arrowRightCirclePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8m-4 4l4-4l-4-4"/></g>`
	arrowRightFromLinePath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5v14m18-7H7m8 6l6-6l-6-6"/>`
	arrowRightLeftPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 3l4 4l-4 4m4-4H4m4 14l-4-4l4-4m-4 4h16"/>`
	arrowRightSquarePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8m-4 4l4-4l-4-4"/></g>`
	arrowRightToLinePath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12H3m8 6l6-6l-6-6m10-1v14"/>`
	arrowUpPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l7-7l7 7m-7 7V5"/>`
	arrowUpAZPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16M20 8h-5m0 2V6.5a2.5 2.5 0 0 1 5 0V10m-5 4h5l-5 6h5"/>`
	arrowUpCirclePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16 12l-4-4l-4 4m4 4V8"/></g>`
	arrowUpDownPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 16l-4 4l-4-4m4 4V4M3 8l4-4l4 4M7 4v16"/>`
	arrowUpFromDotPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m5 9l7-7l7 7m-7 7V2"/><circle cx="12" cy="21" r="1"/></g>`
	arrowUpFromLinePath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 9l-6-6l-6 6m6-6v14m-7 4h14"/>`
	arrowUpLeftPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17V7h10m0 10L7 7"/>`
	arrowUpLeftFromCirclePath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8V2h6M2 2l10 10m0-10A10 10 0 1 1 2 12"/>`
	arrowUpLeftSquarePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 16V8h8m0 8L8 8"/></g>`
	arrowUpNarrowWidePath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16m4-8h4m-4 4h7m-7 4h10"/>`
	arrowUpOneZeroPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 8l4-4l4 4M7 4v16m10-10V4h-2m0 6h4"/><rect width="4" height="6" x="15" y="14" ry="2"/></g>`
	arrowUpRightPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h10v10M7 17L17 7"/>`
	arrowUpRightFromCirclePath          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12A10 10 0 1 1 12 2m10 0L12 12m4-10h6v6"/>`
	arrowUpRightSquarePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 8h8v8m-8 0l8-8"/></g>`
	arrowUpSquarePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 12l-4-4l-4 4m4 4V8"/></g>`
	arrowUpToLinePath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3h14m-1 10l-6-6l-6 6m6-6v14"/>`
	arrowUpWideNarrowPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16m4-8h10m-10 4h7m-7 4h4"/>`
	arrowUpZAPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16m8-16h5l-5 6h5m-5 10v-3.5a2.5 2.5 0 0 1 5 0V20m0-2h-5"/>`
	arrowUpZeroOnePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 8l4-4l4 4M7 4v16"/><rect width="4" height="6" x="15" y="4" ry="2"/><path d="M17 20v-6h-2m0 6h4"/></g>`
	arrowsUpFromLinePath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 6l3-3l3 3M7 17V3m7 3l3-3l3 3m-3 11V3M4 21h16"/>`
	asteriskPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v12m5.196-9L6.804 15m0-6l10.392 6"/>`
	atSignPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8"/></g>`
	atomPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><path d="M20.2 20.2c2.04-2.03.02-7.36-4.5-11.9c-4.54-4.52-9.87-6.54-11.9-4.5c-2.04 2.03-.02 7.36 4.5 11.9c4.54 4.52 9.87 6.54 11.9 4.5Z"/><path d="M15.7 15.7c4.52-4.54 6.54-9.87 4.5-11.9c-2.03-2.04-7.36-.02-11.9 4.5c-4.52 4.54-6.54 9.87-4.5 11.9c2.03 2.04 7.36.02 11.9-4.5Z"/></g>`
	awardPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="6"/><path d="M15.477 12.89L17 22l-5-3l-5 3l1.523-9.11"/></g>`
	axePath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 12l-8.5 8.5a2.12 2.12 0 1 1-3-3L11 9"/><path d="M15 13L9 7l4-4l6 6h3a8 8 0 0 1-7 7z"/></g>`
	axisThreeDPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16h16M4 20l7-7"/>`
	babyPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 12h.01M15 12h.01M10 16c.5.3 1.2.5 2 .5s1.5-.2 2-.5"/><path d="M19 6.3a9 9 0 0 1 1.8 3.9a2 2 0 0 1 0 3.6a9 9 0 0 1-17.6 0a2 2 0 0 1 0-3.6A9 9 0 0 1 12 3c2 0 3.5 1.1 3.5 2.5s-.9 2.5-2 2.5c-.8 0-1.5-.4-1.5-1"/></g>`
	backpackPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20V10a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2ZM9 6V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2"/><path d="M8 21v-5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v5M8 10h8m-8 8h8"/></g>`
	badgePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/>`
	badgeAlertPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 8v4m0 4h.01"/>`
	badgeCentPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 7v10"/><path d="M15.4 10a4 4 0 1 0 0 4"/></g>`
	badgeCheckPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="m9 12l2 2l4-4"/></g>`
	badgeDollarSignPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 2V6"/></g>`
	badgeEuroPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM7 12h5"/><path d="M15 9.4a4 4 0 1 0 0 5.2"/></g>`
	badgeHelpPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3m.08 4h.01"/></g>`
	badgeIndianRupeePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM8 8h8m-8 4h8"/><path d="m13 17l-5-1h1a4 4 0 0 0 0-8"/></g>`
	badgeInfoPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 16v-4m0-4h.01"/>`
	badgeJapaneseYenPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="m9 8l3 3v7m0-7l3-3m-6 4h6m-6 4h6"/></g>`
	badgeMinusPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM8 12h8"/>`
	badgePercentPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM15 9l-6 6m0-6h.01M15 15h.01"/>`
	badgePlusPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 8v8m-4-4h8"/>`
	badgePoundSterlingPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM8 12h4"/><path d="M10 16V9.5a2.5 2.5 0 0 1 5 0M8 16h7"/></g>`
	badgeRussianRublePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM9 16h5"/><path d="M9 12h5a2 2 0 1 0 0-4h-3v9"/></g>`
	badgeSwissFrancPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="M11 17V8h4m-4 4h3m-5 4h4"/></g>`
	badgeXPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM15 9l-6 6m0-6l6 6"/>`
	baggageClaimPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 18H6a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2"/><path d="M17 14V4a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v10"/><rect width="13" height="8" x="8" y="6" rx="1"/><circle cx="18" cy="20" r="2"/><circle cx="9" cy="20" r="2"/></g>`
	banPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m4.9 4.9l14.2 14.2"/></g>`
	bananaPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 13c3.5-2 8-2 10 2a5.5 5.5 0 0 1 8 5"/><path d="M5.15 17.89c5.52-1.52 8.65-6.89 7-12C11.55 4 11.5 2 13 2c3.22 0 5 5.5 5 8c0 6.5-4.2 12-10.49 12C5.11 22 2 22 2 20c0-1.5 1.14-1.55 3.15-2.11Z"/></g>`
	banknotePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="2"/><circle cx="12" cy="12" r="2"/><path d="M6 12h.01M18 12h.01"/></g>`
	barChartPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V10m6 10V4M6 20v-4"/>`
	barChartBigPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><rect width="4" height="7" x="7" y="10" rx="1"/><rect width="4" height="12" x="15" y="5" rx="1"/></g>`
	barChartFourPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18m-8-4V9m5 8V5M8 17v-3"/>`
	barChartHorizontalPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18M7 16h8m-8-5h12M7 6h3"/>`
	barChartHorizontalBigPath           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><rect width="12" height="4" x="7" y="5" rx="1"/><rect width="7" height="4" x="7" y="13" rx="1"/></g>`
	barChartThreePath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18m-3-4V9m-5 8V5M8 17v-3"/>`
	barChartTwoPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V10m-6 10V4M6 20v-6"/>`
	baselinePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16M6 16l6-12l6 12M8 12h8"/>`
	bathPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6L6.5 3.5a1.5 1.5 0 0 0-1-.5C4.683 3 4 3.683 4 4.5V17a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M10 5L8 7m-6 5h20M7 19v2m10-2v2"/>`
	batteryPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2"/></g>`
	batteryChargingPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2M6 7H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h1m6-10l-3 5h4l-3 5m13-6v2"/>`
	batteryFullPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2M6 11v2m4-2v2m4-2v2"/></g>`
	batteryLowPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2M6 11v2"/></g>`
	batteryMediumPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2M6 11v2m4-2v2"/></g>`
	batteryWarningPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 7h2a2 2 0 0 1 2 2v6c0 1-1 2-2 2h-2M6 7H4a2 2 0 0 0-2 2v6c0 1 1 2 2 2h2m16-6v2M10 7v6m0 4v.01"/>`
	beakerPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.5 3h15M6 3v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3M6 14h12"/>`
	beanPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.165 6.598C9.954 7.478 9.64 8.36 9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22c7.732 0 14-6.268 14-14a6 6 0 0 0-11.835-1.402Z"/><path d="M5.341 10.62a4 4 0 1 0 5.279-5.28"/></g>`
	beanOffPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22a13.96 13.96 0 0 0 9.9-4.1M10.75 5.093A6 6 0 0 1 22 8c0 2.411-.61 4.68-1.683 6.66"/><path d="M5.341 10.62a4 4 0 0 0 6.487 1.208M10.62 5.341a4.015 4.015 0 0 1 2.039 2.04M2 2l20 20"/></g>`
	bedPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 4v16M2 8h18a2 2 0 0 1 2 2v10M2 17h20M6 8v9"/>`
	bedDoublePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20v-8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8M4 10V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4m-8-6v6M2 18h20"/>`
	bedSinglePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 20v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8M5 10V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4M3 18h18"/>`
	beefPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12.5" cy="8.5" r="2.5"/><path d="M12.5 2a6.5 6.5 0 0 0-6.22 4.6c-1.1 3.13-.78 3.9-3.18 6.08A3 3 0 0 0 5 18c4 0 8.4-1.8 11.4-4.3A6.5 6.5 0 0 0 12.5 2Z"/><path d="m18.5 6l2.19 4.5a6.48 6.48 0 0 1 .31 2a6.49 6.49 0 0 1-2.6 5.2C15.4 20.2 11 22 7 22a3 3 0 0 1-2.68-1.66L2.4 16.5"/></g>`
	beerPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 11h1a3 3 0 0 1 0 6h-1m-8-5v6m4-6v6m1-10.5c-1 0-1.44.5-3 .5s-2-.5-3-.5s-1.72.5-2.5.5a2.5 2.5 0 0 1 0-5c.78 0 1.57.5 2.5.5S9.44 2 11 2s2 1.5 3 1.5s1.72-.5 2.5-.5a2.5 2.5 0 0 1 0 5c-.78 0-1.5-.5-2.5-.5Z"/><path d="M5 8v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8"/></g>`
	bellPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9m4.3 13a1.94 1.94 0 0 0 3.4 0"/>`
	bellDotPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.4 14.9C20.2 16.4 21 17 21 17H3s3-2 3-9c0-3.3 2.7-6 6-6c.7 0 1.3.1 1.9.3M10.3 21a1.94 1.94 0 0 0 3.4 0"/><circle cx="18" cy="8" r="3"/></g>`
	bellMinusPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.4 12c.8 3.8 2.6 5 2.6 5H3s3-2 3-9c0-3.3 2.7-6 6-6c1.8 0 3.4.8 4.5 2m-6.2 17a1.94 1.94 0 0 0 3.4 0M15 8h6"/>`
	bellOffPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.7 3A6 6 0 0 1 18 8a21.3 21.3 0 0 0 .6 5M17 17H3s3-2 3-9a4.67 4.67 0 0 1 .3-1.7m4 14.7a1.94 1.94 0 0 0 3.4 0M2 2l20 20"/>`
	bellPlusPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.3 14.8C20.1 16.4 21 17 21 17H3s3-2 3-9c0-3.3 2.7-6 6-6c1 0 1.9.2 2.8.7M10.3 21a1.94 1.94 0 0 0 3.4 0M15 8h6m-3-3v6"/>`
	bellRingPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9m4.3 13a1.94 1.94 0 0 0 3.4 0M4 2C2.8 3.7 2 5.7 2 8m20 0c0-2.3-.8-4.3-2-6"/>`
	bikePath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18.5" cy="17.5" r="3.5"/><circle cx="5.5" cy="17.5" r="3.5"/><circle cx="15" cy="5" r="1"/><path d="M12 17.5V14l-3-3l4-3l2 3h2"/></g>`
	binaryPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="4" height="6" x="14" y="14" rx="2"/><rect width="4" height="6" x="6" y="4" rx="2"/><path d="M6 20h4m4-10h4M6 14h2v6m6-16h2v6"/></g>`
	biohazardPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11.9" r="2"/><path d="M6.7 3.4c-.9 2.5 0 5.2 2.2 6.7C6.5 9 3.7 9.6 2 11.6m6.9-1.5l1.4.8m7-7.5c.9 2.5 0 5.2-2.2 6.7c2.4-1.2 5.2-.6 6.9 1.5m-6.9-1.5l-1.4.8m3 9.9c-2.6-.4-4.6-2.6-4.7-5.3c-.2 2.6-2.1 4.8-4.7 5.2m4.7-6.8v1.6m1.5-10.1c-1-.2-2-.2-3 0m6.5 11c.7-.7 1.2-1.6 1.5-2.5m-13 0c.3.9.8 1.8 1.5 2.5"/></g>`
	birdPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 7h.01M3.4 18H12a8 8 0 0 0 8-8V7a4 4 0 0 0-7.28-2.3L2 20"/><path d="m20 7l2 .5l-2 .5M10 18v3m4-3.25V21m-7-3a6 6 0 0 0 3.84-10.61"/></g>`
	bitcoinPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.767 19.089c4.924.868 6.14-6.025 1.216-6.894m-1.216 6.894L5.86 18.047m5.908 1.042l-.347 1.97m1.563-8.864c4.924.869 6.14-6.025 1.215-6.893m-1.215 6.893l-3.94-.694m5.155-6.2L8.29 4.26m5.908 1.042l.348-1.97M7.48 20.364l3.126-17.727"/>`
	blindsPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3h18m-1 4H8m12 4H8m2 8h10M8 15h12M4 3v14"/><circle cx="4" cy="19" r="2"/></g>`
	blocksPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="14" y="3" rx="1"/><path d="M10 21V8a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H3"/></g>`
	bluetoothPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10l-5 5V2l5 5L7 17"/>`
	bluetoothConnectedPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10l-5 5V2l5 5L7 17m11-5h3M3 12h3"/>`
	bluetoothOffPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 17l-5 5V12l-5 5M2 2l20 20M14.5 9.5L17 7l-5-5v4.5"/>`
	bluetoothSearchingPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10l-5 5V2l5 5L7 17m13.83-2.17a4 4 0 0 0 0-5.66M18 12h.01"/>`
	boldPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 12a4 4 0 0 0 0-8H6v8m9 8a4 4 0 0 0 0-8H6v8Z"/>`
	bombPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="13" r="9"/><path d="m19.5 9.5l1.8-1.8a2.4 2.4 0 0 0 0-3.4l-1.6-1.6a2.41 2.41 0 0 0-3.4 0l-1.8 1.8M22 2l-1.5 1.5"/></g>`
	bonePath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 10c.7-.7 1.69 0 2.5 0a2.5 2.5 0 1 0 0-5a.5.5 0 0 1-.5-.5a2.5 2.5 0 1 0-5 0c0 .81.7 1.8 0 2.5l-7 7c-.7.7-1.69 0-2.5 0a2.5 2.5 0 0 0 0 5c.28 0 .5.22.5.5a2.5 2.5 0 1 0 5 0c0-.81-.7-1.8 0-2.5Z"/>`
	bookPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/>`
	bookCopyPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 16V4a2 2 0 0 1 2-2h11"/><path d="M5 14H4a2 2 0 1 0 0 4h1m17 0H11a2 2 0 1 0 0 4h11V6H11a2 2 0 0 0-2 2v12"/></g>`
	bookDownPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3 3l3-3"/></g>`
	bookKeyPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H14"/><path d="M20 8v14H6.5a2.5 2.5 0 0 1 0-5H20"/><circle cx="14" cy="8" r="2"/><path d="m20 2l-4.5 4.5M19 3l1 1"/></g>`
	bookLockPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H10"/><path d="M20 15v7H6.5a2.5 2.5 0 0 1 0-5H20"/><rect width="8" height="5" x="12" y="6" rx="1"/><path d="M18 6V4a2 2 0 1 0-4 0v2"/></g>`
	bookMarkedPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><path d="M10 2v8l3-3l3 3V2"/></g>`
	bookMinusPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M9 10h6"/>`
	bookOpenPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2zm20 0h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/>`
	bookOpenCheckPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 3H2v15h7c1.7 0 3 1.3 3 3V7c0-2.2-1.8-4-4-4Zm8 9l2 2l4-4"/><path d="M22 6V3h-6c-2.2 0-4 1.8-4 4v14c0-1.7 1.3-3 3-3h7v-2.3"/></g>`
	bookPlusPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M9 10h6m-3-3v6"/>`
	bookTemplatePath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 22h-2m2-7v2h-2M4 19.5V15m16-7v3m-2-9h2v2M4 11V9m8-7h2m-2 20h2m-2-5h2m-6 5H6.5a2.5 2.5 0 0 1 0-5H8M4 5v-.5A2.5 2.5 0 0 1 6.5 2H8"/>`
	bookUpPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3-3l3 3"/></g>`
	bookUpTwoPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2"/><path d="M18 2h2v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3-3l3 3M9 5l3-3l3 3"/></g>`
	bookXPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M14.5 7l-5 5m0-5l5 5"/>`
	bookmarkPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z"/>`
	bookmarkMinusPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16zm-4-11H9"/>`
	bookmarkPlusPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16zM12 7v6m3-3H9"/>`
	boomBoxPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 9V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4M8 8v1m4-1v1m4-1v1"/><rect width="20" height="12" x="2" y="9" rx="2"/><circle cx="8" cy="15" r="2"/><circle cx="16" cy="15" r="2"/></g>`
	botPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="10" x="3" y="11" rx="2"/><circle cx="12" cy="5" r="2"/><path d="M12 7v4m-4 5h0m8 0h0"/></g>`
	boxPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"/><path d="m3.3 7l8.7 5l8.7-5M12 22V12"/></g>`
	boxSelectPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3a2 2 0 0 0-2 2m16-2a2 2 0 0 1 2 2m0 14a2 2 0 0 1-2 2M5 21a2 2 0 0 1-2-2M9 3h1M9 21h1m4-18h1m-1 18h1M3 9v1m18-1v1M3 14v1m18-1v1"/>`
	boxesPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.97 12.92A2 2 0 0 0 2 14.63v3.24a2 2 0 0 0 .97 1.71l3 1.8a2 2 0 0 0 2.06 0L12 19v-5.5l-5-3l-4.03 2.42ZM7 16.5l-4.74-2.85M7 16.5l5-3m-5 3v5.17m5-8.17V19l3.97 2.38a2 2 0 0 0 2.06 0l3-1.8a2 2 0 0 0 .97-1.71v-3.24a2 2 0 0 0-.97-1.71L17 10.5l-5 3Zm5 3l-5-3m5 3l4.74-2.85M17 16.5v5.17"/><path d="M7.97 4.42A2 2 0 0 0 7 6.13v4.37l5 3l5-3V6.13a2 2 0 0 0-.97-1.71l-3-1.8a2 2 0 0 0-2.06 0l-3 1.8ZM12 8L7.26 5.15M12 8l4.74-2.85M12 13.5V8"/></g>`
	bracesPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2a2 2 0 0 1 2 2v5c0 1.1.9 2 2 2h1m8 0h1a2 2 0 0 0 2-2v-5c0-1.1.9-2 2-2a2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1"/>`
	bracketsPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3h3v18h-3m-8 0H5V3h3"/>`
	brainPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 2A2.5 2.5 0 0 1 12 4.5v15a2.5 2.5 0 0 1-4.96.44a2.5 2.5 0 0 1-2.96-3.08a3 3 0 0 1-.34-5.58a2.5 2.5 0 0 1 1.32-4.24a2.5 2.5 0 0 1 1.98-3A2.5 2.5 0 0 1 9.5 2Zm5 0A2.5 2.5 0 0 0 12 4.5v15a2.5 2.5 0 0 0 4.96.44a2.5 2.5 0 0 0 2.96-3.08a3 3 0 0 0 .34-5.58a2.5 2.5 0 0 0-1.32-4.24a2.5 2.5 0 0 0-1.98-3A2.5 2.5 0 0 0 14.5 2Z"/>`
	brainCircuitPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4.5a2.5 2.5 0 0 0-4.96-.46a2.5 2.5 0 0 0-1.98 3a2.5 2.5 0 0 0-1.32 4.24a3 3 0 0 0 .34 5.58a2.5 2.5 0 0 0 2.96 3.08a2.5 2.5 0 0 0 4.91.05L12 20V4.5ZM16 8V5c0-1.1.9-2 2-2m-6 10h4"/><path d="M12 18h6a2 2 0 0 1 2 2v1M12 8h8m.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-4 5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/><path d="M20.5 21a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-2-18a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/></g>`
	brainCogPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M12 4.5a2.5 2.5 0 0 0-4.96-.46a2.5 2.5 0 0 0-1.98 3a2.5 2.5 0 0 0-1.32 4.24a3 3 0 0 0 .34 5.58a2.5 2.5 0 0 0 2.96 3.08A2.5 2.5 0 0 0 12 19.5a2.5 2.5 0 0 0 4.96.44a2.5 2.5 0 0 0 2.96-3.08a3 3 0 0 0 .34-5.58a2.5 2.5 0 0 0-1.32-4.24a2.5 2.5 0 0 0-1.98-3A2.5 2.5 0 0 0 12 4.5m3.7 5.9l-.9.4m-5.6 2.4l-.9.4m5.3 2.1l-.4-.9m-2.4-5.6l-.4-.9m5.3 5.2l-.9-.4m-5.6-2.2l-.9-.4m2.2 5.2l.4-.9m2.2-5.6l.4-.9"/></g>`
	briefcasePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="7" rx="2" ry="2"/><path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"/></g>`
	bringToFrontPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="8" y="8" rx="2"/><path d="M4 10a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2m4 16a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2"/></g>`
	brushPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.06 11.9l8.07-8.06a2.85 2.85 0 1 1 4.03 4.03l-8.06 8.08m-6.03-1.01c-1.66 0-3 1.35-3 3.02c0 1.33-2.5 1.52-2 2.02c1.08 1.1 2.49 2.02 4 2.02c2.2 0 4-1.8 4-4.04a3.01 3.01 0 0 0-3-3.02z"/>`
	bugPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 2l1.88 1.88m4.24 0L16 2M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"/><path d="M12 20c-3.3 0-6-2.7-6-6v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v3c0 3.3-2.7 6-6 6m0 0v-9"/><path d="M6.53 9C4.6 8.8 3 7.1 3 5m3 8H2m1 8c0-2.1 1.7-3.9 3.8-4M20.97 5c0 2.1-1.6 3.8-3.5 4M22 13h-4m-.8 4c2.1.1 3.8 1.9 3.8 4"/></g>`
	bugOffPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 7.13V6a3 3 0 0 0-5.14-2.1L8 2m6.12 1.88L16 2"/><path d="M22 13h-4v-2a4 4 0 0 0-4-4h-1.3"/><path d="M20.97 5c0 2.1-1.6 3.8-3.5 4M2 2l20 20M7.7 7.7A4 4 0 0 0 6 11v3a6 6 0 0 0 11.13 3.13M12 20v-8m-6 1H2"/><path d="M3 21c0-2.1 1.7-3.9 3.8-4"/></g>`
	bugPlayPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 2l1.88 1.88m4.24 0L16 2M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"/><path d="M18 11a4 4 0 0 0-4-4h-4a4 4 0 0 0-4 4v3a6.1 6.1 0 0 0 2 4.5"/><path d="M6.53 9C4.6 8.8 3 7.1 3 5m3 8H2m1 8c0-2.1 1.7-3.9 3.8-4M20.97 5c0 2.1-1.6 3.8-3.5 4M12 12l8 5l-8 5Z"/></g>`
	buildingPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><path d="M9 22v-4h6v4M8 6h.01M16 6h.01M12 6h.01M12 10h.01M12 14h.01M16 10h.01M16 14h.01M8 10h.01M8 14h.01"/></g>`
	buildingTwoPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 22V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v18Zm0-10H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2M18 9h2a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-2M10 6h4m-4 4h4m-4 4h4m-4 4h4"/>`
	busPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 6v6m7-6v6M2 12h19.6M18 18h3s.5-1.7.8-2.8c.1-.4.2-.8.2-1.2c0-.4-.1-.8-.2-1.2l-1.4-5C20.1 6.8 19.1 6 18 6H4a2 2 0 0 0-2 2v10h3"/><circle cx="7" cy="18" r="2"/><path d="M9 18h5"/><circle cx="16" cy="18" r="2"/></g>`
	busFrontPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6L2 7m8-1h4m8 1l-2-1"/><rect width="16" height="16" x="4" y="3" rx="2"/><path d="M4 11h16M8 15h.01M16 15h.01M6 19v2m12 0v-2"/></g>`
	cablePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 9a2 2 0 0 1-2-2V5h6v2a2 2 0 0 1-2 2ZM3 5V3m4 2V3"/><path d="M19 15V6.5a3.5 3.5 0 0 0-7 0v11a3.5 3.5 0 0 1-7 0V9m12 12v-2m4 2v-2"/><path d="M22 19h-6v-2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2Z"/></g>`
	cableCarPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 3h.01M14 2h.01M2 9l20-5m-10 8V6.5"/><rect width="16" height="10" x="4" y="12" rx="3"/><path d="M9 12v5m6-5v5M4 17h16"/></g>`
	cakePath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 21v-8a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v8"/><path d="M4 16s.5-1 2-1s2.5 2 4 2s2.5-2 4-2s2.5 2 4 2s2-1 2-1M2 21h20M7 8v2m5-2v2m5-2v2M7 4h.01M12 4h.01M17 4h.01"/></g>`
	cakeSlicePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="7" r="2"/><path d="M7.2 7.9L3 11v9c0 .6.4 1 1 1h16c.6 0 1-.4 1-1v-9c0-2-3-6-7-8l-3.6 2.6M16 13H3m13 4H3"/></g>`
	calculatorPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M8 6h8m0 8v4m0-8h.01M12 10h.01M8 10h.01M12 14h.01M8 14h.01M12 18h.01M8 18h.01"/></g>`
	calendarPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18"/></g>`
	calendarCheckPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18M9 16l2 2l4-4"/></g>`
	calendarCheckTwoPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 14V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-5 10l2 2l4-4"/>`
	calendarClockPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 7.5V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.5M16 2v4M8 2v4m-5 4h5m9.5 7.5L16 16.25V14"/><path d="M22 16a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z"/></g>`
	calendarDaysPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18M8 14h.01M12 14h.01M16 14h.01M8 18h.01M12 18h.01M16 18h.01"/></g>`
	calendarHeartPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h7m4-20v4M8 2v4m-5 4h18"/><path d="M21.29 14.7a2.43 2.43 0 0 0-2.65-.52c-.3.12-.57.3-.8.53l-.34.34l-.35-.34a2.43 2.43 0 0 0-2.65-.53c-.3.12-.56.3-.79.53c-.95.94-1 2.53.2 3.74L17.5 22l3.6-3.55c1.2-1.21 1.14-2.8.19-3.74Z"/></g>`
	calendarMinusPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-5 9h6"/>`
	calendarOffPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.18 4.18A2 2 0 0 0 3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 1.82-1.18M21 15.5V6a2 2 0 0 0-2-2H9.5M16 2v4M3 10h7m11 0h-5.5M2 2l20 20"/>`
	calendarPlusPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-2 6v6m-3-3h6"/>`
	calendarRangePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18m-4 4h-6m2 4H7m0-4h.01M17 18h.01"/></g>`
	calendarSearchPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h7.5M16 2v4M8 2v4m-5 4h18"/><path d="M18 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6v0Zm4 1l-1.5-1.5"/></g>`
	calendarXPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18m-11 4l4 4m0-4l-4 4"/></g>`
	calendarXTwoPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-4 7l5 5m-5 0l5-5"/>`
	cameraPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 4h-5L7 7H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3l-2.5-3z"/><circle cx="12" cy="13" r="3"/></g>`
	cameraOffPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20M7 7H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16M9.5 4h5L17 7h3a2 2 0 0 1 2 2v7.5"/><path d="M14.121 15.121A3 3 0 1 1 9.88 10.88"/></g>`
	candlestickChartPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5v4"/><rect width="4" height="6" x="7" y="9" rx="1"/><path d="M9 15v2m8-14v2"/><rect width="4" height="8" x="15" y="5" rx="1"/><path d="M17 13v3M3 3v18h18"/></g>`
	candyPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9.5 7.5l-2 2a4.95 4.95 0 1 0 7 7l2-2a4.95 4.95 0 1 0-7-7Zm4.5-1v10m-4-9v10"/><path d="m16 7l1-5l1.37.68A3 3 0 0 0 19.7 3H21v1.3c0 .46.1.92.32 1.33L22 7l-5 1m-9 9l-1 5l-1.37-.68A3 3 0 0 0 4.3 21H3v-1.3a3 3 0 0 0-.32-1.33L2 17l5-1"/></g>`
	candyCanePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.7 21a2 2 0 0 1-3.5-2l8.6-14a6 6 0 0 1 10.4 6a2 2 0 1 1-3.464-2a2 2 0 1 0-3.464-2ZM17.75 7L15 2.1m-4.1 2.7L13 9m-5.1.7l2 4.4m-5 .6L7 18.9"/>`
	candyOffPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8.5 8.5l-1 1a4.95 4.95 0 0 0 7 7l1-1m-3.657-9.313A4.947 4.947 0 0 1 16.5 7.5a4.947 4.947 0 0 1 1.313 4.657M14 16.5V14m0-7.5v1.843M10 10v7.5"/><path d="m16 7l1-5l1.367.683A3 3 0 0 0 19.708 3H21v1.292a3 3 0 0 0 .317 1.341L22 7l-5 1m-9 9l-1 5l-1.367-.683A3 3 0 0 0 4.292 21H3v-1.292a3 3 0 0 0-.317-1.341L2 17l5-1M2 2l20 20"/></g>`
	carPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 17h2c.6 0 1-.4 1-1v-3c0-.9-.7-1.7-1.5-1.9C18.7 10.6 16 10 16 10s-1.3-1.4-2.2-2.3c-.5-.4-1.1-.7-1.8-.7H5c-.6 0-1.1.4-1.4.9l-1.4 2.9A3.7 3.7 0 0 0 2 12v4c0 .6.4 1 1 1h2"/><circle cx="7" cy="17" r="2"/><path d="M9 17h6"/><circle cx="17" cy="17" r="2"/></g>`
	carFrontPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 8l-2 2l-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10L3 8m4 6h.01M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2m14-2v2"/></g>`
	carTaxiFrontPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2h4m7 6l-2 2l-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10L3 8m4 6h.01M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2m14-2v2"/></g>`
	carrotPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.27 21.7s9.87-3.5 12.73-6.36a4.5 4.5 0 0 0-6.36-6.37C5.77 11.84 2.27 21.7 2.27 21.7zM8.64 14l-2.05-2.04M15.34 15l-2.46-2.46"/><path d="M22 9s-1.33-2-3.5-2C16.86 7 15 9 15 9s1.33 2 3.5 2S22 9 22 9z"/><path d="M15 2s-2 1.33-2 3.5S15 9 15 9s2-1.84 2-3.5C17 3.33 15 2 15 2z"/></g>`
	caseLowerPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="12" r="3"/><path d="M10 9v6"/><circle cx="17" cy="12" r="3"/><path d="M14 7v8"/></g>`
	caseSensitivePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 15l4-8l4 8m-7-2h6"/><circle cx="18" cy="12" r="3"/><path d="M21 9v6"/></g>`
	caseUpperPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15l4-8l4 8m-7-2h6m5-2h4.5a2 2 0 0 1 0 4H15V7h4a2 2 0 0 1 0 4"/>`
	cassetteTapePath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><circle cx="8" cy="10" r="2"/><path d="M8 12h8"/><circle cx="16" cy="10" r="2"/><path d="m6 20l.7-2.9A1.4 1.4 0 0 1 8.1 16h7.8a1.4 1.4 0 0 1 1.4 1l.7 3"/></g>`
	castPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6M2 12a9 9 0 0 1 8 8m-8-4a5 5 0 0 1 4 4m-4 0h.01"/>`
	castlePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 20v-9H2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2Zm-4-9V4H6v7"/><path d="M15 22v-4a3 3 0 0 0-3-3v0a3 3 0 0 0-3 3v4m13-11V9M2 11V9m4-5V2m12 2V2m-8 2V2m4 2V2"/></g>`
	catPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 5c.67 0 1.35.09 2 .26c1.78-2 5.03-2.84 6.42-2.26c1.4.58-.42 7-.42 7c.57 1.07 1 2.24 1 3.44C21 17.9 16.97 21 12 21s-9-3-9-7.56c0-1.25.5-2.4 1-3.44c0 0-1.89-6.42-.5-7c1.39-.58 4.72.23 6.5 2.23A9.04 9.04 0 0 1 12 5Zm-4 9v.5m8-.5v.5"/><path d="M11.25 16.25h1.5L12 17l-.75-.75Z"/></g>`
	checkPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6L9 17l-5-5"/>`
	checkCheckPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6L7 17l-5-5m20-2l-7.5 7.5L13 16"/>`
	checkCirclePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><path d="M22 4L12 14.01l-3-3"/></g>`
	checkCircleTwoPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="m9 12l2 2l4-4"/></g>`
	checkSquarePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 11l3 3L22 4"/><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/></g>`
	chefHatPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 13.87A4 4 0 0 1 7.41 6a5.11 5.11 0 0 1 1.05-1.54a5 5 0 0 1 7.08 0A5.11 5.11 0 0 1 16.59 6A4 4 0 0 1 18 13.87V21H6ZM6 17h12"/>`
	cherryPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 17a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3c-2.5-2-5 .24-5 3Zm10 0a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3c-2.5-2-5 .24-5 3Z"/><path d="M7 14c3.22-2.91 4.29-8.75 5-12c1.66 2.38 4.94 9 5 12"/><path d="M22 9c-4.29 0-7.14-2.33-10-7c5.71 0 10 4.67 10 7Z"/></g>`
	chevronDownPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 9l6 6l6-6"/>`
	chevronDownCirclePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16 10l-4 4l-4-4"/></g>`
	chevronDownSquarePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 10l-4 4l-4-4"/></g>`
	chevronFirstPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 18l-6-6l6-6M7 6v12"/>`
	chevronLastPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 18l6-6l-6-6m10 0v12"/>`
	chevronLeftPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 18l-6-6l6-6"/>`
	chevronLeftCirclePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m14 16l-4-4l4-4"/></g>`
	chevronLeftSquarePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m14 16l-4-4l4-4"/></g>`
	chevronRightPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 18l6-6l-6-6"/>`
	chevronRightCirclePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m10 8l4 4l-4 4"/></g>`
	chevronRightSquarePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m10 8l4 4l-4 4"/></g>`
	chevronUpPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 15l-6-6l-6 6"/>`
	chevronUpCirclePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m8 14l4-4l4 4"/></g>`
	chevronUpSquarePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m8 14l4-4l4 4"/></g>`
	chevronsDownPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 6l5 5l5-5M7 13l5 5l5-5"/>`
	chevronsDownUpPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 20l5-5l5 5M7 4l5 5l5-5"/>`
	chevronsLeftPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 17l-5-5l5-5m7 10l-5-5l5-5"/>`
	chevronsLeftRightPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 7l-5 5l5 5m6-10l5 5l-5 5"/>`
	chevronsRightPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 17l5-5l-5-5m7 10l5-5l-5-5"/>`
	chevronsRightLeftPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 17l-5-5l5-5M4 17l5-5l-5-5"/>`
	chevronsUpPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 11l-5-5l-5 5m10 7l-5-5l-5 5"/>`
	chevronsUpDownPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 15l5 5l5-5M7 9l5-5l5 5"/>`
	chromePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="M21.17 8H12M3.95 6.06L8.54 14m2.34 7.94L15.46 14"/></g>`
	churchPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18 7l4 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9l4-2"/><path d="M14 22v-4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v4"/><path d="M18 22V5l-6-3l-6 3v17m6-15v5m-2-3h4"/></g>`
	cigarettePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 12H2v4h16m4-4v4M7 12v4m11-8c0-2.5-2-2.5-2-5m6 5c0-2.5-2-2.5-2-5"/>`
	cigaretteOffPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M12 12H2v4h14m6-4v4m-4-4h-.5M7 12v4m11-8c0-2.5-2-2.5-2-5m6 5c0-2.5-2-2.5-2-5"/>`
	circlePath                          = `<circle cx="12" cy="12" r="10" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`
	circleDashedPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.1 2.18a9.93 9.93 0 0 1 3.8 0m3.7 1.53a9.95 9.95 0 0 1 2.69 2.7m1.53 3.69a9.93 9.93 0 0 1 0 3.8m-1.53 3.7a9.95 9.95 0 0 1-2.7 2.69m-3.69 1.53a9.94 9.94 0 0 1-3.8 0m-3.7-1.53a9.95 9.95 0 0 1-2.69-2.7M2.18 13.9a9.93 9.93 0 0 1 0-3.8m1.53-3.7a9.95 9.95 0 0 1 2.7-2.69"/>`
	circleDollarSignPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 2V6"/></g>`
	circleDotPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="1"/></g>`
	circleDotDashedPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.1 2.18a9.93 9.93 0 0 1 3.8 0m3.7 1.53a9.95 9.95 0 0 1 2.69 2.7m1.53 3.69a9.93 9.93 0 0 1 0 3.8m-1.53 3.7a9.95 9.95 0 0 1-2.7 2.69m-3.69 1.53a9.94 9.94 0 0 1-3.8 0m-3.7-1.53a9.95 9.95 0 0 1-2.69-2.7M2.18 13.9a9.93 9.93 0 0 1 0-3.8m1.53-3.7a9.95 9.95 0 0 1 2.7-2.69"/><circle cx="12" cy="12" r="1"/></g>`
	circleEllipsisPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M17 12h.01M12 12h.01M7 12h.01"/></g>`
	circleEqualPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10h10M7 14h10"/><circle cx="12" cy="12" r="10"/></g>`
	circleOffPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M8.35 2.69A10 10 0 0 1 21.3 15.65m-2.22 3.43A10 10 0 1 1 4.92 4.92"/>`
	circleSlashPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 15l6-6"/><circle cx="12" cy="12" r="10"/></g>`
	circleSlashTwoPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M22 2L2 22"/></g>`
	circuitBoardPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M11 9h4a2 2 0 0 0 2-2V3"/><circle cx="9" cy="9" r="2"/><path d="M7 21v-4a2 2 0 0 1 2-2h4"/><circle cx="15" cy="15" r="2"/></g>`
	citrusPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.66 17.67a1.08 1.08 0 0 1-.04 1.6A12 12 0 0 1 4.73 2.38a1.1 1.1 0 0 1 1.61-.04z"/><path d="M19.65 15.66A8 8 0 0 1 8.35 4.34M14 10l-5.5 5.5"/><path d="M14 17.85V10H6.15"/></g>`
	clapperboardPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.2 6L3 11l-.9-2.4c-.3-1.1.3-2.2 1.3-2.5l13.5-4c1.1-.3 2.2.3 2.5 1.3Zm-14-.7l3.1 3.9m3.1-5.8l3.1 4M3 11h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/>`
	clipboardPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/></g>`
	clipboardCheckPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="m9 14l2 2l4-4"/></g>`
	clipboardCopyPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2M16 4h2a2 2 0 0 1 2 2v4m1 4H11"/><path d="m15 10l-4 4l4 4"/></g>`
	clipboardEditPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M10.42 12.61a2.1 2.1 0 1 1 2.97 2.97L7.95 21L4 22l.99-3.95l5.43-5.44Z"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-5.5M4 13.5V6a2 2 0 0 1 2-2h2"/></g>`
	clipboardListPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2m4 7h4m-4 5h4m-8-5h.01M8 16h.01"/></g>`
	clipboardPastePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 2H9a1 1 0 0 0-1 1v2c0 .6.4 1 1 1h6c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1Z"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2M16 4h2a2 2 0 0 1 2 2v2m-9 6h10"/><path d="m17 10l4 4l-4 4"/></g>`
	clipboardSignaturePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-.5M16 4h2a2 2 0 0 1 1.73 1"/><path d="M18.42 9.61a2.1 2.1 0 1 1 2.97 2.97L16.95 17L13 18l.99-3.95l4.43-4.44ZM8 18h1"/></g>`
	clipboardTypePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M9 12v-1h6v1m-4 5h2m-1-6v6"/></g>`
	clipboardXPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2m7 7l-6 6m0-6l6 6"/></g>`
	clockPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></g>`
	clockEightPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l-4 2"/></g>`
	clockElevenPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6L9.5 8"/></g>`
	clockFivePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l2.5 4"/></g>`
	clockFourPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></g>`
	clockNinePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6H7.5"/></g>`
	clockOnePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l2.5-4"/></g>`
	clockSevenPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l-2.5 4"/></g>`
	clockSixPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v10.5"/></g>`
	clockTenPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l-4-2"/></g>`
	clockThreePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6h4.5"/></g>`
	clockTwelvePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6"/></g>`
	clockTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4-2"/></g>`
	cloudPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"/>`
	cloudCogPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="17" r="3"/><path d="M4.2 15.1A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.2m-4.3 2.2l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`
	cloudDrizzlePath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M8 19v1m0-6v1m8 4v1m0-6v1m-4 6v1m0-6v1"/>`
	cloudFogPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M16 17H7m10 4H9"/>`
	cloudHailPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M16 14v2m-8-2v2m8 4h.01M8 20h.01M12 16v2m0 4h.01"/>`
	cloudLightningPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 16.326A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 .5 8.973"/><path d="m13 12l-3 5h4l-3 5"/></g>`
	cloudMoonPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16a3 3 0 1 1 0 6H7a5 5 0 1 1 4.9-6Zm-2.9-7A6 6 0 0 1 16 4a4.24 4.24 0 0 0 6 6a6 6 0 0 1-3 5.197"/>`
	cloudMoonRainPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.083 9A6.002 6.002 0 0 1 16 4a4.243 4.243 0 0 0 6 6c0 2.22-1.206 4.16-3 5.197M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24M11 20v2m-4-3v2"/>`
	cloudOffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M5.782 5.782A7 7 0 0 0 9 19h8.5a4.5 4.5 0 0 0 1.307-.193m2.725-2.307A4.5 4.5 0 0 0 17.5 10h-1.79A7.008 7.008 0 0 0 10 5.07"/>`
	cloudRainPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M16 14v6m-8-6v6m4-4v6"/>`
	cloudRainWindPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M9.2 22l3-7M9 13l-3 7m11-7l-3 7"/>`
	cloudSnowPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M8 15h.01M8 19h.01M12 17h.01M12 21h.01M16 15h.01M16 19h.01"/>`
	cloudSunPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v2m-7.07.93l1.41 1.41M20 12h2m-2.93-7.07l-1.41 1.41m-1.713 6.31a4 4 0 0 0-5.925-4.128M13 22H7a5 5 0 1 1 4.9-6H13a3 3 0 0 1 0 6Z"/>`
	cloudSunRainPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v2m-7.07.93l1.41 1.41M20 12h2m-2.93-7.07l-1.41 1.41m-1.713 6.31a4 4 0 0 0-5.925-4.128M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24M11 20v2m-4-3v2"/>`
	cloudyPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 21H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"/><path d="M22 10a3 3 0 0 0-3-3h-2.207a5.502 5.502 0 0 0-10.702.5"/></g>`
	cloverPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.2 3.8a2.7 2.7 0 0 0-3.81 0l-.4.38l-.4-.4a2.7 2.7 0 0 0-3.82 0C6.73 4.85 6.67 6.64 8 8l4 4l4-4c1.33-1.36 1.27-3.15.2-4.2zM8 8c-1.36-1.33-3.15-1.27-4.2-.2a2.7 2.7 0 0 0 0 3.81l.38.4l-.4.4a2.7 2.7 0 0 0 0 3.82C4.85 17.27 6.64 17.33 8 16m8 0c1.36 1.33 3.15 1.27 4.2.2a2.7 2.7 0 0 0 0-3.81l-.38-.4l.4-.4a2.7 2.7 0 0 0 0-3.82C19.15 6.73 17.36 6.67 16 8"/><path d="M7.8 20.2a2.7 2.7 0 0 0 3.81 0l.4-.38l.4.4a2.7 2.7 0 0 0 3.82 0c1.06-1.06 1.12-2.85-.21-4.21l-4-4l-4 4c-1.33 1.36-1.27 3.15-.2 4.2zM7 17l-5 5"/></g>`
	clubPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.28 9.05a5.5 5.5 0 1 0-10.56 0A5.5 5.5 0 1 0 12 17.66a5.5 5.5 0 1 0 5.28-8.6ZM12 17.66V22"/>`
	codePath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 18l6-6l-6-6M8 6l-6 6l6 6"/>`
	codeTwoPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 16l4-4l-4-4M6 8l-4 4l4 4m8.5-12l-5 16"/>`
	codepenPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 2l10 6.5v7L12 22L2 15.5v-7L12 2zm0 20v-6.5"/><path d="m22 8.5l-10 7l-10-7"/><path d="m2 15.5l10-7l10 7M12 2v6.5"/></g>`
	codesandboxPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><path d="m7.5 4.21l4.5 2.6l4.5-2.6m-9 15.58V14.6L3 12m18 0l-4.5 2.6v5.19M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12"/></g>`
	coffeePath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h1a4 4 0 1 1 0 8h-1M3 8h14v9a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4Zm3-6v2m4-2v2m4-2v2"/>`
	cogPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-12v2m0 18v-2m5 .66l-1-1.73m-5-8.66L7 3.34M20.66 17l-1.73-1M3.34 7l1.73 1M14 12h8M2 12h2m16.66-5l-1.73 1M3.34 17l1.73-1M17 3.34l-1 1.73m-5 8.66l-4 6.93"/></g>`
	coinsPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="8" r="6"/><path d="M18.09 10.37A6 6 0 1 1 10.34 18M7 6h1v4"/><path d="m16.71 13.88l.7.71l-2.82 2.82"/></g>`
	columnsPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M12 3v18"/></g>`
	combinePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="2" y="2" rx="2"/><path d="M14 2c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2m6-8c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2m-10 8H5c-1.7 0-3-1.3-3-3v-1"/><path d="m7 21l3-3l-3-3"/><rect width="8" height="8" x="14" y="14" rx="2"/></g>`
	commandPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"/>`
	compassPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16.24 7.76l-2.12 6.36l-6.36 2.12l2.12-6.36l6.36-2.12z"/></g>`
	componentPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.5 8.5L9 12l-3.5 3.5L2 12l3.5-3.5ZM12 2l3.5 3.5L12 9L8.5 5.5L12 2Zm6.5 6.5L22 12l-3.5 3.5L15 12l3.5-3.5ZM12 15l3.5 3.5L12 22l-3.5-3.5L12 15Z"/>`
	computerPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="8" x="5" y="2" rx="2"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h2m4 0h6"/></g>`
	conciergeBellPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2H2v-2Zm18-2a8 8 0 1 0-16 0m8-12v4m-2-4h4"/>`
	constructionPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="6" rx="1"/><path d="M17 14v7M7 14v7M17 3v3M7 3v3m3 8L2.3 6.3M14 6l7.7 7.7M8 6l8 8"/></g>`
	contactPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 18a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2"/><rect width="18" height="18" x="3" y="4" rx="2"/><circle cx="12" cy="10" r="2"/><path d="M8 2v2m8-2v2"/></g>`
	contactTwoPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 18a4 4 0 0 0-8 0"/><circle cx="12" cy="11" r="3"/><rect width="18" height="18" x="3" y="4" rx="2"/><path d="M8 2v2m8-2v2"/></g>`
	containerPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 7.7c0-.6-.4-1.2-.8-1.5l-6.3-3.9a1.72 1.72 0 0 0-1.7 0l-10.3 6c-.5.2-.9.8-.9 1.4v6.6c0 .5.4 1.2.8 1.5l6.3 3.9a1.72 1.72 0 0 0 1.7 0l10.3-6c.5-.3.9-1 .9-1.5Z"/><path d="M10 21.9V14L2.1 9.1M10 14l11.9-6.9M14 19.8v-8.1m4 5.8V9.4"/></g>`
	contrastPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 18a6 6 0 0 0 0-12v12z"/></g>`
	cookiePath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2a10 10 0 1 0 10 10a4 4 0 0 1-5-5a4 4 0 0 1-5-5M8.5 8.5v.01M16 15.5v.01M12 12v.01M11 17v.01M7 14v.01"/>`
	copyPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`
	copyCheckPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 15l2 2l4-4"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`
	copyMinusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 15h6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`
	copyPlusPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 12v6m-3-3h6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`
	copySlashPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 18l6-6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`
	copyXPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 12l6 6m-6 0l6-6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`
	copyleftPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 9.35a4 4 0 1 1 0 5.3"/></g>`
	copyrightPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M15 9.354a4 4 0 1 0 0 5.292"/></g>`
	cornerDownLeftPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l-5 5l5 5"/><path d="M20 4v7a4 4 0 0 1-4 4H4"/></g>`
	cornerDownRightPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 10l5 5l-5 5"/><path d="M4 4v7a4 4 0 0 0 4 4h12"/></g>`
	cornerLeftDownPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 15l-5 5l-5-5"/><path d="M20 4h-7a4 4 0 0 0-4 4v12"/></g>`
	cornerLeftUpPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 9L9 4L4 9"/><path d="M20 20h-7a4 4 0 0 1-4-4V4"/></g>`
	cornerRightDownPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 15l5 5l5-5"/><path d="M4 4h7a4 4 0 0 1 4 4v12"/></g>`
	cornerRightUpPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 9l5-5l5 5"/><path d="M4 20h7a4 4 0 0 0 4-4V4"/></g>`
	cornerUpLeftPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 14L4 9l5-5"/><path d="M20 20v-7a4 4 0 0 0-4-4H4"/></g>`
	cornerUpRightPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 14l5-5l-5-5"/><path d="M4 20v-7a4 4 0 0 1 4-4h12"/></g>`
	cpuPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" rx="2"/><path d="M9 9h6v6H9zm6-7v2m0 16v2M2 15h2M2 9h2m16 6h2m-2-6h2M9 2v2m0 16v2"/></g>`
	creativeCommonsPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 9.3a2.8 2.8 0 0 0-3.5 1a3.1 3.1 0 0 0 0 3.4a2.7 2.7 0 0 0 3.5 1m7-5.4a2.8 2.8 0 0 0-3.5 1a3.1 3.1 0 0 0 0 3.4a2.7 2.7 0 0 0 3.5 1"/></g>`
	creditCardPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="5" rx="2"/><path d="M2 10h20"/></g>`
	croissantPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4.6 13.11l5.79-3.21c1.89-1.05 4.79 1.78 3.71 3.71l-3.22 5.81C8.8 23.16.79 15.23 4.6 13.11Zm5.9-3.61l-1-2.29C9.2 6.48 8.8 6 8 6H4.5C2.79 6 2 6.5 2 8.5a7.71 7.71 0 0 0 2 4.83M8 6c0-1.55.24-4-2-4c-2 0-2.5 2.17-2.5 4m11 7.5l2.29 1c.73.3 1.21.7 1.21 1.5v3.5c0 1.71-.5 2.5-2.5 2.5a7.71 7.71 0 0 1-4.83-2M18 16c1.55 0 4-.24 4 2c0 2-2.17 2.5-4 2.5"/>`
	cropPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 2v14a2 2 0 0 0 2 2h14"/><path d="M18 22V8a2 2 0 0 0-2-2H2"/></g>`
	crossPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 2a2 2 0 0 0-2 2v5H4a2 2 0 0 0-2 2v2c0 1.1.9 2 2 2h5v5c0 1.1.9 2 2 2h2a2 2 0 0 0 2-2v-5h5a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-5V4a2 2 0 0 0-2-2h-2z"/>`
	crosshairPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M22 12h-4M6 12H2m10-6V2m0 20v-4"/></g>`
	crownPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 4l3 12h14l3-12l-6 7l-4-7l-4 7l-6-7zm3 16h14"/>`
	cupSodaPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 8l1.75 12.28a2 2 0 0 0 2 1.72h4.54a2 2 0 0 0 2-1.72L18 8M5 8h14"/><path d="M7 15a6.47 6.47 0 0 1 5 0a6.47 6.47 0 0 0 5 0m-5-7l1-6h2"/></g>`
	currencyPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="m3 3l3 3m15-3l-3 3M3 21l3-3m15 3l-3-3"/></g>`
	databasePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 18 0V5"/><path d="M3 12a9 3 0 0 0 18 0"/></g>`
	databaseBackupPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 12a9 3 0 0 0 5 2.69M21 9.3V5"/><path d="M3 5v14a9 3 0 0 0 6.47 2.88M12 12v4h4"/><path d="M13 20a5 5 0 0 0 9-3a4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L12 16"/></g>`
	databaseZapPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 12 2.84M21 5v3m0 4l-3 5h4l-3 5"/><path d="M3 12a9 3 0 0 0 11.59 2.87"/></g>`
	deletePath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 5H9l-7 7l7 7h11a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Zm-2 4l-6 6m0-6l6 6"/>`
	dessertPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="4" r="2"/><path d="M10.2 3.2C5.5 4 2 8.1 2 13a2 2 0 0 0 4 0v-1a2 2 0 0 1 4 0v4a2 2 0 0 0 4 0v-4a2 2 0 0 1 4 0v1a2 2 0 0 0 4 0c0-4.9-3.5-9-8.2-9.8"/><path d="M3.2 14.8a9 9 0 0 0 17.6 0"/></g>`
	diamondPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Z"/>`
	diceFivePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M8 8h.01M8 16h.01M16 16h.01M12 12h.01"/></g>`
	diceFourPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M8 8h.01M8 16h.01M16 16h.01"/></g>`
	diceOnePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M12 12h.01"/></g>`
	diceSixPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M16 12h.01M16 16h.01M8 8h.01M8 12h.01M8 16h.01"/></g>`
	diceThreePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M12 12h.01M8 16h.01"/></g>`
	diceTwoPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 9h.01M9 15h.01"/></g>`
	dicesPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="12" height="12" x="2" y="10" rx="2" ry="2"/><path d="m17.92 14l3.5-3.5a2.24 2.24 0 0 0 0-3l-5-4.92a2.24 2.24 0 0 0-3 0L10 6M6 18h.01M10 14h.01M15 6h.01M18 9h.01"/></g>`
	diffPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v14m-7-7h14M5 21h14"/>`
	discPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="2"/></g>`
	discThreePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M6 12c0-1.7.7-3.2 1.8-4.2"/><circle cx="12" cy="12" r="2"/><path d="M18 12c0 1.7-.7 3.2-1.8 4.2"/></g>`
	discTwoPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="M12 12h.01"/></g>`
	dividePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="6" r="1"/><path d="M5 12h14"/><circle cx="12" cy="18" r="1"/></g>`
	divideCirclePath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 12h8m-4 4h0m0-8h0"/><circle cx="12" cy="12" r="10"/></g>`
	divideSquarePath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M8 12h8m-4 4h0m0-8h0"/></g>`
	dnaPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 15c6.667-6 13.333 0 20-6M9 22c1.798-1.998 2.518-3.995 2.807-5.993M15 2c-1.798 1.998-2.518 3.995-2.807 5.993M17 6l-2.5-2.5M14 8l-1-1M7 18l2.5 2.5m-6-6l.5.5m16-6l.5.5m-14 3l1 1m9-3l1 1M10 16l1.5 1.5"/>`
	dnaOffPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 2c-1.35 1.5-2.092 3-2.5 4.5M9 22c1.35-1.5 2.092-3 2.5-4.5M2 15c3.333-3 6.667-3 10-3m10-3c-1.5 1.35-3 2.092-4.5 2.5M17 6l-2.5-2.5M14 8l-1.5-1.5M7 18l2.5 2.5m-6-6l.5.5m16-6l.5.5m-14 3l1 1m9-3l1 1M10 16l1.5 1.5M2 2l20 20"/>`
	dogPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 5.172C10 3.782 8.423 2.679 6.5 3c-2.823.47-4.113 6.006-4 7c.08.703 1.725 1.722 3.656 1c1.261-.472 1.96-1.45 2.344-2.5m5.767-3.328c0-1.39 1.577-2.493 3.5-2.172c2.823.47 4.113 6.006 4 7c-.08.703-1.725 1.722-3.656 1c-1.261-.472-1.855-1.45-2.239-2.5M8 14v.5m8-.5v.5m-4.75 1.75h1.5L12 17l-.75-.75Z"/><path d="M4.42 11.247A13.152 13.152 0 0 0 4 14.556C4 18.728 7.582 21 12 21s8-2.272 8-6.444c0-1.061-.162-2.2-.493-3.309m-9.243-6.082A8.801 8.801 0 0 1 12 5c.78 0 1.5.108 2.161.306"/></g>`
	dollarSignPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20m5-17H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>`
	donutPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.5 10a2.5 2.5 0 0 1-2.4-3H18a2.95 2.95 0 0 1-2.6-4.4a10 10 0 1 0 6.3 7.1c-.3.2-.8.3-1.2.3"/><circle cx="12" cy="12" r="3"/></g>`
	doorClosedPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14m-4 0h20m-8-8v.01"/>`
	doorOpenPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4h3a2 2 0 0 1 2 2v14M2 20h3m8 0h9m-12-8v.01m3-7.448v16.157a1 1 0 0 1-1.242.97L5 20V5.562a2 2 0 0 1 1.515-1.94l4-1A2 2 0 0 1 13 4.561Z"/>`
	dotPath                             = `<circle cx="12.1" cy="12.1" r="1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`
	downloadPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4m4-5l5 5l5-5m-5 5V3"/>`
	downloadCloudPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M12 12v9m-4-4l4 4l4-4"/>`
	dribbblePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M19.13 5.09C15.22 9.14 10 10.44 2.25 10.94m19.5 1.9c-6.62-1.41-12.14 1-16.38 6.32"/><path d="M8.56 2.75c4.37 6 6 9.42 8 17.72"/></g>`
	dropletPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22a7 7 0 0 0 7-7c0-2-1-3.9-3-5.5s-3.5-4-4-6.5c-.5 2.5-2 4.9-4 6.5C6 11.1 5 13 5 15a7 7 0 0 0 7 7z"/>`
	dropletsPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 16.3c2.2 0 4-1.83 4-4.05c0-1.16-.57-2.26-1.71-3.19S7.29 6.75 7 5.3c-.29 1.45-1.14 2.84-2.29 3.76S3 11.1 3 12.25c0 2.22 1.8 4.05 4 4.05z"/><path d="M12.56 6.6A10.97 10.97 0 0 0 14 3.02c.5 2.5 2 4.9 4 6.5s3 3.5 3 5.5a6.98 6.98 0 0 1-11.91 4.97"/></g>`
	drumstickPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.45 15.4c-2.13.65-4.3.32-5.7-1.1c-2.29-2.27-1.76-6.5 1.17-9.42c2.93-2.93 7.15-3.46 9.43-1.18c1.41 1.41 1.74 3.57 1.1 5.71c-1.4-.51-3.26-.02-4.64 1.36c-1.38 1.38-1.87 3.23-1.36 4.63z"/><path d="m11.25 15.6l-2.16 2.16a2.5 2.5 0 1 1-4.56 1.73a2.49 2.49 0 0 1-1.41-4.24a2.5 2.5 0 0 1 3.14-.32l2.16-2.16"/></g>`
	dumbbellPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.5 6.5l11 11M21 21l-1-1M3 3l1 1m14 18l4-4M2 6l4-4m-3 8l7-7m4 18l7-7"/>`
	earPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 8.5a6.5 6.5 0 1 1 13 0c0 6-6 6-6 10a3.5 3.5 0 1 1-7 0"/><path d="M15 8.5a2.5 2.5 0 0 0-5 0v1a2 2 0 1 1 0 4"/></g>`
	earOffPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18.5a3.5 3.5 0 1 0 7 0c0-1.57.92-2.52 2.04-3.46M6 8.5c0-.75.13-1.47.36-2.14M8.8 3.15A6.5 6.5 0 0 1 19 8.5c0 1.63-.44 2.81-1.09 3.76"/><path d="M12.5 6A2.5 2.5 0 0 1 15 8.5M10 13a2 2 0 0 0 1.82-1.18M2 2l20 20"/></g>`
	editPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1l1-4l9.5-9.5z"/></g>`
	editThreePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1l1-4L16.5 3.5z"/>`
	editTwoPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5L17 3z"/>`
	eggPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22c6.23-.05 7.87-5.57 7.5-10c-.36-4.34-3.95-9.96-7.5-10c-3.55.04-7.14 5.66-7.5 10c-.37 4.43 1.27 9.95 7.5 10z"/>`
	eggFriedPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11.5" cy="12.5" r="3.5"/><path d="M3 8c0-3.5 2.5-6 6.5-6c5 0 4.83 3 7.5 5s5 2 5 6c0 4.5-2.5 6.5-7 6.5c-2.5 0-2.5 2.5-6 2.5s-7-2-7-5.5c0-3 1.5-3 1.5-5C3.5 10 3 9 3 8Z"/></g>`
	eggOffPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.399 6.399C5.362 8.157 4.65 10.189 4.5 12c-.37 4.43 1.27 9.95 7.5 10c3.256-.026 5.259-1.547 6.375-3.625m1.157-4.5A14.07 14.07 0 0 0 19.5 12c-.36-4.34-3.95-9.96-7.5-10c-1.04.012-2.082.502-3.046 1.297M2 2l20 20"/>`
	equalPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9h14M5 15h14"/>`
	equalNotPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9h14M5 15h14m0-10L5 19"/>`
	eraserPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 21l-4.3-4.3c-1-1-1-2.5 0-3.4l9.6-9.6c1-1 2.5-1 3.4 0l5.6 5.6c1 1 1 2.5 0 3.4L13 21m9 0H7M5 11l9 9"/>`
	euroPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10h12M4 14h9m6-8a7.7 7.7 0 0 0-5.2-2A7.9 7.9 0 0 0 6 12c0 4.4 3.5 8 7.8 8c2 0 3.8-.8 5.2-2"/>`
	expandPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-6-6m6 6v-4.8m0 4.8h-4.8M3 16.2V21m0 0h4.8M3 21l6-6m12-7.2V3m0 0h-4.8M21 3l-6 6M3 7.8V3m0 0h4.8M3 3l6 6"/>`
	externalLinkPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6m4-3h6v6m-11 5L21 3"/>`
	eyePath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12s3-7 10-7s10 7 10 7s-3 7-10 7s-10-7-10-7Z"/><circle cx="12" cy="12" r="3"/></g>`
	eyeOffPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.88 9.88a3 3 0 1 0 4.24 4.24m-3.39-9.04A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"/><path d="M6.61 6.61A13.526 13.526 0 0 0 2 12s3 7 10 7a9.74 9.74 0 0 0 5.39-1.61M2 2l20 20"/></g>`
	facebookPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>`
	factoryPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8l-7 5V8l-7 5V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Zm15-2h1m-6 0h1m-6 0h1"/>`
	fanPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.827 16.379a6.082 6.082 0 0 1-8.618-7.002l5.412 1.45a6.082 6.082 0 0 1 7.002-8.618l-1.45 5.412a6.082 6.082 0 0 1 8.618 7.002l-5.412-1.45a6.082 6.082 0 0 1-7.002 8.618l1.45-5.412ZM12 12v.01"/>`
	fastForwardPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 19l9-7l-9-7v14zM2 19l9-7l-9-7v14z"/>`
	featherPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5zM16 8L2 22m15.5-7H9"/>`
	ferrisWheelPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="2"/><path d="M12 2v4m-5.2 9l-3.5 2M20.7 7l-3.5 2M6.8 9L3.3 7m17.4 10l-3.5-2M9 22l3-8l3 8m-7 0h8"/><path d="M18 18.7a9 9 0 1 0-12 0"/></g>`
	figmaPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 5.5A3.5 3.5 0 0 1 8.5 2H12v7H8.5A3.5 3.5 0 0 1 5 5.5zM12 2h3.5a3.5 3.5 0 1 1 0 7H12V2z"/><path d="M12 12.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 1 1-7 0zm-7 7A3.5 3.5 0 0 1 8.5 16H12v3.5a3.5 3.5 0 1 1-7 0zm0-7A3.5 3.5 0 0 1 8.5 9H12v7H8.5A3.5 3.5 0 0 1 5 12.5z"/></g>`
	filePath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/></g>`
	fileArchivePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22V4c0-.5.2-1 .6-1.4C5 2.2 5.5 2 6 2h8.5L20 7.5V20c0 .5-.2 1-.6 1.4c-.4.4-.9.6-1.4.6h-2"/><path d="M14 2v6h6"/><circle cx="10" cy="20" r="2"/><path d="M10 7V6m0 6v-1m0 7v-2"/></g>`
	fileAudioPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 22h.5c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V7.5L14.5 2H6c-.5 0-1 .2-1.4.6C4.2 3 4 3.5 4 4v3"/><path d="M14 2v6h6M10 20v-1a2 2 0 1 1 4 0v1a2 2 0 1 1-4 0Zm-4 0v-1a2 2 0 1 0-4 0v1a2 2 0 1 0 4 0Z"/><path d="M2 19v-3a6 6 0 0 1 12 0v3"/></g>`
	fileAudioTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v2"/><path d="M14 2v6h6M2 17v-3a4 4 0 0 1 8 0v3"/><circle cx="9" cy="17" r="1"/><circle cx="3" cy="17" r="1"/></g>`
	fileAxisThreeDPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M8 10v8h8m-8 0l4-4"/></g>`
	fileBadgePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 7V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2h-6"/><path d="M14 2v6h6M5 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M7 16.5L8 22l-3-1l-3 1l1-5.5"/></g>`
	fileBadgeTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M12 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="m14 12.5l1 5.5l-3-1l-3 1l1-5.5"/></g>`
	fileBarChartPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-4m-4 4v-2m8 2v-6"/></g>`
	fileBarChartTwoPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-6m-4 6v-1m8 1v-3"/></g>`
	fileBoxPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 22H18a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2.97 13.12c-.6.36-.97 1.02-.97 1.74v3.28c0 .72.37 1.38.97 1.74l3 1.83c.63.39 1.43.39 2.06 0l3-1.83c.6-.36.97-1.02.97-1.74v-3.28c0-.72-.37-1.38-.97-1.74l-3-1.83a1.97 1.97 0 0 0-2.06 0l-3 1.83ZM7 17l-4.74-2.85M7 17l4.74-2.85M7 17v5"/></g>`
	fileCheckPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9 15l2 2l4-4"/></g>`
	fileCheckTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 15l2 2l4-4"/></g>`
	fileClockPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h2c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V7.5L14.5 2H6c-.5 0-1 .2-1.4.6C4.2 3 4 3.5 4 4v3"/><path d="M14 2v6h6"/><circle cx="8" cy="16" r="6"/><path d="M9.5 17.5L8 16.25V14"/></g>`
	fileCodePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-10 5l-2 2l2 2m4 0l2-2l-2-2"/></g>`
	fileCodeTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M9 18l3-3l-3-3m-4 0l-3 3l3 3"/></g>`
	fileCogPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="13" r="3"/><path d="m9.7 14.4l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.7.9l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4M7.4 9.3l-.3.9M14 2v6h6"/><path d="M4 5.5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-1.5"/></g>`
	fileCogTwoPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/><circle cx="12" cy="15" r="2"/><path d="M12 12v1m0 4v1m2.6-4.5l-.87.5m-3.46 2l-.87.5m5.2 0l-.87-.5m-3.46-2l-.87-.5"/></g>`
	fileDiffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2zM12 13V7m-3 3h6m-6 7h6"/>`
	fileDigitPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="4" height="6" x="2" y="12" rx="2"/><path d="M14 2v6h6"/><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M10 12h2v6m-2 0h4"/></g>`
	fileDownPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-6m-3 3l3 3l3-3"/></g>`
	fileEditPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 13.5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2h-5.5"/><path d="M14 2v6h6m-9.58 4.61a2.1 2.1 0 1 1 2.97 2.97L7.95 21L4 22l.99-3.95l5.43-5.44Z"/></g>`
	fileHeartPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6m-9.71 2.7a2.43 2.43 0 0 0-2.66-.52c-.29.12-.56.3-.78.53l-.35.34l-.35-.34a2.43 2.43 0 0 0-2.65-.53c-.3.12-.56.3-.79.53c-.95.94-1 2.53.2 3.74L6.5 18l3.6-3.55c1.2-1.21 1.14-2.8.19-3.74Z"/></g>`
	fileImagePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/><circle cx="10" cy="13" r="2"/><path d="m20 17l-1.09-1.09a2 2 0 0 0-2.82 0L10 22"/></g>`
	fileInputPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2 15h10m-3 3l3-3l-3-3"/></g>`
	fileJsonPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-10 4a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1a1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"/></g>`
	fileJsonTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M4 12a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1a1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"/></g>`
	fileKeyPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><circle cx="10" cy="16" r="2"/><path d="m16 10l-4.5 4.5M15 11l1 1"/></g>`
	fileKeyTwoPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6"/><circle cx="4" cy="16" r="2"/><path d="m10 10l-4.5 4.5M9 11l1 1"/></g>`
	fileLineChartPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-4 5l-3.5 3.5l-2-2L8 17"/></g>`
	fileLockPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><rect width="8" height="6" x="8" y="12" rx="1"/><path d="M15 12v-2a3 3 0 1 0-6 0v2"/></g>`
	fileLockTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6"/><rect width="8" height="5" x="2" y="13" rx="1"/><path d="M8 13v-2a2 2 0 1 0-4 0v2"/></g>`
	fileMinusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9 15h6"/></g>`
	fileMinusTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 15h6"/></g>`
	fileOutputPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2 15h10m-7-3l-3 3l3 3"/></g>`
	filePieChartPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h2a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v3"/><path d="M14 2v6h6M4.04 11.71a5.84 5.84 0 1 0 8.2 8.29"/><path d="M13.83 16A5.83 5.83 0 0 0 8 10.17V16h5.83Z"/></g>`
	filePlusPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-6m-3 3h6"/></g>`
	filePlusTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 15h6m-3-3v6"/></g>`
	fileQuestionPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M10 10.3c.2-.4.5-.8.9-1a2.1 2.1 0 0 1 2.6.4c.3.4.5.8.5 1.3c0 1.3-2 2-2 2m0 4h.01"/></g>`
	fileScanPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10V7.5L14.5 2H6a2 2 0 0 0-2 2v16c0 1.1.9 2 2 2h4.5"/><path d="M14 2v6h6m-4 14a2 2 0 0 1-2-2m6 2a2 2 0 0 0 2-2m-2-6a2 2 0 0 1 2 2m-6-2a2 2 0 0 0-2 2"/></g>`
	fileSearchPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v3"/><path d="M14 2v6h6M5 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm4 1l-1.5-1.5"/></g>`
	fileSearchTwoPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/><circle cx="11.5" cy="14.5" r="2.5"/><path d="M13.25 16.25L15 18"/></g>`
	fileSignaturePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 19.5v.5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8.5L18 5.5M8 18h1"/><path d="M18.42 9.61a2.1 2.1 0 1 1 2.97 2.97L16.95 17L13 18l.99-3.95l4.43-4.44Z"/></g>`
	fileSpreadsheetPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M8 13h2m-2 4h2m4-4h2m-2 4h2"/></g>`
	fileStackPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 2v5h5"/><path d="M21 6v6.5c0 .8-.7 1.5-1.5 1.5h-7c-.8 0-1.5-.7-1.5-1.5v-9c0-.8.7-1.5 1.5-1.5H17l4 4z"/><path d="M7 8v8.8c0 .3.2.6.4.8c.2.2.5.4.8.4H15"/><path d="M3 12v8.8c0 .3.2.6.4.8c.2.2.5.4.8.4H11"/></g>`
	fileSymlinkPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v7"/><path d="M14 2v6h6M10 18l3-3l-3-3"/><path d="M4 18v-1a2 2 0 0 1 2-2h6"/></g>`
	fileTerminalPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M8 16l2-2l-2-2m4 6h4"/></g>`
	fileTextPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-4 5H8m8 4H8m2-8H8"/></g>`
	fileTypePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9 13v-1h6v1m-4 5h2m-1-6v6"/></g>`
	fileTypeTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2 13v-1h6v1m-4 5h2m-1-6v6"/></g>`
	fileUpPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 4v6m3-3l-3-3l-3 3"/></g>`
	fileVideoPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-10 3l5 3l-5 3v-6Z"/></g>`
	fileVideoTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 8V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6m-10 7.5l4 2.5v-6l-4 2.5"/><rect width="8" height="6" x="2" y="12" rx="1"/></g>`
	fileVolumePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v3"/><path d="M14 2v6h6M7 10l-3 2H2v4h2l3 2v-8Zm4 1c.64.8 1 1.87 1 3s-.36 2.2-1 3"/></g>`
	fileVolumeTwoPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8.5 5.5c.32.4.5.94.5 1.5s-.18 1.1-.5 1.5M15 12c.64.8 1 1.87 1 3s-.36 2.2-1 3m-7-3h.01"/></g>`
	fileWarningPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2zM12 9v4m0 4h.01"/>`
	fileXPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9.5 12.5l5 5m0-5l-5 5"/></g>`
	fileXTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 12.5l5 5m0-5l-5 5"/></g>`
	filesPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.5 2H8.6c-.4 0-.8.2-1.1.5c-.3.3-.5.7-.5 1.1v12.8c0 .4.2.8.5 1.1c.3.3.7.5 1.1.5h9.8c.4 0 .8-.2 1.1-.5c.3-.3.5-.7.5-1.1V6.5L15.5 2z"/><path d="M3 7.6v12.8c0 .4.2.8.5 1.1c.3.3.7.5 1.1.5h9.8M15 2v5h5"/></g>`
	filmPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 3v18M3 7.5h4M3 12h18M3 16.5h4M17 3v18m0-13.5h4m-4 9h4"/></g>`
	filterPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 3H2l8 9.46V19l4 2v-8.54L22 3z"/>`
	filterXPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.013 3H2l8 9.46V19l4 2v-8.54l.9-1.055M22 3l-5 5m0-5l5 5"/>`
	fingerprintPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12C2 6.5 6.5 2 12 2a10 10 0 0 1 8 4"/><path d="M5 19.5C5.5 18 6 15 6 12c0-.7.12-1.37.34-2m10.95 11.02c.12-.6.43-2.3.5-3.02M12 10a2 2 0 0 0-2 2c0 1.02-.1 2.51-.26 4m-1.09 6c.21-.66.45-1.32.57-2M14 13.12c0 2.38 0 6.38-1 8.88M2 16h.01m19.79 0c.2-2 .131-5.354 0-6M9 6.8a6 6 0 0 1 9 5.2c0 .47 0 1.17-.02 2"/></g>`
	fishPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.5 12c.94-3.46 4.94-6 8.5-6c3.56 0 6.06 2.54 7 6c-.94 3.47-3.44 6-7 6s-7.56-2.53-8.5-6ZM18 12v.5"/><path d="M16 17.93a9.77 9.77 0 0 1 0-11.86m-9 4.6C7 8 5.58 5.97 2.73 5.5c-1 1.5-1 5 .23 6.5c-1.24 1.5-1.24 5-.23 6.5C5.58 18.03 7 16 7 13.33"/><path d="M10.46 7.26C10.2 5.88 9.17 4.24 8 3h5.8a2 2 0 0 1 1.98 1.67l.23 1.4m0 11.86l-.23 1.4A2 2 0 0 1 13.8 21H9.5a5.96 5.96 0 0 0 1.49-3.98"/></g>`
	fishOffPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 12.47v.03m0-.5v.47m-.475 5.056A6.744 6.744 0 0 1 15 18c-3.56 0-7.56-2.53-8.5-6c.348-1.28 1.114-2.433 2.121-3.38m3.444-2.088A8.802 8.802 0 0 1 15 6c3.56 0 6.06 2.54 7 6c-.309 1.14-.786 2.177-1.413 3.058"/><path d="M7 10.67C7 8 5.58 5.97 2.73 5.5c-1 1.5-1 5 .23 6.5c-1.24 1.5-1.24 5-.23 6.5C5.58 18.03 7 16 7 13.33m7.48-4.372A9.77 9.77 0 0 1 16 6.07m0 11.86a9.77 9.77 0 0 1-1.728-3.618"/><path d="m16.01 17.93l-.23 1.4A2 2 0 0 1 13.8 21H9.5a5.96 5.96 0 0 0 1.49-3.98M8.53 3h5.27a2 2 0 0 1 1.98 1.67l.23 1.4M2 2l20 20"/></g>`
	fishSymbolPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 16s9-15 20-4C11 23 2 8 2 8"/>`
	flagPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15s1-1 4-1s5 2 8 2s4-1 4-1V3s-1 1-4 1s-5-2-8-2s-4 1-4 1zm0 7v-7"/>`
	flagOffPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2c3 0 5 2 8 2s4-1 4-1v11M4 22V4m0 11s1-1 4-1s5 2 8 2M2 2l20 20"/>`
	flagTriangleLeftPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 22V2L7 7l10 5"/>`
	flagTriangleRightPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 22V2l10 5l-10 5"/>`
	flamePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3c-1.072-2.143-.224-4.054 2-6c.5 2.5 2 4.9 4 6.5c2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"/>`
	flashlightPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6c0 2-2 2-2 4v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4V2h12zM6 6h12m-6 6h0"/>`
	flashlightOffPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 16v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4m1-4h11v4c0 2-2 2-2 4v1m-5-5h7M2 2l20 20"/>`
	flaskConicalPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2v7.527a2 2 0 0 1-.211.896L4.72 20.55a1 1 0 0 0 .9 1.45h12.76a1 1 0 0 0 .9-1.45l-5.069-10.127A2 2 0 0 1 14 9.527V2M8.5 2h7M7 16h10"/>`
	flaskConicalOffPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10L4.72 20.55a1 1 0 0 0 .9 1.45h12.76a1 1 0 0 0 .9-1.45l-1.272-2.542M10 2v2.343M14 2v6.343M8.5 2h7M7 16h9M2 2l20 20"/>`
	flaskRoundPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2v7.31m4-.01V1.99M8.5 2h7M14 9.3a6.5 6.5 0 1 1-4 0M5.52 16h12.96"/>`
	flipHorizontalPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h3m8-18h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-3m-4-1v2m0-8v2m0-8v2m0-8v2"/>`
	flipHorizontalTwoPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 7l5 5l-5 5V7m18 0l-5 5l5 5V7m-9 13v2m0-8v2m0-8v2m0-8v2"/>`
	flipVerticalPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3m18 8v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3m1-4H2m8 0H8m8 0h-2m8 0h-2"/>`
	flipVerticalTwoPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 3l-5 5l-5-5h10m0 18l-5-5l-5 5h10M4 12H2m8 0H8m8 0h-2m8 0h-2"/>`
	flowerPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 7.5a4.5 4.5 0 1 1 4.5 4.5M12 7.5A4.5 4.5 0 1 0 7.5 12M12 7.5V9m-4.5 3a4.5 4.5 0 1 0 4.5 4.5M7.5 12H9m7.5 0a4.5 4.5 0 1 1-4.5 4.5m4.5-4.5H15m-3 4.5V15"/><circle cx="12" cy="12" r="3"/><path d="m8 16l1.5-1.5m5-5L16 8M8 8l1.5 1.5m5 5L16 16"/></g>`
	flowerTwoPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 5a3 3 0 1 1 3 3m-3-3a3 3 0 1 0-3 3m3-3v1M9 8a3 3 0 1 0 3 3M9 8h1m5 0a3 3 0 1 1-3 3m3-3h-1m-2 3v-1"/><circle cx="12" cy="8" r="2"/><path d="M12 10v12m0 0c4.2 0 7-1.667 7-5c-4.2 0-7 1.667-7 5Zm0 0c-4.2 0-7-1.667-7-5c4.2 0 7 1.667 7 5Z"/></g>`
	focusPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/></g>`
	foldHorizontalPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h6m14 0h-6M12 2v2m0 4v2m0 4v2m0 4v2m7-13l-3 3l3 3M5 15l3-3l-3-3"/>`
	foldVerticalPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22v-6m0-8V2M4 12H2m8 0H8m8 0h-2m8 0h-2m-5 7l-3-3l-3 3m6-14l-3 3l-3-3"/>`
	folderPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/>`
	folderArchivePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 20V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2h6"/><circle cx="16" cy="19" r="2"/><path d="M16 11v-1m0 7v-2"/></g>`
	folderCheckPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><path d="m9 13l2 2l4-4"/></g>`
	folderClockPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 20H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2"/><circle cx="16" cy="16" r="6"/><path d="M16 14v2l1 1"/></g>`
	folderClosedPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2ZM2 10h20"/>`
	folderCogPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><path d="M10.5 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v3.5m-.3 7.9l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`
	folderCogTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="2"/><path d="M12 10v1m0 4v1m2.6-4.5l-.87.5m-3.46 2l-.87.5m5.2 0l-.87-.5m-3.46-2l-.87-.5"/></g>`
	folderDotPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="1"/></g>`
	folderDownPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm8-10v6"/><path d="m15 13l-3 3l-3-3"/></g>`
	folderEditPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.42 10.61a2.1 2.1 0 1 1 2.97 2.97L5.95 19L2 20l.99-3.95l5.43-5.44Z"/><path d="M2 11.5V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-9.5"/></g>`
	folderGitPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="2"/><path d="M14 13h3M7 13h3"/></g>`
	folderGitTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v5"/><circle cx="13" cy="12" r="2"/><path d="M18 19c-2.8 0-5-2.2-5-5v8"/><circle cx="20" cy="19" r="2"/></g>`
	folderHeartPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v1.5"/><path d="M21.29 13.7a2.43 2.43 0 0 0-2.65-.52c-.3.12-.57.3-.8.53l-.34.34l-.35-.34a2.43 2.43 0 0 0-2.65-.53c-.3.12-.56.3-.79.53c-.95.94-1 2.53.2 3.74L17.5 21l3.6-3.55c1.2-1.21 1.14-2.8.19-3.74Z"/></g>`
	folderInputPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1m0-4h10"/><path d="m9 16l3-3l-3-3"/></g>`
	folderKanbanPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm4-10v4m4-4v2m4-2v6"/>`
	folderKeyPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 20H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v2"/><circle cx="16" cy="20" r="2"/><path d="m22 14l-4.5 4.5M21 15l1 1"/></g>`
	folderLockPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 20H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v2.5"/><rect width="8" height="5" x="14" y="17" rx="1"/><path d="M20 17v-2a2 2 0 1 0-4 0v2"/></g>`
	folderMinusPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm5-7h6"/>`
	folderOpenPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 14l1.45-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.55 6a2 2 0 0 1-1.94 1.5H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H18a2 2 0 0 1 2 2v2"/>`
	folderOpenDotPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 14l1.45-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.55 6a2 2 0 0 1-1.94 1.5H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H18a2 2 0 0 1 2 2v2"/><circle cx="14" cy="15" r="1"/></g>`
	folderOutputPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7.5V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2m0-7h10"/><path d="m5 10l-3 3l3 3"/></g>`
	folderPlusPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm8-10v6m-3-3h6"/>`
	folderRootPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="2"/><path d="M12 15v5"/></g>`
	folderSearchPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v4"/><circle cx="17" cy="17" r="3"/><path d="m21 21l-1.5-1.5"/></g>`
	folderSearchTwoPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="11.5" cy="12.5" r="2.5"/><path d="M13.27 14.27L15 16"/></g>`
	folderSymlinkPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2"/><path d="m8 16l3-3l-3-3"/><path d="M2 16v-1a2 2 0 0 1 2-2h6"/></g>`
	folderSyncPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v1"/><path d="M12 10v4h4"/><path d="m12 14l1.5-1.5c.9-.9 2.2-1.5 3.5-1.5s2.6.6 3.5 1.5c.4.4.8 1 1 1.5m.5 8v-4h-4"/><path d="m22 18l-1.5 1.5c-.9.9-2.1 1.5-3.5 1.5s-2.6-.6-3.5-1.5c-.4-.4-.8-1-1-1.5"/></g>`
	folderTreePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 10h7a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-2.5a1 1 0 0 1-.8-.4l-.9-1.2A1 1 0 0 0 15 3h-2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1Zm0 11h7a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-2.88a1 1 0 0 1-.9-.55l-.44-.9a1 1 0 0 0-.9-.55H13a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1ZM3 3v2c0 1.1.9 2 2 2h3"/><path d="M3 3v13c0 1.1.9 2 2 2h3"/></g>`
	folderUpPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm8-10v6"/><path d="m9 13l3-3l3 3"/></g>`
	folderXPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm5.5-9.5l5 5m0-5l-5 5"/>`
	foldersPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 17h12a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3.93a2 2 0 0 1-1.66-.9l-.82-1.2a2 2 0 0 0-1.66-.9H8a2 2 0 0 0-2 2v9c0 1.1.9 2 2 2Z"/><path d="M2 8v11c0 1.1.9 2 2 2h14"/></g>`
	footprintsPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v-2.38C4 11.5 2.97 10.5 3 8c.03-2.72 1.49-6 4.5-6C9.37 2 10 3.8 10 5.5c0 3.11-2 5.66-2 8.68V16a2 2 0 1 1-4 0Zm16 4v-2.38c0-2.12 1.03-3.12 1-5.62c-.03-2.72-1.49-6-4.5-6C14.63 6 14 7.8 14 9.5c0 3.11 2 5.66 2 8.68V20a2 2 0 1 0 4 0Zm-4-3h4M4 13h4"/>`
	forkliftPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12H5a2 2 0 0 0-2 2v5"/><circle cx="13" cy="19" r="2"/><circle cx="5" cy="19" r="2"/><path d="M8 19h3m5-17v17h6M6 12V7c0-1.1.9-2 2-2h3l5 5"/></g>`
	formInputPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="2"/><path d="M12 12h.01M17 12h.01M7 12h.01"/></g>`
	forwardPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 17l5-5l-5-5"/><path d="M4 18v-2a4 4 0 0 1 4-4h12"/></g>`
	framePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 6H2m20 12H2M6 2v20M18 2v20"/>`
	framerPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16V9h14V2H5l14 14h-7m-7 0l7 7v-7m-7 0h7"/>`
	frownPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 16s-1.5-2-4-2s-4 2-4 2m1-7h.01M15 9h.01"/></g>`
	fuelPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 22h12M4 9h10m0 13V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v18m10-9h2a2 2 0 0 1 2 2v2a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2V9.83a2 2 0 0 0-.59-1.42L18 5"/>`
	functionSquarePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 17c2 0 2.8-1 2.8-2.8V10c0-2 1-3.3 3.2-3m-6 4.2h5.7"/></g>`
	galleryHorizontalPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 3v18"/><rect width="12" height="18" x="6" y="3" rx="2"/><path d="M22 3v18"/></g>`
	galleryHorizontalEndPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7v10M6 5v14"/><rect width="12" height="18" x="10" y="3" rx="2"/></g>`
	galleryThumbnailsPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="14" x="3" y="3" rx="2"/><path d="M4 21h1m4 0h1m4 0h1m4 0h1"/></g>`
	galleryVerticalPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 2h18"/><rect width="18" height="12" x="3" y="6" rx="2"/><path d="M3 22h18"/></g>`
	galleryVerticalEndPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 2h10M5 6h14"/><rect width="18" height="12" x="3" y="10" rx="2"/></g>`
	gamepadPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 12h4m-2-2v4m7-1h.01M18 11h.01"/><rect width="20" height="12" x="2" y="6" rx="2"/></g>`
	gamepadTwoPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11h4M8 9v4m7-1h.01M18 10h.01m-.69-5H6.68a4 4 0 0 0-3.978 3.59c-.006.052-.01.101-.017.152C2.604 9.416 2 14.456 2 16a3 3 0 0 0 3 3c1 0 1.5-.5 2-1l1.414-1.414A2 2 0 0 1 9.828 16h4.344a2 2 0 0 1 1.414.586L17 18c.5.5 1 1 2 1a3 3 0 0 0 3-3c0-1.545-.604-6.584-.685-7.258c-.007-.05-.011-.1-.017-.151A4 4 0 0 0 17.32 5z"/>`
	ganttChartPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6h10M6 12h9m-4 6h7"/>`
	ganttChartSquarePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 8h7m-8 4h6m-3 4h5"/></g>`
	gaugePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 14l4-4M3.34 19a10 10 0 1 1 17.32 0"/>`
	gaugeCirclePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.6 2.7a10 10 0 1 0 5.7 5.7"/><circle cx="12" cy="12" r="2"/><path d="M13.4 10.6L19 5"/></g>`
	gavelPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 13l-7.5 7.5c-.83.83-2.17.83-3 0c0 0 0 0 0 0a2.12 2.12 0 0 1 0-3L11 10m5 6l6-6M8 8l6-6M9 7l8 8m4-4l-8-8"/>`
	gemPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3h12l4 6l-10 13L2 9Z"/><path d="M11 3L8 9l4 13l4-13l-3-6M2 9h20"/></g>`
	ghostPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 10h.01M15 10h.01M12 2a8 8 0 0 0-8 8v12l3-3l2.5 2.5L12 19l2.5 2.5L17 19l3 3V10a8 8 0 0 0-8-8z"/>`
	giftPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12v10H4V12M2 7h20v5H2zm10 15V7m0 0H7.5a2.5 2.5 0 0 1 0-5C11 2 12 7 12 7zm0 0h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7z"/>`
	gitBranchPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v12"/><circle cx="18" cy="6" r="3"/><circle cx="6" cy="18" r="3"/><path d="M18 9a9 9 0 0 1-9 9"/></g>`
	gitBranchPlusPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v12m12-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6zM6 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/><path d="M15 6a9 9 0 0 0-9 9m12 0v6m3-3h-6"/></g>`
	gitCommitPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M3 12h6m6 0h6"/></g>`
	gitComparePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7m-7 3H8a2 2 0 0 1-2-2V9"/></g>`
	gitForkPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><circle cx="18" cy="6" r="3"/><path d="M18 9v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V9m6 3v3"/></g>`
	gitMergePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M6 21V9a9 9 0 0 0 9 9"/></g>`
	gitPullRequestPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7M6 9v12"/></g>`
	gitPullRequestClosedPath            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M18 11.5V15m3-12l-6 6m6 0l-6-6M6 9v12"/></g>`
	gitPullRequestDraftPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M18 6V5m0 6v-1M6 9v12"/></g>`
	githubPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5c.08-1.25-.27-2.48-1-3.5c.28-1.15.28-2.35 0-3.5c0 0-1 0-3 1.5c-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5c-.39.49-.68 1.05-.85 1.65c-.17.6-.22 1.23-.15 1.85v4"/><path d="M9 18c-4.51 2-5-2-7-2"/></g>`
	gitlabPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 13.29l-3.33-10a.42.42 0 0 0-.14-.18a.38.38 0 0 0-.22-.11a.39.39 0 0 0-.23.07a.42.42 0 0 0-.14.18l-2.26 6.67H8.32L6.1 3.26a.42.42 0 0 0-.1-.18a.38.38 0 0 0-.26-.08a.39.39 0 0 0-.23.07a.42.42 0 0 0-.14.18L2 13.29a.74.74 0 0 0 .27.83L12 21l9.69-6.88a.71.71 0 0 0 .31-.83Z"/>`
	glassWaterPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.2 22H8.8a2 2 0 0 1-2-1.79L5 3h14l-1.81 17.21A2 2 0 0 1 15.2 22Z"/><path d="M6 12a5 5 0 0 1 6 0a5 5 0 0 0 6 0"/></g>`
	glassesPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="15" r="4"/><circle cx="18" cy="15" r="4"/><path d="M14 15a2 2 0 0 0-2-2a2 2 0 0 0-2 2m-7.5-2L5 7c.7-1.3 1.4-2 3-2m13.5 8L19 7c-.7-1.3-1.5-2-3-2"/></g>`
	globePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10a15.3 15.3 0 0 1-4 10a15.3 15.3 0 0 1-4-10a15.3 15.3 0 0 1 4-10z"/></g>`
	globeTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.54 15H17a2 2 0 0 0-2 2v4.54M7 3.34V5a3 3 0 0 0 3 3v0a2 2 0 0 1 2 2v0c0 1.1.9 2 2 2v0a2 2 0 0 0 2-2v0c0-1.1.9-2 2-2h3.17M11 21.95V18a2 2 0 0 0-2-2v0a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"/><circle cx="12" cy="12" r="10"/></g>`
	goalPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 13V2l8 4l-8 4"/><path d="M20.55 10.23A9 9 0 1 1 8 4.94"/><path d="M8 10a5 5 0 1 0 8.9 2.02"/></g>`
	grabPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 11.5V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1.4m0-.4V8a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v2m0-.1V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5m0 0v0a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0"/><path d="M18 11v0a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8a2 2 0 1 1 4 0"/></g>`
	graduationCapPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10v6M2 10l10-5l10 5l-10 5z"/><path d="M6 12v5c3 3 9 3 12 0v-5"/></g>`
	grapePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 5V2l-5.89 5.89"/><circle cx="16.6" cy="15.89" r="3"/><circle cx="8.11" cy="7.4" r="3"/><circle cx="12.35" cy="11.65" r="3"/><circle cx="13.91" cy="5.85" r="3"/><circle cx="18.15" cy="10.09" r="3"/><circle cx="6.56" cy="13.2" r="3"/><circle cx="10.8" cy="17.44" r="3"/><circle cx="5" cy="19" r="3"/></g>`
	gridPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M3 15h18M9 3v18m6-18v18"/></g>`
	gridThreeXThreePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18M3 15h18M9 3v18m6-18v18"/></g>`
	gridTwoXTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 12h18m-9-9v18"/></g>`
	gripPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="1"/><circle cx="19" cy="5" r="1"/><circle cx="5" cy="5" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/><circle cx="12" cy="19" r="1"/><circle cx="19" cy="19" r="1"/><circle cx="5" cy="19" r="1"/></g>`
	gripHorizontalPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="9" r="1"/><circle cx="19" cy="9" r="1"/><circle cx="5" cy="9" r="1"/><circle cx="12" cy="15" r="1"/><circle cx="19" cy="15" r="1"/><circle cx="5" cy="15" r="1"/></g>`
	gripVerticalPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="12" r="1"/><circle cx="9" cy="5" r="1"/><circle cx="9" cy="19" r="1"/><circle cx="15" cy="12" r="1"/><circle cx="15" cy="5" r="1"/><circle cx="15" cy="19" r="1"/></g>`
	groupPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7V5c0-1.1.9-2 2-2h2m10 0h2c1.1 0 2 .9 2 2v2m0 10v2c0 1.1-.9 2-2 2h-2M7 21H5c-1.1 0-2-.9-2-2v-2"/><rect width="7" height="5" x="7" y="7" rx="1"/><rect width="7" height="5" x="10" y="12" rx="1"/></g>`
	hammerPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 12l-8.5 8.5c-.83.83-2.17.83-3 0c0 0 0 0 0 0a2.12 2.12 0 0 1 0-3L12 9m5.64 6L22 10.64"/><path d="m20.91 11.7l-1.25-1.25c-.6-.6-.93-1.4-.93-2.25v-.86L16.01 4.6a5.56 5.56 0 0 0-3.94-1.64H9l.92.82A6.18 6.18 0 0 1 12 8.4v1.56l2 2h2.47l2.26 1.91"/></g>`
	handPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 11V6a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0m0 4V4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v2m0 4.5V6a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v8"/><path d="M18 8a2 2 0 1 1 4 0v6a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"/></g>`
	handMetalPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 12.5V10a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1.4m0-.4V9a2 2 0 1 0-4 0v2m0-.5V5a2 2 0 1 0-4 0v9"/><path d="m7 15l-1.76-1.76a2 2 0 0 0-2.83 2.82l3.6 3.6C7.5 21.14 9.2 22 12 22h2a8 8 0 0 0 8-8V7a2 2 0 1 0-4 0v5"/></g>`
	hardDrivePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12H2m3.45-6.89L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11zM6 16h.01M10 16h.01"/>`
	hardDriveDownloadPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2v8m4-4l-4 4l-4-4"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h.01M10 18h.01"/></g>`
	hardDriveUploadPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 6l-4-4l-4 4m4-4v8"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h.01M10 18h.01"/></g>`
	hardHatPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v2zm8-8V5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5M4 15v-3a6 6 0 0 1 6-6h0m4 0h0a6 6 0 0 1 6 6v3"/>`
	hashPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9h16M4 15h16M10 3L8 21m8-18l-2 18"/>`
	hazePath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.2 6.2l1.4 1.4M2 13h2m16 0h2m-4.6-5.4l1.4-1.4M22 17H2m20 4H2m14-8a4 4 0 0 0-8 0m4-8V2.5"/>`
	hdmiPortPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1l2 2h12l2-2h1a1 1 0 0 0 1-1ZM7.5 12h9"/>`
	headingPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12h12M6 20V4m12 16V4"/>`
	headingFivePath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5 7v-3h4m-4 7.7c.4.2.8.3 1.3.3c1.5 0 2.7-1.1 2.7-2.5S19.8 13 18.3 13H17"/>`
	headingFourPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5 4v4h4m0-4v8"/>`
	headingOnePath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5 6l3-2v8"/>`
	headingSixPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 12h8m-8 6V6m8 12V6"/><circle cx="19" cy="16" r="2"/><path d="M20 10c-2 2-3 3.5-3 6"/></g>`
	headingThreePath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5.5 4.5c1.7-1 3.5 0 3.5 1.5a2 2 0 0 1-2 2m-2 3.5c2 1.5 4 .3 4-1.5a2 2 0 0 0-2-2"/>`
	headingTwoPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m9 12h-4c0-4 4-3 4-6c0-1.5-2-2.5-4-1"/>`
	headphonesPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3"/>`
	heartPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/>`
	heartCrackPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/><path d="m12 13l-1-1l2-2l-3-3l2-2"/></g>`
	heartHandshakePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/><path d="M12 5L9.04 7.96a2.17 2.17 0 0 0 0 3.08v0c.82.82 2.13.85 3 .07l2.07-1.9a2.82 2.82 0 0 1 3.79 0l2.96 2.66M18 15l-2-2m-1 5l-2-2"/></g>`
	heartOffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20m-5.5-5.5L12 21l-7-7c-1.5-1.45-3-3.2-3-5.5a5.5 5.5 0 0 1 2.14-4.35M8.76 3.1c1.15.22 2.13.78 3.24 1.9c1.5-1.5 2.74-2 4.5-2A5.5 5.5 0 0 1 22 8.5c0 2.12-1.3 3.78-2.67 5.17"/>`
	heartPulsePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/><path d="M3.22 12H9.5l.5-1l2 4.5l2-7l1.5 3.5h5.27"/></g>`
	helpCirclePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3m.08 4h.01"/></g>`
	helpingHandPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15l5.12-5.12A3 3 0 0 1 10.24 9H13a2 2 0 1 1 0 4h-2.5m4-.68l4.17-4.89a1.88 1.88 0 0 1 2.92 2.36l-4.2 5.94A3 3 0 0 1 14.96 17H9.83a2 2 0 0 0-1.42.59L7 19m-5-5l6 6"/>`
	hexagonPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>`
	highlighterPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 11l-6 6v3h9l3-3"/><path d="m22 12l-4.6 4.6a2 2 0 0 1-2.8 0l-5.2-5.2a2 2 0 0 1 0-2.8L14 4"/></g>`
	historyPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5m4-1v5l4 2"/></g>`
	homePath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 9l9-7l9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M9 22V12h6v10"/></g>`
	hopPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 5.5C19 7 20.5 9 21 11c-2.5.5-5 .5-8.5-1m-7 7.5C7 19 9 20.5 11 21c.5-2.5.5-5-1-8.5m6.5-1c1 2 1 3.5 1 6c-2.5 0-4 0-6-1m8.5-5c1 1.5 2 3.5 2 4.5c-1.5.5-3 0-4.5-.5m-6 4.5c1.5 1 3.5 2 4.5 2c.5-1.5 0-3-.5-4.5m5-1c1 2 1.5 3.5 1.5 5.5c-2 0-3.5-.5-5.5-1.5"/><path d="M4.783 4.782C8.493 1.072 14.5 1 18 5c-1 1-4.5 2-6.5 1.5c1 1.5 1 4 .5 5.5c-1.5.5-4 .5-5.5-.5C7 13.5 6 17 5 18c-4-3.5-3.927-9.508-.217-13.218ZM4.5 4.5L3 3c-.184-.185-.184-.816 0-1"/></g>`
	hopOffPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 5.5C19 7 20.5 9 21 11c-1.323.265-2.646.39-4.118.226M5.5 17.5C7 19 9 20.5 11 21c.5-2.5.5-5-1-8.5m7.5 5c-2.5 0-4 0-6-1m8.5-5c1 1.5 2 3.5 2 4.5m-10.5 4c1.5 1 3.5 2 4.5 2c.5-1.5 0-3-.5-4.5M22 22c-2 0-3.5-.5-5.5-1.5"/><path d="M4.783 4.782C1.073 8.492 1 14.5 5 18c1-1 2-4.5 1.5-6.5c1.5 1 4 1 5.5.5M8.227 2.57C11.578 1.335 15.453 2.089 18 5c-.88.88-3.7 1.761-5.726 1.618M2 2l20 20"/></g>`
	hotelPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Z"/><path d="m9 16l.348-.24c1.465-1.013 3.84-1.013 5.304 0L15 16M8 7h.01M16 7h.01M12 7h.01M12 11h.01M16 11h.01M8 11h.01M10 22v-6.5m4 0V22"/></g>`
	hourglassPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 22h14M5 2h14m-2 20v-4.172a2 2 0 0 0-.586-1.414L12 12l-4.414 4.414A2 2 0 0 0 7 17.828V22M7 2v4.172a2 2 0 0 0 .586 1.414L12 12l4.414-4.414A2 2 0 0 0 17 6.172V2"/>`
	iceCreamPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 11l4.08 10.35a1 1 0 0 0 1.84 0L17 11m0-4A5 5 0 0 0 7 7m10 0a2 2 0 0 1 0 4H7a2 2 0 0 1 0-4"/>`
	iceCreamTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 17c5 0 8-2.69 8-6H4c0 3.31 3 6 8 6Zm-4 4h8m-4-3v3M5.14 11a3.5 3.5 0 1 1 6.71 0"/><path d="M12.14 11a3.5 3.5 0 1 1 6.71 0"/><path d="M15.5 6.5a3.5 3.5 0 1 0-7 0"/></g>`
	imagePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><circle cx="9" cy="9" r="2"/><path d="m21 15l-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></g>`
	imageMinusPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7m4 2h6"/><circle cx="9" cy="9" r="2"/><path d="m21 15l-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></g>`
	imageOffPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20M10.41 10.41a2 2 0 1 1-2.83-2.83m5.92 5.92L6 21m12-9l3 3"/><path d="M3.59 3.59A1.99 1.99 0 0 0 3 5v14a2 2 0 0 0 2 2h14c.55 0 1.052-.22 1.41-.59M21 15V5a2 2 0 0 0-2-2H9"/></g>`
	imagePlusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7m4 2h6m-3-3v6"/><circle cx="9" cy="9" r="2"/><path d="m21 15l-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></g>`
	importPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3v12m-4-4l4 4l4-4"/><path d="M8 5H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-4"/></g>`
	inboxPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12h-6l-2 3h-4l-2-3H2"/><path d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"/></g>`
	indentPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4 4l-4 4m18-4H11m10-6H11m10 12H11"/>`
	indianRupeePath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3h12M6 8h12M6 13l8.5 8M6 13h3m0 0c6.667 0 6.667-10 0-10"/>`
	infinityPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12c-2-2.67-4-4-6-4a4 4 0 1 0 0 8c2 0 4-1.33 6-4Zm0 0c2 2.67 4 4 6 4a4 4 0 0 0 0-8c-2 0-4 1.33-6 4Z"/>`
	infoPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4m0-4h.01"/></g>`
	instagramPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="5" ry="5"/><path d="M16 11.37A4 4 0 1 1 12.63 8A4 4 0 0 1 16 11.37zm1.5-4.87h.01"/></g>`
	italicPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4h-9m4 16H5M15 4L9 20"/>`
	iterationCcwPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10c0-4.4-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8h8"/><path d="m16 14l4 4l-4 4"/></g>`
	iterationCwPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10c0-4.4 3.6-8 8-8s8 3.6 8 8s-3.6 8-8 8H4"/><path d="m8 22l-4-4l4-4"/></g>`
	japaneseYenPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9.5V21m0-11.5L6 3m6 6.5L18 3M6 15h12M6 11h12"/>`
	joystickPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2ZM6 15v-2m6 2V9"/><circle cx="12" cy="6" r="3"/></g>`
	kanbanPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 5v11m6-11v6m6-6v14"/>`
	kanbanSquarePath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 7v7m4-7v4m4-4v9"/></g>`
	kanbanSquareDashedPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v7m4-7v4m4-4v9M5 3a2 2 0 0 0-2 2m6-2h1m4 0h1m4 0a2 2 0 0 1 2 2m0 4v1m0 4v1m0 4a2 2 0 0 1-2 2m-5 0h1m-6 0h1m-5 0a2 2 0 0 1-2-2m0-5v1m0-6v1"/>`
	keyPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7.5" cy="15.5" r="5.5"/><path d="m21 2l-9.6 9.6m4.1-4.1l3 3L22 7l-3-3"/></g>`
	keyRoundPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 18v3c0 .6.4 1 1 1h4v-3h3v-3h2l1.4-1.4a6.5 6.5 0 1 0-4-4Z"/><circle cx="16.5" cy="7.5" r=".5"/></g>`
	keySquarePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.4 2.7c.9-.9 2.5-.9 3.4 0l5.5 5.5c.9.9.9 2.5 0 3.4l-3.7 3.7c-.9.9-2.5.9-3.4 0L8.7 9.8c-.9-.9-.9-2.5 0-3.4ZM14 7l3 3m-7.6.6L2 18v3c0 .6.4 1 1 1h4v-3h3v-3h2l1.4-1.4"/>`
	keyboardPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2" ry="2"/><path d="M6 8h.001M10 8h.001M14 8h.001M18 8h.001M8 12h.001M12 12h.001M16 12h.001M7 16h10"/></g>`
	lampPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2h8l4 10H4L8 2Zm4 10v6m-4 4v-2c0-1.1.9-2 2-2h4a2 2 0 0 1 2 2v2H8Z"/>`
	lampCeilingPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v5M6 7h12l4 9H2l4-9Zm3.17 9a3 3 0 1 0 5.66 0"/>`
	lampDeskPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 5l-3 3l2 7l8-8l-7-2Z"/><path d="m14 5l-3 3l-3-3l3-3l3 3Z"/><path d="M9.5 6.5L4 12l3 6m-4 4v-2c0-1.1.9-2 2-2h4a2 2 0 0 1 2 2v2H3Z"/></g>`
	lampFloorPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 2h6l3 7H6l3-7Zm3 7v13m-3 0h6"/>`
	lampWallDownPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 13h6l3 7H8l3-7Zm3 0V8a2 2 0 0 0-2-2H8M4 9h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4v6Z"/>`
	lampWallUpPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4h6l3 7H8l3-7Zm3 7v5a2 2 0 0 1-2 2H8m-4-3h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H4v-6Z"/>`
	landmarkPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 22h18M6 18v-7m4 7v-7m4 7v-7m4 7v-7m-6-9l8 5H4z"/>`
	languagesPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 8l6 6m-7 0l6-6l2-3M2 5h12M7 2h1m14 20l-5-10l-5 10m2-4h6"/>`
	laptopPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 16V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v9m16 0H4m16 0l1.28 2.55a1 1 0 0 1-.9 1.45H3.62a1 1 0 0 1-.9-1.45L4 16"/>`
	laptopTwoPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="4" rx="2" ry="2"/><path d="M2 20h20"/></g>`
	lassoPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 22a5 5 0 0 1-2-4m-1.7-4A6.8 6.8 0 0 1 2 10c0-4.4 4.5-8 10-8s10 3.6 10 8s-4.5 8-10 8a12 12 0 0 1-5-1"/><path d="M5 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4z"/></g>`
	lassoSelectPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 22a5 5 0 0 1-2-4m2-1.07c.96.43 1.96.74 2.99.91M3.34 14A6.8 6.8 0 0 1 2 10c0-4.42 4.48-8 10-8s10 3.58 10 8a7.19 7.19 0 0 1-.33 2"/><path d="M5 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm9.33 4h-.09a.35.35 0 0 1-.24-.32v-10a.34.34 0 0 1 .33-.34c.08 0 .15.03.21.08l7.34 6a.33.33 0 0 1-.21.59h-4.49l-2.57 3.85a.35.35 0 0 1-.28.14v0z"/></g>`
	laughPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M18 13a6 6 0 0 1-6 5a6 6 0 0 1-6-5h12ZM9 9h.01M15 9h.01"/></g>`
	layersPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2L2 7l10 5l10-5l-10-5zM2 17l10 5l10-5M2 12l10 5l10-5"/>`
	layoutPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M9 21V9"/></g>`
	layoutDashboardPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="9" x="3" y="3" rx="1"/><rect width="7" height="5" x="14" y="3" rx="1"/><rect width="7" height="9" x="14" y="12" rx="1"/><rect width="7" height="5" x="3" y="16" rx="1"/></g>`
	layoutGridPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/></g>`
	layoutListPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><path d="M14 4h7m-7 5h7m-7 6h7m-7 5h7"/></g>`
	layoutPanelLeftPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="18" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/></g>`
	layoutPanelTopPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/></g>`
	layoutTemplatePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="7" x="3" y="3" rx="1"/><rect width="9" height="7" x="3" y="14" rx="1"/><rect width="5" height="7" x="16" y="14" rx="1"/></g>`
	leafPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20A7 7 0 0 1 9.8 6.1C15.5 5 17 4.48 19 2c1 2 2 4.18 2 8c0 5.5-4.78 10-10 10Z"/><path d="M2 21c0-3 1.85-5.36 5.08-6C9.5 14.52 12 13 13 12"/></g>`
	leafyGreenPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22c1.25-.987 2.27-1.975 3.9-2.2a5.56 5.56 0 0 1 3.8 1.5a4 4 0 0 0 6.187-2.353a3.5 3.5 0 0 0 3.69-5.116A3.5 3.5 0 0 0 20.95 8A3.5 3.5 0 1 0 16 3.05a3.5 3.5 0 0 0-5.831 1.373a3.5 3.5 0 0 0-5.116 3.69a4 4 0 0 0-2.348 6.155C3.499 15.42 4.409 16.712 4.2 18.1C3.926 19.743 3.014 20.732 2 22m0 0L17 7"/>`
	libraryPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 6l4 14M12 6v14M8 8v12M4 4v16"/>`
	lifeBuoyPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m4.93 4.93l4.24 4.24m5.66 0l4.24-4.24m-4.24 9.9l4.24 4.24m-9.9-4.24l-4.24 4.24"/><circle cx="12" cy="12" r="4"/></g>`
	ligaturePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20V8c0-2.2 1.8-4 4-4c1.5 0 2.8.8 3.5 2M6 12h4m4 0h2v8M6 20h4m4 0h4"/>`
	lightbulbPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 14c.2-1 .7-1.7 1.5-2.5c1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5c.7.7 1.3 1.5 1.5 2.5m0 4h6m-5 4h4"/>`
	lightbulbOffPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.8 11.2c.8-.9 1.2-2 1.2-3.2a6 6 0 0 0-9.3-5M2 2l20 20M6.3 6.3a4.67 4.67 0 0 0 1.2 5.2c.7.7 1.3 1.5 1.5 2.5m0 4h6m-5 4h4"/>`
	lineChartPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><path d="m19 9l-5 5l-4-4l-3 3"/></g>`
	linkPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></g>`
	linkTwoPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17H7A5 5 0 0 1 7 7h2m6 0h2a5 5 0 1 1 0 10h-2m-7-5h8"/>`
	linkTwoOffPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17H7A5 5 0 0 1 7 7m8 0h2a5 5 0 0 1 4 8M8 12h4M2 2l20 20"/>`
	linkedinPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2a2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6zM2 9h4v12H2z"/><circle cx="4" cy="4" r="2"/></g>`
	listPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6h13M8 12h13M8 18h13M3 6h.01M3 12h.01M3 18h.01"/>`
	listChecksPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 17l2 2l4-4M3 7l2 2l4-4m4 1h8m-8 6h8m-8 6h8"/>`
	listEndPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 12H3m13-6H3m7 12H3M21 6v10a2 2 0 0 1-2 2h-5"/><path d="m16 16l-2 2l2 2"/></g>`
	listFilterPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18M7 12h10m-7 6h4"/>`
	listMinusPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 12H3m13-6H3m13 12H3m18-6h-6"/>`
	listMusicPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15V6m-2.5 12a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM12 12H3m13-6H3m9 12H3"/>`
	listOrderedPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6h11m-11 6h11m-11 6h11M4 6h1v4m-1 0h2m0 8H4c0-1 2-2 2-3s-1-1.5-2-1"/>`
	listPlusPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 12H3m13-6H3m13 12H3m15-9v6m3-3h-6"/>`
	listRestartPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 6H3m4 6H3m4 6H3m9 0a5 5 0 0 0 9-3a4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L11 14"/><path d="M11 10v4h4"/></g>`
	listStartPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 12H3m13 6H3m7-12H3m18 12V8a2 2 0 0 0-2-2h-5"/><path d="m16 8l-2-2l2-2"/></g>`
	listTodoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="6" x="3" y="5" rx="1"/><path d="m3 17l2 2l4-4m4-9h8m-8 6h8m-8 6h8"/></g>`
	listTreePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12h-8m8-6H8m13 12h-8M3 6v4c0 1.1.9 2 2 2h3"/><path d="M3 10v6c0 1.1.9 2 2 2h3"/></g>`
	listVideoPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12H3m13-6H3m9 12H3m13-6l5 3l-5 3v-6Z"/>`
	listXPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 12H3m13-6H3m13 12H3m16-8l-4 4m0-4l4 4"/>`
	loaderPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v4m0 12v4M4.93 4.93l2.83 2.83m8.48 8.48l2.83 2.83M2 12h4m12 0h4M4.93 19.07l2.83-2.83m8.48-8.48l2.83-2.83"/>`
	loaderTwoPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 1 1-6.219-8.56"/>`
	locatePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12h3m14 0h3M12 2v3m0 14v3"/><circle cx="12" cy="12" r="7"/></g>`
	locateFixedPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12h3m14 0h3M12 2v3m0 14v3"/><circle cx="12" cy="12" r="7"/><circle cx="12" cy="12" r="3"/></g>`
	locateOffPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h3m14 0h3M12 2v3m0 14v3M7.11 7.11C5.83 8.39 5 10.1 5 12c0 3.87 3.13 7 7 7c1.9 0 3.61-.83 4.89-2.11m1.82-2.93c.19-.63.29-1.29.29-1.96c0-3.87-3.13-7-7-7c-.67 0-1.33.1-1.96.29M2 2l20 20"/>`
	lockPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></g>`
	logInPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4m-5-4l5-5l-5-5m5 5H3"/>`
	logOutPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4m7 14l5-5l-5-5m5 5H9"/>`
	lollipopPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3M11 11a2 2 0 0 0 4 0a4 4 0 0 0-8 0a6 6 0 0 0 12 0"/></g>`
	luggagePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 20h0a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h0"/><path d="M8 18V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14m-6 2h4"/><circle cx="16" cy="20" r="2"/><circle cx="8" cy="20" r="2"/></g>`
	mSquarePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 16V8l4 4l4-4v8"/></g>`
	magnetPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 15l-4-4l6.75-6.77a7.79 7.79 0 0 1 11 11L13 22l-4-4l6.39-6.36a2.14 2.14 0 0 0-3-3L6 15M5 8l4 4m3 3l4 4"/>`
	mailPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></g>`
	mailCheckPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m14 12l2 2l4-4"/></g>`
	mailMinusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 15V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m14 12h6"/></g>`
	mailOpenPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.2 8.4c.5.38.8.97.8 1.6v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V10a2 2 0 0 1 .8-1.6l8-6a2 2 0 0 1 2.4 0l8 6Z"/><path d="m22 10l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 10"/></g>`
	mailPlusPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m17 9v6m-3-3h6"/></g>`
	mailQuestionPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m16 8.28c.2-.4.5-.8.9-1a2.1 2.1 0 0 1 2.6.4c.3.4.5.8.5 1.3c0 1.3-2 2-2 2M20 22v.01"/></g>`
	mailSearchPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h7.5"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m16 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6v0Z"/><circle cx="18" cy="18" r="3"/><path d="m22 22l-1.5-1.5"/></g>`
	mailWarningPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m18 7v4m0 4v.01"/></g>`
	mailXPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h9"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m15 10l4 4m0-4l-4 4"/></g>`
	mailboxPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.5C2 7 4 5 6.5 5H18c2.2 0 4 1.8 4 4v8Z"/><path d="M15 9h3v2M6.5 5C9 5 11 7 11 9.5V17a2 2 0 0 1-2 2v0m-3-9h1"/></g>`
	mailsPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="13" x="6" y="4" rx="2"/><path d="m22 7l-7.1 3.78c-.57.3-1.23.3-1.8 0L6 7M2 8v11c0 1.1.9 2 2 2h14"/></g>`
	mapPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 6l6-3l6 3l6-3v15l-6 3l-6-3l-6 3zm6-3v15m6-12v15"/>`
	mapPinPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></g>`
	mapPinOffPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.43 5.43A8.06 8.06 0 0 0 4 10c0 6 8 12 8 12a29.94 29.94 0 0 0 5-5m2.18-3.48A8.66 8.66 0 0 0 20 10a8 8 0 0 0-8-8a7.88 7.88 0 0 0-3.52.82"/><path d="M9.13 9.13A2.78 2.78 0 0 0 9 10a3 3 0 0 0 3 3a2.78 2.78 0 0 0 .87-.13m2.03-3.62a3 3 0 0 0-2.15-2.16M2 2l20 20"/></g>`
	martiniPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22h8m-4-11v11m7-19l-7 8l-7-8Z"/>`
	maximizePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H5a2 2 0 0 0-2 2v3m18 0V5a2 2 0 0 0-2-2h-3M3 16v3a2 2 0 0 0 2 2h3m8 0h3a2 2 0 0 0 2-2v-3"/>`
	maximizeTwoPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h6v6M9 21H3v-6M21 3l-7 7M3 21l7-7"/>`
	medalPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.21 15L2.66 7.14a2 2 0 0 1 .13-2.2L4.4 2.8A2 2 0 0 1 6 2h12a2 2 0 0 1 1.6.8l1.6 2.14a2 2 0 0 1 .14 2.2L16.79 15M11 12L5.12 2.2M13 12l5.88-9.8M8 7h8"/><circle cx="12" cy="17" r="5"/><path d="M12 18v-2h-.5"/></g>`
	megaphonePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l18-5v12L3 14v-3zm8.6 5.8a3 3 0 1 1-5.8-1.6"/>`
	megaphoneOffPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.26 9.26L3 11v3l14.14 3.14m3.86-1.8V6l-7.31 2.03M11.6 16.8a3 3 0 1 1-5.8-1.6M2 2l20 20"/>`
	mehPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 15h8M9 9h.01M15 9h.01"/></g>`
	memoryStickPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19v-3m4 3v-3m4 3v-3m4 3v-3M8 11V9m8 2V9m-4 2V9M2 15h20M2 7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1.1a2 2 0 0 0 0 3.837V17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5.1a2 2 0 0 0 0-3.837Z"/>`
	menuPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h16M4 6h16M4 18h16"/>`
	menuSquarePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 8h10M7 12h10M7 16h10"/></g>`
	mergePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 6l4-4l4 4"/><path d="M12 2v10.3a4 4 0 0 1-1.172 2.872L4 22m16 0l-5-5"/></g>`
	messageCirclePath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 21l1.9-5.7a8.5 8.5 0 1 1 3.8 3.8z"/>`
	messageSquarePath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>`
	messageSquareDashedPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6V5c0-1.1.9-2 2-2h2m4 0h3m4 0h1c1.1 0 2 .9 2 2m0 4v2m0 4c0 1.1-.9 2-2 2h-1m-4 0h-3m-4 0l-4 4v-5m0-4v-2"/>`
	messageSquarePlusPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2zM9 10h6m-3-3v6"/>`
	messagesSquarePath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 9a2 2 0 0 1-2 2H6l-4 4V4c0-1.1.9-2 2-2h8a2 2 0 0 1 2 2v5Zm4 0h2a2 2 0 0 1 2 2v11l-4-4h-6a2 2 0 0 1-2-2v-1"/>`
	micPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2a3 3 0 0 0-3 3v7a3 3 0 0 0 6 0V5a3 3 0 0 0-3-3Z"/><path d="M19 10v2a7 7 0 0 1-14 0v-2m7 9v3"/></g>`
	micOffPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20m-3.11-8.77A7.12 7.12 0 0 0 19 12v-2M5 10v2a7 7 0 0 0 12 5m-2-7.66V5a3 3 0 0 0-5.68-1.33"/><path d="M9 9v3a3 3 0 0 0 5.12 2.12M12 19v3"/></g>`
	micTwoPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 8l-9.04 9.06a2.82 2.82 0 1 0 3.98 3.98L16 12"/><circle cx="17" cy="7" r="5"/></g>`
	microscopePath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18h8M3 22h18m-7 0a7 7 0 1 0 0-14h-1m-4 6h2m-2-2a2 2 0 0 1-2-2V6h6v4a2 2 0 0 1-2 2Zm3-6V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"/>`
	microwavePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="15" x="2" y="4" rx="2"/><rect width="8" height="7" x="6" y="8" rx="1"/><path d="M18 8v7M6 19v2m12-2v2"/></g>`
	milestonePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6H5a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h13l4-3.5L18 6Zm-6 7v8m0-18v3"/>`
	milkPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2h8M9 2v2.789a4 4 0 0 1-.672 2.219l-.656.984A4 4 0 0 0 7 10.212V20a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-9.789a4 4 0 0 0-.672-2.219l-.656-.984A4 4 0 0 1 15 4.788V2"/><path d="M7 15a6.472 6.472 0 0 1 5 0a6.47 6.47 0 0 0 5 0"/></g>`
	milkOffPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2h8M9 2v1.343M15 2v2.789a4 4 0 0 0 .672 2.219l.656.984a4 4 0 0 1 .672 2.22v1.131M7.8 7.8l-.128.192A4 4 0 0 0 7 10.212V20a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-3"/><path d="M7 15a6.47 6.47 0 0 1 5 0a6.472 6.472 0 0 0 3.435.435M2 2l20 20"/></g>`
	minimizePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v3a2 2 0 0 1-2 2H3m18 0h-3a2 2 0 0 1-2-2V3M3 16h3a2 2 0 0 1 2 2v3m8 0v-3a2 2 0 0 1 2-2h3"/>`
	minimizeTwoPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14h6v6m10-10h-6V4m0 6l7-7M3 21l7-7"/>`
	minusPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>`
	minusCirclePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8"/></g>`
	minusSquarePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8"/></g>`
	monitorPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M8 21h8m-4-4v4"/></g>`
	monitorCheckPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l2 2l4-4"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	monitorDotPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="6" r="3"/><path d="M22 12v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h9m-1 14v4m-4 0h8"/></g>`
	monitorDownPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 13V7m3 3l-3 3l-3-3"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	monitorOffPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17H4a2 2 0 0 1-2-2V5c0-1.5 1-2 1-2m19 12V5a2 2 0 0 0-2-2H9M8 21h8m-4-4v4M2 2l20 20"/>`
	monitorPausePath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 13V7m4 6V7"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	monitorPlayPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 7l5 3l-5 3Z"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	monitorSmartphonePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 8V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8m-2 4v-3.96v3.15M7 19h5"/><rect width="6" height="10" x="16" y="12" rx="2"/></g>`
	monitorSpeakerPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.5 20H8m9-11h.01"/><rect width="10" height="16" x="12" y="4" rx="2"/><path d="M8 6H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4"/><circle cx="17" cy="15" r="1"/></g>`
	monitorStopPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 7h6v6H9z"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	monitorUpPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l3-3l3 3m-3 3V7"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	monitorXPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14.5 12.5l-5-5m0 5l5-5"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`
	moonPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a6 6 0 0 0 9 9a9 9 0 1 1-9-9Z"/>`
	moonStarPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a6 6 0 0 0 9 9a9 9 0 1 1-9-9Zm7 0v4m2-2h-4"/>`
	moreHorizontalPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></g>`
	moreVerticalPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="12" cy="5" r="1"/><circle cx="12" cy="19" r="1"/></g>`
	mountainPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 3l4 8l5-5l5 15H2L8 3z"/>`
	mountainSnowPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 3l4 8l5-5l5 15H2L8 3z"/><path d="M4.14 15.08c2.62-1.57 5.24-1.43 7.86.42c2.74 1.94 5.49 2 8.23.19"/></g>`
	mousePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="7"/><path d="M12 6v4"/></g>`
	mousePointerPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l7.07 16.97l2.51-7.39l7.39-2.51L3 3zm10 10l6 6"/>`
	mousePointerClickPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 9l5 12l1.774-5.226L21 14L9 9zm7.071 7.071l4.243 4.243M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122"/>`
	mousePointerSquarePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"/><path d="m12 12l4 10l1.7-4.3L22 16Z"/></g>`
	mousePointerSquareDashedPath        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3a2 2 0 0 0-2 2m16-2a2 2 0 0 1 2 2m-9 7l4 10l1.7-4.3L22 16Zm-7 9a2 2 0 0 1-2-2M9 3h1M9 21h2m3-18h1M3 9v1m18-1v2M3 14v1"/>`
	mousePointerTwoPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l7.07 17l2.51-7.39L21 11.07z"/>`
	movePath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 9l-3 3l3 3M9 5l3-3l3 3m0 14l-3 3l-3-3M19 9l3 3l-3 3M2 12h20M12 2v20"/>`
	moveDiagonalPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5h6v6m-8 8H5v-6m14-8L5 19"/>`
	moveDiagonalTwoPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11V5h6m8 8v6h-6M5 5l14 14"/>`
	moveDownPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 18l4 4l4-4M12 2v20"/>`
	moveDownLeftPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 19H5v-6m14-8L5 19"/>`
	moveDownRightPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 13v6h-6M5 5l14 14"/>`
	moveHorizontalPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 8l4 4l-4 4M6 8l-4 4l4 4m-4-4h20"/>`
	moveLeftPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 8l-4 4l4 4m-4-4h20"/>`
	moveRightPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 8l4 4l-4 4M2 12h20"/>`
	moveThreeDPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 3v16h16M5 19l6-6"/><path d="m2 6l3-3l3 3m10 10l3 3l-3 3"/></g>`
	moveUpPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 6l4-4l4 4m-4-4v20"/>`
	moveUpLeftPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11V5h6M5 5l14 14"/>`
	moveUpRightPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5h6v6m0-6L5 19"/>`
	moveVerticalPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 18l4 4l4-4M8 6l4-4l4 4m-4-4v20"/>`
	musicPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 18V5l12-2v13"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="16" r="3"/></g>`
	musicFourPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 18V5l12-2v13M9 9l12-2"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="16" r="3"/></g>`
	musicThreePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="18" r="4"/><path d="M16 18V2"/></g>`
	musicTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="18" r="4"/><path d="M12 18V2l7 4"/></g>`
	navigationPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l19-9l-9 19l-2-8l-8-2z"/>`
	navigationOffPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.43 8.43L3 11l8 2l2 8l2.57-5.43m1.82-3.84L22 2l-9.73 4.61M2 2l20 20"/>`
	navigationTwoPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l7 19l-7-4l-7 4l7-19z"/>`
	navigationTwoOffPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.31 9.31L5 21l7-4l7 4l-1.17-3.17m-3.3-8.95L12 2l-1.17 3.17M2 2l20 20"/>`
	networkPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="6" x="16" y="16" rx="1"/><rect width="6" height="6" x="2" y="16" rx="1"/><rect width="6" height="6" x="9" y="2" rx="1"/><path d="M5 16v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3m-7-4V8"/></g>`
	newspaperPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-2 2Zm0 0a2 2 0 0 1-2-2v-9c0-1.1.9-2 2-2h2m12 5h-8m5 4h-5"/><path d="M10 6h8v4h-8V6Z"/></g>`
	nfcPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8.32a7.43 7.43 0 0 1 0 7.36m3.46-9.47a11.76 11.76 0 0 1 0 11.58M12.91 4.1a15.91 15.91 0 0 1 .01 15.8M16.37 2a20.16 20.16 0 0 1 0 20"/>`
	nutPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4V2m-7 8v4a7.004 7.004 0 0 0 5.277 6.787c.412.104.802.292 1.102.592L12 22l.621-.621c.3-.3.69-.488 1.102-.592A7.003 7.003 0 0 0 19 14v-4"/><path d="M12 4C8 4 4.5 6 4 8c-.243.97-.919 1.952-2 3c1.31-.082 1.972-.29 3-1c.54.92.982 1.356 2 2c1.452-.647 1.954-1.098 2.5-2c.595.995 1.151 1.427 2.5 2c1.31-.621 1.862-1.058 2.5-2c.629.977 1.162 1.423 2.5 2c1.209-.548 1.68-.967 2-2c1.032.916 1.683 1.157 3 1c-1.297-1.036-1.758-2.03-2-3c-.5-2-4-4-8-4Z"/></g>`
	nutOffPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4V2m-7 8v4a7.004 7.004 0 0 0 5.277 6.787c.412.104.802.292 1.102.592L12 22l.621-.621c.3-.3.69-.488 1.102-.592a7.01 7.01 0 0 0 4.125-2.939M19 10v3.343"/><path d="M12 12c-1.349-.573-1.905-1.005-2.5-2c-.546.902-1.048 1.353-2.5 2c-1.018-.644-1.46-1.08-2-2c-1.028.71-1.69.918-3 1c1.081-1.048 1.757-2.03 2-3c.194-.776.84-1.551 1.79-2.21m11.654 5.997c.887-.457 1.28-.891 1.556-1.787c1.032.916 1.683 1.157 3 1c-1.297-1.036-1.758-2.03-2-3c-.5-2-4-4-8-4c-.74 0-1.461.068-2.15.192M2 2l20 20"/></g>`
	octagonPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2z"/>`
	optionPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h6l6 18h6M14 3h7"/>`
	orbitPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><path d="M10.4 21.9a10 10 0 0 0 9.941-15.416M13.5 2.1a10 10 0 0 0-9.841 15.416"/></g>`
	outdentPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 8l-4 4l4 4m14-4H11m10-6H11m10 12H11"/>`
	packagePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7.5 4.27l9 5.15M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"/><path d="m3.3 7l8.7 5l8.7-5M12 22V12"/></g>`
	packageCheckPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 16l2 2l4-4"/><path d="M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/></g>`
	packageMinusPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 16h6m-1-6V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/></g>`
	packageOpenPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.91 8.84L8.56 2.23a1.93 1.93 0 0 0-1.81 0L3.1 4.13a2.12 2.12 0 0 0-.05 3.69l12.22 6.93a2 2 0 0 0 1.94 0L21 12.51a2.12 2.12 0 0 0-.09-3.67Z"/><path d="m3.09 8.84l12.35-6.61a1.93 1.93 0 0 1 1.81 0l3.65 1.9a2.12 2.12 0 0 1 .1 3.69L8.73 14.75a2 2 0 0 1-1.94 0L3 12.51a2.12 2.12 0 0 1 .09-3.67ZM12 22v-9"/><path d="M20 13.5v3.37a2.06 2.06 0 0 1-1.11 1.83l-6 3.08a1.93 1.93 0 0 1-1.78 0l-6-3.08A2.06 2.06 0 0 1 4 16.87V13.5"/></g>`
	packagePlusPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 16h6m-3-3v6m2-9V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/></g>`
	packageSearchPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/><circle cx="18.5" cy="15.5" r="2.5"/><path d="M20.27 17.27L22 19"/></g>`
	packageTwoPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Zm0 0l2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9m-9-6v6"/>`
	packageXPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12m5 1l5 5m-5 0l5-5"/></g>`
	paintBucketPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 11l-8-8l-8.6 8.6a2 2 0 0 0 0 2.8l5.2 5.2c.8.8 2 .8 2.8 0L19 11ZM5 2l5 5m-8 6h15m5 7a2 2 0 1 1-4 0c0-1.6 1.7-2.4 2-4c.3 1.6 2 2.4 2 4Z"/>`
	paintbrushPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.37 2.63L14 7l-1.59-1.59a2 2 0 0 0-2.82 0L8 7l9 9l1.59-1.59a2 2 0 0 0 0-2.82L17 10l4.37-4.37a2.12 2.12 0 1 0-3-3Z"/><path d="M9 8c-2 3-4 3.5-7 4l8 10c2-1 6-5 6-7m-1.5 2.5L4.5 15"/></g>`
	paintbrushTwoPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 19.9V16h3a2 2 0 0 0 2-2v-2H5v2c0 1.1.9 2 2 2h3v3.9a2 2 0 1 0 4 0ZM6 12V2h12v10M14 2v4m-4-4v2"/>`
	palettePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="13.5" cy="6.5" r=".5"/><circle cx="17.5" cy="10.5" r=".5"/><circle cx="8.5" cy="7.5" r=".5"/><circle cx="6.5" cy="12.5" r=".5"/><path d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c.926 0 1.648-.746 1.648-1.688c0-.437-.18-.835-.437-1.125c-.29-.289-.438-.652-.438-1.125a1.64 1.64 0 0 1 1.668-1.668h1.996c3.051 0 5.555-2.503 5.555-5.554C21.965 6.012 17.461 2 12 2z"/></g>`
	palmtreePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 8c0-2.76-2.46-5-5.5-5S2 5.24 2 8h2l1-1l1 1h4m3-.86A5.82 5.82 0 0 1 16.5 6c3.04 0 5.5 2.24 5.5 5h-3l-1-1l-1 1h-3"/><path d="M5.89 9.71c-2.15 2.15-2.3 5.47-.35 7.43l4.24-4.25l.7-.7l.71-.71l2.12-2.12c-1.95-1.96-5.27-1.8-7.42.35z"/><path d="M11 15.5c.5 2.5-.17 4.5-1 6.5h4c2-5.5-.5-12-1-14"/></g>`
	panelBottomPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 15h18"/></g>`
	panelBottomClosePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 15h18m-6-7l-3 3l-3-3"/></g>`
	panelBottomInactivePath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M14 15h1m4 0h2M3 15h2m4 0h1"/></g>`
	panelBottomOpenPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 15h18M9 10l3-3l3 3"/></g>`
	panelLeftPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18"/></g>`
	panelLeftClosePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18m7-6l-3-3l3-3"/></g>`
	panelLeftInactivePath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 14v1m0 4v2M9 3v2m0 4v1"/></g>`
	panelLeftOpenPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18m5-12l3 3l-3 3"/></g>`
	panelRightPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 3v18"/></g>`
	panelRightClosePath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 3v18M8 9l3 3l-3 3"/></g>`
	panelRightInactivePath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 14v1m0 4v2m0-18v2m0 4v1"/></g>`
	panelRightOpenPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 3v18m-5-6l-3-3l3-3"/></g>`
	panelTopPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18"/></g>`
	panelTopClosePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M9 16l3-3l3 3"/></g>`
	panelTopInactivePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M14 9h1m4 0h2M3 9h2m4 0h1"/></g>`
	panelTopOpenPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18m-6 5l-3 3l-3-3"/></g>`
	paperclipPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.44 11.05l-9.19 9.19a6 6 0 0 1-8.49-8.49l8.57-8.57A4 4 0 1 1 18 8.84l-8.59 8.57a2 2 0 0 1-2.83-2.83l8.49-8.48"/>`
	parenthesesPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 21s-4-3-4-9s4-9 4-9m8 0s4 3 4 9s-4 9-4 9"/>`
	parkingCirclePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 17V7h4a3 3 0 0 1 0 6H9"/></g>`
	parkingCircleOffPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m5 5l14 14m-6-6a3 3 0 1 0 0-6H9v2m0 8v-2.34"/></g>`
	parkingMeterPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 9a3 3 0 1 1 6 0m-3 3v3m-1 0h2"/><path d="M19 9a7 7 0 1 0-13.6 2.3C6.4 14.4 8 19 8 19h8s1.6-4.6 2.6-7.7c.3-.8.4-1.5.4-2.3m-7 10v3"/></g>`
	parkingSquarePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 17V7h4a3 3 0 0 1 0 6H9"/></g>`
	parkingSquareOffPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.6 3.6A2 2 0 0 1 5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-.59 1.41M3 8.7V19a2 2 0 0 0 2 2h10.3M2 2l20 20"/><path d="M13 13a3 3 0 1 0 0-6H9v2m0 8v-2.3"/></g>`
	partyPopperPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.8 11.3L2 22l10.7-3.79M4 3h.01M22 8h.01M15 2h.01M22 20h.01M22 2l-2.24.75a2.9 2.9 0 0 0-1.96 3.12v0c.1.86-.57 1.63-1.45 1.63h-.38c-.86 0-1.6.6-1.76 1.44L14 10m8 3l-.82-.33c-.86-.34-1.82.2-1.98 1.11v0c-.11.7-.72 1.22-1.43 1.22H17M11 2l.33.82c.34.86-.2 1.82-1.11 1.98v0C9.52 4.9 9 5.52 9 6.23V7"/><path d="M11 13c1.93 1.93 2.83 4.17 2 5c-.83.83-3.07-.07-5-2c-1.93-1.93-2.83-4.17-2-5c.83-.83 3.07.07 5 2Z"/></g>`
	pausePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4h4v16H6zm8 0h4v16h-4z"/>`
	pauseCirclePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 15V9m4 6V9"/></g>`
	pauseOctagonPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15V9m4 6V9M7.714 2h8.572L22 7.714v8.572L16.286 22H7.714L2 16.286V7.714L7.714 2z"/>`
	pawPrintPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="4" r="2"/><circle cx="18" cy="8" r="2"/><circle cx="20" cy="16" r="2"/><path d="M9 10a5 5 0 0 1 5 5v3.5a3.5 3.5 0 0 1-6.84 1.045q-.64-2.065-2.7-2.705A3.5 3.5 0 0 1 5.5 10Z"/></g>`
	pcCasePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2"/><path d="M15 14h.01M9 6h6m-6 4h6"/></g>`
	penPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5Z"/>`
	penLinePath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1l1-4Z"/>`
	penSquarePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.12 2.12 0 0 1 3 3L12 15l-4 1l1-4Z"/></g>`
	penToolPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 19l7-7l3 3l-7 7l-3-3z"/><path d="m18 13l-1.5-7.5L2 2l3.5 14.5L13 18l5-5zM2 2l7.586 7.586"/><circle cx="11" cy="11" r="2"/></g>`
	pencilPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5Zm-2 2l4 4"/>`
	pencilLinePath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1l1-4ZM15 5l3 3"/>`
	pencilRulerPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 5l4 4m-6-2L8.7 2.7a2.41 2.41 0 0 0-3.4 0L2.7 5.3a2.41 2.41 0 0 0 0 3.4L7 13m1-7l2-2M2 22l5.5-1.5L21.17 6.83a2.82 2.82 0 0 0-4-4L3.5 16.5Zm16-6l2-2"/><path d="m17 11l4.3 4.3c.94.94.94 2.46 0 3.4l-2.6 2.6c-.94.94-2.46.94-3.4 0L11 17"/></g>`
	percentPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 5L5 19"/><circle cx="6.5" cy="6.5" r="2.5"/><circle cx="17.5" cy="17.5" r="2.5"/></g>`
	percentCirclePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m15 9l-6 6m0-6h.01M15 15h.01"/></g>`
	percentDiamondPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Zm6.5-1.1h.01m5.29.3l-5 5m5.2.3h.01"/>`
	percentSquarePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m15 9l-6 6m0-6h.01M15 15h.01"/></g>`
	personStandingPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="1"/><path d="m9 20l3-6l3 6M6 8l6 2l6-2m-6 2v4"/></g>`
	phonePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 16.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneCallPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 16.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92zM14.05 2a9 9 0 0 1 8 7.94m-8-3.94A5 5 0 0 1 18 10"/>`
	phoneForwardedPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 2l4 4l-4 4m-4-4h8m0 10.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneIncomingPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2v6h6m0-6l-6 6m6 8.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneMissedPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 2l-6 6m0-6l6 6m0 8.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	phoneOffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.68 13.31a16 16 0 0 0 3.41 2.6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7a2 2 0 0 1 1.72 2v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.42 19.42 0 0 1-3.33-2.67m-2.67-3.34a19.79 19.79 0 0 1-3.07-8.63A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91M22 2L2 22"/>`
	phoneOutgoingPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 8V2h-6m0 6l6-6m0 14.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`
	piPath                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 4v16M4 7c0-1.7 1.3-3 3-3h13"/><path d="M18 20c-1.7 0-3-1.3-3-3V4"/></g>`
	piSquarePath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7h10m-7 0v10m6 0a2 2 0 0 1-2-2V7"/></g>`
	pictureInPicturePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4.5v5H3m-1-6l6 6m13 0v-3c0-1.16-.84-2-2-2h-7m-9 9v2c0 1.05.95 2 2 2h3"/><rect width="10" height="7" x="12" y="13.5" ry="2"/></g>`
	pictureInPictureTwoPath             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 9V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10c0 1.1.9 2 2 2h4"/><rect width="10" height="7" x="12" y="13" rx="2"/></g>`
	pieChartPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.21 15.89A10 10 0 1 1 8 2.83"/><path d="M22 12A10 10 0 0 0 12 2v10z"/></g>`
	piggyBankPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5c-1.5 0-2.8 1.4-3 2c-3.5-1.5-11-.3-11 5c0 1.8 0 3 2 4.5V20h4v-2h3v2h4v-4c1-.5 1.7-1 2-2h2v-4h-2c0-1-.5-1.5-1-2h0V5zM2 9v1c0 1.1.9 2 2 2h1m11-1h0"/>`
	pilcrowPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4v16m4-16v16m2-16H9.5a4.5 4.5 0 0 0 0 9H13"/>`
	pilcrowSquarePath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 12H9.5a2.5 2.5 0 0 1 0-5H17m-5 0v10m4-10v10"/></g>`
	pillPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.5 20.5l10-10a4.95 4.95 0 1 0-7-7l-10 10a4.95 4.95 0 1 0 7 7Zm-2-12l7 7"/>`
	pinPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17v5m-7-5h14v-1.76a2 2 0 0 0-1.11-1.79l-1.78-.9A2 2 0 0 1 15 10.76V6h1a2 2 0 0 0 0-4H8a2 2 0 0 0 0 4h1v4.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24Z"/>`
	pinOffPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20m-10-5v5M9 9v1.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V17h12m-2-7.66V6h1a2 2 0 0 0 0-4H7.89"/>`
	pipettePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 22l1-1h3l9-9M3 21v-3l9-9"/><path d="m15 6l3.4-3.4a2.1 2.1 0 1 1 3 3L18 9l.4.4a2.1 2.1 0 1 1-3 3l-3.8-3.8a2.1 2.1 0 1 1 3-3l.4.4Z"/></g>`
	pizzaPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 11h.01M11 15h.01M16 16h.01M2 16l20 6l-6-20A20 20 0 0 0 2 16"/><path d="M5.71 17.11a17.04 17.04 0 0 1 11.4-11.4"/></g>`
	planePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.8 19.2L16 11l3.5-3.5C21 6 21.5 4 21 3c-1-.5-3 0-4.5 1.5L13 8L4.8 6.2c-.5-.1-.9.1-1.1.5l-.3.5c-.2.5-.1 1 .3 1.3L9 12l-2 3H4l-1 1l3 2l2 3l1-1v-3l3-2l3.5 5.3c.3.4.8.5 1.3.3l.5-.2c.4-.3.6-.7.5-1.2z"/>`
	planeLandingPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22h20M3.77 10.77L2 9l2-4.5l1.1.55c.55.28.9.84.9 1.45s.35 1.17.9 1.45L8 8.5l3-6l1.05.53a2 2 0 0 1 1.09 1.52l.72 5.4a2 2 0 0 0 1.09 1.52l4.4 2.2c.42.22.78.55 1.01.96l.6 1.03c.49.88-.06 1.98-1.06 2.1l-1.18.15c-.47.06-.95-.02-1.37-.24L4.29 11.15a2 2 0 0 1-.52-.38Z"/>`
	planeTakeoffPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22h20M6.36 17.4L4 17l-2-4l1.1-.55a2 2 0 0 1 1.8 0l.17.1a2 2 0 0 0 1.8 0L8 12L5 6l.9-.45a2 2 0 0 1 2.09.2l4.02 3a2 2 0 0 0 2.1.2l4.19-2.06a2.41 2.41 0 0 1 1.73-.17L21 7a1.4 1.4 0 0 1 .87 1.99l-.38.76c-.23.46-.6.84-1.07 1.08L7.58 17.2a2 2 0 0 1-1.22.18Z"/>`
	playPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 3l14 9l-14 9V3z"/>`
	playCirclePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m10 8l6 4l-6 4V8z"/></g>`
	playSquarePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m9 8l6 4l-6 4Z"/></g>`
	plugPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22v-5M9 8V2m6 6V2m3 6v5a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4V8Z"/>`
	plugTwoPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 2v6m6-6v6m-3 9v5M5 8h14M6 11V8h12v3a6 6 0 1 1-12 0v0Z"/>`
	plugZapPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6l-2.3 2.3a2.4 2.4 0 0 0 0 3.4ZM2 22l3-3m2.5-5.5L10 11m.5 5.5L13 14m5-11l-4 4h6l-4 4"/>`
	plugZapTwoPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 2l-2 2.5h3L12 7m-2 7v-3m4 3v-3m-3 8c-1.7 0-3-1.3-3-3v-2h8v2c0 1.7-1.3 3-3 3Zm1 3v-3"/>`
	plusPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7-7v14"/>`
	plusCirclePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8m-4-4v8"/></g>`
	plusSquarePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8m-4-4v8"/></g>`
	pocketPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3h16a2 2 0 0 1 2 2v6a10 10 0 0 1-10 10A10 10 0 0 1 2 11V5a2 2 0 0 1 2-2z"/><path d="m8 10l4 4l4-4"/></g>`
	pocketKnifePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 2v1c0 1 2 1 2 2S3 6 3 7s2 1 2 2s-2 1-2 2s2 1 2 2m13-7h.01M6 18h.01m14.82-9.17a4 4 0 0 0-5.66-5.66l-12 12a4 4 0 1 0 5.66 5.66Z"/><path d="M18 11.66V22a4 4 0 0 0 4-4V6"/></g>`
	podcastPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11" r="1"/><path d="M11 17a1 1 0 0 1 2 0c0 .5-.34 3-.5 4.5a.5.5 0 0 1-1 0c-.16-1.5-.5-4-.5-4.5Zm-3-3a5 5 0 1 1 8 0"/><path d="M17 18.5a9 9 0 1 0-10 0"/></g>`
	pointerPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 14a8 8 0 0 1-8 8m4-11v-1a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0m0 0V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1m0-.5V4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v10"/><path d="M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"/></g>`
	popcornPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 8a2 2 0 0 0 0-4a2 2 0 0 0-4 0a2 2 0 0 0-4 0a2 2 0 0 0-4 0a2 2 0 0 0 0 4m4 14L9 8m5 14l1-14"/><path d="M20 8c.5 0 .9.4.8 1l-2.6 12c-.1.5-.7 1-1.2 1H7c-.6 0-1.1-.4-1.2-1L3.2 9c-.1-.6.3-1 .8-1Z"/></g>`
	popsiclePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.6 14.4c.8-.8.8-2 0-2.8l-8.1-8.1a4.95 4.95 0 1 0-7.1 7.1l8.1 8.1c.9.7 2.1.7 2.9-.1ZM22 22l-5.5-5.5"/>`
	poundSterlingPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 7c0-5.333-8-5.333-8 0m0 0v14m-4 0h12M6 13h10"/>`
	powerPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.36 6.64a9 9 0 1 1-12.73 0M12 2v10"/>`
	powerOffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.36 6.64A9 9 0 0 1 20.77 15M6.16 6.16a9 9 0 1 0 12.68 12.68M12 2v4M2 2l20 20"/>`
	presentationPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3h20m-1 0v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3m4 18l5-5l5 5"/>`
	printerPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9V2h12v7M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><path d="M6 14h12v8H6z"/></g>`
	projectorPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 7L3 5m6 1V3m4 4l2-2"/><circle cx="9" cy="13" r="3"/><path d="M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17M16 16h2"/></g>`
	puzzlePath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.439 7.85c-.049.322.059.648.289.878l1.568 1.568c.47.47.706 1.087.706 1.704s-.235 1.233-.706 1.704l-1.611 1.611a.98.98 0 0 1-.837.276c-.47-.07-.802-.48-.968-.925a2.501 2.501 0 1 0-3.214 3.214c.446.166.855.497.925.968a.979.979 0 0 1-.276.837l-1.61 1.61a2.404 2.404 0 0 1-1.705.707a2.402 2.402 0 0 1-1.704-.706l-1.568-1.568a1.026 1.026 0 0 0-.877-.29c-.493.074-.84.504-1.02.968a2.5 2.5 0 1 1-3.237-3.237c.464-.18.894-.527.967-1.02a1.026 1.026 0 0 0-.289-.877l-1.568-1.568A2.402 2.402 0 0 1 1.998 12c0-.617.236-1.234.706-1.704L4.23 8.77c.24-.24.581-.353.917-.303c.515.077.877.528 1.073 1.01a2.5 2.5 0 1 0 3.259-3.259c-.482-.196-.933-.558-1.01-1.073c-.05-.336.062-.676.303-.917l1.525-1.525A2.402 2.402 0 0 1 12 1.998c.617 0 1.234.236 1.704.706l1.568 1.568c.23.23.556.338.877.29c.493-.074.84-.504 1.02-.968a2.5 2.5 0 1 1 3.237 3.237c-.464.18-.894.527-.967 1.02Z"/>`
	qrCodePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="5" height="5" x="3" y="3" rx="1"/><rect width="5" height="5" x="16" y="3" rx="1"/><rect width="5" height="5" x="3" y="16" rx="1"/><path d="M21 16h-3a2 2 0 0 0-2 2v3m5 0v.01M12 7v3a2 2 0 0 1-2 2H7m-4 0h.01M12 3h.01M12 16v.01M16 12h1m4 0v.01M12 21v-1"/></g>`
	quotePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21c3 0 7-1 7-8V5c0-1.25-.756-2.017-2-2H4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2c1 0 1 0 1 1v1c0 1-1 2-2 2s-1 .008-1 1.031V20c0 1 0 1 1 1zm12 0c3 0 7-1 7-8V5c0-1.25-.757-2.017-2-2h-4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2h.75c0 2.25.25 4-2.75 4v3c0 1 0 1 1 1z"/>`
	rabbitPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 8.54V4a2 2 0 1 0-4 0v3"/><path d="M18 21h-8a4 4 0 0 1-4-4a7 7 0 0 1 7-7h.2L9.6 6.4a1.93 1.93 0 1 1 2.8-2.8L15.8 7h.2c3.3 0 6 2.7 6 6v1a2 2 0 0 1-2 2h-1c-1.7 0-3 1.3-3 3"/><path d="M7.61 12.53a3 3 0 1 0-1.6 4.3M13 16a3 3 0 0 1 2.24 5M18 12h.01"/></g>`
	radarPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.07 4.93A10 10 0 0 0 6.99 3.34M4 6h.01M2.29 9.62a10 10 0 1 0 19.02-1.27"/><path d="M16.24 7.76a6 6 0 1 0-8.01 8.91M12 18h.01m5.98-6.34a6 6 0 0 1-2.22 5.01"/><circle cx="12" cy="12" r="2"/><path d="m13.41 10.59l5.66-5.66"/></g>`
	radiationPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h0M7.5 4.2c-.3-.5-.9-.7-1.3-.4C3.9 5.5 2.3 8.1 2 11c-.1.5.4 1 1 1h5c0-1.5.8-2.8 2-3.4c-1.1-1.9-2-3.5-2.5-4.4zM21 12c.6 0 1-.4 1-1c-.3-2.9-1.8-5.5-4.1-7.1c-.4-.3-1.1-.2-1.3.3c-.6.9-1.5 2.5-2.6 4.3c1.2.7 2 2 2 3.5h5zM7.5 19.8c-.3.5-.1 1.1.4 1.3c2.6 1.2 5.6 1.2 8.2 0c.5-.2.7-.8.4-1.3c-.5-.9-1.4-2.5-2.5-4.3c-1.2.7-2.8.7-4 0c-1.1 1.8-2 3.4-2.5 4.3z"/>`
	radioPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.9 19.1C1 15.2 1 8.8 4.9 4.9m2.9 11.3c-2.3-2.3-2.3-6.1 0-8.5"/><circle cx="12" cy="12" r="2"/><path d="M16.2 7.8c2.3 2.3 2.3 6.1 0 8.5m2.9-11.4C23 8.8 23 15.1 19.1 19"/></g>`
	radioReceiverPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 16v2m14-2v2"/><rect width="20" height="8" x="2" y="8" rx="2"/><path d="M18 12h0"/></g>`
	radioTowerPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.9 16.1C1 12.2 1 5.8 4.9 1.9m2.9 2.8a6.14 6.14 0 0 0-.8 7.5"/><circle cx="12" cy="9" r="2"/><path d="M16.2 4.8c2 2 2.26 5.11.8 7.47M19.1 1.9a9.96 9.96 0 0 1 0 14.1m-9.6 2h5M8 22l4-11l4 11"/></g>`
	railSymbolPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15h14M5 9h14m-5 11l-5-5l6-6l-5-5"/>`
	rainbowPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 17a10 10 0 0 0-20 0"/><path d="M6 17a6 6 0 0 1 12 0"/><path d="M10 17a2 2 0 0 1 4 0"/></g>`
	ratPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 5c0-1.7-1.3-3-3-3s-3 1.3-3 3c0 .8.3 1.5.8 2H11c-3.9 0-7 3.1-7 7v0c0 2.2 1.8 4 4 4"/><path d="M16.8 3.9c.3-.3.6-.5 1-.7c1.5-.6 3.3.1 3.9 1.6c.6 1.5-.1 3.3-1.6 3.9l1.6 2.8c.2.3.2.7.2 1c-.2.8-.9 1.2-1.7 1.1c0 0-1.6-.3-2.7-.6H17c-1.7 0-3 1.3-3 3"/><path d="M13.2 18a3 3 0 0 0-2.2-5m2 9H4a2 2 0 0 1 0-4h12m0-9h.01"/></g>`
	ratioPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="12" height="20" x="6" y="2" rx="2"/><rect width="20" height="12" x="2" y="6" rx="2"/></g>`
	receiptPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 2v20l2-1l2 1l2-1l2 1l2-1l2 1l2-1l2 1V2l-2 1l-2-1l-2 1l-2-1l-2 1l-2-1l-2 1l-2-1Z"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 1V7"/></g>`
	rectangleHorizontalPath             = `<rect width="20" height="12" x="2" y="6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`
	rectangleVerticalPath               = `<rect width="12" height="20" x="6" y="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`
	recyclePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 19H4.815a1.83 1.83 0 0 1-1.57-.881a1.785 1.785 0 0 1-.004-1.784L7.196 9.5M11 19h8.203a1.83 1.83 0 0 0 1.556-.89a1.784 1.784 0 0 0 0-1.775l-1.226-2.12"/><path d="m14 16l-3 3l3 3m-5.707-8.404L7.196 9.5L3.1 10.598m6.244-4.787l1.093-1.892A1.83 1.83 0 0 1 11.985 3a1.784 1.784 0 0 1 1.546.888l3.943 6.843"/><path d="m13.378 9.633l4.096 1.098l1.097-4.096"/></g>`
	redoPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 7v6h-6"/><path d="M3 17a9 9 0 0 1 9-9a9 9 0 0 1 6 2.3l3 2.7"/></g>`
	redoDotPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="17" r="1"/><path d="M21 7v6h-6"/><path d="M3 17a9 9 0 0 1 9-9a9 9 0 0 1 6 2.3l3 2.7"/></g>`
	redoTwoPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 14l5-5l-5-5"/><path d="M20 9H9.5A5.5 5.5 0 0 0 4 14.5v0A5.5 5.5 0 0 0 9.5 20H13"/></g>`
	refreshCcwPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 0 0-9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5m-5 4a9 9 0 0 0 9 9a9.75 9.75 0 0 0 6.74-2.74L21 16"/><path d="M16 16h5v5"/></g>`
	refreshCcwDotPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 2v6h6"/><path d="M21 12A9 9 0 0 0 6 5.3L3 8m18 14v-6h-6"/><path d="M3 12a9 9 0 0 0 15 6.7l3-2.7"/><circle cx="12" cy="12" r="1"/></g>`
	refreshCwPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 0 1 9-9a9.75 9.75 0 0 1 6.74 2.74L21 8"/><path d="M21 3v5h-5m5 4a9 9 0 0 1-9 9a9.75 9.75 0 0 1-6.74-2.74L3 16"/><path d="M8 16H3v5"/></g>`
	refreshCwOffPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 8l-2.26-2.26A9.75 9.75 0 0 0 12 3c-1 0-1.97.16-2.87.47M8 16H3v5m0-9c0-2.49 1-4.74 2.64-6.36"/><path d="m3 16l2.26 2.26A9.75 9.75 0 0 0 12 21c2.49 0 4.74-1 6.36-2.64M21 12c0 1-.16 1.97-.47 2.87M21 3v5h-5m6 14L2 2"/></g>`
	refrigeratorPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6Zm0 4h14m-4-3v6"/>`
	regexPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3v10m-4.33-7.5l8.66 5m-8.66 0l8.66-5M9 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2z"/>`
	removeFormattingPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7V4h16v3M5 20h6m2-16L8 20m7-5l5 5m0-5l-5 5"/>`
	repeatPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 2l4 4l-4 4"/><path d="M3 11v-1a4 4 0 0 1 4-4h14M7 22l-4-4l4-4"/><path d="M21 13v1a4 4 0 0 1-4 4H3"/></g>`
	repeatOnePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 2l4 4l-4 4"/><path d="M3 11v-1a4 4 0 0 1 4-4h14M7 22l-4-4l4-4"/><path d="M21 13v1a4 4 0 0 1-4 4H3m8-8h1v4"/></g>`
	repeatTwoPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 9l3-3l3 3"/><path d="M13 18H7a2 2 0 0 1-2-2V6m17 9l-3 3l-3-3"/><path d="M11 6h6a2 2 0 0 1 2 2v10"/></g>`
	replacePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 4c0-1.1.9-2 2-2m4 0c1.1 0 2 .9 2 2m0 4c0 1.1-.9 2-2 2m-4 0c-1.1 0-2-.9-2-2M3 7l3 3l3-3"/><path d="M6 10V5c0-1.7 1.3-3 3-3h1"/><rect width="8" height="8" x="2" y="14" rx="2"/></g>`
	replaceAllPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 4c0-1.1.9-2 2-2m4 0c1.1 0 2 .9 2 2m0 4c0 1.1-.9 2-2 2m-4 0c-1.1 0-2-.9-2-2M3 7l3 3l3-3"/><path d="M6 10V5c0-1.7 1.3-3 3-3h1"/><rect width="8" height="8" x="2" y="14" rx="2"/><path d="M14 14c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2m6-8c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2"/></g>`
	replyPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 17l-5-5l5-5"/><path d="M20 18v-2a4 4 0 0 0-4-4H4"/></g>`
	replyAllPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 17l-5-5l5-5m5 10l-5-5l5-5"/><path d="M22 18v-2a4 4 0 0 0-4-4H7"/></g>`
	rewindPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 19l-9-7l9-7v14zm11 0l-9-7l9-7v14z"/>`
	rocketPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09zM12 15l-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z"/><path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0m1 7v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"/></g>`
	rockingChairPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3.5 2l3 10.5H18m-8.5 0l-4 7.5m9.5-7.5l3.5 7.5M2.75 18a13 13 0 0 0 18.5 0"/>`
	rollerCoasterPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19V5m4 14V6.8M14 19v-7.8M18 5v4m0 10v-6m4 6V9M2 19V9a4 4 0 0 1 4-4c2 0 4 1.33 6 4s4 4 6 4a4 4 0 1 0-3-6.65"/>`
	rotateCcwPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/></g>`
	rotateCwPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/></g>`
	rotateThreeDPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.466 7.5C15.643 4.237 13.952 2 12 2C9.239 2 7 6.477 7 12s2.239 10 5 10c.342 0 .677-.069 1-.2m2.194-8.093l3.814 1.86l-1.86 3.814"/><path d="M19 15.57c-1.804.885-4.274 1.43-7 1.43c-5.523 0-10-2.239-10-5s4.477-5 10-5c4.838 0 8.873 1.718 9.8 4"/></g>`
	routerPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6.01 18H6m4.01 0H10m5-8v4m2.84-6.83a4 4 0 0 0-5.66 0m8.48-2.83a8 8 0 0 0-11.31 0"/></g>`
	rowsPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 12h18"/></g>`
	rssPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 11a9 9 0 0 1 9 9M4 4a16 16 0 0 1 16 16"/><circle cx="5" cy="19" r="1"/></g>`
	rulerPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21.3 15.3a2.4 2.4 0 0 1 0 3.4l-2.6 2.6a2.4 2.4 0 0 1-3.4 0L2.7 8.7a2.41 2.41 0 0 1 0-3.4l2.6-2.6a2.41 2.41 0 0 1 3.4 0Zm-6.8-2.8l2-2m-5-1l2-2m-5-1l2-2m7 11l2-2"/>`
	russianRublePath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11h8a4 4 0 0 0 0-8H9v18m-3-6h8"/>`
	sailboatPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 18H2a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4Zm-1-4L10 2L3 14h18ZM10 2v16"/>`
	saladPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21h10m-5 0a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Z"/><path d="M11.38 12a2.4 2.4 0 0 1-.4-4.77a2.4 2.4 0 0 1 3.2-2.77a2.4 2.4 0 0 1 3.47-.63a2.4 2.4 0 0 1 3.37 3.37a2.4 2.4 0 0 1-1.1 3.7a2.51 2.51 0 0 1 .03 1.1M13 12l4-4"/><path d="M10.9 7.25A3.99 3.99 0 0 0 4 10c0 .73.2 1.41.54 2"/></g>`
	sandwichPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 11v3a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-3m-9 8H4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-3.83M3 11l7.77-6.04a2 2 0 0 1 2.46 0L21 11H3Z"/><path d="M12.97 19.77L7 15h12.5l-3.75 4.5a2 2 0 0 1-2.78.27Z"/></g>`
	satellitePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 7L9 3L5 7l4 4m8 0l4 4l-4 4l-4-4"/><path d="m8 12l4 4l6-6l-4-4Zm8-4l3-3M9 21a6 6 0 0 0-6-6"/></g>`
	satelliteDishPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10a7.31 7.31 0 0 0 10 10Zm5 5l3-3m5 1a6 6 0 0 0-6-6m10 6A10 10 0 0 0 11 3"/>`
	savePath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/><path d="M17 21v-8H7v8M7 3v5h8"/></g>`
	saveAllPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 4a2 2 0 0 1 2-2h10l4 4v10.2a2 2 0 0 1-2 1.8H8a2 2 0 0 1-2-2Z"/><path d="M10 2v4h6m2 12v-7h-8v7"/><path d="M18 22H4a2 2 0 0 1-2-2V6"/></g>`
	scalePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 16l3-8l3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1ZM2 16l3-8l3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Zm5 5h10M12 3v18M3 7h2c2 0 5-1 7-2c2 1 5 2 7 2h2"/>`
	scaleThreeDPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="19" r="2"/><circle cx="5" cy="5" r="2"/><path d="M5 7v12h12M5 19l6-6"/></g>`
	scalingPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 3L9 15m3-12H3v18h18v-9m-5-9h5v5"/><path d="M14 15H9v-5"/></g>`
	scanPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/>`
	scanFacePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2m5-3s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01"/>`
	scanLinePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2m4-5h10"/>`
	scatterChartPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7.5" cy="7.5" r=".5"/><circle cx="18.5" cy="5.5" r=".5"/><circle cx="11.5" cy="11.5" r=".5"/><circle cx="7.5" cy="16.5" r=".5"/><circle cx="17.5" cy="14.5" r=".5"/><path d="M3 3v18h18"/></g>`
	schoolPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m4 6l8-4l8 4m-2 4l4 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8l4-2"/><path d="M14 22v-4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v4m8-17v17M6 5v17"/><circle cx="12" cy="9" r="2"/></g>`
	schoolTwoPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="1"/><path d="M22 20V8h-4l-6-4l-6 4H2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2ZM6 17v.01M6 13v.01M18 17v.01M18 13v.01"/><path d="M14 22v-5a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5"/></g>`
	scissorsPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="6" r="3"/><path d="M8.12 8.12L12 12m8-8L8.12 15.88"/><circle cx="6" cy="18" r="3"/><path d="M14.8 14.8L20 20"/></g>`
	scissorsLineDashedPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.42 9.42L8 12"/><circle cx="4" cy="8" r="2"/><path d="m14 6l-8.58 8.58"/><circle cx="4" cy="16" r="2"/><path d="M10.8 14.8L14 18m2-6h-2m8 0h-2"/></g>`
	scissorsSquarePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><circle cx="8" cy="8" r="2"/><path d="M9.414 9.414L12 12m2.8 2.8L18 18"/><circle cx="8" cy="16" r="2"/><path d="m18 6l-8.586 8.586"/></g>`
	scissorsSquareDashedBottomPath      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2m-10 0H8m8 0h-2"/><circle cx="8" cy="8" r="2"/><path d="M9.414 9.414L12 12m2.8 2.8L18 18"/><circle cx="8" cy="16" r="2"/><path d="m18 6l-8.586 8.586"/></g>`
	screenSharePath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3M8 21h8m-4-4v4m5-13l5-5m-5 0h5v5"/>`
	screenShareOffPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3M8 21h8m-4-4v4M22 3l-5 5m0-5l5 5"/>`
	scrollPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 21h12a2 2 0 0 0 2-2v-2H10v2a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v3h4"/><path d="M19 17V5a2 2 0 0 0-2-2H4"/></g>`
	scrollTextPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 21h12a2 2 0 0 0 2-2v-2H10v2a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v3h4"/><path d="M19 17V5a2 2 0 0 0-2-2H4m11 5h-5m5 4h-5"/></g>`
	searchPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`
	searchCheckPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 11l2 2l4-4"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`
	searchCodePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 9l-2 2l2 2m4 0l2-2l-2-2"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`
	searchLargePath                     = `<g fill="none"><path clip-rule="evenodd" d="M18.874 19.581a6 6 0 1 1 .707-.707l4.273 4.272l-.708.708zM20 15a5 5 0 1 1-10 0a5 5 0 0 1 10 0z" fill="currentColor" fill-rule="evenodd"/></g>`
	searchSlashPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.5 8.5l-5 5"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`
	searchXPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.5 8.5l-5 5m0-5l5 5"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`
	sendPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 2l-7 20l-4-9l-9-4Zm0 0L11 13"/>`
	sendHorizontalPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l3 9l-3 9l19-9Zm3 9h16"/>`
	sendToBackPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="14" y="14" rx="2"/><rect width="8" height="8" x="2" y="2" rx="2"/><path d="M7 14v1a2 2 0 0 0 2 2h1m4-10h1a2 2 0 0 1 2 2v1"/></g>`
	separatorHorizontalPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M8 8l4-4l4 4m0 8l-4 4l-4-4"/>`
	separatorVerticalPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v18M8 8l-4 4l4 4m8 0l4-4l-4-4"/>`
	serverPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="2" rx="2" ry="2"/><rect width="20" height="8" x="2" y="14" rx="2" ry="2"/><path d="M6 6h.01M6 18h.01"/></g>`
	serverCogPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M4.5 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.5m-15 4H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-.5M6 6h.01M6 18h.01m9.69-4.6l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.7.9l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4m-2.3-2.1l-.3.9"/></g>`
	serverCrashPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2M6 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2M6 6h.01M6 18h.01"/><path d="m13 6l-4 6h6l-4 6"/></g>`
	serverOffPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 2h13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-5m-5 0L2.5 2.5C2 2 2 2.5 2 5v3a2 2 0 0 0 2 2h6zm12 7v-1a2 2 0 0 0-2-2h-1M4 14a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16.5l1-.5l.5.5l-8-8H4zm2 4h.01M2 2l20 20"/>`
	settingsPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z"/><circle cx="12" cy="12" r="3"/></g>`
	settingsTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 7h-9m3 10H5"/><circle cx="17" cy="17" r="3"/><circle cx="7" cy="7" r="3"/></g>`
	shapesPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.3 10a.7.7 0 0 1-.626-1.079L11.4 3a.7.7 0 0 1 1.198-.043L16.3 8.9a.7.7 0 0 1-.572 1.1Z"/><rect width="7" height="7" x="3" y="14" rx="1"/><circle cx="17.5" cy="17.5" r="3.5"/></g>`
	sharePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8m-4-6l-4-4l-4 4m4-4v13"/>`
	shareTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/><path d="m8.59 13.51l6.83 3.98m-.01-10.98l-6.82 3.98"/></g>`
	sheetPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M3 15h18M9 9v12m6-12v12"/></g>`
	shellPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 11a2 2 0 1 1-4 0a4 4 0 0 1 8 0a6 6 0 0 1-12 0a8 8 0 0 1 16 0a10 10 0 1 1-20 0a11.93 11.93 0 0 1 2.42-7.22a2 2 0 1 1 3.16 2.44"/>`
	shieldPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10"/>`
	shieldAlertPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10m0-14v4m0 4h.01"/>`
	shieldBanPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M4 5l14 12"/>`
	shieldCheckPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10"/><path d="m9 12l2 2l4-4"/></g>`
	shieldClosePath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10zM9.5 9l5 5m0-5l-5 5"/>`
	shieldEllipsisPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M8 11h.01M12 11h.01M16 11h.01"/>`
	shieldHalfPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10m0 0V2"/>`
	shieldMinusPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M8 11h8"/>`
	shieldOffPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.7 14a6.9 6.9 0 0 0 .3-2V5l-8-3l-3.2 1.2M2 2l20 20M4.7 4.7L4 5v7c0 6 8 10 8 10a20.3 20.3 0 0 0 5.62-4.38"/>`
	shieldPlusPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M8 11h8m-4 4V7"/>`
	shieldQuestionPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10"/><path d="M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3m.08 4h.01"/></g>`
	shieldXPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10m2.5-13l-5 5m0-5l5 5"/>`
	shipPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 21c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/><path d="M19.38 20A11.6 11.6 0 0 0 21 14l-9-4l-9 4c0 2.9.94 5.34 2.81 7.76"/><path d="M19 13V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v6m7-3v4m0-12v3"/></g>`
	shipWheelPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="M12 2v7.5M19 5l-5.23 5.23M22 12h-7.5m4.5 7l-5.23-5.23M12 14.5V22m-1.77-8.23L5 19m4.5-7H2m8.23-1.77L5 5"/><circle cx="12" cy="12" r="2.5"/></g>`
	shirtPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.38 3.46L16 2a4 4 0 0 1-8 0L3.62 3.46a2 2 0 0 0-1.34 2.23l.58 3.47a1 1 0 0 0 .99.84H6v10c0 1.1.9 2 2 2h8a2 2 0 0 0 2-2V10h2.15a1 1 0 0 0 .99-.84l.58-3.47a2 2 0 0 0-1.34-2.23z"/>`
	shoppingBagPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4ZM3 6h18"/><path d="M16 10a4 4 0 0 1-8 0"/></g>`
	shoppingBasketPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 11l4-7m10 7l-4-7M2 11h20M3.5 11l1.6 7.4a2 2 0 0 0 2 1.6h9.8c.9 0 1.8-.7 2-1.6l1.7-7.4M9 11l1 9m-5.5-4.5h15M15 11l-1 9"/>`
	shoppingCartPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="21" r="1"/><circle cx="19" cy="21" r="1"/><path d="M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12"/></g>`
	shovelPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22v-5l5-5l5 5l-5 5zm7.5-7.5L16 8m1-6l5 5l-.5.5a3.53 3.53 0 0 1-5 0s0 0 0 0a3.53 3.53 0 0 1 0-5L17 2"/>`
	showerHeadPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l2.5 2.5m7 0a4.95 4.95 0 0 0-7 7M15 5L5 15m9 2v.01M10 16v.01M13 13v.01M16 10v.01M11 20v.01M17 14v.01M20 11v.01"/>`
	shrinkPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 15l6 6m-6-6v4.8m0-4.8h4.8M9 19.8V15m0 0H4.2M9 15l-6 6M15 4.2V9m0 0h4.8M15 9l6-6M9 4.2V9m0 0H4.2M9 9L3 3"/>`
	shrubPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22v-7l-2-2"/><path d="M17 8v.8A6 6 0 0 1 13.8 20v0H10v0A6.5 6.5 0 0 1 7 8h0a5 5 0 0 1 10 0Zm-3 6l-2 2"/></g>`
	shufflePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 18h1.4c1.3 0 2.5-.6 3.3-1.7l6.1-8.6c.7-1.1 2-1.7 3.3-1.7H22"/><path d="m18 2l4 4l-4 4M2 6h1.9c1.5 0 2.9.9 3.6 2.2M22 18h-5.9c-1.3 0-2.6-.7-3.3-1.8l-.5-.8"/><path d="m18 14l4 4l-4 4"/></g>`
	sigmaPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 7V4H6l6 8l-6 8h12v-3"/>`
	sigmaSquarePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M16 8.9V7H8l4 5l-4 5h8v-1.9"/></g>`
	signalPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4m5 4v-8m5 8V8m5-4v16"/>`
	signalHighPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4m5 4v-8m5 8V8"/>`
	signalLowPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4"/>`
	signalMediumPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4m5 4v-8"/>`
	signalZeroPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01"/>`
	sirenPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v6H7v-6Zm-2 8a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2H5v-2Zm16-8h1m-3.5-7.5L18 5M2 12h1m9-10v1M4.929 4.929l.707.707M12 12v6"/>`
	skipBackPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20L9 12l10-8v16zM5 19V5"/>`
	skipForwardPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 4l10 8l-10 8V4zm14 1v14"/>`
	skullPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="12" r="1"/><circle cx="15" cy="12" r="1"/><path d="M8 20v2h8v-2m-3.5-3l-.5-1l-.5 1h1z"/><path d="M16 20a2 2 0 0 0 1.56-3.25a8 8 0 1 0-11.12 0A2 2 0 0 0 8 20"/></g>`
	slackPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="3" height="8" x="13" y="2" rx="1.5"/><path d="M19 8.5V10h1.5A1.5 1.5 0 1 0 19 8.5"/><rect width="3" height="8" x="8" y="14" rx="1.5"/><path d="M5 15.5V14H3.5A1.5 1.5 0 1 0 5 15.5"/><rect width="8" height="3" x="14" y="13" rx="1.5"/><path d="M15.5 19H14v1.5a1.5 1.5 0 1 0 1.5-1.5"/><rect width="8" height="3" x="2" y="8" rx="1.5"/><path d="M8.5 5H10V3.5A1.5 1.5 0 1 0 8.5 5"/></g>`
	slashPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 2L2 22"/>`
	slicePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 14l-6 6h9v-3"/><path d="M18.37 3.63L8 14l3 3L21.37 6.63a2.12 2.12 0 1 0-3-3Z"/></g>`
	slidersPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 21v-7m0-4V3m8 18v-9m0-4V3m8 18v-5m0-4V3M2 14h4m4-6h4m4 8h4"/>`
	slidersHorizontalPath               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 4h-7m-4 0H3m18 8h-9m-4 0H3m18 8h-5m-4 0H3M14 2v4m-6 4v4m8 4v4"/>`
	smartphonePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12 18h.01"/></g>`
	smartphoneChargingPath              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12.667 8L10 12h4l-2.667 4"/></g>`
	smartphoneNfcPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="12" x="2" y="6" rx="1"/><path d="M13 8.32a7.43 7.43 0 0 1 0 7.36m3.46-9.47a11.76 11.76 0 0 1 0 11.58M19.91 4.1a15.91 15.91 0 0 1 .01 15.8"/></g>`
	smilePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 14s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01"/></g>`
	smilePlusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 11v1a10 10 0 1 1-9-10"/><path d="M8 14s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01M16 5h6m-3-3v6"/></g>`
	snailPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 13a6 6 0 1 0 12 0a4 4 0 1 0-8 0a2 2 0 0 0 4 0"/><circle cx="10" cy="13" r="8"/><path d="M2 21h12c4.4 0 8-3.6 8-8V7a2 2 0 1 0-4 0v6m0-10l1.1 2.2M22 3l-1.1 2.2"/></g>`
	snowflakePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h20M12 2v20m8-6l-4-4l4-4M4 8l4 4l-4 4M16 4l-4 4l-4-4m0 16l4-4l4 4"/>`
	sofaPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 9V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v3"/><path d="M2 11v5a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v2H6v-2a2 2 0 0 0-4 0Zm2 7v2m16-2v2M12 4v9"/></g>`
	sortAscPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 11h4m-4 4h7m-7 4h10M9 7L6 4L3 7m3-1v14"/>`
	sortDescPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5h10M11 9h7m-7 4h4M3 17l3 3l3-3m-3 1V4"/>`
	soupPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Zm-5 0h10m2.5-9L22 6m-5.75-3c.27.1.8.53.75 1.36c-.06.83-.93 1.2-1 2.02c-.05.78.34 1.24.73 1.62m-5.48-5c.27.1.8.53.74 1.36c-.05.83-.93 1.2-.98 2.02c-.06.78.33 1.24.72 1.62M6.25 3c.27.1.8.53.75 1.36c-.06.83-.93 1.2-1 2.02c-.05.78.34 1.24.74 1.62"/>`
	spacePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"/>`
	spadePath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9c-1.5 1.5-3 3.2-3 5.5A5.5 5.5 0 0 0 7.5 20c1.8 0 3-.5 4.5-2c1.5 1.5 2.7 2 4.5 2a5.5 5.5 0 0 0 5.5-5.5c0-2.3-1.5-4-3-5.5l-7-7l-7 7Zm7 9v4"/>`
	sparklePath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l-1.9 5.8a2 2 0 0 1-1.287 1.288L3 12l5.8 1.9a2 2 0 0 1 1.288 1.287L12 21l1.9-5.8a2 2 0 0 1 1.287-1.288L21 12l-5.8-1.9a2 2 0 0 1-1.288-1.287Z"/>`
	sparklesPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l-1.912 5.813a2 2 0 0 1-1.275 1.275L3 12l5.813 1.912a2 2 0 0 1 1.275 1.275L12 21l1.912-5.813a2 2 0 0 1 1.275-1.275L21 12l-5.813-1.912a2 2 0 0 1-1.275-1.275L12 3ZM5 3v4m14 10v4M3 5h4m10 14h4"/>`
	speakerPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><circle cx="12" cy="14" r="4"/><path d="M12 6h.01"/></g>`
	spellCheckPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 16l6-12l6 12M8 12h8m0 8l2 2l4-4"/>`
	spellCheckTwoPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 16l6-12l6 12M8 12h8M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1c1.1 0 1.1 1 2.3 1c1.1 0 1.1-1 2.3-1c1.1 0 1.1 1 2.3 1c1.1 0 1.1-1 2.3-1"/>`
	splinePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><path d="M5 17A12 12 0 0 1 17 5"/></g>`
	splitPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 3h5v5M8 3H3v5"/><path d="M12 22v-8.3a4 4 0 0 0-1.172-2.872L3 3m12 6l6-6"/></g>`
	splitSquareHorizontalPath           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19H5c-1 0-2-1-2-2V7c0-1 1-2 2-2h3m8 0h3c1 0 2 1 2 2v10c0 1-1 2-2 2h-3M12 4v16"/>`
	splitSquareVerticalPath             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8V5c0-1 1-2 2-2h10c1 0 2 1 2 2v3m0 8v3c0 1-1 2-2 2H7c-1 0-2-1-2-2v-3m-1-4h16"/>`
	sprayCanPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h.01M7 5h.01M11 7h.01M3 7h.01M7 9h.01M3 11h.01M15 5h4v4h-4zm4 4l2 2v10c0 .6-.4 1-1 1h-6c-.6 0-1-.4-1-1V11l2-2m-2 5l8-2m-8 7l8-2"/>`
	sproutPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 20h10m-7 0c5.5-2.5.8-6.4 3-10"/><path d="M9.5 9.4c1.1.8 1.8 2.2 2.3 3.7c-2 .4-3.5.4-4.8-.3c-1.2-.6-2.3-1.9-3-4.2c2.8-.5 4.4 0 5.5.8zM14.1 6a7 7 0 0 0-1.1 4c1.9-.1 3.3-.6 4.3-1.4c1-1 1.6-2.3 1.7-4.6c-2.7.1-4 1-4.9 2z"/></g>`
	squarePath                          = `<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`
	squareAsteriskPath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 8v8m-3.5-2l7-4m-7 0l7 4"/></g>`
	squareCodePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m10 10l-2 2l2 2m4 0l2-2l-2-2"/></g>`
	squareDashedBottomPath              = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2M9 21h1m4 0h1"/>`
	squareDashedBottomCodePath          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 10l-2 2l2 2m4 0l2-2l-2-2"/><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2M9 21h1m4 0h1"/></g>`
	squareDotPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="12" r="1"/></g>`
	squareEqualPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 10h10M7 14h10"/></g>`
	squareSlashPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m9 15l6-6"/></g>`
	squareStackPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2m0 12c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2"/><rect width="8" height="8" x="14" y="14" rx="2"/></g>`
	squirrelPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 6a4 4 0 0 0-4 4a7 7 0 0 0-7 7c0-5 4-5 4-10.5a4.5 4.5 0 1 0-9 0a2.5 2.5 0 0 0 5 0C7 10 3 11 3 17c0 2.8 2.2 5 5 5h10"/><path d="M16 20c0-1.7 1.3-3 3-3h1a2 2 0 0 0 2-2v-2a4 4 0 0 0-4-4V4"/><path d="M15.2 22a3 3 0 0 0-2.2-5m5-4h.01"/></g>`
	stampPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 22h14m.27-8.27A2.5 2.5 0 0 0 17.5 13h-11A2.5 2.5 0 0 0 4 15.5V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-1.5c0-.66-.26-1.3-.73-1.77Z"/><path d="M14 13V8.5C14 7 15 7 15 5a3 3 0 0 0-3-3c-1.66 0-3 1-3 3s1 2 1 3.5V13"/></g>`
	starPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l3.09 6.26L22 9.27l-5 4.87l1.18 6.88L12 17.77l-6.18 3.25L7 14.14L2 9.27l6.91-1.01L12 2z"/>`
	starHalfPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.8L5.8 21L7 14.1L2 9.3l7-1L12 2"/>`
	starOffPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.34 8.34L2 9.27l5 4.87L5.82 21L12 17.77L18.18 21l-.59-3.43m.83-4.81L22 9.27l-6.91-1L12 2l-1.44 2.91M2 2l20 20"/>`
	stepBackPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V4m-4 16L4 12l10-8z"/>`
	stepForwardPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4v16m4-16l10 8l-10 8z"/>`
	stethoscopePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.8 2.3A.3.3 0 1 0 5 2H4a2 2 0 0 0-2 2v5a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6V4a2 2 0 0 0-2-2h-1a.2.2 0 1 0 .3.3"/><path d="M8 15v1a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6v-4"/><circle cx="20" cy="10" r="2"/></g>`
	stickerPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z"/><path d="M15 3v6h6m-11 7s.8 1 2 1c1.3 0 2-1 2-1m-6-3h0m8 0h0"/></g>`
	stickyNotePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z"/><path d="M15 3v6h6"/></g>`
	stopCirclePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 9h6v6H9z"/></g>`
	storePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 7l4.41-4.41A2 2 0 0 1 7.83 2h8.34a2 2 0 0 1 1.42.59L22 7M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path d="M15 22v-4a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4M2 7h20m0 0v3a2 2 0 0 1-2 2v0a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 16 12a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 12 12a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 8 12a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 4 12v0a2 2 0 0 1-2-2V7"/></g>`
	stretchHorizontalPath               = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="6" x="2" y="4" rx="2"/><rect width="20" height="6" x="2" y="14" rx="2"/></g>`
	stretchVerticalPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="20" x="4" y="2" rx="2"/><rect width="6" height="20" x="14" y="2" rx="2"/></g>`
	strikethroughPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4H9a3 3 0 0 0-2.83 4M14 12a4 4 0 0 1 0 8H6m-2-8h16"/>`
	subscriptPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 5l8 8m0-8l-8 8m16 6h-4c0-1.5.44-2 1.5-2.5S20 15.33 20 14c0-.47-.17-.93-.48-1.29a2.11 2.11 0 0 0-2.62-.44c-.42.24-.74.62-.9 1.07"/>`
	subtitlesPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13h4m4 0h2M7 9h2m4 0h4m4 6a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10Z"/>`
	sunPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 2v2m0 16v2M4.93 4.93l1.41 1.41m11.32 11.32l1.41 1.41M2 12h2m16 0h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41"/></g>`
	sunDimPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 4h.01M20 12h.01M12 20h.01M4 12h.01m13.647-5.657h.01m-.01 11.314h.01m-11.324 0h.01m-.01-11.314h.01"/></g>`
	sunMediumPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 3v1m0 16v1m-9-9h1m16 0h1m-2.636-6.364l-.707.707M6.343 17.657l-.707.707m0-12.728l.707.707m11.314 11.314l.707.707"/></g>`
	sunMoonPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 8a2 2 0 1 0 4 4M12 2v2m0 16v2M4.93 4.93l1.41 1.41m11.32 11.32l1.41 1.41M2 12h2m16 0h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41"/></g>`
	sunSnowPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9a3 3 0 1 0 0 6m-8-3h1m11 9V3m-4 1V3m0 18v-1m-6.36-1.64l.7-.7m0-11.32l-.7-.7M14 12h8m-5-8l-3 3m0 10l3 3m4-5l-3-3l3-3"/>`
	sunrisePath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v8m-7.07.93l1.41 1.41M2 18h2m16 0h2m-2.93-7.07l-1.41 1.41M22 22H2M8 6l4-4l4 4m0 12a4 4 0 0 0-8 0"/>`
	sunsetPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10V2m-7.07 8.93l1.41 1.41M2 18h2m16 0h2m-2.93-7.07l-1.41 1.41M22 22H2M16 6l-4 4l-4-4m8 12a4 4 0 0 0-8 0"/>`
	superscriptPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 19l8-8m0 8l-8-8m16 1h-4c0-1.5.442-2 1.5-2.5S20 8.334 20 7.002c0-.472-.17-.93-.484-1.29a2.105 2.105 0 0 0-2.617-.436c-.42.239-.738.614-.899 1.06"/>`
	swissFrancPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 21V3h8M6 16h9m-5-6.5h7"/>`
	switchCameraPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 19H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h5m4 0h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-5"/><circle cx="12" cy="12" r="3"/><path d="m18 22l-3-3l3-3M6 2l3 3l-3 3"/></g>`
	swordPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 17.5L3 6V3h3l11.5 11.5M13 19l6-6m-3 3l4 4m-1 1l2-2"/>`
	swordsPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 17.5L3 6V3h3l11.5 11.5M13 19l6-6m-3 3l4 4m-1 1l2-2M14.5 6.5L18 3h3v3l-3.5 3.5M5 14l4 4m-2-1l-3 3m-1-1l2 2"/>`
	syringePath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 2l4 4m-5 1l3-3m-1 5L8.7 19.3c-1 1-2.5 1-3.4 0l-.6-.6c-1-1-1-2.5 0-3.4L15 5m-6 6l4 4m-8 4l-3 3M14 4l6 6"/>`
	tablePath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3v18"/><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18M3 15h18"/></g>`
	tablePropertiesPath                 = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 3v18"/><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M21 9H3m18 6H3"/></g>`
	tableTwoPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3H5a2 2 0 0 0-2 2v4m6-6h10a2 2 0 0 1 2 2v4M9 3v18m0 0h10a2 2 0 0 0 2-2V9M9 21H5a2 2 0 0 1-2-2V9m0 0h18"/>`
	tabletPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><path d="M12 18h.01"/></g>`
	tabletSmartphonePath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="10" height="14" x="3" y="8" rx="2"/><path d="M5 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2.4M8 18h.01"/></g>`
	tabletsPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="7" r="5"/><circle cx="17" cy="17" r="5"/><path d="M12 17h10M3.46 10.54l7.08-7.08"/></g>`
	tagPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2H2v10l9.29 9.29c.94.94 2.48.94 3.42 0l6.58-6.58c.94-.94.94-2.48 0-3.42L12 2ZM7 7h.01"/>`
	tagsPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5H2v7l6.29 6.29c.94.94 2.48.94 3.42 0l3.58-3.58c.94-.94.94-2.48 0-3.42L9 5ZM6 9.01V9"/><path d="m15 5l6.3 6.3a2.4 2.4 0 0 1 0 3.4L17 19"/></g>`
	tallyFivePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16m5-16v16m5-16v16m3-14L2 18"/>`
	tallyFourPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16m5-16v16m5-16v16"/>`
	tallyOnePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16"/>`
	tallyThreePath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16m5-16v16"/>`
	tallyTwoPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16"/>`
	targetPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="6"/><circle cx="12" cy="12" r="2"/></g>`
	tentPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20L10 4M5 20l9-16M3 20h18m-9-5l-3 5m3-5l3 5"/>`
	terminalPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 17l6-6l-6-6m8 14h8"/>`
	terminalSquarePath                  = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 11l2-2l-2-2m4 6h4"/><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/></g>`
	testTubePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5h0c-1.4 0-2.5-1.1-2.5-2.5V2m-1 0h7m-1 14h-5"/>`
	testTubeTwoPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7L6.82 21.18a2.83 2.83 0 0 1-3.99-.01v0a2.83 2.83 0 0 1 0-4L17 3m-1-1l6 6m-10 8H4"/>`
	testTubesPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 2v17.5A2.5 2.5 0 0 1 6.5 22v0A2.5 2.5 0 0 1 4 19.5V2m16 0v17.5a2.5 2.5 0 0 1-2.5 2.5v0a2.5 2.5 0 0 1-2.5-2.5V2M3 2h7m4 0h7M9 16H4m16 0h-5"/>`
	textPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 6.1H3m18 6H3M15.1 18H3"/>`
	textCursorPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 22h-1a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h1M7 22h1a4 4 0 0 0 4-4v-1M7 2h1a4 4 0 0 1 4 4v1"/>`
	textCursorInputPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h1a3 3 0 0 1 3 3a3 3 0 0 1 3-3h1m0 16h-1a3 3 0 0 1-3-3a3 3 0 0 1-3 3H5m0-4H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h1m8 0h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-7M9 7v10"/>`
	textQuotePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 6H3m18 6H8m13 6H8m-5-6v6"/>`
	textSelectPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3a2 2 0 0 0-2 2m16-2a2 2 0 0 1 2 2m0 14a2 2 0 0 1-2 2M5 21a2 2 0 0 1-2-2M9 3h1M9 21h1m4-18h1m-1 18h1M3 9v1m18-1v1M3 14v1m18-1v1M7 8h8m-8 4h10M7 16h6"/>`
	thermometerPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"/>`
	thermometerSnowflakePath            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h10M9 4v16M3 9l3 3l-3 3m9-9L9 9L6 6m0 12l3-3l1.5 1.5M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"/>`
	thermometerSunPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9a4 4 0 0 0-2 7.5M12 3v2M6.6 18.4l-1.4 1.4M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0ZM4 13H2m4.34-5.66L4.93 5.93"/>`
	thumbsDownPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 14V2M9 18.12L10 14H4.17a2 2 0 0 1-1.92-2.56l2.33-8A2 2 0 0 1 6.5 2H20a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.76a2 2 0 0 0-1.79 1.11L12 22h0a3.13 3.13 0 0 1-3-3.88Z"/>`
	thumbsUpPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10v12m8-16.12L14 10h5.83a2 2 0 0 1 1.92 2.56l-2.33 8A2 2 0 0 1 17.5 22H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h2.76a2 2 0 0 0 1.79-1.11L12 2h0a3.13 3.13 0 0 1 3 3.88Z"/>`
	ticketPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Zm11-4v2m0 10v2m0-8v2"/>`
	timerPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2h4m-2 12l3-3"/><circle cx="12" cy="14" r="8"/></g>`
	timerOffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2h4m-9.4 9a8 8 0 0 0 1.7 8.7a8 8 0 0 0 8.7 1.7m-7.6-14a8 8 0 0 1 10.3 1a8 8 0 0 1 .9 10.2M2 2l20 20M12 12v-2"/>`
	timerResetPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2h4m-2 12v-4m-8 3a8 8 0 0 1 8-7a8 8 0 1 1-5.3 14L4 17.6"/><path d="M9 17H4v5"/></g>`
	toggleLeftPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="6" ry="6"/><circle cx="8" cy="12" r="2"/></g>`
	toggleRightPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="6" ry="6"/><circle cx="16" cy="12" r="2"/></g>`
	tornadoPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 4H3m15 4H6m13 4H9m7 4h-6m1 4H9"/>`
	touchpadPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M2 14h20m-10 6v-6"/></g>`
	touchpadOffPath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16M2 14h12m8 0h-2m-8 6v-6M2 2l20 20m0-6V6a2 2 0 0 0-2-2H10"/>`
	towerControlPath                    = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.2 12.27L20 6H4l1.8 6.27a1 1 0 0 0 .95.73h10.5a1 1 0 0 0 .96-.73ZM8 13v9m8 0v-9M9 6l1 7m5-7l-1 7m-2-7V2m1 0h-2"/>`
	toyBrickPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="8" rx="1"/><path d="M10 8V5c0-.6-.4-1-1-1H6a1 1 0 0 0-1 1v3m14 0V5c0-.6-.4-1-1-1h-3a1 1 0 0 0-1 1v3"/></g>`
	tractorPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4h9l1 7m-9 0V4m4 6V4m10 1c-.6 0-1 .4-1 1v5.6"/><path d="m10 11l11 .9c.6 0 .9.5.8 1.1l-.8 5h-1"/><circle cx="7" cy="15" r=".5"/><circle cx="7" cy="15" r="5"/><path d="M16 18h-5"/><circle cx="18" cy="18" r="2"/></g>`
	trafficConePath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.3 6.2a4.55 4.55 0 0 0 5.4 0m-6.8 4.5c.9.8 2.4 1.3 4.1 1.3s3.2-.5 4.1-1.3"/><path d="M13.9 3.5a1.93 1.93 0 0 0-3.8-.1l-3 10c-.1.2-.1.4-.1.6c0 1.7 2.2 3 5 3s5-1.3 5-3c0-.2 0-.4-.1-.5Z"/><path d="m7.5 12.2l-4.7 2.7c-.5.3-.8.7-.8 1.1s.3.8.8 1.1l7.6 4.5c.9.5 2.1.5 3 0l7.6-4.5c.7-.3 1-.7 1-1.1s-.3-.8-.8-1.1l-4.7-2.8"/></g>`
	trainFrontPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 3.1V7a4 4 0 0 0 8 0V3.1M9 15l-1-1m7 1l1-1"/><path d="M9 19c-2.8 0-5-2.2-5-5v-4a8 8 0 0 1 16 0v4c0 2.8-2.2 5-5 5Zm-1 0l-2 3m10-3l2 3"/></g>`
	trainFrontTunnelPath                = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 22V12a10 10 0 1 1 20 0v10"/><path d="M15 6.8v1.4a3 2.8 0 1 1-6 0V6.8m1 8.2h.01M14 15h.01"/><path d="M10 19a4 4 0 0 1-4-4v-3a6 6 0 1 1 12 0v3a4 4 0 0 1-4 4Zm-1 0l-2 3m8-3l2 3"/></g>`
	trainTrackPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 17L17 2M2 14l8 8M5 11l8 8M8 8l8 8M11 5l8 8M14 2l8 8M7 22L22 7"/>`
	tramFrontPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="3" rx="2"/><path d="M4 11h16m-8-8v8m-4 8l-2 3m12 0l-2-3m-8-4h0m8 0h0"/></g>`
	trashPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6m3 0V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/>`
	trashTwoPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6m3 0V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2m-6 5v6m4-6v6"/>`
	treeDeciduousPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19h8a4 4 0 0 0 3.8-2.8a4 4 0 0 0-1.6-4.5c1-1.1 1-2.7.4-4c-.7-1.2-2.2-2-3.6-1.7a3 3 0 0 0-3-3a3 3 0 0 0-3 3c-1.4-.2-2.9.5-3.6 1.7c-.7 1.3-.5 2.9.4 4a4 4 0 0 0-1.6 4.5A4 4 0 0 0 8 19Zm4 0v3"/>`
	treePinePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 14l3 3.3a1 1 0 0 1-.7 1.7H4.7a1 1 0 0 1-.7-1.7L7 14h-.3a1 1 0 0 1-.7-1.7L9 9h-.2A1 1 0 0 1 8 7.3L12 3l4 4.3a1 1 0 0 1-.8 1.7H15l3 3.3a1 1 0 0 1-.7 1.7H17Zm-5 8v-3"/>`
	treesPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 10v.2A3 3 0 0 1 8.9 16v0H5v0h0a3 3 0 0 1-1-5.8V10a3 3 0 0 1 6 0Zm-3 6v6m6-3v3"/><path d="M12 19h8.3a1 1 0 0 0 .7-1.7L18 14h.3a1 1 0 0 0 .7-1.7L16 9h.2a1 1 0 0 0 .8-1.7L13 3l-1.4 1.5"/></g>`
	trelloPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M7 7h3v9H7zm7 0h3v5h-3z"/></g>`
	trendingDownPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 17l-8.5-8.5l-5 5L2 7"/><path d="M16 17h6v-6"/></g>`
	trendingUpPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 7l-8.5 8.5l-5-5L2 17"/><path d="M16 7h6v6"/></g>`
	trianglePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.73 18l-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z"/>`
	triangleRightPath                   = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 18a2 2 0 0 1-2 2H3c-1.1 0-1.3-.6-.4-1.3L20.4 4.3c.9-.7 1.6-.4 1.6.7Z"/>`
	trophyPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9H4.5a2.5 2.5 0 0 1 0-5H6m12 5h1.5a2.5 2.5 0 0 0 0-5H18M4 22h16m-10-7.34V17c0 .55-.47.98-.97 1.21C7.85 18.75 7 20.24 7 22m7-7.34V17c0 .55.47.98.97 1.21C16.15 18.75 17 20.24 17 22"/><path d="M18 2H6v7a6 6 0 0 0 12 0V2Z"/></g>`
	truckPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 18H3c-.6 0-1-.4-1-1V7c0-.6.4-1 1-1h10c.6 0 1 .4 1 1v11m0-9h4l4 4v4c0 .6-.4 1-1 1h-2"/><circle cx="7" cy="18" r="2"/><path d="M15 18H9"/><circle cx="17" cy="18" r="2"/></g>`
	turtlePath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 10l2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a8 8 0 1 0-16 0v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3l2-4h4ZM4.82 7.9L8 10m7.18-2.1L12 10"/><path d="M16.93 10H20a2 2 0 0 1 0 4H2"/></g>`
	tvPath                              = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="15" x="2" y="7" rx="2" ry="2"/><path d="m17 2l-5 5l-5-5"/></g>`
	tvTwoPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21h10"/><rect width="20" height="14" x="2" y="3" rx="2"/></g>`
	twitchPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 2H3v16h5v4l4-4h5l4-4V2zm-10 9V7m5 4V7"/>`
	twitterPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6c2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4c-.9-4.2 4-6.6 7-3.8c1.1 0 3-1.2 3-1.2z"/>`
	typePath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7V4h16v3M9 20h6M12 4v16"/>`
	umbrellaPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12a10.06 10.06 1 0 0-20 0Zm-10 0v8a2 2 0 0 0 4 0M12 2v1"/>`
	underlinePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4v6a6 6 0 0 0 12 0V4M4 20h16"/>`
	undoPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7v6h6"/><path d="M21 17a9 9 0 0 0-9-9a9 9 0 0 0-6 2.3L3 13"/></g>`
	undoDotPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="17" r="1"/><path d="M3 7v6h6"/><path d="M21 17a9 9 0 0 0-9-9a9 9 0 0 0-6 2.3L3 13"/></g>`
	undoTwoPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 14L4 9l5-5"/><path d="M4 9h10.5a5.5 5.5 0 0 1 5.5 5.5v0a5.5 5.5 0 0 1-5.5 5.5H11"/></g>`
	unfoldHorizontalPath                = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h6M8 12H2M12 2v2m0 4v2m0 4v2m0 4v2m7-7l3-3l-3-3M5 9l-3 3l3 3"/>`
	unfoldVerticalPath                  = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22v-6m0-8V2M4 12H2m8 0H8m8 0h-2m8 0h-2m-5 7l-3 3l-3-3m6-14l-3-3l-3 3"/>`
	ungroupPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="6" x="5" y="4" rx="1"/><rect width="8" height="6" x="11" y="14" rx="1"/></g>`
	unlinkPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18.84 12.25l1.72-1.71h-.02a5.004 5.004 0 0 0-.12-7.07a5.006 5.006 0 0 0-6.95 0l-1.72 1.71m-6.58 6.57l-1.71 1.71a5.004 5.004 0 0 0 .12 7.07a5.006 5.006 0 0 0 6.95 0l1.71-1.71M8 2v3M2 8h3m11 11v3m3-6h3"/>`
	unlinkTwoPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7h2a5 5 0 0 1 0 10h-2m-6 0H7A5 5 0 0 1 7 7h2"/>`
	unlockPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 9.9-1"/></g>`
	unplugPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 5l3-3M2 22l3-3m1.3 1.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6l-2.3 2.3a2.4 2.4 0 0 0 0 3.4Zm1.2-6.8L10 11m.5 5.5L13 14m-1-8l6 6l2.3-2.3a2.4 2.4 0 0 0 0-3.4l-2.6-2.6a2.4 2.4 0 0 0-3.4 0Z"/>`
	uploadPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4m14-7l-5-5l-5 5m5-5v12"/>`
	uploadCloudPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M12 12v9"/><path d="m16 16l-4-4l-4 4"/></g>`
	usbPath                             = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="10" cy="7" r="1"/><circle cx="4" cy="20" r="1"/><path d="M4.7 19.3L19 5m2-2l-3 1l2 2ZM9.26 7.68L5 12l2 5m3-3l5 2l3.5-3.5"/><path d="m18 12l1-1l1 1l-1 1Z"/></g>`
	userPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></g>`
	userCheckPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="m16 11l2 2l4-4"/></g>`
	userCheckTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="m16 11l2 2l4-4"/></g>`
	userCirclePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="10" r="3"/><path d="M7 20.662V19a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.662"/></g>`
	userCircleTwoPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 20a6 6 0 0 0-12 0"/><circle cx="12" cy="10" r="4"/><circle cx="12" cy="12" r="10"/></g>`
	userCogPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="15" r="3"/><circle cx="9" cy="7" r="4"/><path d="M10 15H6a4 4 0 0 0-4 4v2m19.7-4.6l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`
	userCogTwoPath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="15" r="3"/><circle cx="8" cy="9" r="4"/><path d="M10.5 13.5A6 6 0 0 0 2 19m19.7-2.6l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`
	userMinusPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 11h-6"/></g>`
	userMinusTwoPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="M22 11h-6"/></g>`
	userPlusPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M19 8v6m3-3h-6"/></g>`
	userPlusTwoPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="M19 8v6m3-3h-6"/></g>`
	userSquarePath                      = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="10" r="3"/><path d="M7 21v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"/></g>`
	userSquareTwoPath                   = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 21a6 6 0 0 0-12 0"/><circle cx="12" cy="11" r="4"/><rect width="18" height="18" x="3" y="3" rx="2"/></g>`
	userTwoPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="5"/><path d="M20 21a8 8 0 1 0-16 0"/></g>`
	userXPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="m17 8l5 5m0-5l-5 5"/></g>`
	userXTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="m17 8l5 5m0-5l-5 5"/></g>`
	usersPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87m-3-12a4 4 0 0 1 0 7.75"/></g>`
	usersTwoPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="M22 19a6 6 0 0 0-6-6a4 4 0 1 0 0-8"/></g>`
	utensilsPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 2v7c0 1.1.9 2 2 2h4a2 2 0 0 0 2-2V2M7 2v20m14-7V2v0a5 5 0 0 0-5 5v6c0 1.1.9 2 2 2h3Zm0 0v7"/>`
	utensilsCrossedPath                 = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 2l-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8m-7 7L3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0l7 7m-19.9-.2l6.4-6.3M19 5l-7 7"/>`
	utilityPolePath                     = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20M2 5h20M3 3v2m4-2v2m10-2v2m4-2v2m-2 0l-7 7l-7-7"/>`
	variablePath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 21s-4-3-4-9s4-9 4-9m8 0s4 3 4 9s-4 9-4 9M15 9l-6 6m0-6l6 6"/>`
	veganPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14m0 0c4 0 6-2 6-6c-4 0-6 2-6 6"/><path d="M17.41 3.6a10 10 0 1 0 3 3"/></g>`
	venetianMaskPath                    = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12a5 5 0 0 0 5 5a8 8 0 0 1 5 2a8 8 0 0 1 5-2a5 5 0 0 0 5-5V7h-5a8 8 0 0 0-5 2a8 8 0 0 0-5-2H2Z"/><path d="M6 11c1.5 0 3 .5 3 2c-2 0-3 0-3-2Zm12 0c-1.5 0-3 .5-3 2c2 0 3 0 3-2Z"/></g>`
	verifiedPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3c-1.2 0-2.4.6-3 1.7A3.6 3.6 0 0 0 4.6 9c-1 .6-1.7 1.8-1.7 3s.7 2.4 1.7 3c-.3 1.2 0 2.5 1 3.4c.8.8 2.1 1.2 3.3 1c.6 1 1.8 1.6 3 1.6s2.4-.6 3-1.7c1.2.3 2.5 0 3.4-1c.8-.8 1.2-2 1-3.3c1-.6 1.6-1.8 1.6-3s-.6-2.4-1.7-3c.3-1.2 0-2.5-1-3.4a3.7 3.7 0 0 0-3.3-1c-.6-1-1.8-1.6-3-1.6Z"/><path d="m9 12l2 2l4-4"/></g>`
	vibratePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 8l2 2l-2 2l2 2l-2 2m20-8l-2 2l2 2l-2 2l2 2"/><rect width="8" height="14" x="8" y="5" rx="1"/></g>`
	vibrateOffPath                      = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 8l2 2l-2 2l2 2l-2 2m20-8l-2 2l2 2l-2 2l2 2M8 8v10c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2m0-5.66V6c0-.55-.45-1-1-1h-4.34M2 2l20 20"/>`
	videoPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 8l-6 4l6 4V8Z"/><rect width="14" height="12" x="2" y="6" rx="2" ry="2"/></g>`
	videoOffPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.66 6H14a2 2 0 0 1 2 2v2.34l1 1L22 8v8m-6 0a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2l10 10ZM2 2l20 20"/>`
	videotapePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M2 8h20"/><circle cx="8" cy="14" r="2"/><path d="M8 12h8"/><circle cx="16" cy="14" r="2"/></g>`
	viewPath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12s2.545-5 7-5c4.454 0 7 5 7 5s-2.546 5-7 5c-4.455 0-7-5-7-5z"/><path d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm9 4v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2M21 7V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2"/></g>`
	voicemailPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="12" r="4"/><circle cx="18" cy="12" r="4"/><path d="M6 16h12"/></g>`
	volumePath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5z"/>`
	volumeOnePath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm4.54 3.46a5 5 0 0 1 0 7.07"/>`
	volumeTwoPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm4.54 3.46a5 5 0 0 1 0 7.07m3.53-10.6a10 10 0 0 1 0 14.14"/>`
	volumeXPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm11 4l-6 6m0-6l6 6"/>`
	votePath                            = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 12l2 2l4-4"/><path d="M5 7c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v12H5V7Zm17 12H2"/></g>`
	walletPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"/><path d="M3 5v14a2 2 0 0 0 2 2h16v-5"/><path d="M18 12a2 2 0 0 0 0 4h4v-4Z"/></g>`
	walletCardsPath                     = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2M3 11h3c.8 0 1.6.3 2.1.9l1.1.9c1.6 1.6 4.1 1.6 5.7 0l1.1-.9c.5-.5 1.3-.9 2.1-.9H21"/></g>`
	walletTwoPath                       = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 14h.01M7 7h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14"/>`
	wallpaperPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="9" r="2"/><path d="m9 17l6.1-6.1a2 2 0 0 1 2.81.01L22 15V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2M8 21h8m-4-4v4"/></g>`
	wandPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4V2m0 14v-2M8 9h2m10 0h2m-4.2 2.8L19 13m-4-4h0m2.8-2.8L19 5M3 21l9-9m.2-5.8L11 5"/>`
	wandTwoPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.64 3.64l-1.28-1.28a1.21 1.21 0 0 0-1.72 0L2.36 18.64a1.21 1.21 0 0 0 0 1.72l1.28 1.28a1.2 1.2 0 0 0 1.72 0L21.64 5.36a1.2 1.2 0 0 0 0-1.72ZM14 7l3 3M5 6v4m14 4v4M10 2v2M7 8H3m18 8h-4M11 3H9"/>`
	warehousePath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 8.35V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8.35A2 2 0 0 1 3.26 6.5l8-3.2a2 2 0 0 1 1.48 0l8 3.2A2 2 0 0 1 22 8.35ZM6 18h12M6 14h12"/><path d="M6 10h12v12H6z"/></g>`
	watchPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="6"/><path d="M12 10v2l1 1m3.13-5.34l-.81-4.05a2 2 0 0 0-2-1.61h-2.68a2 2 0 0 0-2 1.61l-.78 4.05m.02 8.7l.8 4a2 2 0 0 0 2 1.61h2.72a2 2 0 0 0 2-1.61l.81-4.05"/></g>`
	wavesPath                           = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 6c.6.5 1.2 1 2.5 1C7 7 7 5 9.5 5c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1M2 12c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1M2 18c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/>`
	webcamPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="8"/><circle cx="12" cy="10" r="3"/><path d="M7 22h10m-5 0v-4"/></g>`
	webhookPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 16.98h-5.99c-1.1 0-1.95.94-2.48 1.9A4 4 0 0 1 2 17c.01-.7.2-1.4.57-2"/><path d="m6 17l3.13-5.78c.53-.97.1-2.18-.5-3.1a4 4 0 1 1 6.89-4.06"/><path d="m12 6l3.13 5.73C15.66 12.7 16.9 13 18 13a4 4 0 0 1 0 8"/></g>`
	wheatPath                           = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 22L16 8M3.47 12.53L5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Zm4-4L9 7l1.53 1.53a3.5 3.5 0 0 1 0 4.94L9 15l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Zm4-4L13 3l1.53 1.53a3.5 3.5 0 0 1 0 4.94L13 11l-1.53-1.53a3.5 3.5 0 0 1 0-4.94ZM20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"/><path d="M11.47 17.47L13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0Zm4-4L17 15l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.5 3.5 0 0 1 4.94 0Zm4-4L21 11l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L13 11l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"/></g>`
	wheatOffPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 22l10-10m4-4l-1.17 1.17M3.47 12.53L5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94ZM8 8l-.53.53a3.5 3.5 0 0 0 0 4.94L9 15l1.53-1.53c.55-.55.88-1.25.98-1.97m-.6-6.24c.15-.26.34-.51.56-.73L13 3l1.53 1.53a3.5 3.5 0 0 1 .28 4.62M20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"/><path d="M11.47 17.47L13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0ZM16 16l-.53.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.49 3.49 0 0 1 1.97-.98m6.24.6c.26-.15.51-.34.73-.56L21 11l-1.53-1.53a3.5 3.5 0 0 0-4.62-.28M2 2l20 20"/></g>`
	wholeWordPath                       = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="12" r="3"/><path d="M10 9v6"/><circle cx="17" cy="12" r="3"/><path d="M14 7v8m8 2v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"/></g>`
	wifiPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13a10 10 0 0 1 14 0M8.5 16.5a5 5 0 0 1 7 0M2 8.82a15 15 0 0 1 20 0M12 20h.01"/>`
	wifiOffPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M8.5 16.5a5 5 0 0 1 7 0M2 8.82a15 15 0 0 1 4.17-2.65M10.66 5c4.01-.36 8.14.9 11.34 3.76m-5.15 2.49a10 10 0 0 1 2.22 1.68M5 13a10 10 0 0 1 5.24-2.76M12 20h.01"/>`
	windPath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.7 7.7a2.5 2.5 0 1 1 1.8 4.3H2m7.6-7.4A2 2 0 1 1 11 8H2m10.6 11.4A2 2 0 1 0 14 16H2"/>`
	winePath                            = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22h8M7 10h10m-5 5v7m0-7a5 5 0 0 0 5-5c0-2-.5-4-2-8H9c-1.5 4-2 6-2 8a5 5 0 0 0 5 5Z"/>`
	wineOffPath                         = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22h8M7 10h3m7 0h-1.343M12 15v7M7.307 7.307A12.33 12.33 0 0 0 7 10a5 5 0 0 0 7.391 4.391M8.638 2.981C8.75 2.668 8.872 2.34 9 2h6c1.5 4 2 6 2 8c0 .407-.05.809-.145 1.198M2 2l20 20"/>`
	workflowPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="3" y="3" rx="2"/><path d="M7 11v4a2 2 0 0 0 2 2h4"/><rect width="8" height="8" x="13" y="13" rx="2"/></g>`
	wrapTextPath                        = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 6h18M3 12h15a3 3 0 1 1 0 6h-4"/><path d="m16 16l-2 2l2 2M3 18h7"/></g>`
	wrenchPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>`
	xPath                               = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6L6 18M6 6l12 12"/>`
	xCirclePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m15 9l-6 6m0-6l6 6"/></g>`
	xOctagonPath                        = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2zM15 9l-6 6m0-6l6 6"/>`
	xSquarePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="m15 9l-6 6m0-6l6 6"/></g>`
	youtubePath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.5 17a24.12 24.12 0 0 1 0-10a2 2 0 0 1 1.4-1.4a49.56 49.56 0 0 1 16.2 0A2 2 0 0 1 21.5 7a24.12 24.12 0 0 1 0 10a2 2 0 0 1-1.4 1.4a49.55 49.55 0 0 1-16.2 0A2 2 0 0 1 2.5 17"/><path d="m10 15l5-3l-5-3z"/></g>`
	zapPath                             = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 2L3 14h9l-1 8l10-12h-9l1-8z"/>`
	zapOffPath                          = `<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.41 6.75L13 2l-2.43 2.92m8 7.99L21 10h-5.34M8 8l-5 6h9l-1 8l5-6M2 2l20 20"/>`
	zoomInPath                          = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35M11 8v6m-3-3h6"/></g>`
	zoomOutPath                         = `<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35M8 11h6"/></g>`
)

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func IconFromName(name string) g.Node {
	switch name {
	case "accessibility":
		return Accessibility()
	case "activity":
		return Activity()
	case "activitySquare":
		return ActivitySquare()
	case "airVent":
		return AirVent()
	case "airplay":
		return Airplay()
	case "alarmCheck":
		return AlarmCheck()
	case "alarmClock":
		return AlarmClock()
	case "alarmClockOff":
		return AlarmClockOff()
	case "alarmMinus":
		return AlarmMinus()
	case "alarmPlus":
		return AlarmPlus()
	case "album":
		return Album()
	case "alertCircle":
		return AlertCircle()
	case "alertOctagon":
		return AlertOctagon()
	case "alertTriangle":
		return AlertTriangle()
	case "alignCenter":
		return AlignCenter()
	case "alignCenterHorizontal":
		return AlignCenterHorizontal()
	case "alignCenterVertical":
		return AlignCenterVertical()
	case "alignEndHorizontal":
		return AlignEndHorizontal()
	case "alignEndVertical":
		return AlignEndVertical()
	case "alignHorizontalDistributeCenter":
		return AlignHorizontalDistributeCenter()
	case "alignHorizontalDistributeEnd":
		return AlignHorizontalDistributeEnd()
	case "alignHorizontalDistributeStart":
		return AlignHorizontalDistributeStart()
	case "alignHorizontalJustifyCenter":
		return AlignHorizontalJustifyCenter()
	case "alignHorizontalJustifyEnd":
		return AlignHorizontalJustifyEnd()
	case "alignHorizontalJustifyStart":
		return AlignHorizontalJustifyStart()
	case "alignHorizontalSpaceAround":
		return AlignHorizontalSpaceAround()
	case "alignHorizontalSpaceBetween":
		return AlignHorizontalSpaceBetween()
	case "alignJustify":
		return AlignJustify()
	case "alignLeft":
		return AlignLeft()
	case "alignRight":
		return AlignRight()
	case "alignStartHorizontal":
		return AlignStartHorizontal()
	case "alignStartVertical":
		return AlignStartVertical()
	case "alignVerticalDistributeCenter":
		return AlignVerticalDistributeCenter()
	case "alignVerticalDistributeEnd":
		return AlignVerticalDistributeEnd()
	case "alignVerticalDistributeStart":
		return AlignVerticalDistributeStart()
	case "alignVerticalJustifyCenter":
		return AlignVerticalJustifyCenter()
	case "alignVerticalJustifyEnd":
		return AlignVerticalJustifyEnd()
	case "alignVerticalJustifyStart":
		return AlignVerticalJustifyStart()
	case "alignVerticalSpaceAround":
		return AlignVerticalSpaceAround()
	case "alignVerticalSpaceBetween":
		return AlignVerticalSpaceBetween()
	case "ampersand":
		return Ampersand()
	case "ampersands":
		return Ampersands()
	case "anchor":
		return Anchor()
	case "angry":
		return Angry()
	case "annoyed":
		return Annoyed()
	case "antenna":
		return Antenna()
	case "aperture":
		return Aperture()
	case "appWindow":
		return AppWindow()
	case "apple":
		return Apple()
	case "archive":
		return Archive()
	case "archiveRestore":
		return ArchiveRestore()
	case "archiveX":
		return ArchiveX()
	case "areaChart":
		return AreaChart()
	case "armchair":
		return Armchair()
	case "arrowBigDown":
		return ArrowBigDown()
	case "arrowBigDownDash":
		return ArrowBigDownDash()
	case "arrowBigLeft":
		return ArrowBigLeft()
	case "arrowBigLeftDash":
		return ArrowBigLeftDash()
	case "arrowBigRight":
		return ArrowBigRight()
	case "arrowBigRightDash":
		return ArrowBigRightDash()
	case "arrowBigUp":
		return ArrowBigUp()
	case "arrowBigUpDash":
		return ArrowBigUpDash()
	case "arrowDown":
		return ArrowDown()
	case "arrowDownAZ":
		return ArrowDownAZ()
	case "arrowDownCircle":
		return ArrowDownCircle()
	case "arrowDownFromLine":
		return ArrowDownFromLine()
	case "arrowDownLeft":
		return ArrowDownLeft()
	case "arrowDownLeftFromCircle":
		return ArrowDownLeftFromCircle()
	case "arrowDownLeftSquare":
		return ArrowDownLeftSquare()
	case "arrowDownNarrowWide":
		return ArrowDownNarrowWide()
	case "arrowDownOneZero":
		return ArrowDownOneZero()
	case "arrowDownRight":
		return ArrowDownRight()
	case "arrowDownRightFromCircle":
		return ArrowDownRightFromCircle()
	case "arrowDownRightSquare":
		return ArrowDownRightSquare()
	case "arrowDownSquare":
		return ArrowDownSquare()
	case "arrowDownToDot":
		return ArrowDownToDot()
	case "arrowDownToLine":
		return ArrowDownToLine()
	case "arrowDownUp":
		return ArrowDownUp()
	case "arrowDownWideNarrow":
		return ArrowDownWideNarrow()
	case "arrowDownZA":
		return ArrowDownZA()
	case "arrowDownZeroOne":
		return ArrowDownZeroOne()
	case "arrowLeft":
		return ArrowLeft()
	case "arrowLeftCircle":
		return ArrowLeftCircle()
	case "arrowLeftFromLine":
		return ArrowLeftFromLine()
	case "arrowLeftRight":
		return ArrowLeftRight()
	case "arrowLeftSquare":
		return ArrowLeftSquare()
	case "arrowLeftToLine":
		return ArrowLeftToLine()
	case "arrowRight":
		return ArrowRight()
	case "arrowRightCircle":
		return ArrowRightCircle()
	case "arrowRightFromLine":
		return ArrowRightFromLine()
	case "arrowRightLeft":
		return ArrowRightLeft()
	case "arrowRightSquare":
		return ArrowRightSquare()
	case "arrowRightToLine":
		return ArrowRightToLine()
	case "arrowUp":
		return ArrowUp()
	case "arrowUpAZ":
		return ArrowUpAZ()
	case "arrowUpCircle":
		return ArrowUpCircle()
	case "arrowUpDown":
		return ArrowUpDown()
	case "arrowUpFromDot":
		return ArrowUpFromDot()
	case "arrowUpFromLine":
		return ArrowUpFromLine()
	case "arrowUpLeft":
		return ArrowUpLeft()
	case "arrowUpLeftFromCircle":
		return ArrowUpLeftFromCircle()
	case "arrowUpLeftSquare":
		return ArrowUpLeftSquare()
	case "arrowUpNarrowWide":
		return ArrowUpNarrowWide()
	case "arrowUpOneZero":
		return ArrowUpOneZero()
	case "arrowUpRight":
		return ArrowUpRight()
	case "arrowUpRightFromCircle":
		return ArrowUpRightFromCircle()
	case "arrowUpRightSquare":
		return ArrowUpRightSquare()
	case "arrowUpSquare":
		return ArrowUpSquare()
	case "arrowUpToLine":
		return ArrowUpToLine()
	case "arrowUpWideNarrow":
		return ArrowUpWideNarrow()
	case "arrowUpZA":
		return ArrowUpZA()
	case "arrowUpZeroOne":
		return ArrowUpZeroOne()
	case "arrowsUpFromLine":
		return ArrowsUpFromLine()
	case "asterisk":
		return Asterisk()
	case "atSign":
		return AtSign()
	case "atom":
		return Atom()
	case "award":
		return Award()
	case "axe":
		return Axe()
	case "axisThreeD":
		return AxisThreeD()
	case "baby":
		return Baby()
	case "backpack":
		return Backpack()
	case "badge":
		return Badge()
	case "badgeAlert":
		return BadgeAlert()
	case "badgeCent":
		return BadgeCent()
	case "badgeCheck":
		return BadgeCheck()
	case "badgeDollarSign":
		return BadgeDollarSign()
	case "badgeEuro":
		return BadgeEuro()
	case "badgeHelp":
		return BadgeHelp()
	case "badgeIndianRupee":
		return BadgeIndianRupee()
	case "badgeInfo":
		return BadgeInfo()
	case "badgeJapaneseYen":
		return BadgeJapaneseYen()
	case "badgeMinus":
		return BadgeMinus()
	case "badgePercent":
		return BadgePercent()
	case "badgePlus":
		return BadgePlus()
	case "badgePoundSterling":
		return BadgePoundSterling()
	case "badgeRussianRuble":
		return BadgeRussianRuble()
	case "badgeSwissFranc":
		return BadgeSwissFranc()
	case "badgeX":
		return BadgeX()
	case "baggageClaim":
		return BaggageClaim()
	case "ban":
		return Ban()
	case "banana":
		return Banana()
	case "banknote":
		return Banknote()
	case "barChart":
		return BarChart()
	case "barChartBig":
		return BarChartBig()
	case "barChartFour":
		return BarChartFour()
	case "barChartHorizontal":
		return BarChartHorizontal()
	case "barChartHorizontalBig":
		return BarChartHorizontalBig()
	case "barChartThree":
		return BarChartThree()
	case "barChartTwo":
		return BarChartTwo()
	case "baseline":
		return Baseline()
	case "bath":
		return Bath()
	case "battery":
		return Battery()
	case "batteryCharging":
		return BatteryCharging()
	case "batteryFull":
		return BatteryFull()
	case "batteryLow":
		return BatteryLow()
	case "batteryMedium":
		return BatteryMedium()
	case "batteryWarning":
		return BatteryWarning()
	case "beaker":
		return Beaker()
	case "bean":
		return Bean()
	case "beanOff":
		return BeanOff()
	case "bed":
		return Bed()
	case "bedDouble":
		return BedDouble()
	case "bedSingle":
		return BedSingle()
	case "beef":
		return Beef()
	case "beer":
		return Beer()
	case "bell":
		return Bell()
	case "bellDot":
		return BellDot()
	case "bellMinus":
		return BellMinus()
	case "bellOff":
		return BellOff()
	case "bellPlus":
		return BellPlus()
	case "bellRing":
		return BellRing()
	case "bike":
		return Bike()
	case "binary":
		return Binary()
	case "biohazard":
		return Biohazard()
	case "bird":
		return Bird()
	case "bitcoin":
		return Bitcoin()
	case "blinds":
		return Blinds()
	case "blocks":
		return Blocks()
	case "bluetooth":
		return Bluetooth()
	case "bluetoothConnected":
		return BluetoothConnected()
	case "bluetoothOff":
		return BluetoothOff()
	case "bluetoothSearching":
		return BluetoothSearching()
	case "bold":
		return Bold()
	case "bomb":
		return Bomb()
	case "bone":
		return Bone()
	case "book":
		return Book()
	case "bookCopy":
		return BookCopy()
	case "bookDown":
		return BookDown()
	case "bookKey":
		return BookKey()
	case "bookLock":
		return BookLock()
	case "bookMarked":
		return BookMarked()
	case "bookMinus":
		return BookMinus()
	case "bookOpen":
		return BookOpen()
	case "bookOpenCheck":
		return BookOpenCheck()
	case "bookPlus":
		return BookPlus()
	case "bookTemplate":
		return BookTemplate()
	case "bookUp":
		return BookUp()
	case "bookUpTwo":
		return BookUpTwo()
	case "bookX":
		return BookX()
	case "bookmark":
		return Bookmark()
	case "bookmarkMinus":
		return BookmarkMinus()
	case "bookmarkPlus":
		return BookmarkPlus()
	case "boomBox":
		return BoomBox()
	case "bot":
		return Bot()
	case "box":
		return Box()
	case "boxSelect":
		return BoxSelect()
	case "boxes":
		return Boxes()
	case "braces":
		return Braces()
	case "brackets":
		return Brackets()
	case "brain":
		return Brain()
	case "brainCircuit":
		return BrainCircuit()
	case "brainCog":
		return BrainCog()
	case "briefcase":
		return Briefcase()
	case "bringToFront":
		return BringToFront()
	case "brush":
		return Brush()
	case "bug":
		return Bug()
	case "bugOff":
		return BugOff()
	case "bugPlay":
		return BugPlay()
	case "building":
		return Building()
	case "buildingTwo":
		return BuildingTwo()
	case "bus":
		return Bus()
	case "busFront":
		return BusFront()
	case "cable":
		return Cable()
	case "cableCar":
		return CableCar()
	case "cake":
		return Cake()
	case "cakeSlice":
		return CakeSlice()
	case "calculator":
		return Calculator()
	case "calendar":
		return Calendar()
	case "calendarCheck":
		return CalendarCheck()
	case "calendarCheckTwo":
		return CalendarCheckTwo()
	case "calendarClock":
		return CalendarClock()
	case "calendarDays":
		return CalendarDays()
	case "calendarHeart":
		return CalendarHeart()
	case "calendarMinus":
		return CalendarMinus()
	case "calendarOff":
		return CalendarOff()
	case "calendarPlus":
		return CalendarPlus()
	case "calendarRange":
		return CalendarRange()
	case "calendarSearch":
		return CalendarSearch()
	case "calendarX":
		return CalendarX()
	case "calendarXTwo":
		return CalendarXTwo()
	case "camera":
		return Camera()
	case "cameraOff":
		return CameraOff()
	case "candlestickChart":
		return CandlestickChart()
	case "candy":
		return Candy()
	case "candyCane":
		return CandyCane()
	case "candyOff":
		return CandyOff()
	case "car":
		return Car()
	case "carFront":
		return CarFront()
	case "carTaxiFront":
		return CarTaxiFront()
	case "carrot":
		return Carrot()
	case "caseLower":
		return CaseLower()
	case "caseSensitive":
		return CaseSensitive()
	case "caseUpper":
		return CaseUpper()
	case "cassetteTape":
		return CassetteTape()
	case "cast":
		return Cast()
	case "castle":
		return Castle()
	case "cat":
		return Cat()
	case "check":
		return Check()
	case "checkCheck":
		return CheckCheck()
	case "checkCircle":
		return CheckCircle()
	case "checkCircleTwo":
		return CheckCircleTwo()
	case "checkSquare":
		return CheckSquare()
	case "chefHat":
		return ChefHat()
	case "cherry":
		return Cherry()
	case "chevronDown":
		return ChevronDown()
	case "chevronDownCircle":
		return ChevronDownCircle()
	case "chevronDownSquare":
		return ChevronDownSquare()
	case "chevronFirst":
		return ChevronFirst()
	case "chevronLast":
		return ChevronLast()
	case "chevronLeft":
		return ChevronLeft()
	case "chevronLeftCircle":
		return ChevronLeftCircle()
	case "chevronLeftSquare":
		return ChevronLeftSquare()
	case "chevronRight":
		return ChevronRight()
	case "chevronRightCircle":
		return ChevronRightCircle()
	case "chevronRightSquare":
		return ChevronRightSquare()
	case "chevronUp":
		return ChevronUp()
	case "chevronUpCircle":
		return ChevronUpCircle()
	case "chevronUpSquare":
		return ChevronUpSquare()
	case "chevronsDown":
		return ChevronsDown()
	case "chevronsDownUp":
		return ChevronsDownUp()
	case "chevronsLeft":
		return ChevronsLeft()
	case "chevronsLeftRight":
		return ChevronsLeftRight()
	case "chevronsRight":
		return ChevronsRight()
	case "chevronsRightLeft":
		return ChevronsRightLeft()
	case "chevronsUp":
		return ChevronsUp()
	case "chevronsUpDown":
		return ChevronsUpDown()
	case "chrome":
		return Chrome()
	case "church":
		return Church()
	case "cigarette":
		return Cigarette()
	case "cigaretteOff":
		return CigaretteOff()
	case "circle":
		return Circle()
	case "circleDashed":
		return CircleDashed()
	case "circleDollarSign":
		return CircleDollarSign()
	case "circleDot":
		return CircleDot()
	case "circleDotDashed":
		return CircleDotDashed()
	case "circleEllipsis":
		return CircleEllipsis()
	case "circleEqual":
		return CircleEqual()
	case "circleOff":
		return CircleOff()
	case "circleSlash":
		return CircleSlash()
	case "circleSlashTwo":
		return CircleSlashTwo()
	case "circuitBoard":
		return CircuitBoard()
	case "citrus":
		return Citrus()
	case "clapperboard":
		return Clapperboard()
	case "clipboard":
		return Clipboard()
	case "clipboardCheck":
		return ClipboardCheck()
	case "clipboardCopy":
		return ClipboardCopy()
	case "clipboardEdit":
		return ClipboardEdit()
	case "clipboardList":
		return ClipboardList()
	case "clipboardPaste":
		return ClipboardPaste()
	case "clipboardSignature":
		return ClipboardSignature()
	case "clipboardType":
		return ClipboardType()
	case "clipboardX":
		return ClipboardX()
	case "clock":
		return Clock()
	case "clockEight":
		return ClockEight()
	case "clockEleven":
		return ClockEleven()
	case "clockFive":
		return ClockFive()
	case "clockFour":
		return ClockFour()
	case "clockNine":
		return ClockNine()
	case "clockOne":
		return ClockOne()
	case "clockSeven":
		return ClockSeven()
	case "clockSix":
		return ClockSix()
	case "clockTen":
		return ClockTen()
	case "clockThree":
		return ClockThree()
	case "clockTwelve":
		return ClockTwelve()
	case "clockTwo":
		return ClockTwo()
	case "cloud":
		return Cloud()
	case "cloudCog":
		return CloudCog()
	case "cloudDrizzle":
		return CloudDrizzle()
	case "cloudFog":
		return CloudFog()
	case "cloudHail":
		return CloudHail()
	case "cloudLightning":
		return CloudLightning()
	case "cloudMoon":
		return CloudMoon()
	case "cloudMoonRain":
		return CloudMoonRain()
	case "cloudOff":
		return CloudOff()
	case "cloudRain":
		return CloudRain()
	case "cloudRainWind":
		return CloudRainWind()
	case "cloudSnow":
		return CloudSnow()
	case "cloudSun":
		return CloudSun()
	case "cloudSunRain":
		return CloudSunRain()
	case "cloudy":
		return Cloudy()
	case "clover":
		return Clover()
	case "club":
		return Club()
	case "code":
		return Code()
	case "codeTwo":
		return CodeTwo()
	case "codepen":
		return Codepen()
	case "codesandbox":
		return Codesandbox()
	case "coffee":
		return Coffee()
	case "cog":
		return Cog()
	case "coins":
		return Coins()
	case "columns":
		return Columns()
	case "combine":
		return Combine()
	case "command":
		return Command()
	case "compass":
		return Compass()
	case "component":
		return Component()
	case "computer":
		return Computer()
	case "conciergeBell":
		return ConciergeBell()
	case "construction":
		return Construction()
	case "contact":
		return Contact()
	case "contactTwo":
		return ContactTwo()
	case "container":
		return Container()
	case "contrast":
		return Contrast()
	case "cookie":
		return Cookie()
	case "copy":
		return Copy()
	case "copyCheck":
		return CopyCheck()
	case "copyMinus":
		return CopyMinus()
	case "copyPlus":
		return CopyPlus()
	case "copySlash":
		return CopySlash()
	case "copyX":
		return CopyX()
	case "copyleft":
		return Copyleft()
	case "copyright":
		return Copyright()
	case "cornerDownLeft":
		return CornerDownLeft()
	case "cornerDownRight":
		return CornerDownRight()
	case "cornerLeftDown":
		return CornerLeftDown()
	case "cornerLeftUp":
		return CornerLeftUp()
	case "cornerRightDown":
		return CornerRightDown()
	case "cornerRightUp":
		return CornerRightUp()
	case "cornerUpLeft":
		return CornerUpLeft()
	case "cornerUpRight":
		return CornerUpRight()
	case "cpu":
		return Cpu()
	case "creativeCommons":
		return CreativeCommons()
	case "creditCard":
		return CreditCard()
	case "croissant":
		return Croissant()
	case "crop":
		return Crop()
	case "cross":
		return Cross()
	case "crosshair":
		return Crosshair()
	case "crown":
		return Crown()
	case "cupSoda":
		return CupSoda()
	case "currency":
		return Currency()
	case "database":
		return Database()
	case "databaseBackup":
		return DatabaseBackup()
	case "databaseZap":
		return DatabaseZap()
	case "delete":
		return Delete()
	case "dessert":
		return Dessert()
	case "diamond":
		return Diamond()
	case "diceFive":
		return DiceFive()
	case "diceFour":
		return DiceFour()
	case "diceOne":
		return DiceOne()
	case "diceSix":
		return DiceSix()
	case "diceThree":
		return DiceThree()
	case "diceTwo":
		return DiceTwo()
	case "dices":
		return Dices()
	case "diff":
		return Diff()
	case "disc":
		return Disc()
	case "discThree":
		return DiscThree()
	case "discTwo":
		return DiscTwo()
	case "divide":
		return Divide()
	case "divideCircle":
		return DivideCircle()
	case "divideSquare":
		return DivideSquare()
	case "dna":
		return Dna()
	case "dnaOff":
		return DnaOff()
	case "dog":
		return Dog()
	case "dollarSign":
		return DollarSign()
	case "donut":
		return Donut()
	case "doorClosed":
		return DoorClosed()
	case "doorOpen":
		return DoorOpen()
	case "dot":
		return Dot()
	case "download":
		return Download()
	case "downloadCloud":
		return DownloadCloud()
	case "dribbble":
		return Dribbble()
	case "droplet":
		return Droplet()
	case "droplets":
		return Droplets()
	case "drumstick":
		return Drumstick()
	case "dumbbell":
		return Dumbbell()
	case "ear":
		return Ear()
	case "earOff":
		return EarOff()
	case "edit":
		return Edit()
	case "editThree":
		return EditThree()
	case "editTwo":
		return EditTwo()
	case "egg":
		return Egg()
	case "eggFried":
		return EggFried()
	case "eggOff":
		return EggOff()
	case "equal":
		return Equal()
	case "equalNot":
		return EqualNot()
	case "eraser":
		return Eraser()
	case "euro":
		return Euro()
	case "expand":
		return Expand()
	case "externalLink":
		return ExternalLink()
	case "eye":
		return Eye()
	case "eyeOff":
		return EyeOff()
	case "facebook":
		return Facebook()
	case "factory":
		return Factory()
	case "fan":
		return Fan()
	case "fastForward":
		return FastForward()
	case "feather":
		return Feather()
	case "ferrisWheel":
		return FerrisWheel()
	case "figma":
		return Figma()
	case "file":
		return File()
	case "fileArchive":
		return FileArchive()
	case "fileAudio":
		return FileAudio()
	case "fileAudioTwo":
		return FileAudioTwo()
	case "fileAxisThreeD":
		return FileAxisThreeD()
	case "fileBadge":
		return FileBadge()
	case "fileBadgeTwo":
		return FileBadgeTwo()
	case "fileBarChart":
		return FileBarChart()
	case "fileBarChartTwo":
		return FileBarChartTwo()
	case "fileBox":
		return FileBox()
	case "fileCheck":
		return FileCheck()
	case "fileCheckTwo":
		return FileCheckTwo()
	case "fileClock":
		return FileClock()
	case "fileCode":
		return FileCode()
	case "fileCodeTwo":
		return FileCodeTwo()
	case "fileCog":
		return FileCog()
	case "fileCogTwo":
		return FileCogTwo()
	case "fileDiff":
		return FileDiff()
	case "fileDigit":
		return FileDigit()
	case "fileDown":
		return FileDown()
	case "fileEdit":
		return FileEdit()
	case "fileHeart":
		return FileHeart()
	case "fileImage":
		return FileImage()
	case "fileInput":
		return FileInput()
	case "fileJson":
		return FileJson()
	case "fileJsonTwo":
		return FileJsonTwo()
	case "fileKey":
		return FileKey()
	case "fileKeyTwo":
		return FileKeyTwo()
	case "fileLineChart":
		return FileLineChart()
	case "fileLock":
		return FileLock()
	case "fileLockTwo":
		return FileLockTwo()
	case "fileMinus":
		return FileMinus()
	case "fileMinusTwo":
		return FileMinusTwo()
	case "fileOutput":
		return FileOutput()
	case "filePieChart":
		return FilePieChart()
	case "filePlus":
		return FilePlus()
	case "filePlusTwo":
		return FilePlusTwo()
	case "fileQuestion":
		return FileQuestion()
	case "fileScan":
		return FileScan()
	case "fileSearch":
		return FileSearch()
	case "fileSearchTwo":
		return FileSearchTwo()
	case "fileSignature":
		return FileSignature()
	case "fileSpreadsheet":
		return FileSpreadsheet()
	case "fileStack":
		return FileStack()
	case "fileSymlink":
		return FileSymlink()
	case "fileTerminal":
		return FileTerminal()
	case "fileText":
		return FileText()
	case "fileType":
		return FileType()
	case "fileTypeTwo":
		return FileTypeTwo()
	case "fileUp":
		return FileUp()
	case "fileVideo":
		return FileVideo()
	case "fileVideoTwo":
		return FileVideoTwo()
	case "fileVolume":
		return FileVolume()
	case "fileVolumeTwo":
		return FileVolumeTwo()
	case "fileWarning":
		return FileWarning()
	case "fileX":
		return FileX()
	case "fileXTwo":
		return FileXTwo()
	case "files":
		return Files()
	case "film":
		return Film()
	case "filter":
		return Filter()
	case "filterX":
		return FilterX()
	case "fingerprint":
		return Fingerprint()
	case "fish":
		return Fish()
	case "fishOff":
		return FishOff()
	case "fishSymbol":
		return FishSymbol()
	case "flag":
		return Flag()
	case "flagOff":
		return FlagOff()
	case "flagTriangleLeft":
		return FlagTriangleLeft()
	case "flagTriangleRight":
		return FlagTriangleRight()
	case "flame":
		return Flame()
	case "flashlight":
		return Flashlight()
	case "flashlightOff":
		return FlashlightOff()
	case "flaskConical":
		return FlaskConical()
	case "flaskConicalOff":
		return FlaskConicalOff()
	case "flaskRound":
		return FlaskRound()
	case "flipHorizontal":
		return FlipHorizontal()
	case "flipHorizontalTwo":
		return FlipHorizontalTwo()
	case "flipVertical":
		return FlipVertical()
	case "flipVerticalTwo":
		return FlipVerticalTwo()
	case "flower":
		return Flower()
	case "flowerTwo":
		return FlowerTwo()
	case "focus":
		return Focus()
	case "foldHorizontal":
		return FoldHorizontal()
	case "foldVertical":
		return FoldVertical()
	case "folder":
		return Folder()
	case "folderArchive":
		return FolderArchive()
	case "folderCheck":
		return FolderCheck()
	case "folderClock":
		return FolderClock()
	case "folderClosed":
		return FolderClosed()
	case "folderCog":
		return FolderCog()
	case "folderCogTwo":
		return FolderCogTwo()
	case "folderDot":
		return FolderDot()
	case "folderDown":
		return FolderDown()
	case "folderEdit":
		return FolderEdit()
	case "folderGit":
		return FolderGit()
	case "folderGitTwo":
		return FolderGitTwo()
	case "folderHeart":
		return FolderHeart()
	case "folderInput":
		return FolderInput()
	case "folderKanban":
		return FolderKanban()
	case "folderKey":
		return FolderKey()
	case "folderLock":
		return FolderLock()
	case "folderMinus":
		return FolderMinus()
	case "folderOpen":
		return FolderOpen()
	case "folderOpenDot":
		return FolderOpenDot()
	case "folderOutput":
		return FolderOutput()
	case "folderPlus":
		return FolderPlus()
	case "folderRoot":
		return FolderRoot()
	case "folderSearch":
		return FolderSearch()
	case "folderSearchTwo":
		return FolderSearchTwo()
	case "folderSymlink":
		return FolderSymlink()
	case "folderSync":
		return FolderSync()
	case "folderTree":
		return FolderTree()
	case "folderUp":
		return FolderUp()
	case "folderX":
		return FolderX()
	case "folders":
		return Folders()
	case "footprints":
		return Footprints()
	case "forklift":
		return Forklift()
	case "formInput":
		return FormInput()
	case "forward":
		return Forward()
	case "frame":
		return Frame()
	case "framer":
		return Framer()
	case "frown":
		return Frown()
	case "fuel":
		return Fuel()
	case "functionSquare":
		return FunctionSquare()
	case "galleryHorizontal":
		return GalleryHorizontal()
	case "galleryHorizontalEnd":
		return GalleryHorizontalEnd()
	case "galleryThumbnails":
		return GalleryThumbnails()
	case "galleryVertical":
		return GalleryVertical()
	case "galleryVerticalEnd":
		return GalleryVerticalEnd()
	case "gamepad":
		return Gamepad()
	case "gamepadTwo":
		return GamepadTwo()
	case "ganttChart":
		return GanttChart()
	case "ganttChartSquare":
		return GanttChartSquare()
	case "gauge":
		return Gauge()
	case "gaugeCircle":
		return GaugeCircle()
	case "gavel":
		return Gavel()
	case "gem":
		return Gem()
	case "ghost":
		return Ghost()
	case "gift":
		return Gift()
	case "gitBranch":
		return GitBranch()
	case "gitBranchPlus":
		return GitBranchPlus()
	case "gitCommit":
		return GitCommit()
	case "gitCompare":
		return GitCompare()
	case "gitFork":
		return GitFork()
	case "gitMerge":
		return GitMerge()
	case "gitPullRequest":
		return GitPullRequest()
	case "gitPullRequestClosed":
		return GitPullRequestClosed()
	case "gitPullRequestDraft":
		return GitPullRequestDraft()
	case "github":
		return Github()
	case "gitlab":
		return Gitlab()
	case "glassWater":
		return GlassWater()
	case "glasses":
		return Glasses()
	case "globe":
		return Globe()
	case "globeTwo":
		return GlobeTwo()
	case "goal":
		return Goal()
	case "grab":
		return Grab()
	case "graduationCap":
		return GraduationCap()
	case "grape":
		return Grape()
	case "grid":
		return Grid()
	case "gridThreeXThree":
		return GridThreeXThree()
	case "gridTwoXTwo":
		return GridTwoXTwo()
	case "grip":
		return Grip()
	case "gripHorizontal":
		return GripHorizontal()
	case "gripVertical":
		return GripVertical()
	case "group":
		return Group()
	case "hammer":
		return Hammer()
	case "hand":
		return Hand()
	case "handMetal":
		return HandMetal()
	case "hardDrive":
		return HardDrive()
	case "hardDriveDownload":
		return HardDriveDownload()
	case "hardDriveUpload":
		return HardDriveUpload()
	case "hardHat":
		return HardHat()
	case "hash":
		return Hash()
	case "haze":
		return Haze()
	case "hdmiPort":
		return HdmiPort()
	case "heading":
		return Heading()
	case "headingFive":
		return HeadingFive()
	case "headingFour":
		return HeadingFour()
	case "headingOne":
		return HeadingOne()
	case "headingSix":
		return HeadingSix()
	case "headingThree":
		return HeadingThree()
	case "headingTwo":
		return HeadingTwo()
	case "headphones":
		return Headphones()
	case "heart":
		return Heart()
	case "heartCrack":
		return HeartCrack()
	case "heartHandshake":
		return HeartHandshake()
	case "heartOff":
		return HeartOff()
	case "heartPulse":
		return HeartPulse()
	case "helpCircle":
		return HelpCircle()
	case "helpingHand":
		return HelpingHand()
	case "hexagon":
		return Hexagon()
	case "highlighter":
		return Highlighter()
	case "history":
		return History()
	case "home":
		return Home()
	case "hop":
		return Hop()
	case "hopOff":
		return HopOff()
	case "hotel":
		return Hotel()
	case "hourglass":
		return Hourglass()
	case "iceCream":
		return IceCream()
	case "iceCreamTwo":
		return IceCreamTwo()
	case "image":
		return Image()
	case "imageMinus":
		return ImageMinus()
	case "imageOff":
		return ImageOff()
	case "imagePlus":
		return ImagePlus()
	case "import":
		return Import()
	case "inbox":
		return Inbox()
	case "indent":
		return Indent()
	case "indianRupee":
		return IndianRupee()
	case "infinity":
		return Infinity()
	case "info":
		return Info()
	case "instagram":
		return Instagram()
	case "italic":
		return Italic()
	case "iterationCcw":
		return IterationCcw()
	case "iterationCw":
		return IterationCw()
	case "japaneseYen":
		return JapaneseYen()
	case "joystick":
		return Joystick()
	case "kanban":
		return Kanban()
	case "kanbanSquare":
		return KanbanSquare()
	case "kanbanSquareDashed":
		return KanbanSquareDashed()
	case "key":
		return Key()
	case "keyRound":
		return KeyRound()
	case "keySquare":
		return KeySquare()
	case "keyboard":
		return Keyboard()
	case "lamp":
		return Lamp()
	case "lampCeiling":
		return LampCeiling()
	case "lampDesk":
		return LampDesk()
	case "lampFloor":
		return LampFloor()
	case "lampWallDown":
		return LampWallDown()
	case "lampWallUp":
		return LampWallUp()
	case "landmark":
		return Landmark()
	case "languages":
		return Languages()
	case "laptop":
		return Laptop()
	case "laptopTwo":
		return LaptopTwo()
	case "lasso":
		return Lasso()
	case "lassoSelect":
		return LassoSelect()
	case "laugh":
		return Laugh()
	case "layers":
		return Layers()
	case "layout":
		return Layout()
	case "layoutDashboard":
		return LayoutDashboard()
	case "layoutGrid":
		return LayoutGrid()
	case "layoutList":
		return LayoutList()
	case "layoutPanelLeft":
		return LayoutPanelLeft()
	case "layoutPanelTop":
		return LayoutPanelTop()
	case "layoutTemplate":
		return LayoutTemplate()
	case "leaf":
		return Leaf()
	case "leafyGreen":
		return LeafyGreen()
	case "library":
		return Library()
	case "lifeBuoy":
		return LifeBuoy()
	case "ligature":
		return Ligature()
	case "lightbulb":
		return Lightbulb()
	case "lightbulbOff":
		return LightbulbOff()
	case "lineChart":
		return LineChart()
	case "link":
		return Link()
	case "linkTwo":
		return LinkTwo()
	case "linkTwoOff":
		return LinkTwoOff()
	case "linkedin":
		return Linkedin()
	case "list":
		return List()
	case "listChecks":
		return ListChecks()
	case "listEnd":
		return ListEnd()
	case "listFilter":
		return ListFilter()
	case "listMinus":
		return ListMinus()
	case "listMusic":
		return ListMusic()
	case "listOrdered":
		return ListOrdered()
	case "listPlus":
		return ListPlus()
	case "listRestart":
		return ListRestart()
	case "listStart":
		return ListStart()
	case "listTodo":
		return ListTodo()
	case "listTree":
		return ListTree()
	case "listVideo":
		return ListVideo()
	case "listX":
		return ListX()
	case "loader":
		return Loader()
	case "loaderTwo":
		return LoaderTwo()
	case "locate":
		return Locate()
	case "locateFixed":
		return LocateFixed()
	case "locateOff":
		return LocateOff()
	case "lock":
		return Lock()
	case "logIn":
		return LogIn()
	case "logOut":
		return LogOut()
	case "lollipop":
		return Lollipop()
	case "luggage":
		return Luggage()
	case "mSquare":
		return MSquare()
	case "magnet":
		return Magnet()
	case "mail":
		return Mail()
	case "mailCheck":
		return MailCheck()
	case "mailMinus":
		return MailMinus()
	case "mailOpen":
		return MailOpen()
	case "mailPlus":
		return MailPlus()
	case "mailQuestion":
		return MailQuestion()
	case "mailSearch":
		return MailSearch()
	case "mailWarning":
		return MailWarning()
	case "mailX":
		return MailX()
	case "mailbox":
		return Mailbox()
	case "mails":
		return Mails()
	case "map":
		return Map()
	case "mapPin":
		return MapPin()
	case "mapPinOff":
		return MapPinOff()
	case "martini":
		return Martini()
	case "maximize":
		return Maximize()
	case "maximizeTwo":
		return MaximizeTwo()
	case "medal":
		return Medal()
	case "megaphone":
		return Megaphone()
	case "megaphoneOff":
		return MegaphoneOff()
	case "meh":
		return Meh()
	case "memoryStick":
		return MemoryStick()
	case "menu":
		return Menu()
	case "menuSquare":
		return MenuSquare()
	case "merge":
		return Merge()
	case "messageCircle":
		return MessageCircle()
	case "messageSquare":
		return MessageSquare()
	case "messageSquareDashed":
		return MessageSquareDashed()
	case "messageSquarePlus":
		return MessageSquarePlus()
	case "messagesSquare":
		return MessagesSquare()
	case "mic":
		return Mic()
	case "micOff":
		return MicOff()
	case "micTwo":
		return MicTwo()
	case "microscope":
		return Microscope()
	case "microwave":
		return Microwave()
	case "milestone":
		return Milestone()
	case "milk":
		return Milk()
	case "milkOff":
		return MilkOff()
	case "minimize":
		return Minimize()
	case "minimizeTwo":
		return MinimizeTwo()
	case "minus":
		return Minus()
	case "minusCircle":
		return MinusCircle()
	case "minusSquare":
		return MinusSquare()
	case "monitor":
		return Monitor()
	case "monitorCheck":
		return MonitorCheck()
	case "monitorDot":
		return MonitorDot()
	case "monitorDown":
		return MonitorDown()
	case "monitorOff":
		return MonitorOff()
	case "monitorPause":
		return MonitorPause()
	case "monitorPlay":
		return MonitorPlay()
	case "monitorSmartphone":
		return MonitorSmartphone()
	case "monitorSpeaker":
		return MonitorSpeaker()
	case "monitorStop":
		return MonitorStop()
	case "monitorUp":
		return MonitorUp()
	case "monitorX":
		return MonitorX()
	case "moon":
		return Moon()
	case "moonStar":
		return MoonStar()
	case "moreHorizontal":
		return MoreHorizontal()
	case "moreVertical":
		return MoreVertical()
	case "mountain":
		return Mountain()
	case "mountainSnow":
		return MountainSnow()
	case "mouse":
		return Mouse()
	case "mousePointer":
		return MousePointer()
	case "mousePointerClick":
		return MousePointerClick()
	case "mousePointerSquare":
		return MousePointerSquare()
	case "mousePointerSquareDashed":
		return MousePointerSquareDashed()
	case "mousePointerTwo":
		return MousePointerTwo()
	case "move":
		return Move()
	case "moveDiagonal":
		return MoveDiagonal()
	case "moveDiagonalTwo":
		return MoveDiagonalTwo()
	case "moveDown":
		return MoveDown()
	case "moveDownLeft":
		return MoveDownLeft()
	case "moveDownRight":
		return MoveDownRight()
	case "moveHorizontal":
		return MoveHorizontal()
	case "moveLeft":
		return MoveLeft()
	case "moveRight":
		return MoveRight()
	case "moveThreeD":
		return MoveThreeD()
	case "moveUp":
		return MoveUp()
	case "moveUpLeft":
		return MoveUpLeft()
	case "moveUpRight":
		return MoveUpRight()
	case "moveVertical":
		return MoveVertical()
	case "music":
		return Music()
	case "musicFour":
		return MusicFour()
	case "musicThree":
		return MusicThree()
	case "musicTwo":
		return MusicTwo()
	case "navigation":
		return Navigation()
	case "navigationOff":
		return NavigationOff()
	case "navigationTwo":
		return NavigationTwo()
	case "navigationTwoOff":
		return NavigationTwoOff()
	case "network":
		return Network()
	case "newspaper":
		return Newspaper()
	case "nfc":
		return Nfc()
	case "nut":
		return Nut()
	case "nutOff":
		return NutOff()
	case "octagon":
		return Octagon()
	case "option":
		return Option()
	case "orbit":
		return Orbit()
	case "outdent":
		return Outdent()
	case "package":
		return Package()
	case "packageCheck":
		return PackageCheck()
	case "packageMinus":
		return PackageMinus()
	case "packageOpen":
		return PackageOpen()
	case "packagePlus":
		return PackagePlus()
	case "packageSearch":
		return PackageSearch()
	case "packageTwo":
		return PackageTwo()
	case "packageX":
		return PackageX()
	case "paintBucket":
		return PaintBucket()
	case "paintbrush":
		return Paintbrush()
	case "paintbrushTwo":
		return PaintbrushTwo()
	case "palette":
		return Palette()
	case "palmtree":
		return Palmtree()
	case "panelBottom":
		return PanelBottom()
	case "panelBottomClose":
		return PanelBottomClose()
	case "panelBottomInactive":
		return PanelBottomInactive()
	case "panelBottomOpen":
		return PanelBottomOpen()
	case "panelLeft":
		return PanelLeft()
	case "panelLeftClose":
		return PanelLeftClose()
	case "panelLeftInactive":
		return PanelLeftInactive()
	case "panelLeftOpen":
		return PanelLeftOpen()
	case "panelRight":
		return PanelRight()
	case "panelRightClose":
		return PanelRightClose()
	case "panelRightInactive":
		return PanelRightInactive()
	case "panelRightOpen":
		return PanelRightOpen()
	case "panelTop":
		return PanelTop()
	case "panelTopClose":
		return PanelTopClose()
	case "panelTopInactive":
		return PanelTopInactive()
	case "panelTopOpen":
		return PanelTopOpen()
	case "paperclip":
		return Paperclip()
	case "parentheses":
		return Parentheses()
	case "parkingCircle":
		return ParkingCircle()
	case "parkingCircleOff":
		return ParkingCircleOff()
	case "parkingMeter":
		return ParkingMeter()
	case "parkingSquare":
		return ParkingSquare()
	case "parkingSquareOff":
		return ParkingSquareOff()
	case "partyPopper":
		return PartyPopper()
	case "pause":
		return Pause()
	case "pauseCircle":
		return PauseCircle()
	case "pauseOctagon":
		return PauseOctagon()
	case "pawPrint":
		return PawPrint()
	case "pcCase":
		return PcCase()
	case "pen":
		return Pen()
	case "penLine":
		return PenLine()
	case "penSquare":
		return PenSquare()
	case "penTool":
		return PenTool()
	case "pencil":
		return Pencil()
	case "pencilLine":
		return PencilLine()
	case "pencilRuler":
		return PencilRuler()
	case "percent":
		return Percent()
	case "percentCircle":
		return PercentCircle()
	case "percentDiamond":
		return PercentDiamond()
	case "percentSquare":
		return PercentSquare()
	case "personStanding":
		return PersonStanding()
	case "phone":
		return Phone()
	case "phoneCall":
		return PhoneCall()
	case "phoneForwarded":
		return PhoneForwarded()
	case "phoneIncoming":
		return PhoneIncoming()
	case "phoneMissed":
		return PhoneMissed()
	case "phoneOff":
		return PhoneOff()
	case "phoneOutgoing":
		return PhoneOutgoing()
	case "pi":
		return Pi()
	case "piSquare":
		return PiSquare()
	case "pictureInPicture":
		return PictureInPicture()
	case "pictureInPictureTwo":
		return PictureInPictureTwo()
	case "pieChart":
		return PieChart()
	case "piggyBank":
		return PiggyBank()
	case "pilcrow":
		return Pilcrow()
	case "pilcrowSquare":
		return PilcrowSquare()
	case "pill":
		return Pill()
	case "pin":
		return Pin()
	case "pinOff":
		return PinOff()
	case "pipette":
		return Pipette()
	case "pizza":
		return Pizza()
	case "plane":
		return Plane()
	case "planeLanding":
		return PlaneLanding()
	case "planeTakeoff":
		return PlaneTakeoff()
	case "play":
		return Play()
	case "playCircle":
		return PlayCircle()
	case "playSquare":
		return PlaySquare()
	case "plug":
		return Plug()
	case "plugTwo":
		return PlugTwo()
	case "plugZap":
		return PlugZap()
	case "plugZapTwo":
		return PlugZapTwo()
	case "plus":
		return Plus()
	case "plusCircle":
		return PlusCircle()
	case "plusSquare":
		return PlusSquare()
	case "pocket":
		return Pocket()
	case "pocketKnife":
		return PocketKnife()
	case "podcast":
		return Podcast()
	case "pointer":
		return Pointer()
	case "popcorn":
		return Popcorn()
	case "popsicle":
		return Popsicle()
	case "poundSterling":
		return PoundSterling()
	case "power":
		return Power()
	case "powerOff":
		return PowerOff()
	case "presentation":
		return Presentation()
	case "printer":
		return Printer()
	case "projector":
		return Projector()
	case "puzzle":
		return Puzzle()
	case "qrCode":
		return QrCode()
	case "quote":
		return Quote()
	case "rabbit":
		return Rabbit()
	case "radar":
		return Radar()
	case "radiation":
		return Radiation()
	case "radio":
		return Radio()
	case "radioReceiver":
		return RadioReceiver()
	case "radioTower":
		return RadioTower()
	case "railSymbol":
		return RailSymbol()
	case "rainbow":
		return Rainbow()
	case "rat":
		return Rat()
	case "ratio":
		return Ratio()
	case "receipt":
		return Receipt()
	case "rectangleHorizontal":
		return RectangleHorizontal()
	case "rectangleVertical":
		return RectangleVertical()
	case "recycle":
		return Recycle()
	case "redo":
		return Redo()
	case "redoDot":
		return RedoDot()
	case "redoTwo":
		return RedoTwo()
	case "refreshCcw":
		return RefreshCcw()
	case "refreshCcwDot":
		return RefreshCcwDot()
	case "refreshCw":
		return RefreshCw()
	case "refreshCwOff":
		return RefreshCwOff()
	case "refrigerator":
		return Refrigerator()
	case "regex":
		return Regex()
	case "removeFormatting":
		return RemoveFormatting()
	case "repeat":
		return Repeat()
	case "repeatOne":
		return RepeatOne()
	case "repeatTwo":
		return RepeatTwo()
	case "replace":
		return Replace()
	case "replaceAll":
		return ReplaceAll()
	case "reply":
		return Reply()
	case "replyAll":
		return ReplyAll()
	case "rewind":
		return Rewind()
	case "rocket":
		return Rocket()
	case "rockingChair":
		return RockingChair()
	case "rollerCoaster":
		return RollerCoaster()
	case "rotateCcw":
		return RotateCcw()
	case "rotateCw":
		return RotateCw()
	case "rotateThreeD":
		return RotateThreeD()
	case "router":
		return Router()
	case "rows":
		return Rows()
	case "rss":
		return Rss()
	case "ruler":
		return Ruler()
	case "russianRuble":
		return RussianRuble()
	case "sailboat":
		return Sailboat()
	case "salad":
		return Salad()
	case "sandwich":
		return Sandwich()
	case "satellite":
		return Satellite()
	case "satelliteDish":
		return SatelliteDish()
	case "save":
		return Save()
	case "saveAll":
		return SaveAll()
	case "scale":
		return Scale()
	case "scaleThreeD":
		return ScaleThreeD()
	case "scaling":
		return Scaling()
	case "scan":
		return Scan()
	case "scanFace":
		return ScanFace()
	case "scanLine":
		return ScanLine()
	case "scatterChart":
		return ScatterChart()
	case "school":
		return School()
	case "schoolTwo":
		return SchoolTwo()
	case "scissors":
		return Scissors()
	case "scissorsLineDashed":
		return ScissorsLineDashed()
	case "scissorsSquare":
		return ScissorsSquare()
	case "scissorsSquareDashedBottom":
		return ScissorsSquareDashedBottom()
	case "screenShare":
		return ScreenShare()
	case "screenShareOff":
		return ScreenShareOff()
	case "scroll":
		return Scroll()
	case "scrollText":
		return ScrollText()
	case "search":
		return Search()
	case "searchCheck":
		return SearchCheck()
	case "searchCode":
		return SearchCode()
	case "searchLarge":
		return SearchLarge()
	case "searchSlash":
		return SearchSlash()
	case "searchX":
		return SearchX()
	case "send":
		return Send()
	case "sendHorizontal":
		return SendHorizontal()
	case "sendToBack":
		return SendToBack()
	case "separatorHorizontal":
		return SeparatorHorizontal()
	case "separatorVertical":
		return SeparatorVertical()
	case "server":
		return Server()
	case "serverCog":
		return ServerCog()
	case "serverCrash":
		return ServerCrash()
	case "serverOff":
		return ServerOff()
	case "settings":
		return Settings()
	case "settingsTwo":
		return SettingsTwo()
	case "shapes":
		return Shapes()
	case "share":
		return Share()
	case "shareTwo":
		return ShareTwo()
	case "sheet":
		return Sheet()
	case "shell":
		return Shell()
	case "shield":
		return Shield()
	case "shieldAlert":
		return ShieldAlert()
	case "shieldBan":
		return ShieldBan()
	case "shieldCheck":
		return ShieldCheck()
	case "shieldClose":
		return ShieldClose()
	case "shieldEllipsis":
		return ShieldEllipsis()
	case "shieldHalf":
		return ShieldHalf()
	case "shieldMinus":
		return ShieldMinus()
	case "shieldOff":
		return ShieldOff()
	case "shieldPlus":
		return ShieldPlus()
	case "shieldQuestion":
		return ShieldQuestion()
	case "shieldX":
		return ShieldX()
	case "ship":
		return Ship()
	case "shipWheel":
		return ShipWheel()
	case "shirt":
		return Shirt()
	case "shoppingBag":
		return ShoppingBag()
	case "shoppingBasket":
		return ShoppingBasket()
	case "shoppingCart":
		return ShoppingCart()
	case "shovel":
		return Shovel()
	case "showerHead":
		return ShowerHead()
	case "shrink":
		return Shrink()
	case "shrub":
		return Shrub()
	case "shuffle":
		return Shuffle()
	case "sigma":
		return Sigma()
	case "sigmaSquare":
		return SigmaSquare()
	case "signal":
		return Signal()
	case "signalHigh":
		return SignalHigh()
	case "signalLow":
		return SignalLow()
	case "signalMedium":
		return SignalMedium()
	case "signalZero":
		return SignalZero()
	case "siren":
		return Siren()
	case "skipBack":
		return SkipBack()
	case "skipForward":
		return SkipForward()
	case "skull":
		return Skull()
	case "slack":
		return Slack()
	case "slash":
		return Slash()
	case "slice":
		return Slice()
	case "sliders":
		return Sliders()
	case "slidersHorizontal":
		return SlidersHorizontal()
	case "smartphone":
		return Smartphone()
	case "smartphoneCharging":
		return SmartphoneCharging()
	case "smartphoneNfc":
		return SmartphoneNfc()
	case "smile":
		return Smile()
	case "smilePlus":
		return SmilePlus()
	case "snail":
		return Snail()
	case "snowflake":
		return Snowflake()
	case "sofa":
		return Sofa()
	case "sortAsc":
		return SortAsc()
	case "sortDesc":
		return SortDesc()
	case "soup":
		return Soup()
	case "space":
		return Space()
	case "spade":
		return Spade()
	case "sparkle":
		return Sparkle()
	case "sparkles":
		return Sparkles()
	case "speaker":
		return Speaker()
	case "spellCheck":
		return SpellCheck()
	case "spellCheckTwo":
		return SpellCheckTwo()
	case "spline":
		return Spline()
	case "split":
		return Split()
	case "splitSquareHorizontal":
		return SplitSquareHorizontal()
	case "splitSquareVertical":
		return SplitSquareVertical()
	case "sprayCan":
		return SprayCan()
	case "sprout":
		return Sprout()
	case "square":
		return Square()
	case "squareAsterisk":
		return SquareAsterisk()
	case "squareCode":
		return SquareCode()
	case "squareDashedBottom":
		return SquareDashedBottom()
	case "squareDashedBottomCode":
		return SquareDashedBottomCode()
	case "squareDot":
		return SquareDot()
	case "squareEqual":
		return SquareEqual()
	case "squareSlash":
		return SquareSlash()
	case "squareStack":
		return SquareStack()
	case "squirrel":
		return Squirrel()
	case "stamp":
		return Stamp()
	case "star":
		return Star()
	case "starHalf":
		return StarHalf()
	case "starOff":
		return StarOff()
	case "stepBack":
		return StepBack()
	case "stepForward":
		return StepForward()
	case "stethoscope":
		return Stethoscope()
	case "sticker":
		return Sticker()
	case "stickyNote":
		return StickyNote()
	case "stopCircle":
		return StopCircle()
	case "store":
		return Store()
	case "stretchHorizontal":
		return StretchHorizontal()
	case "stretchVertical":
		return StretchVertical()
	case "strikethrough":
		return Strikethrough()
	case "subscript":
		return Subscript()
	case "subtitles":
		return Subtitles()
	case "sun":
		return Sun()
	case "sunDim":
		return SunDim()
	case "sunMedium":
		return SunMedium()
	case "sunMoon":
		return SunMoon()
	case "sunSnow":
		return SunSnow()
	case "sunrise":
		return Sunrise()
	case "sunset":
		return Sunset()
	case "superscript":
		return Superscript()
	case "swissFranc":
		return SwissFranc()
	case "switchCamera":
		return SwitchCamera()
	case "sword":
		return Sword()
	case "swords":
		return Swords()
	case "syringe":
		return Syringe()
	case "table":
		return Table()
	case "tableProperties":
		return TableProperties()
	case "tableTwo":
		return TableTwo()
	case "tablet":
		return Tablet()
	case "tabletSmartphone":
		return TabletSmartphone()
	case "tablets":
		return Tablets()
	case "tag":
		return Tag()
	case "tags":
		return Tags()
	case "tallyFive":
		return TallyFive()
	case "tallyFour":
		return TallyFour()
	case "tallyOne":
		return TallyOne()
	case "tallyThree":
		return TallyThree()
	case "tallyTwo":
		return TallyTwo()
	case "target":
		return Target()
	case "tent":
		return Tent()
	case "terminal":
		return Terminal()
	case "terminalSquare":
		return TerminalSquare()
	case "testTube":
		return TestTube()
	case "testTubeTwo":
		return TestTubeTwo()
	case "testTubes":
		return TestTubes()
	case "text":
		return Text()
	case "textCursor":
		return TextCursor()
	case "textCursorInput":
		return TextCursorInput()
	case "textQuote":
		return TextQuote()
	case "textSelect":
		return TextSelect()
	case "thermometer":
		return Thermometer()
	case "thermometerSnowflake":
		return ThermometerSnowflake()
	case "thermometerSun":
		return ThermometerSun()
	case "thumbsDown":
		return ThumbsDown()
	case "thumbsUp":
		return ThumbsUp()
	case "ticket":
		return Ticket()
	case "timer":
		return Timer()
	case "timerOff":
		return TimerOff()
	case "timerReset":
		return TimerReset()
	case "toggleLeft":
		return ToggleLeft()
	case "toggleRight":
		return ToggleRight()
	case "tornado":
		return Tornado()
	case "touchpad":
		return Touchpad()
	case "touchpadOff":
		return TouchpadOff()
	case "towerControl":
		return TowerControl()
	case "toyBrick":
		return ToyBrick()
	case "tractor":
		return Tractor()
	case "trafficCone":
		return TrafficCone()
	case "trainFront":
		return TrainFront()
	case "trainFrontTunnel":
		return TrainFrontTunnel()
	case "trainTrack":
		return TrainTrack()
	case "tramFront":
		return TramFront()
	case "trash":
		return Trash()
	case "trashTwo":
		return TrashTwo()
	case "treeDeciduous":
		return TreeDeciduous()
	case "treePine":
		return TreePine()
	case "trees":
		return Trees()
	case "trello":
		return Trello()
	case "trendingDown":
		return TrendingDown()
	case "trendingUp":
		return TrendingUp()
	case "triangle":
		return Triangle()
	case "triangleRight":
		return TriangleRight()
	case "trophy":
		return Trophy()
	case "truck":
		return Truck()
	case "turtle":
		return Turtle()
	case "tv":
		return Tv()
	case "tvTwo":
		return TvTwo()
	case "twitch":
		return Twitch()
	case "twitter":
		return Twitter()
	case "type":
		return Type()
	case "umbrella":
		return Umbrella()
	case "underline":
		return Underline()
	case "undo":
		return Undo()
	case "undoDot":
		return UndoDot()
	case "undoTwo":
		return UndoTwo()
	case "unfoldHorizontal":
		return UnfoldHorizontal()
	case "unfoldVertical":
		return UnfoldVertical()
	case "ungroup":
		return Ungroup()
	case "unlink":
		return Unlink()
	case "unlinkTwo":
		return UnlinkTwo()
	case "unlock":
		return Unlock()
	case "unplug":
		return Unplug()
	case "upload":
		return Upload()
	case "uploadCloud":
		return UploadCloud()
	case "usb":
		return Usb()
	case "user":
		return User()
	case "userCheck":
		return UserCheck()
	case "userCheckTwo":
		return UserCheckTwo()
	case "userCircle":
		return UserCircle()
	case "userCircleTwo":
		return UserCircleTwo()
	case "userCog":
		return UserCog()
	case "userCogTwo":
		return UserCogTwo()
	case "userMinus":
		return UserMinus()
	case "userMinusTwo":
		return UserMinusTwo()
	case "userPlus":
		return UserPlus()
	case "userPlusTwo":
		return UserPlusTwo()
	case "userSquare":
		return UserSquare()
	case "userSquareTwo":
		return UserSquareTwo()
	case "userTwo":
		return UserTwo()
	case "userX":
		return UserX()
	case "userXTwo":
		return UserXTwo()
	case "users":
		return Users()
	case "usersTwo":
		return UsersTwo()
	case "utensils":
		return Utensils()
	case "utensilsCrossed":
		return UtensilsCrossed()
	case "utilityPole":
		return UtilityPole()
	case "variable":
		return Variable()
	case "vegan":
		return Vegan()
	case "venetianMask":
		return VenetianMask()
	case "verified":
		return Verified()
	case "vibrate":
		return Vibrate()
	case "vibrateOff":
		return VibrateOff()
	case "video":
		return Video()
	case "videoOff":
		return VideoOff()
	case "videotape":
		return Videotape()
	case "view":
		return View()
	case "voicemail":
		return Voicemail()
	case "volume":
		return Volume()
	case "volumeOne":
		return VolumeOne()
	case "volumeTwo":
		return VolumeTwo()
	case "volumeX":
		return VolumeX()
	case "vote":
		return Vote()
	case "wallet":
		return Wallet()
	case "walletCards":
		return WalletCards()
	case "walletTwo":
		return WalletTwo()
	case "wallpaper":
		return Wallpaper()
	case "wand":
		return Wand()
	case "wandTwo":
		return WandTwo()
	case "warehouse":
		return Warehouse()
	case "watch":
		return Watch()
	case "waves":
		return Waves()
	case "webcam":
		return Webcam()
	case "webhook":
		return Webhook()
	case "wheat":
		return Wheat()
	case "wheatOff":
		return WheatOff()
	case "wholeWord":
		return WholeWord()
	case "wifi":
		return Wifi()
	case "wifiOff":
		return WifiOff()
	case "wind":
		return Wind()
	case "wine":
		return Wine()
	case "wineOff":
		return WineOff()
	case "workflow":
		return Workflow()
	case "wrapText":
		return WrapText()
	case "wrench":
		return Wrench()
	case "x":
		return X()
	case "xCircle":
		return XCircle()
	case "xOctagon":
		return XOctagon()
	case "xSquare":
		return XSquare()
	case "youtube":
		return Youtube()
	case "zap":
		return Zap()
	case "zapOff":
		return ZapOff()
	case "zoomIn":
		return ZoomIn()
	case "zoomOut":
		return ZoomOut()
	default:
		panic("Unknown icon name: " + name)
	}
}

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(accessibilityPath), g.Group(children))
}

func Activity(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(activityPath), g.Group(children))
}

func ActivitySquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(activitySquarePath), g.Group(children))
}

func AirVent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(airVentPath), g.Group(children))
}

func Airplay(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(airplayPath), g.Group(children))
}

func AlarmCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alarmCheckPath), g.Group(children))
}

func AlarmClock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alarmClockPath), g.Group(children))
}

func AlarmClockOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alarmClockOffPath), g.Group(children))
}

func AlarmMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alarmMinusPath), g.Group(children))
}

func AlarmPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alarmPlusPath), g.Group(children))
}

func Album(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(albumPath), g.Group(children))
}

func AlertCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alertCirclePath), g.Group(children))
}

func AlertOctagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alertOctagonPath), g.Group(children))
}

func AlertTriangle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alertTrianglePath), g.Group(children))
}

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignCenterPath), g.Group(children))
}

func AlignCenterHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignCenterHorizontalPath), g.Group(children))
}

func AlignCenterVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignCenterVerticalPath), g.Group(children))
}

func AlignEndHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignEndHorizontalPath), g.Group(children))
}

func AlignEndVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignEndVerticalPath), g.Group(children))
}

func AlignHorizontalDistributeCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalDistributeCenterPath), g.Group(children))
}

func AlignHorizontalDistributeEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalDistributeEndPath), g.Group(children))
}

func AlignHorizontalDistributeStart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalDistributeStartPath), g.Group(children))
}

func AlignHorizontalJustifyCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalJustifyCenterPath), g.Group(children))
}

func AlignHorizontalJustifyEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalJustifyEndPath), g.Group(children))
}

func AlignHorizontalJustifyStart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalJustifyStartPath), g.Group(children))
}

func AlignHorizontalSpaceAround(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalSpaceAroundPath), g.Group(children))
}

func AlignHorizontalSpaceBetween(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignHorizontalSpaceBetweenPath), g.Group(children))
}

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignJustifyPath), g.Group(children))
}

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignLeftPath), g.Group(children))
}

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignRightPath), g.Group(children))
}

func AlignStartHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignStartHorizontalPath), g.Group(children))
}

func AlignStartVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignStartVerticalPath), g.Group(children))
}

func AlignVerticalDistributeCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalDistributeCenterPath), g.Group(children))
}

func AlignVerticalDistributeEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalDistributeEndPath), g.Group(children))
}

func AlignVerticalDistributeStart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalDistributeStartPath), g.Group(children))
}

func AlignVerticalJustifyCenter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalJustifyCenterPath), g.Group(children))
}

func AlignVerticalJustifyEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalJustifyEndPath), g.Group(children))
}

func AlignVerticalJustifyStart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalJustifyStartPath), g.Group(children))
}

func AlignVerticalSpaceAround(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalSpaceAroundPath), g.Group(children))
}

func AlignVerticalSpaceBetween(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(alignVerticalSpaceBetweenPath), g.Group(children))
}

func Ampersand(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ampersandPath), g.Group(children))
}

func Ampersands(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ampersandsPath), g.Group(children))
}

func Anchor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(anchorPath), g.Group(children))
}

func Angry(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(angryPath), g.Group(children))
}

func Annoyed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(annoyedPath), g.Group(children))
}

func Antenna(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(antennaPath), g.Group(children))
}

func Aperture(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(aperturePath), g.Group(children))
}

func AppWindow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(appWindowPath), g.Group(children))
}

func Apple(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(applePath), g.Group(children))
}

func Archive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(archivePath), g.Group(children))
}

func ArchiveRestore(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(archiveRestorePath), g.Group(children))
}

func ArchiveX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(archiveXPath), g.Group(children))
}

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(areaChartPath), g.Group(children))
}

func Armchair(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(armchairPath), g.Group(children))
}

func ArrowBigDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigDownPath), g.Group(children))
}

func ArrowBigDownDash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigDownDashPath), g.Group(children))
}

func ArrowBigLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigLeftPath), g.Group(children))
}

func ArrowBigLeftDash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigLeftDashPath), g.Group(children))
}

func ArrowBigRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigRightPath), g.Group(children))
}

func ArrowBigRightDash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigRightDashPath), g.Group(children))
}

func ArrowBigUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigUpPath), g.Group(children))
}

func ArrowBigUpDash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowBigUpDashPath), g.Group(children))
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownPath), g.Group(children))
}

func ArrowDownAZ(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownAZPath), g.Group(children))
}

func ArrowDownCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownCirclePath), g.Group(children))
}

func ArrowDownFromLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownFromLinePath), g.Group(children))
}

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownLeftPath), g.Group(children))
}

func ArrowDownLeftFromCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownLeftFromCirclePath), g.Group(children))
}

func ArrowDownLeftSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownLeftSquarePath), g.Group(children))
}

func ArrowDownNarrowWide(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownNarrowWidePath), g.Group(children))
}

func ArrowDownOneZero(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownOneZeroPath), g.Group(children))
}

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownRightPath), g.Group(children))
}

func ArrowDownRightFromCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownRightFromCirclePath), g.Group(children))
}

func ArrowDownRightSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownRightSquarePath), g.Group(children))
}

func ArrowDownSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownSquarePath), g.Group(children))
}

func ArrowDownToDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownToDotPath), g.Group(children))
}

func ArrowDownToLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownToLinePath), g.Group(children))
}

func ArrowDownUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownUpPath), g.Group(children))
}

func ArrowDownWideNarrow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownWideNarrowPath), g.Group(children))
}

func ArrowDownZA(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownZAPath), g.Group(children))
}

func ArrowDownZeroOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowDownZeroOnePath), g.Group(children))
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftPath), g.Group(children))
}

func ArrowLeftCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftCirclePath), g.Group(children))
}

func ArrowLeftFromLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftFromLinePath), g.Group(children))
}

func ArrowLeftRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftRightPath), g.Group(children))
}

func ArrowLeftSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftSquarePath), g.Group(children))
}

func ArrowLeftToLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowLeftToLinePath), g.Group(children))
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightPath), g.Group(children))
}

func ArrowRightCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightCirclePath), g.Group(children))
}

func ArrowRightFromLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightFromLinePath), g.Group(children))
}

func ArrowRightLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightLeftPath), g.Group(children))
}

func ArrowRightSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightSquarePath), g.Group(children))
}

func ArrowRightToLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowRightToLinePath), g.Group(children))
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpPath), g.Group(children))
}

func ArrowUpAZ(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpAZPath), g.Group(children))
}

func ArrowUpCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpCirclePath), g.Group(children))
}

func ArrowUpDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpDownPath), g.Group(children))
}

func ArrowUpFromDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpFromDotPath), g.Group(children))
}

func ArrowUpFromLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpFromLinePath), g.Group(children))
}

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpLeftPath), g.Group(children))
}

func ArrowUpLeftFromCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpLeftFromCirclePath), g.Group(children))
}

func ArrowUpLeftSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpLeftSquarePath), g.Group(children))
}

func ArrowUpNarrowWide(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpNarrowWidePath), g.Group(children))
}

func ArrowUpOneZero(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpOneZeroPath), g.Group(children))
}

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpRightPath), g.Group(children))
}

func ArrowUpRightFromCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpRightFromCirclePath), g.Group(children))
}

func ArrowUpRightSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpRightSquarePath), g.Group(children))
}

func ArrowUpSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpSquarePath), g.Group(children))
}

func ArrowUpToLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpToLinePath), g.Group(children))
}

func ArrowUpWideNarrow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpWideNarrowPath), g.Group(children))
}

func ArrowUpZA(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpZAPath), g.Group(children))
}

func ArrowUpZeroOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowUpZeroOnePath), g.Group(children))
}

func ArrowsUpFromLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(arrowsUpFromLinePath), g.Group(children))
}

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(asteriskPath), g.Group(children))
}

func AtSign(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(atSignPath), g.Group(children))
}

func Atom(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(atomPath), g.Group(children))
}

func Award(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(awardPath), g.Group(children))
}

func Axe(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(axePath), g.Group(children))
}

func AxisThreeD(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(axisThreeDPath), g.Group(children))
}

func Baby(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(babyPath), g.Group(children))
}

func Backpack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(backpackPath), g.Group(children))
}

func Badge(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgePath), g.Group(children))
}

func BadgeAlert(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeAlertPath), g.Group(children))
}

func BadgeCent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeCentPath), g.Group(children))
}

func BadgeCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeCheckPath), g.Group(children))
}

func BadgeDollarSign(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeDollarSignPath), g.Group(children))
}

func BadgeEuro(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeEuroPath), g.Group(children))
}

func BadgeHelp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeHelpPath), g.Group(children))
}

func BadgeIndianRupee(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeIndianRupeePath), g.Group(children))
}

func BadgeInfo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeInfoPath), g.Group(children))
}

func BadgeJapaneseYen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeJapaneseYenPath), g.Group(children))
}

func BadgeMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeMinusPath), g.Group(children))
}

func BadgePercent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgePercentPath), g.Group(children))
}

func BadgePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgePlusPath), g.Group(children))
}

func BadgePoundSterling(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgePoundSterlingPath), g.Group(children))
}

func BadgeRussianRuble(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeRussianRublePath), g.Group(children))
}

func BadgeSwissFranc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeSwissFrancPath), g.Group(children))
}

func BadgeX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(badgeXPath), g.Group(children))
}

func BaggageClaim(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(baggageClaimPath), g.Group(children))
}

func Ban(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(banPath), g.Group(children))
}

func Banana(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bananaPath), g.Group(children))
}

func Banknote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(banknotePath), g.Group(children))
}

func BarChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartPath), g.Group(children))
}

func BarChartBig(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartBigPath), g.Group(children))
}

func BarChartFour(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartFourPath), g.Group(children))
}

func BarChartHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartHorizontalPath), g.Group(children))
}

func BarChartHorizontalBig(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartHorizontalBigPath), g.Group(children))
}

func BarChartThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartThreePath), g.Group(children))
}

func BarChartTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(barChartTwoPath), g.Group(children))
}

func Baseline(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(baselinePath), g.Group(children))
}

func Bath(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bathPath), g.Group(children))
}

func Battery(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(batteryPath), g.Group(children))
}

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(batteryChargingPath), g.Group(children))
}

func BatteryFull(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(batteryFullPath), g.Group(children))
}

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(batteryLowPath), g.Group(children))
}

func BatteryMedium(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(batteryMediumPath), g.Group(children))
}

func BatteryWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(batteryWarningPath), g.Group(children))
}

func Beaker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beakerPath), g.Group(children))
}

func Bean(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beanPath), g.Group(children))
}

func BeanOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beanOffPath), g.Group(children))
}

func Bed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bedPath), g.Group(children))
}

func BedDouble(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bedDoublePath), g.Group(children))
}

func BedSingle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bedSinglePath), g.Group(children))
}

func Beef(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beefPath), g.Group(children))
}

func Beer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(beerPath), g.Group(children))
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellPath), g.Group(children))
}

func BellDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellDotPath), g.Group(children))
}

func BellMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellMinusPath), g.Group(children))
}

func BellOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellOffPath), g.Group(children))
}

func BellPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellPlusPath), g.Group(children))
}

func BellRing(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bellRingPath), g.Group(children))
}

func Bike(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bikePath), g.Group(children))
}

func Binary(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(binaryPath), g.Group(children))
}

func Biohazard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(biohazardPath), g.Group(children))
}

func Bird(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(birdPath), g.Group(children))
}

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bitcoinPath), g.Group(children))
}

func Blinds(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(blindsPath), g.Group(children))
}

func Blocks(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(blocksPath), g.Group(children))
}

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothPath), g.Group(children))
}

func BluetoothConnected(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothConnectedPath), g.Group(children))
}

func BluetoothOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothOffPath), g.Group(children))
}

func BluetoothSearching(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bluetoothSearchingPath), g.Group(children))
}

func Bold(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boldPath), g.Group(children))
}

func Bomb(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bombPath), g.Group(children))
}

func Bone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bonePath), g.Group(children))
}

func Book(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookPath), g.Group(children))
}

func BookCopy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookCopyPath), g.Group(children))
}

func BookDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookDownPath), g.Group(children))
}

func BookKey(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookKeyPath), g.Group(children))
}

func BookLock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookLockPath), g.Group(children))
}

func BookMarked(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookMarkedPath), g.Group(children))
}

func BookMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookMinusPath), g.Group(children))
}

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookOpenPath), g.Group(children))
}

func BookOpenCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookOpenCheckPath), g.Group(children))
}

func BookPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookPlusPath), g.Group(children))
}

func BookTemplate(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookTemplatePath), g.Group(children))
}

func BookUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookUpPath), g.Group(children))
}

func BookUpTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookUpTwoPath), g.Group(children))
}

func BookX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookXPath), g.Group(children))
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookmarkPath), g.Group(children))
}

func BookmarkMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookmarkMinusPath), g.Group(children))
}

func BookmarkPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bookmarkPlusPath), g.Group(children))
}

func BoomBox(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boomBoxPath), g.Group(children))
}

func Bot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(botPath), g.Group(children))
}

func Box(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boxPath), g.Group(children))
}

func BoxSelect(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boxSelectPath), g.Group(children))
}

func Boxes(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(boxesPath), g.Group(children))
}

func Braces(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bracesPath), g.Group(children))
}

func Brackets(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bracketsPath), g.Group(children))
}

func Brain(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(brainPath), g.Group(children))
}

func BrainCircuit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(brainCircuitPath), g.Group(children))
}

func BrainCog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(brainCogPath), g.Group(children))
}

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(briefcasePath), g.Group(children))
}

func BringToFront(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bringToFrontPath), g.Group(children))
}

func Brush(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(brushPath), g.Group(children))
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bugPath), g.Group(children))
}

func BugOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bugOffPath), g.Group(children))
}

func BugPlay(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(bugPlayPath), g.Group(children))
}

func Building(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(buildingPath), g.Group(children))
}

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(buildingTwoPath), g.Group(children))
}

func Bus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(busPath), g.Group(children))
}

func BusFront(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(busFrontPath), g.Group(children))
}

func Cable(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cablePath), g.Group(children))
}

func CableCar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cableCarPath), g.Group(children))
}

func Cake(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cakePath), g.Group(children))
}

func CakeSlice(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cakeSlicePath), g.Group(children))
}

func Calculator(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calculatorPath), g.Group(children))
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarPath), g.Group(children))
}

func CalendarCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarCheckPath), g.Group(children))
}

func CalendarCheckTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarCheckTwoPath), g.Group(children))
}

func CalendarClock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarClockPath), g.Group(children))
}

func CalendarDays(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarDaysPath), g.Group(children))
}

func CalendarHeart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarHeartPath), g.Group(children))
}

func CalendarMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarMinusPath), g.Group(children))
}

func CalendarOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarOffPath), g.Group(children))
}

func CalendarPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarPlusPath), g.Group(children))
}

func CalendarRange(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarRangePath), g.Group(children))
}

func CalendarSearch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarSearchPath), g.Group(children))
}

func CalendarX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarXPath), g.Group(children))
}

func CalendarXTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(calendarXTwoPath), g.Group(children))
}

func Camera(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cameraPath), g.Group(children))
}

func CameraOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cameraOffPath), g.Group(children))
}

func CandlestickChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(candlestickChartPath), g.Group(children))
}

func Candy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(candyPath), g.Group(children))
}

func CandyCane(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(candyCanePath), g.Group(children))
}

func CandyOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(candyOffPath), g.Group(children))
}

func Car(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(carPath), g.Group(children))
}

func CarFront(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(carFrontPath), g.Group(children))
}

func CarTaxiFront(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(carTaxiFrontPath), g.Group(children))
}

func Carrot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(carrotPath), g.Group(children))
}

func CaseLower(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(caseLowerPath), g.Group(children))
}

func CaseSensitive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(caseSensitivePath), g.Group(children))
}

func CaseUpper(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(caseUpperPath), g.Group(children))
}

func CassetteTape(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cassetteTapePath), g.Group(children))
}

func Cast(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(castPath), g.Group(children))
}

func Castle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(castlePath), g.Group(children))
}

func Cat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(catPath), g.Group(children))
}

func Check(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkPath), g.Group(children))
}

func CheckCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkCheckPath), g.Group(children))
}

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkCirclePath), g.Group(children))
}

func CheckCircleTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkCircleTwoPath), g.Group(children))
}

func CheckSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(checkSquarePath), g.Group(children))
}

func ChefHat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chefHatPath), g.Group(children))
}

func Cherry(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cherryPath), g.Group(children))
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronDownPath), g.Group(children))
}

func ChevronDownCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronDownCirclePath), g.Group(children))
}

func ChevronDownSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronDownSquarePath), g.Group(children))
}

func ChevronFirst(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronFirstPath), g.Group(children))
}

func ChevronLast(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronLastPath), g.Group(children))
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronLeftPath), g.Group(children))
}

func ChevronLeftCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronLeftCirclePath), g.Group(children))
}

func ChevronLeftSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronLeftSquarePath), g.Group(children))
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronRightPath), g.Group(children))
}

func ChevronRightCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronRightCirclePath), g.Group(children))
}

func ChevronRightSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronRightSquarePath), g.Group(children))
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronUpPath), g.Group(children))
}

func ChevronUpCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronUpCirclePath), g.Group(children))
}

func ChevronUpSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronUpSquarePath), g.Group(children))
}

func ChevronsDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsDownPath), g.Group(children))
}

func ChevronsDownUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsDownUpPath), g.Group(children))
}

func ChevronsLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsLeftPath), g.Group(children))
}

func ChevronsLeftRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsLeftRightPath), g.Group(children))
}

func ChevronsRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsRightPath), g.Group(children))
}

func ChevronsRightLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsRightLeftPath), g.Group(children))
}

func ChevronsUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsUpPath), g.Group(children))
}

func ChevronsUpDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chevronsUpDownPath), g.Group(children))
}

func Chrome(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(chromePath), g.Group(children))
}

func Church(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(churchPath), g.Group(children))
}

func Cigarette(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cigarettePath), g.Group(children))
}

func CigaretteOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cigaretteOffPath), g.Group(children))
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circlePath), g.Group(children))
}

func CircleDashed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleDashedPath), g.Group(children))
}

func CircleDollarSign(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleDollarSignPath), g.Group(children))
}

func CircleDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleDotPath), g.Group(children))
}

func CircleDotDashed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleDotDashedPath), g.Group(children))
}

func CircleEllipsis(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleEllipsisPath), g.Group(children))
}

func CircleEqual(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleEqualPath), g.Group(children))
}

func CircleOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleOffPath), g.Group(children))
}

func CircleSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleSlashPath), g.Group(children))
}

func CircleSlashTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circleSlashTwoPath), g.Group(children))
}

func CircuitBoard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(circuitBoardPath), g.Group(children))
}

func Citrus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(citrusPath), g.Group(children))
}

func Clapperboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clapperboardPath), g.Group(children))
}

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardPath), g.Group(children))
}

func ClipboardCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardCheckPath), g.Group(children))
}

func ClipboardCopy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardCopyPath), g.Group(children))
}

func ClipboardEdit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardEditPath), g.Group(children))
}

func ClipboardList(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardListPath), g.Group(children))
}

func ClipboardPaste(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardPastePath), g.Group(children))
}

func ClipboardSignature(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardSignaturePath), g.Group(children))
}

func ClipboardType(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardTypePath), g.Group(children))
}

func ClipboardX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clipboardXPath), g.Group(children))
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockPath), g.Group(children))
}

func ClockEight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockEightPath), g.Group(children))
}

func ClockEleven(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockElevenPath), g.Group(children))
}

func ClockFive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockFivePath), g.Group(children))
}

func ClockFour(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockFourPath), g.Group(children))
}

func ClockNine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockNinePath), g.Group(children))
}

func ClockOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockOnePath), g.Group(children))
}

func ClockSeven(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockSevenPath), g.Group(children))
}

func ClockSix(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockSixPath), g.Group(children))
}

func ClockTen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockTenPath), g.Group(children))
}

func ClockThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockThreePath), g.Group(children))
}

func ClockTwelve(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockTwelvePath), g.Group(children))
}

func ClockTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clockTwoPath), g.Group(children))
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudPath), g.Group(children))
}

func CloudCog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudCogPath), g.Group(children))
}

func CloudDrizzle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudDrizzlePath), g.Group(children))
}

func CloudFog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudFogPath), g.Group(children))
}

func CloudHail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudHailPath), g.Group(children))
}

func CloudLightning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudLightningPath), g.Group(children))
}

func CloudMoon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudMoonPath), g.Group(children))
}

func CloudMoonRain(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudMoonRainPath), g.Group(children))
}

func CloudOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudOffPath), g.Group(children))
}

func CloudRain(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudRainPath), g.Group(children))
}

func CloudRainWind(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudRainWindPath), g.Group(children))
}

func CloudSnow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudSnowPath), g.Group(children))
}

func CloudSun(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudSunPath), g.Group(children))
}

func CloudSunRain(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudSunRainPath), g.Group(children))
}

func Cloudy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloudyPath), g.Group(children))
}

func Clover(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cloverPath), g.Group(children))
}

func Club(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(clubPath), g.Group(children))
}

func Code(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codePath), g.Group(children))
}

func CodeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codeTwoPath), g.Group(children))
}

func Codepen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codepenPath), g.Group(children))
}

func Codesandbox(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(codesandboxPath), g.Group(children))
}

func Coffee(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(coffeePath), g.Group(children))
}

func Cog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cogPath), g.Group(children))
}

func Coins(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(coinsPath), g.Group(children))
}

func Columns(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(columnsPath), g.Group(children))
}

func Combine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(combinePath), g.Group(children))
}

func Command(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(commandPath), g.Group(children))
}

func Compass(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(compassPath), g.Group(children))
}

func Component(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(componentPath), g.Group(children))
}

func Computer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(computerPath), g.Group(children))
}

func ConciergeBell(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(conciergeBellPath), g.Group(children))
}

func Construction(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(constructionPath), g.Group(children))
}

func Contact(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(contactPath), g.Group(children))
}

func ContactTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(contactTwoPath), g.Group(children))
}

func Container(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(containerPath), g.Group(children))
}

func Contrast(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(contrastPath), g.Group(children))
}

func Cookie(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cookiePath), g.Group(children))
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyPath), g.Group(children))
}

func CopyCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyCheckPath), g.Group(children))
}

func CopyMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyMinusPath), g.Group(children))
}

func CopyPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyPlusPath), g.Group(children))
}

func CopySlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copySlashPath), g.Group(children))
}

func CopyX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyXPath), g.Group(children))
}

func Copyleft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyleftPath), g.Group(children))
}

func Copyright(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(copyrightPath), g.Group(children))
}

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerDownLeftPath), g.Group(children))
}

func CornerDownRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerDownRightPath), g.Group(children))
}

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerLeftDownPath), g.Group(children))
}

func CornerLeftUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerLeftUpPath), g.Group(children))
}

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerRightDownPath), g.Group(children))
}

func CornerRightUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerRightUpPath), g.Group(children))
}

func CornerUpLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerUpLeftPath), g.Group(children))
}

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cornerUpRightPath), g.Group(children))
}

func Cpu(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cpuPath), g.Group(children))
}

func CreativeCommons(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(creativeCommonsPath), g.Group(children))
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(creditCardPath), g.Group(children))
}

func Croissant(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(croissantPath), g.Group(children))
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

func Crown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(crownPath), g.Group(children))
}

func CupSoda(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(cupSodaPath), g.Group(children))
}

func Currency(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(currencyPath), g.Group(children))
}

func Database(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(databasePath), g.Group(children))
}

func DatabaseBackup(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(databaseBackupPath), g.Group(children))
}

func DatabaseZap(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(databaseZapPath), g.Group(children))
}

func Delete(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(deletePath), g.Group(children))
}

func Dessert(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dessertPath), g.Group(children))
}

func Diamond(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diamondPath), g.Group(children))
}

func DiceFive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diceFivePath), g.Group(children))
}

func DiceFour(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diceFourPath), g.Group(children))
}

func DiceOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diceOnePath), g.Group(children))
}

func DiceSix(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diceSixPath), g.Group(children))
}

func DiceThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diceThreePath), g.Group(children))
}

func DiceTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diceTwoPath), g.Group(children))
}

func Dices(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dicesPath), g.Group(children))
}

func Diff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(diffPath), g.Group(children))
}

func Disc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(discPath), g.Group(children))
}

func DiscThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(discThreePath), g.Group(children))
}

func DiscTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(discTwoPath), g.Group(children))
}

func Divide(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dividePath), g.Group(children))
}

func DivideCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(divideCirclePath), g.Group(children))
}

func DivideSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(divideSquarePath), g.Group(children))
}

func Dna(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dnaPath), g.Group(children))
}

func DnaOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dnaOffPath), g.Group(children))
}

func Dog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dogPath), g.Group(children))
}

func DollarSign(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dollarSignPath), g.Group(children))
}

func Donut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(donutPath), g.Group(children))
}

func DoorClosed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(doorClosedPath), g.Group(children))
}

func DoorOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(doorOpenPath), g.Group(children))
}

func Dot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dotPath), g.Group(children))
}

func Download(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(downloadPath), g.Group(children))
}

func DownloadCloud(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(downloadCloudPath), g.Group(children))
}

func Dribbble(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dribbblePath), g.Group(children))
}

func Droplet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropletPath), g.Group(children))
}

func Droplets(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dropletsPath), g.Group(children))
}

func Drumstick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(drumstickPath), g.Group(children))
}

func Dumbbell(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(dumbbellPath), g.Group(children))
}

func Ear(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(earPath), g.Group(children))
}

func EarOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(earOffPath), g.Group(children))
}

func Edit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(editPath), g.Group(children))
}

func EditThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(editThreePath), g.Group(children))
}

func EditTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(editTwoPath), g.Group(children))
}

func Egg(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eggPath), g.Group(children))
}

func EggFried(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eggFriedPath), g.Group(children))
}

func EggOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eggOffPath), g.Group(children))
}

func Equal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(equalPath), g.Group(children))
}

func EqualNot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(equalNotPath), g.Group(children))
}

func Eraser(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eraserPath), g.Group(children))
}

func Euro(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(euroPath), g.Group(children))
}

func Expand(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(expandPath), g.Group(children))
}

func ExternalLink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(externalLinkPath), g.Group(children))
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eyePath), g.Group(children))
}

func EyeOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(eyeOffPath), g.Group(children))
}

func Facebook(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(facebookPath), g.Group(children))
}

func Factory(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(factoryPath), g.Group(children))
}

func Fan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fanPath), g.Group(children))
}

func FastForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fastForwardPath), g.Group(children))
}

func Feather(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(featherPath), g.Group(children))
}

func FerrisWheel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ferrisWheelPath), g.Group(children))
}

func Figma(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(figmaPath), g.Group(children))
}

func File(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePath), g.Group(children))
}

func FileArchive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileArchivePath), g.Group(children))
}

func FileAudio(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileAudioPath), g.Group(children))
}

func FileAudioTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileAudioTwoPath), g.Group(children))
}

func FileAxisThreeD(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileAxisThreeDPath), g.Group(children))
}

func FileBadge(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileBadgePath), g.Group(children))
}

func FileBadgeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileBadgeTwoPath), g.Group(children))
}

func FileBarChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileBarChartPath), g.Group(children))
}

func FileBarChartTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileBarChartTwoPath), g.Group(children))
}

func FileBox(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileBoxPath), g.Group(children))
}

func FileCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCheckPath), g.Group(children))
}

func FileCheckTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCheckTwoPath), g.Group(children))
}

func FileClock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileClockPath), g.Group(children))
}

func FileCode(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCodePath), g.Group(children))
}

func FileCodeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCodeTwoPath), g.Group(children))
}

func FileCog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCogPath), g.Group(children))
}

func FileCogTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileCogTwoPath), g.Group(children))
}

func FileDiff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileDiffPath), g.Group(children))
}

func FileDigit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileDigitPath), g.Group(children))
}

func FileDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileDownPath), g.Group(children))
}

func FileEdit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileEditPath), g.Group(children))
}

func FileHeart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileHeartPath), g.Group(children))
}

func FileImage(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileImagePath), g.Group(children))
}

func FileInput(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileInputPath), g.Group(children))
}

func FileJson(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileJsonPath), g.Group(children))
}

func FileJsonTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileJsonTwoPath), g.Group(children))
}

func FileKey(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileKeyPath), g.Group(children))
}

func FileKeyTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileKeyTwoPath), g.Group(children))
}

func FileLineChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileLineChartPath), g.Group(children))
}

func FileLock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileLockPath), g.Group(children))
}

func FileLockTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileLockTwoPath), g.Group(children))
}

func FileMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileMinusPath), g.Group(children))
}

func FileMinusTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileMinusTwoPath), g.Group(children))
}

func FileOutput(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileOutputPath), g.Group(children))
}

func FilePieChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePieChartPath), g.Group(children))
}

func FilePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePlusPath), g.Group(children))
}

func FilePlusTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filePlusTwoPath), g.Group(children))
}

func FileQuestion(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileQuestionPath), g.Group(children))
}

func FileScan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileScanPath), g.Group(children))
}

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileSearchPath), g.Group(children))
}

func FileSearchTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileSearchTwoPath), g.Group(children))
}

func FileSignature(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileSignaturePath), g.Group(children))
}

func FileSpreadsheet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileSpreadsheetPath), g.Group(children))
}

func FileStack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileStackPath), g.Group(children))
}

func FileSymlink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileSymlinkPath), g.Group(children))
}

func FileTerminal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileTerminalPath), g.Group(children))
}

func FileText(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileTextPath), g.Group(children))
}

func FileType(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileTypePath), g.Group(children))
}

func FileTypeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileTypeTwoPath), g.Group(children))
}

func FileUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileUpPath), g.Group(children))
}

func FileVideo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileVideoPath), g.Group(children))
}

func FileVideoTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileVideoTwoPath), g.Group(children))
}

func FileVolume(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileVolumePath), g.Group(children))
}

func FileVolumeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileVolumeTwoPath), g.Group(children))
}

func FileWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileWarningPath), g.Group(children))
}

func FileX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileXPath), g.Group(children))
}

func FileXTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fileXTwoPath), g.Group(children))
}

func Files(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filesPath), g.Group(children))
}

func Film(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filmPath), g.Group(children))
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filterPath), g.Group(children))
}

func FilterX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(filterXPath), g.Group(children))
}

func Fingerprint(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fingerprintPath), g.Group(children))
}

func Fish(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fishPath), g.Group(children))
}

func FishOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fishOffPath), g.Group(children))
}

func FishSymbol(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fishSymbolPath), g.Group(children))
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flagPath), g.Group(children))
}

func FlagOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flagOffPath), g.Group(children))
}

func FlagTriangleLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flagTriangleLeftPath), g.Group(children))
}

func FlagTriangleRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flagTriangleRightPath), g.Group(children))
}

func Flame(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flamePath), g.Group(children))
}

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flashlightPath), g.Group(children))
}

func FlashlightOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flashlightOffPath), g.Group(children))
}

func FlaskConical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flaskConicalPath), g.Group(children))
}

func FlaskConicalOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flaskConicalOffPath), g.Group(children))
}

func FlaskRound(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flaskRoundPath), g.Group(children))
}

func FlipHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flipHorizontalPath), g.Group(children))
}

func FlipHorizontalTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flipHorizontalTwoPath), g.Group(children))
}

func FlipVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flipVerticalPath), g.Group(children))
}

func FlipVerticalTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flipVerticalTwoPath), g.Group(children))
}

func Flower(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flowerPath), g.Group(children))
}

func FlowerTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(flowerTwoPath), g.Group(children))
}

func Focus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(focusPath), g.Group(children))
}

func FoldHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(foldHorizontalPath), g.Group(children))
}

func FoldVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(foldVerticalPath), g.Group(children))
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderPath), g.Group(children))
}

func FolderArchive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderArchivePath), g.Group(children))
}

func FolderCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderCheckPath), g.Group(children))
}

func FolderClock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderClockPath), g.Group(children))
}

func FolderClosed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderClosedPath), g.Group(children))
}

func FolderCog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderCogPath), g.Group(children))
}

func FolderCogTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderCogTwoPath), g.Group(children))
}

func FolderDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderDotPath), g.Group(children))
}

func FolderDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderDownPath), g.Group(children))
}

func FolderEdit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderEditPath), g.Group(children))
}

func FolderGit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderGitPath), g.Group(children))
}

func FolderGitTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderGitTwoPath), g.Group(children))
}

func FolderHeart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderHeartPath), g.Group(children))
}

func FolderInput(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderInputPath), g.Group(children))
}

func FolderKanban(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderKanbanPath), g.Group(children))
}

func FolderKey(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderKeyPath), g.Group(children))
}

func FolderLock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderLockPath), g.Group(children))
}

func FolderMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderMinusPath), g.Group(children))
}

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderOpenPath), g.Group(children))
}

func FolderOpenDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderOpenDotPath), g.Group(children))
}

func FolderOutput(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderOutputPath), g.Group(children))
}

func FolderPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderPlusPath), g.Group(children))
}

func FolderRoot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderRootPath), g.Group(children))
}

func FolderSearch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderSearchPath), g.Group(children))
}

func FolderSearchTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderSearchTwoPath), g.Group(children))
}

func FolderSymlink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderSymlinkPath), g.Group(children))
}

func FolderSync(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderSyncPath), g.Group(children))
}

func FolderTree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderTreePath), g.Group(children))
}

func FolderUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderUpPath), g.Group(children))
}

func FolderX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(folderXPath), g.Group(children))
}

func Folders(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(foldersPath), g.Group(children))
}

func Footprints(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(footprintsPath), g.Group(children))
}

func Forklift(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(forkliftPath), g.Group(children))
}

func FormInput(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(formInputPath), g.Group(children))
}

func Forward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(forwardPath), g.Group(children))
}

func Frame(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(framePath), g.Group(children))
}

func Framer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(framerPath), g.Group(children))
}

func Frown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(frownPath), g.Group(children))
}

func Fuel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(fuelPath), g.Group(children))
}

func FunctionSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(functionSquarePath), g.Group(children))
}

func GalleryHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(galleryHorizontalPath), g.Group(children))
}

func GalleryHorizontalEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(galleryHorizontalEndPath), g.Group(children))
}

func GalleryThumbnails(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(galleryThumbnailsPath), g.Group(children))
}

func GalleryVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(galleryVerticalPath), g.Group(children))
}

func GalleryVerticalEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(galleryVerticalEndPath), g.Group(children))
}

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gamepadPath), g.Group(children))
}

func GamepadTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gamepadTwoPath), g.Group(children))
}

func GanttChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ganttChartPath), g.Group(children))
}

func GanttChartSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ganttChartSquarePath), g.Group(children))
}

func Gauge(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gaugePath), g.Group(children))
}

func GaugeCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gaugeCirclePath), g.Group(children))
}

func Gavel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gavelPath), g.Group(children))
}

func Gem(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gemPath), g.Group(children))
}

func Ghost(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ghostPath), g.Group(children))
}

func Gift(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(giftPath), g.Group(children))
}

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitBranchPath), g.Group(children))
}

func GitBranchPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitBranchPlusPath), g.Group(children))
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

func GitPullRequest(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitPullRequestPath), g.Group(children))
}

func GitPullRequestClosed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitPullRequestClosedPath), g.Group(children))
}

func GitPullRequestDraft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitPullRequestDraftPath), g.Group(children))
}

func Github(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(githubPath), g.Group(children))
}

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gitlabPath), g.Group(children))
}

func GlassWater(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(glassWaterPath), g.Group(children))
}

func Glasses(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(glassesPath), g.Group(children))
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(globePath), g.Group(children))
}

func GlobeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(globeTwoPath), g.Group(children))
}

func Goal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(goalPath), g.Group(children))
}

func Grab(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(grabPath), g.Group(children))
}

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(graduationCapPath), g.Group(children))
}

func Grape(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(grapePath), g.Group(children))
}

func Grid(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gridPath), g.Group(children))
}

func GridThreeXThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gridThreeXThreePath), g.Group(children))
}

func GridTwoXTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gridTwoXTwoPath), g.Group(children))
}

func Grip(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gripPath), g.Group(children))
}

func GripHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gripHorizontalPath), g.Group(children))
}

func GripVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(gripVerticalPath), g.Group(children))
}

func Group(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(groupPath), g.Group(children))
}

func Hammer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hammerPath), g.Group(children))
}

func Hand(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(handPath), g.Group(children))
}

func HandMetal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(handMetalPath), g.Group(children))
}

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hardDrivePath), g.Group(children))
}

func HardDriveDownload(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hardDriveDownloadPath), g.Group(children))
}

func HardDriveUpload(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hardDriveUploadPath), g.Group(children))
}

func HardHat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hardHatPath), g.Group(children))
}

func Hash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hashPath), g.Group(children))
}

func Haze(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hazePath), g.Group(children))
}

func HdmiPort(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hdmiPortPath), g.Group(children))
}

func Heading(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingPath), g.Group(children))
}

func HeadingFive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingFivePath), g.Group(children))
}

func HeadingFour(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingFourPath), g.Group(children))
}

func HeadingOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingOnePath), g.Group(children))
}

func HeadingSix(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingSixPath), g.Group(children))
}

func HeadingThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingThreePath), g.Group(children))
}

func HeadingTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headingTwoPath), g.Group(children))
}

func Headphones(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(headphonesPath), g.Group(children))
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartPath), g.Group(children))
}

func HeartCrack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartCrackPath), g.Group(children))
}

func HeartHandshake(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartHandshakePath), g.Group(children))
}

func HeartOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartOffPath), g.Group(children))
}

func HeartPulse(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(heartPulsePath), g.Group(children))
}

func HelpCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(helpCirclePath), g.Group(children))
}

func HelpingHand(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(helpingHandPath), g.Group(children))
}

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hexagonPath), g.Group(children))
}

func Highlighter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(highlighterPath), g.Group(children))
}

func History(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(historyPath), g.Group(children))
}

func Home(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(homePath), g.Group(children))
}

func Hop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hopPath), g.Group(children))
}

func HopOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hopOffPath), g.Group(children))
}

func Hotel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hotelPath), g.Group(children))
}

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(hourglassPath), g.Group(children))
}

func IceCream(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(iceCreamPath), g.Group(children))
}

func IceCreamTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(iceCreamTwoPath), g.Group(children))
}

func Image(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(imagePath), g.Group(children))
}

func ImageMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(imageMinusPath), g.Group(children))
}

func ImageOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(imageOffPath), g.Group(children))
}

func ImagePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(imagePlusPath), g.Group(children))
}

func Import(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(importPath), g.Group(children))
}

func Inbox(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(inboxPath), g.Group(children))
}

func Indent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(indentPath), g.Group(children))
}

func IndianRupee(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(indianRupeePath), g.Group(children))
}

func Infinity(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(infinityPath), g.Group(children))
}

func Info(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(infoPath), g.Group(children))
}

func Instagram(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(instagramPath), g.Group(children))
}

func Italic(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(italicPath), g.Group(children))
}

func IterationCcw(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(iterationCcwPath), g.Group(children))
}

func IterationCw(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(iterationCwPath), g.Group(children))
}

func JapaneseYen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(japaneseYenPath), g.Group(children))
}

func Joystick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(joystickPath), g.Group(children))
}

func Kanban(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(kanbanPath), g.Group(children))
}

func KanbanSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(kanbanSquarePath), g.Group(children))
}

func KanbanSquareDashed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(kanbanSquareDashedPath), g.Group(children))
}

func Key(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keyPath), g.Group(children))
}

func KeyRound(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keyRoundPath), g.Group(children))
}

func KeySquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keySquarePath), g.Group(children))
}

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(keyboardPath), g.Group(children))
}

func Lamp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lampPath), g.Group(children))
}

func LampCeiling(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lampCeilingPath), g.Group(children))
}

func LampDesk(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lampDeskPath), g.Group(children))
}

func LampFloor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lampFloorPath), g.Group(children))
}

func LampWallDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lampWallDownPath), g.Group(children))
}

func LampWallUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lampWallUpPath), g.Group(children))
}

func Landmark(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(landmarkPath), g.Group(children))
}

func Languages(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(languagesPath), g.Group(children))
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(laptopPath), g.Group(children))
}

func LaptopTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(laptopTwoPath), g.Group(children))
}

func Lasso(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lassoPath), g.Group(children))
}

func LassoSelect(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lassoSelectPath), g.Group(children))
}

func Laugh(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(laughPath), g.Group(children))
}

func Layers(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layersPath), g.Group(children))
}

func Layout(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutPath), g.Group(children))
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

func LayoutPanelLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutPanelLeftPath), g.Group(children))
}

func LayoutPanelTop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutPanelTopPath), g.Group(children))
}

func LayoutTemplate(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(layoutTemplatePath), g.Group(children))
}

func Leaf(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(leafPath), g.Group(children))
}

func LeafyGreen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(leafyGreenPath), g.Group(children))
}

func Library(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(libraryPath), g.Group(children))
}

func LifeBuoy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lifeBuoyPath), g.Group(children))
}

func Ligature(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ligaturePath), g.Group(children))
}

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lightbulbPath), g.Group(children))
}

func LightbulbOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lightbulbOffPath), g.Group(children))
}

func LineChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lineChartPath), g.Group(children))
}

func Link(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkPath), g.Group(children))
}

func LinkTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkTwoPath), g.Group(children))
}

func LinkTwoOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkTwoOffPath), g.Group(children))
}

func Linkedin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(linkedinPath), g.Group(children))
}

func List(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listPath), g.Group(children))
}

func ListChecks(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listChecksPath), g.Group(children))
}

func ListEnd(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listEndPath), g.Group(children))
}

func ListFilter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listFilterPath), g.Group(children))
}

func ListMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listMinusPath), g.Group(children))
}

func ListMusic(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listMusicPath), g.Group(children))
}

func ListOrdered(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listOrderedPath), g.Group(children))
}

func ListPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listPlusPath), g.Group(children))
}

func ListRestart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listRestartPath), g.Group(children))
}

func ListStart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listStartPath), g.Group(children))
}

func ListTodo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listTodoPath), g.Group(children))
}

func ListTree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listTreePath), g.Group(children))
}

func ListVideo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listVideoPath), g.Group(children))
}

func ListX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(listXPath), g.Group(children))
}

func Loader(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(loaderPath), g.Group(children))
}

func LoaderTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(loaderTwoPath), g.Group(children))
}

func Locate(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(locatePath), g.Group(children))
}

func LocateFixed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(locateFixedPath), g.Group(children))
}

func LocateOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(locateOffPath), g.Group(children))
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lockPath), g.Group(children))
}

func LogIn(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(logInPath), g.Group(children))
}

func LogOut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(logOutPath), g.Group(children))
}

func Lollipop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(lollipopPath), g.Group(children))
}

func Luggage(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(luggagePath), g.Group(children))
}

func MSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mSquarePath), g.Group(children))
}

func Magnet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(magnetPath), g.Group(children))
}

func Mail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailPath), g.Group(children))
}

func MailCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailCheckPath), g.Group(children))
}

func MailMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailMinusPath), g.Group(children))
}

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailOpenPath), g.Group(children))
}

func MailPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailPlusPath), g.Group(children))
}

func MailQuestion(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailQuestionPath), g.Group(children))
}

func MailSearch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailSearchPath), g.Group(children))
}

func MailWarning(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailWarningPath), g.Group(children))
}

func MailX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailXPath), g.Group(children))
}

func Mailbox(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailboxPath), g.Group(children))
}

func Mails(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mailsPath), g.Group(children))
}

func Map(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mapPath), g.Group(children))
}

func MapPin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mapPinPath), g.Group(children))
}

func MapPinOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mapPinOffPath), g.Group(children))
}

func Martini(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(martiniPath), g.Group(children))
}

func Maximize(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(maximizePath), g.Group(children))
}

func MaximizeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(maximizeTwoPath), g.Group(children))
}

func Medal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(medalPath), g.Group(children))
}

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(megaphonePath), g.Group(children))
}

func MegaphoneOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(megaphoneOffPath), g.Group(children))
}

func Meh(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mehPath), g.Group(children))
}

func MemoryStick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(memoryStickPath), g.Group(children))
}

func Menu(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(menuPath), g.Group(children))
}

func MenuSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(menuSquarePath), g.Group(children))
}

func Merge(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mergePath), g.Group(children))
}

func MessageCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messageCirclePath), g.Group(children))
}

func MessageSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messageSquarePath), g.Group(children))
}

func MessageSquareDashed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messageSquareDashedPath), g.Group(children))
}

func MessageSquarePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messageSquarePlusPath), g.Group(children))
}

func MessagesSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(messagesSquarePath), g.Group(children))
}

func Mic(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(micPath), g.Group(children))
}

func MicOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(micOffPath), g.Group(children))
}

func MicTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(micTwoPath), g.Group(children))
}

func Microscope(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(microscopePath), g.Group(children))
}

func Microwave(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(microwavePath), g.Group(children))
}

func Milestone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(milestonePath), g.Group(children))
}

func Milk(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(milkPath), g.Group(children))
}

func MilkOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(milkOffPath), g.Group(children))
}

func Minimize(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minimizePath), g.Group(children))
}

func MinimizeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minimizeTwoPath), g.Group(children))
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minusPath), g.Group(children))
}

func MinusCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minusCirclePath), g.Group(children))
}

func MinusSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(minusSquarePath), g.Group(children))
}

func Monitor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorPath), g.Group(children))
}

func MonitorCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorCheckPath), g.Group(children))
}

func MonitorDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorDotPath), g.Group(children))
}

func MonitorDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorDownPath), g.Group(children))
}

func MonitorOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorOffPath), g.Group(children))
}

func MonitorPause(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorPausePath), g.Group(children))
}

func MonitorPlay(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorPlayPath), g.Group(children))
}

func MonitorSmartphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorSmartphonePath), g.Group(children))
}

func MonitorSpeaker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorSpeakerPath), g.Group(children))
}

func MonitorStop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorStopPath), g.Group(children))
}

func MonitorUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorUpPath), g.Group(children))
}

func MonitorX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(monitorXPath), g.Group(children))
}

func Moon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moonPath), g.Group(children))
}

func MoonStar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moonStarPath), g.Group(children))
}

func MoreHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moreHorizontalPath), g.Group(children))
}

func MoreVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moreVerticalPath), g.Group(children))
}

func Mountain(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mountainPath), g.Group(children))
}

func MountainSnow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mountainSnowPath), g.Group(children))
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePath), g.Group(children))
}

func MousePointer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePointerPath), g.Group(children))
}

func MousePointerClick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePointerClickPath), g.Group(children))
}

func MousePointerSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePointerSquarePath), g.Group(children))
}

func MousePointerSquareDashed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePointerSquareDashedPath), g.Group(children))
}

func MousePointerTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(mousePointerTwoPath), g.Group(children))
}

func Move(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(movePath), g.Group(children))
}

func MoveDiagonal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveDiagonalPath), g.Group(children))
}

func MoveDiagonalTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveDiagonalTwoPath), g.Group(children))
}

func MoveDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveDownPath), g.Group(children))
}

func MoveDownLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveDownLeftPath), g.Group(children))
}

func MoveDownRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveDownRightPath), g.Group(children))
}

func MoveHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveHorizontalPath), g.Group(children))
}

func MoveLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveLeftPath), g.Group(children))
}

func MoveRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveRightPath), g.Group(children))
}

func MoveThreeD(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveThreeDPath), g.Group(children))
}

func MoveUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveUpPath), g.Group(children))
}

func MoveUpLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveUpLeftPath), g.Group(children))
}

func MoveUpRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveUpRightPath), g.Group(children))
}

func MoveVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(moveVerticalPath), g.Group(children))
}

func Music(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(musicPath), g.Group(children))
}

func MusicFour(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(musicFourPath), g.Group(children))
}

func MusicThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(musicThreePath), g.Group(children))
}

func MusicTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(musicTwoPath), g.Group(children))
}

func Navigation(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(navigationPath), g.Group(children))
}

func NavigationOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(navigationOffPath), g.Group(children))
}

func NavigationTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(navigationTwoPath), g.Group(children))
}

func NavigationTwoOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(navigationTwoOffPath), g.Group(children))
}

func Network(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(networkPath), g.Group(children))
}

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(newspaperPath), g.Group(children))
}

func Nfc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(nfcPath), g.Group(children))
}

func Nut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(nutPath), g.Group(children))
}

func NutOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(nutOffPath), g.Group(children))
}

func Octagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(octagonPath), g.Group(children))
}

func Option(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(optionPath), g.Group(children))
}

func Orbit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(orbitPath), g.Group(children))
}

func Outdent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(outdentPath), g.Group(children))
}

func Package(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packagePath), g.Group(children))
}

func PackageCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packageCheckPath), g.Group(children))
}

func PackageMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packageMinusPath), g.Group(children))
}

func PackageOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packageOpenPath), g.Group(children))
}

func PackagePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packagePlusPath), g.Group(children))
}

func PackageSearch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packageSearchPath), g.Group(children))
}

func PackageTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packageTwoPath), g.Group(children))
}

func PackageX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(packageXPath), g.Group(children))
}

func PaintBucket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paintBucketPath), g.Group(children))
}

func Paintbrush(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paintbrushPath), g.Group(children))
}

func PaintbrushTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paintbrushTwoPath), g.Group(children))
}

func Palette(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(palettePath), g.Group(children))
}

func Palmtree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(palmtreePath), g.Group(children))
}

func PanelBottom(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelBottomPath), g.Group(children))
}

func PanelBottomClose(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelBottomClosePath), g.Group(children))
}

func PanelBottomInactive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelBottomInactivePath), g.Group(children))
}

func PanelBottomOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelBottomOpenPath), g.Group(children))
}

func PanelLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelLeftPath), g.Group(children))
}

func PanelLeftClose(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelLeftClosePath), g.Group(children))
}

func PanelLeftInactive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelLeftInactivePath), g.Group(children))
}

func PanelLeftOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelLeftOpenPath), g.Group(children))
}

func PanelRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelRightPath), g.Group(children))
}

func PanelRightClose(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelRightClosePath), g.Group(children))
}

func PanelRightInactive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelRightInactivePath), g.Group(children))
}

func PanelRightOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelRightOpenPath), g.Group(children))
}

func PanelTop(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelTopPath), g.Group(children))
}

func PanelTopClose(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelTopClosePath), g.Group(children))
}

func PanelTopInactive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelTopInactivePath), g.Group(children))
}

func PanelTopOpen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(panelTopOpenPath), g.Group(children))
}

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(paperclipPath), g.Group(children))
}

func Parentheses(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(parenthesesPath), g.Group(children))
}

func ParkingCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(parkingCirclePath), g.Group(children))
}

func ParkingCircleOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(parkingCircleOffPath), g.Group(children))
}

func ParkingMeter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(parkingMeterPath), g.Group(children))
}

func ParkingSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(parkingSquarePath), g.Group(children))
}

func ParkingSquareOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(parkingSquareOffPath), g.Group(children))
}

func PartyPopper(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(partyPopperPath), g.Group(children))
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pausePath), g.Group(children))
}

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pauseCirclePath), g.Group(children))
}

func PauseOctagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pauseOctagonPath), g.Group(children))
}

func PawPrint(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pawPrintPath), g.Group(children))
}

func PcCase(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pcCasePath), g.Group(children))
}

func Pen(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(penPath), g.Group(children))
}

func PenLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(penLinePath), g.Group(children))
}

func PenSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(penSquarePath), g.Group(children))
}

func PenTool(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(penToolPath), g.Group(children))
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pencilPath), g.Group(children))
}

func PencilLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pencilLinePath), g.Group(children))
}

func PencilRuler(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pencilRulerPath), g.Group(children))
}

func Percent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(percentPath), g.Group(children))
}

func PercentCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(percentCirclePath), g.Group(children))
}

func PercentDiamond(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(percentDiamondPath), g.Group(children))
}

func PercentSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(percentSquarePath), g.Group(children))
}

func PersonStanding(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(personStandingPath), g.Group(children))
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phonePath), g.Group(children))
}

func PhoneCall(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneCallPath), g.Group(children))
}

func PhoneForwarded(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneForwardedPath), g.Group(children))
}

func PhoneIncoming(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneIncomingPath), g.Group(children))
}

func PhoneMissed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneMissedPath), g.Group(children))
}

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneOffPath), g.Group(children))
}

func PhoneOutgoing(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(phoneOutgoingPath), g.Group(children))
}

func Pi(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(piPath), g.Group(children))
}

func PiSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(piSquarePath), g.Group(children))
}

func PictureInPicture(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pictureInPicturePath), g.Group(children))
}

func PictureInPictureTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pictureInPictureTwoPath), g.Group(children))
}

func PieChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pieChartPath), g.Group(children))
}

func PiggyBank(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(piggyBankPath), g.Group(children))
}

func Pilcrow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pilcrowPath), g.Group(children))
}

func PilcrowSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pilcrowSquarePath), g.Group(children))
}

func Pill(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pillPath), g.Group(children))
}

func Pin(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pinPath), g.Group(children))
}

func PinOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pinOffPath), g.Group(children))
}

func Pipette(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pipettePath), g.Group(children))
}

func Pizza(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pizzaPath), g.Group(children))
}

func Plane(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(planePath), g.Group(children))
}

func PlaneLanding(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(planeLandingPath), g.Group(children))
}

func PlaneTakeoff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(planeTakeoffPath), g.Group(children))
}

func Play(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(playPath), g.Group(children))
}

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(playCirclePath), g.Group(children))
}

func PlaySquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(playSquarePath), g.Group(children))
}

func Plug(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plugPath), g.Group(children))
}

func PlugTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plugTwoPath), g.Group(children))
}

func PlugZap(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plugZapPath), g.Group(children))
}

func PlugZapTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plugZapTwoPath), g.Group(children))
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plusPath), g.Group(children))
}

func PlusCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plusCirclePath), g.Group(children))
}

func PlusSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(plusSquarePath), g.Group(children))
}

func Pocket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pocketPath), g.Group(children))
}

func PocketKnife(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pocketKnifePath), g.Group(children))
}

func Podcast(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(podcastPath), g.Group(children))
}

func Pointer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(pointerPath), g.Group(children))
}

func Popcorn(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(popcornPath), g.Group(children))
}

func Popsicle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(popsiclePath), g.Group(children))
}

func PoundSterling(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(poundSterlingPath), g.Group(children))
}

func Power(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(powerPath), g.Group(children))
}

func PowerOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(powerOffPath), g.Group(children))
}

func Presentation(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(presentationPath), g.Group(children))
}

func Printer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(printerPath), g.Group(children))
}

func Projector(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(projectorPath), g.Group(children))
}

func Puzzle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(puzzlePath), g.Group(children))
}

func QrCode(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(qrCodePath), g.Group(children))
}

func Quote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(quotePath), g.Group(children))
}

func Rabbit(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rabbitPath), g.Group(children))
}

func Radar(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(radarPath), g.Group(children))
}

func Radiation(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(radiationPath), g.Group(children))
}

func Radio(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(radioPath), g.Group(children))
}

func RadioReceiver(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(radioReceiverPath), g.Group(children))
}

func RadioTower(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(radioTowerPath), g.Group(children))
}

func RailSymbol(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(railSymbolPath), g.Group(children))
}

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rainbowPath), g.Group(children))
}

func Rat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ratPath), g.Group(children))
}

func Ratio(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ratioPath), g.Group(children))
}

func Receipt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(receiptPath), g.Group(children))
}

func RectangleHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rectangleHorizontalPath), g.Group(children))
}

func RectangleVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rectangleVerticalPath), g.Group(children))
}

func Recycle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(recyclePath), g.Group(children))
}

func Redo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(redoPath), g.Group(children))
}

func RedoDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(redoDotPath), g.Group(children))
}

func RedoTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(redoTwoPath), g.Group(children))
}

func RefreshCcw(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(refreshCcwPath), g.Group(children))
}

func RefreshCcwDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(refreshCcwDotPath), g.Group(children))
}

func RefreshCw(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(refreshCwPath), g.Group(children))
}

func RefreshCwOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(refreshCwOffPath), g.Group(children))
}

func Refrigerator(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(refrigeratorPath), g.Group(children))
}

func Regex(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(regexPath), g.Group(children))
}

func RemoveFormatting(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(removeFormattingPath), g.Group(children))
}

func Repeat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(repeatPath), g.Group(children))
}

func RepeatOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(repeatOnePath), g.Group(children))
}

func RepeatTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(repeatTwoPath), g.Group(children))
}

func Replace(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(replacePath), g.Group(children))
}

func ReplaceAll(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(replaceAllPath), g.Group(children))
}

func Reply(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(replyPath), g.Group(children))
}

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(replyAllPath), g.Group(children))
}

func Rewind(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rewindPath), g.Group(children))
}

func Rocket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rocketPath), g.Group(children))
}

func RockingChair(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rockingChairPath), g.Group(children))
}

func RollerCoaster(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rollerCoasterPath), g.Group(children))
}

func RotateCcw(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rotateCcwPath), g.Group(children))
}

func RotateCw(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rotateCwPath), g.Group(children))
}

func RotateThreeD(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rotateThreeDPath), g.Group(children))
}

func Router(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(routerPath), g.Group(children))
}

func Rows(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rowsPath), g.Group(children))
}

func Rss(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rssPath), g.Group(children))
}

func Ruler(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(rulerPath), g.Group(children))
}

func RussianRuble(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(russianRublePath), g.Group(children))
}

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sailboatPath), g.Group(children))
}

func Salad(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(saladPath), g.Group(children))
}

func Sandwich(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sandwichPath), g.Group(children))
}

func Satellite(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(satellitePath), g.Group(children))
}

func SatelliteDish(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(satelliteDishPath), g.Group(children))
}

func Save(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(savePath), g.Group(children))
}

func SaveAll(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(saveAllPath), g.Group(children))
}

func Scale(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scalePath), g.Group(children))
}

func ScaleThreeD(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scaleThreeDPath), g.Group(children))
}

func Scaling(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scalingPath), g.Group(children))
}

func Scan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scanPath), g.Group(children))
}

func ScanFace(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scanFacePath), g.Group(children))
}

func ScanLine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scanLinePath), g.Group(children))
}

func ScatterChart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scatterChartPath), g.Group(children))
}

func School(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(schoolPath), g.Group(children))
}

func SchoolTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(schoolTwoPath), g.Group(children))
}

func Scissors(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scissorsPath), g.Group(children))
}

func ScissorsLineDashed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scissorsLineDashedPath), g.Group(children))
}

func ScissorsSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scissorsSquarePath), g.Group(children))
}

func ScissorsSquareDashedBottom(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scissorsSquareDashedBottomPath), g.Group(children))
}

func ScreenShare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(screenSharePath), g.Group(children))
}

func ScreenShareOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(screenShareOffPath), g.Group(children))
}

func Scroll(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scrollPath), g.Group(children))
}

func ScrollText(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(scrollTextPath), g.Group(children))
}

func Search(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchPath), g.Group(children))
}

func SearchCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchCheckPath), g.Group(children))
}

func SearchCode(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchCodePath), g.Group(children))
}

func SearchLarge(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchLargePath), g.Group(children))
}

func SearchSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchSlashPath), g.Group(children))
}

func SearchX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(searchXPath), g.Group(children))
}

func Send(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sendPath), g.Group(children))
}

func SendHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sendHorizontalPath), g.Group(children))
}

func SendToBack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sendToBackPath), g.Group(children))
}

func SeparatorHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(separatorHorizontalPath), g.Group(children))
}

func SeparatorVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(separatorVerticalPath), g.Group(children))
}

func Server(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(serverPath), g.Group(children))
}

func ServerCog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(serverCogPath), g.Group(children))
}

func ServerCrash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(serverCrashPath), g.Group(children))
}

func ServerOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(serverOffPath), g.Group(children))
}

func Settings(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(settingsPath), g.Group(children))
}

func SettingsTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(settingsTwoPath), g.Group(children))
}

func Shapes(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shapesPath), g.Group(children))
}

func Share(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sharePath), g.Group(children))
}

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shareTwoPath), g.Group(children))
}

func Sheet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sheetPath), g.Group(children))
}

func Shell(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shellPath), g.Group(children))
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldPath), g.Group(children))
}

func ShieldAlert(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldAlertPath), g.Group(children))
}

func ShieldBan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldBanPath), g.Group(children))
}

func ShieldCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldCheckPath), g.Group(children))
}

func ShieldClose(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldClosePath), g.Group(children))
}

func ShieldEllipsis(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldEllipsisPath), g.Group(children))
}

func ShieldHalf(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldHalfPath), g.Group(children))
}

func ShieldMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldMinusPath), g.Group(children))
}

func ShieldOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldOffPath), g.Group(children))
}

func ShieldPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldPlusPath), g.Group(children))
}

func ShieldQuestion(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldQuestionPath), g.Group(children))
}

func ShieldX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shieldXPath), g.Group(children))
}

func Ship(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shipPath), g.Group(children))
}

func ShipWheel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shipWheelPath), g.Group(children))
}

func Shirt(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shirtPath), g.Group(children))
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shoppingBagPath), g.Group(children))
}

func ShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shoppingBasketPath), g.Group(children))
}

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shoppingCartPath), g.Group(children))
}

func Shovel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shovelPath), g.Group(children))
}

func ShowerHead(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(showerHeadPath), g.Group(children))
}

func Shrink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shrinkPath), g.Group(children))
}

func Shrub(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shrubPath), g.Group(children))
}

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(shufflePath), g.Group(children))
}

func Sigma(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sigmaPath), g.Group(children))
}

func SigmaSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sigmaSquarePath), g.Group(children))
}

func Signal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signalPath), g.Group(children))
}

func SignalHigh(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signalHighPath), g.Group(children))
}

func SignalLow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signalLowPath), g.Group(children))
}

func SignalMedium(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signalMediumPath), g.Group(children))
}

func SignalZero(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(signalZeroPath), g.Group(children))
}

func Siren(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sirenPath), g.Group(children))
}

func SkipBack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(skipBackPath), g.Group(children))
}

func SkipForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(skipForwardPath), g.Group(children))
}

func Skull(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(skullPath), g.Group(children))
}

func Slack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(slackPath), g.Group(children))
}

func Slash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(slashPath), g.Group(children))
}

func Slice(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(slicePath), g.Group(children))
}

func Sliders(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(slidersPath), g.Group(children))
}

func SlidersHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(slidersHorizontalPath), g.Group(children))
}

func Smartphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smartphonePath), g.Group(children))
}

func SmartphoneCharging(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smartphoneChargingPath), g.Group(children))
}

func SmartphoneNfc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smartphoneNfcPath), g.Group(children))
}

func Smile(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smilePath), g.Group(children))
}

func SmilePlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(smilePlusPath), g.Group(children))
}

func Snail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(snailPath), g.Group(children))
}

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(snowflakePath), g.Group(children))
}

func Sofa(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sofaPath), g.Group(children))
}

func SortAsc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sortAscPath), g.Group(children))
}

func SortDesc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sortDescPath), g.Group(children))
}

func Soup(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(soupPath), g.Group(children))
}

func Space(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(spacePath), g.Group(children))
}

func Spade(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(spadePath), g.Group(children))
}

func Sparkle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sparklePath), g.Group(children))
}

func Sparkles(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sparklesPath), g.Group(children))
}

func Speaker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(speakerPath), g.Group(children))
}

func SpellCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(spellCheckPath), g.Group(children))
}

func SpellCheckTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(spellCheckTwoPath), g.Group(children))
}

func Spline(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(splinePath), g.Group(children))
}

func Split(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(splitPath), g.Group(children))
}

func SplitSquareHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(splitSquareHorizontalPath), g.Group(children))
}

func SplitSquareVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(splitSquareVerticalPath), g.Group(children))
}

func SprayCan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sprayCanPath), g.Group(children))
}

func Sprout(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sproutPath), g.Group(children))
}

func Square(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squarePath), g.Group(children))
}

func SquareAsterisk(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareAsteriskPath), g.Group(children))
}

func SquareCode(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareCodePath), g.Group(children))
}

func SquareDashedBottom(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareDashedBottomPath), g.Group(children))
}

func SquareDashedBottomCode(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareDashedBottomCodePath), g.Group(children))
}

func SquareDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareDotPath), g.Group(children))
}

func SquareEqual(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareEqualPath), g.Group(children))
}

func SquareSlash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareSlashPath), g.Group(children))
}

func SquareStack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squareStackPath), g.Group(children))
}

func Squirrel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(squirrelPath), g.Group(children))
}

func Stamp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stampPath), g.Group(children))
}

func Star(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(starPath), g.Group(children))
}

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(starHalfPath), g.Group(children))
}

func StarOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(starOffPath), g.Group(children))
}

func StepBack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stepBackPath), g.Group(children))
}

func StepForward(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stepForwardPath), g.Group(children))
}

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stethoscopePath), g.Group(children))
}

func Sticker(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stickerPath), g.Group(children))
}

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stickyNotePath), g.Group(children))
}

func StopCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stopCirclePath), g.Group(children))
}

func Store(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(storePath), g.Group(children))
}

func StretchHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stretchHorizontalPath), g.Group(children))
}

func StretchVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(stretchVerticalPath), g.Group(children))
}

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(strikethroughPath), g.Group(children))
}

func Subscript(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(subscriptPath), g.Group(children))
}

func Subtitles(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(subtitlesPath), g.Group(children))
}

func Sun(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunPath), g.Group(children))
}

func SunDim(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunDimPath), g.Group(children))
}

func SunMedium(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunMediumPath), g.Group(children))
}

func SunMoon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunMoonPath), g.Group(children))
}

func SunSnow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunSnowPath), g.Group(children))
}

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunrisePath), g.Group(children))
}

func Sunset(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(sunsetPath), g.Group(children))
}

func Superscript(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(superscriptPath), g.Group(children))
}

func SwissFranc(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swissFrancPath), g.Group(children))
}

func SwitchCamera(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(switchCameraPath), g.Group(children))
}

func Sword(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swordPath), g.Group(children))
}

func Swords(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(swordsPath), g.Group(children))
}

func Syringe(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(syringePath), g.Group(children))
}

func Table(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tablePath), g.Group(children))
}

func TableProperties(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tablePropertiesPath), g.Group(children))
}

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tableTwoPath), g.Group(children))
}

func Tablet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tabletPath), g.Group(children))
}

func TabletSmartphone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tabletSmartphonePath), g.Group(children))
}

func Tablets(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tabletsPath), g.Group(children))
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tagPath), g.Group(children))
}

func Tags(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tagsPath), g.Group(children))
}

func TallyFive(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tallyFivePath), g.Group(children))
}

func TallyFour(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tallyFourPath), g.Group(children))
}

func TallyOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tallyOnePath), g.Group(children))
}

func TallyThree(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tallyThreePath), g.Group(children))
}

func TallyTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tallyTwoPath), g.Group(children))
}

func Target(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(targetPath), g.Group(children))
}

func Tent(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tentPath), g.Group(children))
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(terminalPath), g.Group(children))
}

func TerminalSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(terminalSquarePath), g.Group(children))
}

func TestTube(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(testTubePath), g.Group(children))
}

func TestTubeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(testTubeTwoPath), g.Group(children))
}

func TestTubes(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(testTubesPath), g.Group(children))
}

func Text(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textPath), g.Group(children))
}

func TextCursor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textCursorPath), g.Group(children))
}

func TextCursorInput(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textCursorInputPath), g.Group(children))
}

func TextQuote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textQuotePath), g.Group(children))
}

func TextSelect(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(textSelectPath), g.Group(children))
}

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thermometerPath), g.Group(children))
}

func ThermometerSnowflake(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thermometerSnowflakePath), g.Group(children))
}

func ThermometerSun(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thermometerSunPath), g.Group(children))
}

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thumbsDownPath), g.Group(children))
}

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(thumbsUpPath), g.Group(children))
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ticketPath), g.Group(children))
}

func Timer(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(timerPath), g.Group(children))
}

func TimerOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(timerOffPath), g.Group(children))
}

func TimerReset(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(timerResetPath), g.Group(children))
}

func ToggleLeft(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(toggleLeftPath), g.Group(children))
}

func ToggleRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(toggleRightPath), g.Group(children))
}

func Tornado(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tornadoPath), g.Group(children))
}

func Touchpad(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(touchpadPath), g.Group(children))
}

func TouchpadOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(touchpadOffPath), g.Group(children))
}

func TowerControl(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(towerControlPath), g.Group(children))
}

func ToyBrick(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(toyBrickPath), g.Group(children))
}

func Tractor(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tractorPath), g.Group(children))
}

func TrafficCone(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trafficConePath), g.Group(children))
}

func TrainFront(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trainFrontPath), g.Group(children))
}

func TrainFrontTunnel(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trainFrontTunnelPath), g.Group(children))
}

func TrainTrack(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trainTrackPath), g.Group(children))
}

func TramFront(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tramFrontPath), g.Group(children))
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trashPath), g.Group(children))
}

func TrashTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trashTwoPath), g.Group(children))
}

func TreeDeciduous(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(treeDeciduousPath), g.Group(children))
}

func TreePine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(treePinePath), g.Group(children))
}

func Trees(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(treesPath), g.Group(children))
}

func Trello(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trelloPath), g.Group(children))
}

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trendingDownPath), g.Group(children))
}

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trendingUpPath), g.Group(children))
}

func Triangle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trianglePath), g.Group(children))
}

func TriangleRight(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(triangleRightPath), g.Group(children))
}

func Trophy(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(trophyPath), g.Group(children))
}

func Truck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(truckPath), g.Group(children))
}

func Turtle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(turtlePath), g.Group(children))
}

func Tv(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tvPath), g.Group(children))
}

func TvTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(tvTwoPath), g.Group(children))
}

func Twitch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(twitchPath), g.Group(children))
}

func Twitter(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(twitterPath), g.Group(children))
}

func Type(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(typePath), g.Group(children))
}

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(umbrellaPath), g.Group(children))
}

func Underline(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(underlinePath), g.Group(children))
}

func Undo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(undoPath), g.Group(children))
}

func UndoDot(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(undoDotPath), g.Group(children))
}

func UndoTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(undoTwoPath), g.Group(children))
}

func UnfoldHorizontal(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unfoldHorizontalPath), g.Group(children))
}

func UnfoldVertical(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unfoldVerticalPath), g.Group(children))
}

func Ungroup(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(ungroupPath), g.Group(children))
}

func Unlink(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unlinkPath), g.Group(children))
}

func UnlinkTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unlinkTwoPath), g.Group(children))
}

func Unlock(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unlockPath), g.Group(children))
}

func Unplug(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(unplugPath), g.Group(children))
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(uploadPath), g.Group(children))
}

func UploadCloud(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(uploadCloudPath), g.Group(children))
}

func Usb(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(usbPath), g.Group(children))
}

func User(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userPath), g.Group(children))
}

func UserCheck(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userCheckPath), g.Group(children))
}

func UserCheckTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userCheckTwoPath), g.Group(children))
}

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userCirclePath), g.Group(children))
}

func UserCircleTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userCircleTwoPath), g.Group(children))
}

func UserCog(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userCogPath), g.Group(children))
}

func UserCogTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userCogTwoPath), g.Group(children))
}

func UserMinus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userMinusPath), g.Group(children))
}

func UserMinusTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userMinusTwoPath), g.Group(children))
}

func UserPlus(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userPlusPath), g.Group(children))
}

func UserPlusTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userPlusTwoPath), g.Group(children))
}

func UserSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userSquarePath), g.Group(children))
}

func UserSquareTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userSquareTwoPath), g.Group(children))
}

func UserTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userTwoPath), g.Group(children))
}

func UserX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userXPath), g.Group(children))
}

func UserXTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(userXTwoPath), g.Group(children))
}

func Users(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(usersPath), g.Group(children))
}

func UsersTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(usersTwoPath), g.Group(children))
}

func Utensils(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(utensilsPath), g.Group(children))
}

func UtensilsCrossed(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(utensilsCrossedPath), g.Group(children))
}

func UtilityPole(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(utilityPolePath), g.Group(children))
}

func Variable(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(variablePath), g.Group(children))
}

func Vegan(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(veganPath), g.Group(children))
}

func VenetianMask(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(venetianMaskPath), g.Group(children))
}

func Verified(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(verifiedPath), g.Group(children))
}

func Vibrate(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(vibratePath), g.Group(children))
}

func VibrateOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(vibrateOffPath), g.Group(children))
}

func Video(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(videoPath), g.Group(children))
}

func VideoOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(videoOffPath), g.Group(children))
}

func Videotape(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(videotapePath), g.Group(children))
}

func View(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(viewPath), g.Group(children))
}

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(voicemailPath), g.Group(children))
}

func Volume(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(volumePath), g.Group(children))
}

func VolumeOne(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(volumeOnePath), g.Group(children))
}

func VolumeTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(volumeTwoPath), g.Group(children))
}

func VolumeX(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(volumeXPath), g.Group(children))
}

func Vote(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(votePath), g.Group(children))
}

func Wallet(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(walletPath), g.Group(children))
}

func WalletCards(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(walletCardsPath), g.Group(children))
}

func WalletTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(walletTwoPath), g.Group(children))
}

func Wallpaper(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wallpaperPath), g.Group(children))
}

func Wand(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wandPath), g.Group(children))
}

func WandTwo(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wandTwoPath), g.Group(children))
}

func Warehouse(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(warehousePath), g.Group(children))
}

func Watch(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(watchPath), g.Group(children))
}

func Waves(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wavesPath), g.Group(children))
}

func Webcam(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(webcamPath), g.Group(children))
}

func Webhook(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(webhookPath), g.Group(children))
}

func Wheat(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wheatPath), g.Group(children))
}

func WheatOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wheatOffPath), g.Group(children))
}

func WholeWord(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wholeWordPath), g.Group(children))
}

func Wifi(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiPath), g.Group(children))
}

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wifiOffPath), g.Group(children))
}

func Wind(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(windPath), g.Group(children))
}

func Wine(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(winePath), g.Group(children))
}

func WineOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wineOffPath), g.Group(children))
}

func Workflow(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(workflowPath), g.Group(children))
}

func WrapText(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wrapTextPath), g.Group(children))
}

func Wrench(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(wrenchPath), g.Group(children))
}

func X(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(xPath), g.Group(children))
}

func XCircle(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(xCirclePath), g.Group(children))
}

func XOctagon(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(xOctagonPath), g.Group(children))
}

func XSquare(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(xSquarePath), g.Group(children))
}

func Youtube(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(youtubePath), g.Group(children))
}

func Zap(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(zapPath), g.Group(children))
}

func ZapOff(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(zapOffPath), g.Group(children))
}

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(zoomInPath), g.Group(children))
}

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(viewbox, hAttr, g.Raw(zoomOutPath), g.Group(children))
}
