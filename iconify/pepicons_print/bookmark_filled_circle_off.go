package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkFilledCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="m15 17.676l-4.645 4.676A.5.5 0 0 1 9.5 22V8a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.855.352L15 17.676Z" opacity=".2"/><path fill-rule="evenodd" d="M8.355 20.352L13 15.676l4.645 4.676A.5.5 0 0 0 18.5 20V6a.5.5 0 0 0-.5-.5H8a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .855.352Zm.145-1.565V6.5h9v12.287l-4.145-4.172a.5.5 0 0 0-.71 0L8.5 18.787Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}