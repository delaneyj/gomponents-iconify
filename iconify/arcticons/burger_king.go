package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BurgerKing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 14.55a.78.78 0 0 1-.77-.9C9 8.43 16.42 4.5 24 4.5s15 3.93 15.81 9.15a.78.78 0 0 1-.77.9Zm30.1 23.27a.73.73 0 0 1 .71.85c-.76 4.81-8.22 4.83-15.81 4.83s-15 0-15.81-4.83a.73.73 0 0 1 .71-.85ZM17.91 24.36v11.55m21.91-7.7a3.86 3.86 0 0 0-4.07-3.86a4 4 0 0 0-3.65 4.08V32a3.86 3.86 0 0 0 3.9 3.89h0A3.86 3.86 0 0 0 39.82 32H36m-14.74 3.89V24.34l7.65 11.55V24.34m-20.73.01v11.54m6.2 0l-4.75-5.77l4.75-5.74m-4.75 5.74H8.18m22.86-7.69h2.98m-2.98-5.97h2.98m-2.98 2.99h1.95m-1.95-2.99v5.97m-1.74-3.97a2 2 0 0 0-2.11-2a2.08 2.08 0 0 0-1.88 2.12v1.85a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2h-2m-7.7 2v-6h1.95a2 2 0 0 1 0 4h-1.95m1.95.04l1.96 1.96m12.35 0v-6h2a2 2 0 0 1 0 4h-2m1.95.04l1.96 1.96M13.9 16.46v4a2 2 0 1 0 4 0v-4m-7.26 2.99a1.49 1.49 0 0 1 0 3H8.18v-6h2.46a1.5 1.5 0 0 1 0 3Zm0 0H8.18"/>`),
		g.Group(children),
	)
}