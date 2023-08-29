package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="160" r="128" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M143.65 227.82L48 400l86.86-.42a16 16 0 0 1 13.82 7.8L192 480l88.33-194.32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M366.54 224L464 400l-86.86-.42a16 16 0 0 0-13.82 7.8L320 480l-64-140.8"/><circle cx="256" cy="160" r="64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/>`),
		g.Group(children),
	)
}