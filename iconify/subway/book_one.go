package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M277.3 50.2v426.7c64-53.3 170.7-53.3 234.7 0V50.2c-64-53.3-170.7-53.3-234.7 0zM0 50.2v426.7c64-53.3 170.7-53.3 234.7 0V50.2C170.7-3.1 64-3.1 0 50.2z"/>`),
		g.Group(children),
	)
}