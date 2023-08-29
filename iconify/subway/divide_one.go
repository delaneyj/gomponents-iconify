package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivideOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 228.3v85.3h512v-85.3H0zM298.7 57.7h-85.3V143h85.3V57.7zm-85.4 426.6h85.3V399h-85.3v85.3z"/>`),
		g.Group(children),
	)
}