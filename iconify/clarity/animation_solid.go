package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnimationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityAnimationSolid0" fill="currentColor"><path d="M3.5 23.77a6.41 6.41 0 0 0 9.33 8.67a11.65 11.65 0 0 1-9.33-8.67Z"/><path d="M7.68 14.53a9.6 9.6 0 0 0 13.4 13.7a14.11 14.11 0 0 1-13.4-13.7Z"/><path d="M21.78 2A12.12 12.12 0 1 1 9.66 14.15A12.12 12.12 0 0 1 21.78 2"/></g>`),
		g.Group(children),
	)
}