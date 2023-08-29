package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mahjonggbuilder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.3 5.54L31 7.73a2.62 2.62 0 0 1 2.08 2.63v31.37A1.61 1.61 0 0 1 31 43.44l-17.36-4.63a2.27 2.27 0 0 1-2-2.39V7.05c.04-.68.36-1.4 1.66-1.51Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.56 43.13l3.21-2.66a1.55 1.55 0 0 0 .39-1.24l.16-30.31a2.81 2.81 0 0 0-2.69-3L21 4.56a12.3 12.3 0 0 0-3.59.13l-4.07.85m7.16 28.1L22.25 37l1.9-1.44Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.55 23.31A10.5 10.5 0 0 1 18 18.49c0-2.42 4.25-2.62 4.25-2.62m.03 0c2.62-.2 5.8.18 5.71 2.3s-1.39 4.1-2.33 5.48m-5.34-9.4a7.94 7.94 0 0 1 5-1.11m-4.45 7.5l3.69-.71m-2.17 3.38a18.18 18.18 0 0 1-7.24 7.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.39 23.31l4.76 7.14l2.65-1.35c-2.58-1.48-5.05-4.03-7.41-5.79Zm-.11-12.66l.17 24.01"/>`),
		g.Group(children),
	)
}