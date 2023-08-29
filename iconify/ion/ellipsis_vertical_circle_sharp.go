package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisVerticalCircleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 256c0-114.87-93.13-208-208-208S48 141.13 48 256s93.13 208 208 208s208-93.13 208-208Zm-234-90a26 26 0 1 1 26 26a26 26 0 0 1-26-26Zm0 90a26 26 0 1 1 26 26a26 26 0 0 1-26-26Zm0 90a26 26 0 1 1 26 26a26 26 0 0 1-26-26Z"/>`),
		g.Group(children),
	)
}