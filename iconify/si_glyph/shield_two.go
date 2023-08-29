package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.99.062c-3.307 0-5.988 1.944-5.988 1.944l.006 6.054c0 4.999 5.982 7.888 5.982 7.888s5.983-2.722 5.983-7.903c0-5.18.002-6.038.002-6.038S12.295.062 8.99.062zm.063 13.994s-5.1-2.26-5.1-6.17l-.004-5.419S6.234.946 9.053.946v13.11z"/>`),
		g.Group(children),
	)
}