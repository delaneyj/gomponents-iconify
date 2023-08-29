package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSlantUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1.6 19.278a1.5 1.5 0 0 1 0-2.121L17.158 1.6a1.5 1.5 0 1 1 2.121 2.122L3.722 19.278a1.5 1.5 0 0 1-2.121 0Z" opacity=".2"/><path d="M.454 19.546a.5.5 0 0 1 0-.707L18.839.454a.5.5 0 1 1 .707.707L1.16 19.546a.5.5 0 0 1-.707 0Z"/></g>`),
		g.Group(children),
	)
}