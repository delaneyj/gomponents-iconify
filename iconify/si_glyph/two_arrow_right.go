package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.83 7.999L9.086 13h3.939l3.944-5.001l-4.042-4.969H9.009l-.011.011l3.832 4.958zm-7.783 0L1.096 13h3.938L9 7.999L4.935 3.03H1.018l-.01.011l4.039 4.958z"/>`),
		g.Group(children),
	)
}