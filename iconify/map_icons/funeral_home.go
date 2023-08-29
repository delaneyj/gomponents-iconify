package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FuneralHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25.242 2.089L1 23h6v25h36V23h6L25.242 2.089zM34 30h-7v12.088L26.954 42h-4.775l-.179.088V30h-6v-4h6v-9h5v9h7v4z"/>`),
		g.Group(children),
	)
}