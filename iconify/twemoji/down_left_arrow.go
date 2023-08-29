package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 4a4 4 0 0 0-4-4H4a4 4 0 0 0-4 4v28a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V4z"/><path fill="#FFF" d="M9 11v16h16z"/><path fill="#FFF" d="M9.53 20.814L23.343 7L29 12.657L15.186 26.471z"/>`),
		g.Group(children),
	)
}