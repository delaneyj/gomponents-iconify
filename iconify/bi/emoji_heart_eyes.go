package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiHeartEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/><path d="M11.315 10.014a.5.5 0 0 1 .548.736A4.498 4.498 0 0 1 7.965 13a4.498 4.498 0 0 1-3.898-2.25a.5.5 0 0 1 .548-.736h.005l.017.005l.067.015l.252.055c.215.046.515.108.857.169c.693.124 1.522.242 2.152.242c.63 0 1.46-.118 2.152-.242a26.58 26.58 0 0 0 1.109-.224l.067-.015l.017-.004l.005-.002zM4.756 4.566c.763-1.424 4.02-.12.952 3.434c-4.496-1.596-2.35-4.298-.952-3.434zm6.488 0c1.398-.864 3.544 1.838-.952 3.434c-3.067-3.554.19-4.858.952-3.434z"/></g>`),
		g.Group(children),
	)
}