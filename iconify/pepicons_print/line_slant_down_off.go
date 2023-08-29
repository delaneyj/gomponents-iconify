package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSlantDownOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M1.6 1.6a1.5 1.5 0 0 1 2.122 0l15.556 15.557a1.5 1.5 0 1 1-2.121 2.121L1.6 3.722a1.5 1.5 0 0 1 0-2.121Z" opacity=".2"/><path d="M.454.454a.5.5 0 0 1 .707 0l18.385 18.385a.5.5 0 0 1-.707.707L.454 1.16a.5.5 0 0 1 0-.707Z"/></g><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}