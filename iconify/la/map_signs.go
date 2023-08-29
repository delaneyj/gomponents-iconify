package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSigns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 5v2H5v9h20.469l.281-.344l3.563-4.156l-3.563-4.156L25.469 7H17V5zM7 9h17.531l2.157 2.5L24.53 14H7zm8 8v10h2V17z"/>`),
		g.Group(children),
	)
}