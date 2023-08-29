package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unbalanced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M404.5 162.5c-26.1 0-47 20.9-47 47s20.9 47 47 47s47-20.9 47-47s-20.9-47-47-47zm72.2 89.1l-447.99 176l6.58 16.8l448.01-176l-6.6-16.8zM121.2 287l-87.28 35l29 72.4l87.28-35l-29-72.4zM256 383l-78 104h156l-78-104z"/>`),
		g.Group(children),
	)
}