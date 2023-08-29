package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoolDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 40v24a8 8 0 0 1-8 8H72a8 8 0 0 1-8-8V40a8 8 0 0 1 8-8h112a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M200 64V40a16 16 0 0 0-16-16H72a16 16 0 0 0-16 16v24a16 16 0 0 0 16 16h6.64L56.1 222.75a8 8 0 0 0 6.65 9.15a7.82 7.82 0 0 0 1.25.1a8 8 0 0 0 7.89-6.75L79.68 176h96.64l7.78 49.25A8 8 0 0 0 192 232a7.82 7.82 0 0 0 1.26-.1a8 8 0 0 0 6.65-9.15L177.36 80H184a16 16 0 0 0 16-16ZM72 40h112v24H72Zm101.79 120H82.21l12.63-80h66.32Z"/></g>`),
		g.Group(children),
	)
}