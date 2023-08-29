package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-6 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm3-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm4.28-7.12a1 1 0 0 0-1.4-.16A2.78 2.78 0 0 0 14 14h-3.65a2.81 2.81 0 0 0-.88-1.31a1 1 0 0 0-1.36.2a1 1 0 0 0 .14 1.39a1 1 0 0 1 .25.72a1.09 1.09 0 0 1-.25.72a1 1 0 1 0 1.25 1.56a2.89 2.89 0 0 0 .84-1.28H14a2.72 2.72 0 0 0 .89 1.31a1 1 0 0 0 .57.18a1 1 0 0 0 .78-.38a1 1 0 0 0-.14-1.39a1 1 0 0 1-.25-.72a1.09 1.09 0 0 1 .25-.72a1 1 0 0 0 .18-1.4Z"/>`),
		g.Group(children),
	)
}