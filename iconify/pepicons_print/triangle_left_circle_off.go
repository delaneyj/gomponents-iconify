package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="m10.986 13.75l6.514 3.778V9.972l-6.514 3.778Zm-2.495-.865a1 1 0 0 0 0 1.73l9.507 5.514a1 1 0 0 0 1.502-.865V8.236a1 1 0 0 0-1.502-.865l-9.507 5.514Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M7 13a.5.5 0 0 1 .243-.429l10-6A.5.5 0 0 1 18 7v12a.5.5 0 0 1-.757.429l-10-6A.5.5 0 0 1 7 13Zm10 5.117V7.883L8.472 13L17 18.117Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}