package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurlingStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#BE1931" d="M28 12H13V8a1 1 0 0 1 1-1h11a2 2 0 0 0 0-4H14a5 5 0 0 0-5 5v4H8a2 2 0 0 0 0 4h20a2 2 0 0 0 0-4z"/><path fill="#66757F" d="M36 25a6 6 0 0 1-6 6H6a6 6 0 0 1-6-6v-5a6 6 0 0 1 6-6h24a6 6 0 0 1 6 6v5z"/><path fill="#99AAB5" d="M0 20h36v5H0z"/>`),
		g.Group(children),
	)
}