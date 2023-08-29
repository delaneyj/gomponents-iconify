package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 9a8 8 0 1 1 16 0v2h-6.981v9.5a2.5 2.5 0 0 1-5 0v-2.643h2V20.5a.5.5 0 1 0 1 0V11H4V9Zm8-6a6 6 0 0 1 6 6H6a6 6 0 0 1 6-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}