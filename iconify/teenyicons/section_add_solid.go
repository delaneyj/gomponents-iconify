package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SectionAddSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0H4v1H1.5a.5.5 0 0 0-.5.5V4H0V1.5ZM9 1H6V0h3v1Zm4.5 0H11V0h2.5A1.5 1.5 0 0 1 15 1.5V4h-1V1.5a.5.5 0 0 0-.5-.5ZM7 7V4h1v3h3v1H8v3H7V8H4V7h3ZM0 9V6h1v3H0Zm14 0V6h1v3h-1ZM0 13.5V11h1v2.5a.5.5 0 0 0 .5.5H4v1H1.5A1.5 1.5 0 0 1 0 13.5Zm14 0V11h1v2.5a1.5 1.5 0 0 1-1.5 1.5H11v-1h2.5a.5.5 0 0 0 .5-.5ZM9 15H6v-1h3v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}