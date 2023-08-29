package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisHorizontalCircleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48C141.13 48 48 141.13 48 256s93.13 208 208 208s208-93.13 208-208S370.87 48 256 48Zm-90 234a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm90 0a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm90 0a26 26 0 1 1 26-26a26 26 0 0 1-26 26Z"/>`),
		g.Group(children),
	)
}