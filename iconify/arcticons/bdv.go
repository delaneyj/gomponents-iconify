package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bdv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.16 42.5l15.415-8.23l3.56-26.758c.172-1.289-1.195-2.225-2.335-1.599l-12.842 7.055a2.604 2.604 0 0 0-1.34 2.05L22.16 42.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.067 14.515L22.16 42.5l15.415-8.23L22.848 6.465a1.814 1.814 0 0 0-2.42-.77l-12.58 6.34a1.814 1.814 0 0 0-.78 2.48Z"/>`),
		g.Group(children),
	)
}