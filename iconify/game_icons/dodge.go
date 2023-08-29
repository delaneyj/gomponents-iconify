package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dodge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M465.074 481.646c-24.928 1.186-334.495-.667-334.495-.667c-188.07-188.857 283.16-251.37-8.1-406.484l-20.382 21.487l-27.14-71.216l87.582 6.818l-17.534 19.7c364.276 106.15-48.98 305.105 320.068 430.36z"/>`),
		g.Group(children),
	)
}