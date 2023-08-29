package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wpr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m3.993 5.425l4.854 17.382l.035.15h4.186l1.87-6.932l1.87 6.932h4.185l.036-.15l1.96-7.078l3.435 12.41A15.937 15.937 0 0 1 16 32C7.163 32 0 24.837 0 16A15.94 15.94 0 0 1 3.993 5.425zm2.65-2.405A15.928 15.928 0 0 1 16 0c8.837 0 16 7.163 16 16a15.93 15.93 0 0 1-3.025 9.364L24.901 10.65l-.035-.15h-3.75l-2.336 8.438l-1.973-8.438h-3.739l-1.973 8.438L6.642 3.02z"/>`),
		g.Group(children),
	)
}