package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NegativeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-30-1a1 1 0 1 0 0 2h20a1 1 0 1 0 0-2H14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}