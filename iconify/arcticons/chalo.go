package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chalo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.506 25.868c0 1.511 1.209 2.72 2.62 2.72a2.709 2.709 0 0 0 2.72-2.72v-2.72c0-1.51-1.21-2.72-2.72-2.72s-2.62 1.21-2.62 2.72v2.72Zm-5.369-5.339v8.059h4.029m-5.369 0l-2.619-8.059l-2.72 8.059m.907-2.72h3.526M14.78 20.529v8.059m5.339-8.059v8.059m-5.339-4.029h5.339m-6.679 1.309h0a2.709 2.709 0 0 1-2.72 2.72h0A2.709 2.709 0 0 1 8 25.868v-2.72a2.709 2.709 0 0 1 2.72-2.72h0c1.511 0 2.62 1.21 2.62 2.72h0m24.645-3.736H40v2.015"/>`),
		g.Group(children),
	)
}