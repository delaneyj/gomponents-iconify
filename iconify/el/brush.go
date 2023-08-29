package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1157.602.013c-46.711 2.677-736.479 591.498-793.123 798.838l130.736 136.944C868.899 624.199 988.915 294.221 1200 .649L1157.617 0l-.015.013zM323.267 840.562C87.09 927.418 235.147 1099.273 0 1183.352c266.294 59.953 384.296-49.748 454.003-205.421L323.267 840.548v.014z"/>`),
		g.Group(children),
	)
}