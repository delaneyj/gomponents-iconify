package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldingStoolCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path fill-rule="evenodd" d="M8.25 14.5a1.5 1.5 0 0 1 1.5-1.5h10a1.5 1.5 0 0 1 0 3h-10a1.5 1.5 0 0 1-1.5-1.5Zm2.5-10a1 1 0 0 1 1 1v5a1 1 0 0 1-2 0v-5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.75 9.5a1 1 0 0 1 1 1V13a1 1 0 1 1-2 0v-2.5a1 1 0 0 1 1-1Zm8 0a1 1 0 0 1 1 1V13a1 1 0 1 1-2 0v-2.5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18.75 4.5a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1Zm1.825 17.565a1 1 0 0 1-1.39.26l-9.5-6.5a1 1 0 1 1 1.13-1.65l9.5 6.5a1 1 0 0 1 .26 1.39Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.925 22.065a1 1 0 0 0 1.39.26l9.5-6.5a1 1 0 1 0-1.13-1.65l-9.5 6.5a1 1 0 0 0-.26 1.39ZM18.75 8h-8V6h8v2Zm0 2.5h-8v-2h8v2Zm-12 1a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1Zm12 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M10.75 10.5h8v3h-8v-3Zm0-4h8v3h-8v-3Z"/></g><path fill-rule="evenodd" d="M7 14a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Zm2-9.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9 9.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V10a.5.5 0 0 1 .5-.5Zm8 0a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V10a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 4.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Zm1.413 16.782a.5.5 0 0 1-.695.13l-9.5-6.5a.5.5 0 0 1 .564-.825l9.5 6.5a.5.5 0 0 1 .13.695Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.587 21.282a.5.5 0 0 0 .695.13l9.5-6.5a.5.5 0 0 0-.564-.825l-9.5 6.5a.5.5 0 0 0-.13.695ZM17 7H9V6h8v1Zm0 3H9V9h8v1ZM5.5 11a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Zm11 0a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}