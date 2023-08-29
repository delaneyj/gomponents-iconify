package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CountdownCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M22 14a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z" opacity=".2"/><path fill-rule="evenodd" d="M13 6a.5.5 0 0 1 .5-.5A7.5 7.5 0 1 1 6 13a.5.5 0 0 1 1 0a6.5 6.5 0 1 0 6.5-6.5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M11 6.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-2.5 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM7 10.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path fill-rule="evenodd" d="M8.947 14.224a.5.5 0 0 0-.223-.671l-2-1a.5.5 0 0 0-.448.894l2 1a.5.5 0 0 0 .671-.223Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.854 12.646a.5.5 0 0 1 0 .708l-1.5 1.5a.5.5 0 1 1-.708-.707l1.5-1.5a.5.5 0 0 1 .708 0ZM13.5 9.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 13a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}