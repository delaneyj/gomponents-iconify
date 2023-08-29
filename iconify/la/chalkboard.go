package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chalkboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 7v16H3v2h26v-2h-2V7zm2 2h18v14H7zm14.281 3.281L17 16.562l-3.281-3.28l-.719-.688l-.719.687l-3 3l1.438 1.438L13 15.437l3.281 3.282l.719.687l.719-.687l5-5zM20 20l-1 1l1 1h4v-2z"/>`),
		g.Group(children),
	)
}