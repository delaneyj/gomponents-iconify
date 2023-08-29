package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="M20.45 18.777L24 11.512l3.55 7.265l.898-.44l-.899.44c.289.59.849 1.004 1.503 1.1l7.946 1.167v.001a.031.031 0 0 1 .002.008l-.003.004l-5.753 5.665a2.011 2.011 0 0 0-.57 1.77l1.358 7.998v.003l-7.1-3.77a1.986 1.986 0 0 0-1.864 0l-7.1 3.77v-.003l1.358-7.999a2.011 2.011 0 0 0-.57-1.77l-5.753-5.664l-.003-.003c0-.003 0-.006.002-.009v-.002l7.946-1.166a1.996 1.996 0 0 0 1.503-1.1Z"/>`),
		g.Group(children),
	)
}