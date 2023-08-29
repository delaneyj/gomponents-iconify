package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundPushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#D3D3D3" d="M15.93 29.92c-.56 0-1.01-.45-1.01-1.01V15.94h2.03V28.9c-.01.57-.46 1.02-1.02 1.02Z"/><path fill="#F70A8D" d="M15.93 17.86a7.93 7.93 0 1 0 0-15.86a7.93 7.93 0 0 0 0 15.86Z"/><path fill="#fff" d="M16.5 5.41c-1.04 1.04-1.09 2.66-.12 3.63c.97.97 2.6.92 3.63-.12c1.04-1.04 1.09-2.66.12-3.63c-.97-.97-2.6-.92-3.63.12Z"/></g>`),
		g.Group(children),
	)
}