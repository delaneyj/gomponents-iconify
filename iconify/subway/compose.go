package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 28.3H0v469.3h85.3V263H64V28.3zm149.3 0h-64V263H128v234.7h106.7V263h-21.3V28.3zm234.7 0V263h-21.3v234.7H512V28.3h-64zm-85.3 0h-64V263h-21.3v234.7H384V263h-21.3V28.3z"/>`),
		g.Group(children),
	)
}