package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13zM6 11a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 11zm4-5.5c0-.828.448-1.5 1-1.5s1 .672 1 1.5S11.552 7 11 7s-1-.672-1-1.5zm-6 0C4 4.672 4.448 4 5 4s1 .672 1 1.5S5.552 7 5 7s-1-.672-1-1.5z"/>`),
		g.Group(children),
	)
}