package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="iconoirEmojiSingRight0" fill="currentColor" d="M8.5 9a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm7 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><use href="#iconoirEmojiSingRight0"/><use href="#iconoirEmojiSingRight0"/></g>`),
		g.Group(children),
	)
}