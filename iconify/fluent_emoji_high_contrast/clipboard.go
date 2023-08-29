package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M9 12.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13ZM9 21.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8a.5.5 0 0 1-.5-.5Z"/><path d="M18 3a2 2 0 1 0-4 0h-1a2 2 0 0 0-2 2H8a1 1 0 0 0-1 1v21a1 1 0 0 0 1 1h10.293a1 1 0 0 0 .707-.293L24.707 22a1 1 0 0 0 .293-.707V6a1 1 0 0 0-1-1h-3a2 2 0 0 0-2-2h-1Zm-7 3v.5a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V6h3v15h-4.5a1.5 1.5 0 0 0-1.5 1.5V27H8V6h3Zm6-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm6.293 19L19 26.293V22.5a.5.5 0 0 1 .5-.5h3.793Z"/><path d="M25.5 4H22a2 2 0 0 0-2-2h5.5A2.5 2.5 0 0 1 28 4.5v24a2.5 2.5 0 0 1-2.5 2.5h-19A2.5 2.5 0 0 1 4 28.5v-24A2.5 2.5 0 0 1 6.5 2H12a2 2 0 0 0-2 2H6.5a.5.5 0 0 0-.5.5v24a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 .5-.5v-24a.5.5 0 0 0-.5-.5Z"/></g>`),
		g.Group(children),
	)
}