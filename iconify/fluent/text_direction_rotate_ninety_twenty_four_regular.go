package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRotateNinetyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6.75 3a.75.75 0 0 0-.75.75v14.69l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 1 0-1.06-1.06l-.72.72V3.75A.75.75 0 0 0 6.75 3z" fill="currentColor"/><path d="M11.03 3.053l9.496 3.753c.595.236.63 1.043.104 1.345l-.105.05l-9.5 3.747a.75.75 0 0 1-.643-1.352l.092-.043L13 9.556V5.443l-2.52-.995a.75.75 0 0 1-.454-.876l.031-.097a.75.75 0 0 1 .876-.453l.098.03zm7.178 4.45L14.5 6.036v2.928l3.708-1.461z" fill="currentColor"/><path d="M15.5 12.75a.75.75 0 0 1 1.5 0v5.69l.72-.72a.75.75 0 1 1 1.06 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l.72.72v-5.69z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}