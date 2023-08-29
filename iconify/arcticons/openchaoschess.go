package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openchaoschess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="15.371" fill="none" stroke="currentColor" stroke-linecap="square" stroke-miterlimit="15.118"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 2.5l1.953 3.889h-3.906L24 2.5zm0 3.889v3.83m15.203-1.422l-1.369 4.131l-2.762-2.762l4.131-1.369zm-2.75 2.75l-2.708 2.708M45.5 24l-3.889 1.953v-3.906L45.5 24zm-3.889 0h-3.83m1.422 15.203l-4.131-1.369l2.762-2.762l1.369 4.131zm-2.75-2.75l-2.708-2.708M24 45.5l-1.953-3.889h3.906L24 45.5zm0-3.889v-3.83M8.797 39.203l1.369-4.131l2.762 2.762l-4.131 1.369zm2.75-2.75l2.708-2.708M2.5 24l3.889-1.953v3.906L2.5 24zm3.889 0h3.83M8.797 8.797l4.131 1.369l-2.762 2.762l-1.369-4.131zm2.75 2.75l2.708 2.708m7.362 5.863v0c0-1.165.552-2.109 1.231-2.109h2.327c.68 0 1.231.944 1.231 2.11v0c0 1.164-.551 2.108-1.23 2.108h-2.328c-.68 0-1.23-.944-1.23-2.109Zm1.547 2.109h1.695v1.834h-1.695z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.671 34.5H20.35a.893.893 0 0 1-.979-.962a3.278 3.278 0 0 1 2.548-3.222l1.018-.309h0v-5.946h2.15v5.946h0l1.016.31a3.257 3.257 0 0 1 2.526 3.33c-.027.584-.341.811-.958.853Zm-3.659-21v4.509m-2.031-2.342h4.062m-4.426 8.394h4.789"/>`),
		g.Group(children),
	)
}