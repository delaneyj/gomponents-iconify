package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FotScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.9 418.9V279.3l-46.5 46.5l-46.5-46.5l-46.5 46.5l46.5 46.5l-46.5 46.5h139.5zM0 0v512h512V0H0zm465.5 465.5h-419v-419h418.9v419zM186.2 232.7l46.5-46.5l-46.5-46.5l46.5-46.5H93.1v139.6l46.5-46.5l46.6 46.4z"/>`),
		g.Group(children),
	)
}