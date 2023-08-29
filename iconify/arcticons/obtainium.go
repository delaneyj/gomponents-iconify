package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Obtainium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.862 41.676L43.5 11.618L31.232 6.324l-8.681 3.448l3.56 8.965l5.504-2.186l-3.778 12.471l-11.306-6.48l5.504-2.186l-3.56-8.965l-8.682 3.448L4.5 27.107l28.362 14.569z"/>`),
		g.Group(children),
	)
}