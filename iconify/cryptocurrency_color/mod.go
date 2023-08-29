package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#09547d"/><g fill="#fff"><path d="M22.985 21.007V8.08l-6.312 6.449z" opacity=".5"/><path d="m9 7l.304.312l8.467 8.675L9 24.985z"/></g></g>`),
		g.Group(children),
	)
}