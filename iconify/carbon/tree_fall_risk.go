package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeFallRisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M23.75 10h-1.5V6h1.5ZM23 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M29.912 13.935L23.628 2.371a.718.718 0 0 0-1.256 0l-6.284 11.564A.72.72 0 0 0 16.72 15h12.56a.72.72 0 0 0 .631-1.065zM22.25 6h1.5v4h-1.5zm.75 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm3 6v-2h-9v11a2.003 2.003 0 0 0 2 2h3v-2h-3v-9zM12 30H9v-2h3V15.566l-3.515-2.109l1.03-1.714l3.514 2.108A2.011 2.011 0 0 1 14 15.566V28a2.002 2.002 0 0 1-2 2z"/><path fill="currentColor" d="m18.664 5.006l.96-1.767A8.932 8.932 0 0 0 15 2a8.027 8.027 0 0 0-6.921 4.001L8 6a6 6 0 0 0 0 12v-2a4 4 0 0 1 0-8a2.7 2.7 0 0 1 .387.04l.863.113l.309-.667A6.02 6.02 0 0 1 15 4a6.892 6.892 0 0 1 3.664 1.006Z"/>`),
		g.Group(children),
	)
}