package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundEco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.95 5.97a2.018 2.018 0 0 0-1.92-1.92c-.7-.03-1.37-.05-2.02-.05c-5.03 0-8.52.97-10.46 2.91c-3.68 3.68-3.15 8.9.09 11.9h.01a16 16 0 0 1 7.67-8.8c-.21.18-4.7 3.58-5.51 10.25c1.05.48 2.2.75 3.36.75c2.05 0 4.16-.8 5.92-2.55c2.19-2.2 3.14-6.36 2.86-12.49z"/>`),
		g.Group(children),
	)
}