package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18H4v-8h5v8Zm-2-2v-4H6v4h1Zm6 0V8h-1v8h1Zm2 2h-5V6h5v12Zm4-2V4h-1v12h1Zm2 2h-5V2h5v16Zm1 4H3v-2h19v2Z"/>`),
		g.Group(children),
	)
}