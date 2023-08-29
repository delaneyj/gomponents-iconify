package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrustumOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7.72 3.728l3.484-1.558a1.95 1.95 0 0 1 1.59 0l4.496 2.01c.554.246.963.736 1.112 1.328l2.538 10.158c.103.412.07.832-.075 1.206m-2.299 1.699l-5.725 2.738a1.945 1.945 0 0 1-1.682 0l-7.035-3.365a1.99 1.99 0 0 1-1.064-2.278l2.52-10.08"/><path d="m18 4.82l-5.198 2.324a1.963 1.963 0 0 1-1.602 0m.8.176V8m0 4v9.5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}