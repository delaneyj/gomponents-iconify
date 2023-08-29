package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m7.986 10.75l6.514 3.778V6.972L7.986 10.75Zm-2.495-.865a1 1 0 0 0 0 1.73l9.507 5.514a1 1 0 0 0 1.502-.865V5.236a1 1 0 0 0-1.502-.865L5.491 9.885Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M4 10a.5.5 0 0 1 .243-.429l10-6A.5.5 0 0 1 15 4v12a.5.5 0 0 1-.757.429l-10-6A.5.5 0 0 1 4 10Zm10 5.117V4.883L5.472 10L14 15.117Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}