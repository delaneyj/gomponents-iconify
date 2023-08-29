package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchCircleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 64C150.13 64 64 150.13 64 256s86.13 192 192 192s192-86.13 192-192S361.87 64 256 64Zm80 294.63l-54.15-54.15a88.08 88.08 0 1 1 22.63-22.63L358.63 336Z"/><circle cx="232" cy="232" r="56" fill="currentColor"/>`),
		g.Group(children),
	)
}