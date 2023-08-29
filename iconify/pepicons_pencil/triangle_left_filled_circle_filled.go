package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTriangleLeftFilledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.501 13.866a1 1 0 0 1 0-1.732l10-5.769A1 1 0 0 1 19 7.231V18.77a1 1 0 0 1-1.5.866l-9.999-5.769Z"/><path d="M17.5 6.365a1 1 0 0 1 1.5.866V18.77a1 1 0 0 1-1.5.866l-9.999-5.769a1 1 0 0 1 0-1.732l10-5.769ZM12.003 13L16 15.306v-4.612L12.003 13Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTriangleLeftFilledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}