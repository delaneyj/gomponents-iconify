package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoilMoistureField(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.5 30a5.202 5.202 0 0 1-4.626-8.08l3.616-5.382a1.217 1.217 0 0 1 2.02 0l3.55 5.277a5.492 5.492 0 0 1 .94 2.936A5.385 5.385 0 0 1 24.5 30zm0-11.38l-2.936 4.367A3.208 3.208 0 0 0 24.5 28a3.385 3.385 0 0 0 3.5-3.249a3.435 3.435 0 0 0-.63-1.867zM11 16v-5h1a4.005 4.005 0 0 0 4-4V4h-3a3.978 3.978 0 0 0-2.747 1.107A6.003 6.003 0 0 0 5 2H2v3a6.007 6.007 0 0 0 6 6h1v5H2v2h14v-2zm2-10h1v1a2.002 2.002 0 0 1-2 2h-1V8a2.002 2.002 0 0 1 2-2zM8 9a4.005 4.005 0 0 1-4-4V4h1a4.005 4.005 0 0 1 4 4v1zM2 21h14v2H2zm0 5h14v2H2z"/>`),
		g.Group(children),
	)
}