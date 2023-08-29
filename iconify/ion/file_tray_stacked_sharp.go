package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTrayStackedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 16H64L32 176v144h448V176Zm-12 160H320a64 64 0 0 1-128 0H76L98 58h316ZM320 352a64 64 0 0 1-128 0H32v144h448V352Z"/>`),
		g.Group(children),
	)
}