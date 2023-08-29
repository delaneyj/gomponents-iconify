package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plunger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#FBB8AB" d="M29.414 2.586a2 2 0 0 1 0 2.828L17.328 17.5L14 18l.5-3.328L26.586 2.586a2 2 0 0 1 2.828 0Z"/><path fill="#F8312F" d="M14.828 24.95a4.502 4.502 0 0 0 .629-5.578l1.957-1.958l-2.828-2.828l-1.958 1.957a4.502 4.502 0 0 0-5.578.629l-1.525 1.525l-.11-.111a2 2 0 0 0-2.83 2.828l8 8a2 2 0 1 0 2.83-2.828l-.112-.111l1.525-1.525Z"/></g>`),
		g.Group(children),
	)
}