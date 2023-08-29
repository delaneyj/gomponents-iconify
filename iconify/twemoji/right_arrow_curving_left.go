package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M18 5h-8v4h8.01A7 7 0 0 1 18 23h-4v-4l-7.2 6l7.2 6v-4h4c6.074 0 11-4.926 11-11c0-6.075-4.926-11-11-11z"/>`),
		g.Group(children),
	)
}