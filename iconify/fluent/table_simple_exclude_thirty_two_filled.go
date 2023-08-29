package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleExcludeThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 2H12v10H2V6.5A4.5 4.5 0 0 1 6.5 2ZM2 14v5a4.5 4.5 0 0 0 4.5 4.5H12V14H2Zm21.5-2V6.5A4.5 4.5 0 0 0 19 2h-5v10h9.5ZM16 19a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3v-8Z"/>`),
		g.Group(children),
	)
}