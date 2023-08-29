package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M424 496h-35.25L256.008 381.19L123.467 496H88V16h336ZM120 48v408.667l135.992-117.8L392 456.5V48Z"/>`),
		g.Group(children),
	)
}