package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M13.75 11.486L9.972 18h7.556l-3.778-6.514Zm.865-2.495a1 1 0 0 0-1.73 0l-5.514 9.507A1 1 0 0 0 8.236 20h11.028a1 1 0 0 0 .865-1.502l-5.514-9.507Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M13 7a.5.5 0 0 1 .429.243l6 10A.5.5 0 0 1 19 18H7a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 13 7ZM7.883 17h10.234L13 8.472L7.883 17Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}