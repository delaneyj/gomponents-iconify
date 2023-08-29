package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 4H9v5h10V5a1 1 0 0 0-1-1zM1 15c0 .553.248 1 .8 1H7v-5H1v4zm8 1h9a1 1 0 0 0 1-1v-4H9v5zM1 5v4h6V4H1.8c-.552 0-.8.447-.8 1z"/>`),
		g.Group(children),
	)
}