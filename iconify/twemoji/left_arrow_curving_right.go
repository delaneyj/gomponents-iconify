package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrowCurvingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M0 32a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4a4 4 0 0 0-4 4v28z"/><path fill="#FFF" d="M29.2 25L22 19v4h-4a7 7 0 0 1 0-14h8V5h-8C11.926 5 7 9.925 7 16c0 6.074 4.926 11 11 11h4v4l7.2-6z"/>`),
		g.Group(children),
	)
}