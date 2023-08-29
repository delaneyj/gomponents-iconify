package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M339 99a83 83 0 1 0-102 80.8V464l19 32l19-32V179.8A83.28 83.28 0 0 0 339 99Zm-59-6a21 21 0 1 1 21-21a21 21 0 0 1-21 21Z"/>`),
		g.Group(children),
	)
}