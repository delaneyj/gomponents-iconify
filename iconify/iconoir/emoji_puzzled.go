package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiPuzzled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"/><path d="M11.5 15.5s1.5-2 4.5-2s4.5 2 4.5 2M3 4c0-2.754 4-2.754 4 0c0 1.967-2 1.64-2 4m0 3.01l.01-.011"/><path fill="currentColor" d="M17.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm-7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`),
		g.Group(children),
	)
}