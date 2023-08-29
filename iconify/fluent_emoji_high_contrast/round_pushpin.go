package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundPushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16.695 9.445c-.97-.97-.92-2.59.12-3.63c1.03-1.04 2.66-1.09 3.63-.12c.97.97.92 2.59-.12 3.63c-1.03 1.04-2.66 1.09-3.63.12Z"/><path d="M7 9.93a8.93 8.93 0 1 1 9.95 8.872V28.9c-.01.57-.46 1.02-1.02 1.02s-1.01-.45-1.01-1.01V18.803A8.931 8.931 0 0 1 7 9.93ZM15.93 3a6.93 6.93 0 1 0 0 13.86a6.93 6.93 0 0 0 0-13.86Z"/></g>`),
		g.Group(children),
	)
}