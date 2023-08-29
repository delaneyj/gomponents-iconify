package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undertake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.861 8.777a1.379 1.379 0 0 0-.976.402l-2.184 2.184V15.5h6.013l6.148-1.537l3.741-1.596a.645.645 0 0 0-.48-1.188l-.02.006l-6.773 1.557h-3.757v-2h3.247a.983.983 0 0 0 0-1.965H8.861Zm7.919 1.35l3.836-.883a2.647 2.647 0 0 1 1.86 4.924l-.027.013l-3.948 1.684l-6.54 1.635H0V9.95h4.286l2.187-2.187a3.38 3.38 0 0 1 2.392-.986h.001h-.002h4.956a2.983 2.983 0 0 1 2.96 3.35ZM3.7 11.949H2V15.5h1.7v-3.55Z"/>`),
		g.Group(children),
	)
}