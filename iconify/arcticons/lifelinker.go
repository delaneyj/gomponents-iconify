package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifelinker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="19.984" height="29.383" x="9.539" y="14.385" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="19.984" height="29.383" x="9.539" y="14.385" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.72 15.557l-.485-.92a2 2 0 0 1 .833-2.703L24.2 4.464a2 2 0 0 1 2.703.833l11.863 22.44a2 2 0 0 1-.833 2.703l-8.408 4.445"/>`),
		g.Group(children),
	)
}