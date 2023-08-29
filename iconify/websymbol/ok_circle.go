package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OkCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M500 0q136 0 251 67t182 182t67 251t-67 251t-182 182t-251 67t-251-67T67 751T0 500t67-251T249 67T500 0zM260 447l-54 53q71 71 129 133.5t81 89.5l22 26q8-13 22-35.5t56.5-85.5T600 509.5T695 390t99-104q-18-35-36-35q-10 6-32 22t-113.5 97.5T402 571z"/>`),
		g.Group(children),
	)
}