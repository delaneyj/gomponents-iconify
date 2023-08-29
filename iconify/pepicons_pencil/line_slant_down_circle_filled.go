package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSlantDownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilLineSlantDownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M3.454 3.454a.5.5 0 0 1 .707 0l18.385 18.385a.5.5 0 0 1-.707.707L3.454 4.16a.5.5 0 0 1 0-.707Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilLineSlantDownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}