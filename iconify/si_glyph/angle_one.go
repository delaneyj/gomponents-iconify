package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(2 1)"><ellipse cx="13.479" cy="13.458" rx="1.479" ry="1.458"/><ellipse cx="13.479" cy="9.458" rx="1.479" ry="1.458"/><ellipse cx="9.479" cy="13.458" rx="1.479" ry="1.458"/><ellipse cx="5.479" cy="13.458" rx="1.479" ry="1.458"/><ellipse cx="9.479" cy="9.458" rx="1.479" ry="1.458"/><ellipse cx="13.479" cy="5.458" rx="1.479" ry="1.458"/><ellipse cx="13.479" cy="1.458" rx="1.479" ry="1.458"/><ellipse cx="9.479" cy="5.458" rx="1.479" ry="1.458"/><ellipse cx="5.479" cy="9.458" rx="1.479" ry="1.458"/><ellipse cx="1.479" cy="13.458" rx="1.479" ry="1.458"/></g>`),
		g.Group(children),
	)
}