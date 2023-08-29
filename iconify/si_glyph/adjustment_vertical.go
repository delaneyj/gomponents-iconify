package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustmentVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 0v5.007h1.95V0H2zm0 11v4.958h2.01V11H2zm6 3v1.976h1.966V14H8zM8 0v8h2V0H8zm6 9v6.942h1.977V9H14zm0-9v2.933h2.009V0H14zm.917 8.049c1.059 0 2.094-.994 2.094-2.081S15.976 4 14.917 4S13 4.881 13 5.968s.858 2.081 1.917 2.081zm-6 4.961c1.059 0 2.04-1.051 2.04-2.104C10.958 9.853 9.977 9 8.918 9A1.912 1.912 0 0 0 7 10.906c0 1.053.858 2.103 1.917 2.103zm-6-3.021c1.059 0 2.057-1.013 2.057-2.072C4.974 6.858 3.976 6 2.917 6A1.917 1.917 0 0 0 1 7.917c0 1.059.858 2.072 1.917 2.072z"/>`),
		g.Group(children),
	)
}