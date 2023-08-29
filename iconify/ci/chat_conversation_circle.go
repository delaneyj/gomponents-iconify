package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatConversationCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.338 15.99a6 6 0 1 0-5.408-2.78l-.424 1.272v.002c-.163.487-.244.73-.187.893c.05.141.163.253.304.303c.162.058.404-.023.888-.184l.007-.002l1.272-.424a5.971 5.971 0 0 0 3.548.92Zm0 0s0 0 0 0Zm0 0a6.003 6.003 0 0 0 8.872 3.08l1.272.424h.003c.487.163.73.244.893.186a.5.5 0 0 0 .302-.303c.058-.162-.023-.406-.186-.895l-.424-1.272l.142-.235A6 6 0 0 0 15 8l-.225.004l-.113.006"/>`),
		g.Group(children),
	)
}