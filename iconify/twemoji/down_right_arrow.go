package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M0 4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4z"/><path fill="#FFF" d="M27 11v16H11z"/><path fill="#FFF" d="M7 12.657L12.658 7l13.814 13.814l-5.656 5.657z"/>`),
		g.Group(children),
	)
}