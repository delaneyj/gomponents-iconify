package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.016 3.984h4.981L11.016.016v3.968zm-4.809 7.102L4.5 12.793l-1.707-1.707l-.707.707L3.793 13.5l-1.707 1.707l.707.707L4.5 14.207l1.707 1.707l.707-.707L5.207 13.5l1.707-1.707l-.707-.707z"/><path d="M9.969 5.016V.011H4.034v10.4l.466.466l1.908-1.906l2.623 2.623L7.125 13.5l2.47 2.469h6.382V5.016H9.969z"/></g>`),
		g.Group(children),
	)
}