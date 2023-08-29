package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 172a28 28 0 1 1-28-28a28 28 0 0 1 28 28Zm-156 4a28 28 0 1 0 28 28a28 28 0 0 0-28-28Z" opacity=".2"/><path d="M212.92 25.69a8 8 0 0 0-6.86-1.45l-128 32A8 8 0 0 0 72 64v110.08A36 36 0 1 0 88 204V70.25l112-28v99.83A36 36 0 1 0 216 172V32a8 8 0 0 0-3.08-6.31ZM52 224a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm128-32a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/></g>`),
		g.Group(children),
	)
}