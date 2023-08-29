package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2v1.083A6.002 6.002 0 0 0 6 9a5 5 0 0 0-5 5v8h22v-8a5 5 0 0 0-5-5a6.002 6.002 0 0 0-5-5.917V2h-2Zm7 9a3 3 0 0 1 3 3v6h-3v-9Zm-2 9h-3v-5h-2v5H8V10h8v10ZM6 20H3v-6a3 3 0 0 1 3-3v9ZM8.126 8a4.002 4.002 0 0 1 7.748 0H8.126Z"/>`),
		g.Group(children),
	)
}