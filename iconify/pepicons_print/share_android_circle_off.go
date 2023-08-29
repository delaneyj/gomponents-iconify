package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAndroidCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M10.175 17.207a2.5 2.5 0 1 1 .233-4.274l5.093-3.01a2.5 2.5 0 1 1 .434 1.487l-4.593 2.714a2.494 2.494 0 0 1-.09 1.963l4.658 2.54a2.5 2.5 0 1 1-.408 1.486l-5.327-2.905Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M8 15.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm9-1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0 14a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="m9.754 12.18l-.508-.86l5.5-3.25l.508.86l-5.5 3.25ZM15 17.878l.479-.878l-5.5-3l-.479.878l5.5 3Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}