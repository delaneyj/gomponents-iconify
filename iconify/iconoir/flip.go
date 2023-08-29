package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 20H2L9.5 4v16Zm10.625 0H22l-.938-2m-4.687 2H14.5v-2m0-6v2m3.75-2l.938 2m-2.813-6L14.5 4v4"/>`),
		g.Group(children),
	)
}