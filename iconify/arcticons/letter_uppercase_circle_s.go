package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterUppercaseCircleS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M21.943 11.85c-3.348 0-6.062 2.72-6.062 6.075h0c0 3.355 2.714 6.075 6.062 6.075h4.114c3.348 0 6.062 2.72 6.062 6.075h0c0 3.355-2.714 6.075-6.062 6.075"/><path d="M31.542 13.9c-1.675-1.406-3.482-2.05-7.542-2.05h-2.057M16.458 34.1c1.675 1.406 3.482 2.05 7.542 2.05h2.057"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}