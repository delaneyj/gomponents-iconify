package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentPdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.964 0H2.016v15.958h11.942V5.033H8.964V0zM6.047 10.036H5.029v.985h-.998v1.025H2.969V7.954l2.061.014v.984h1.018v1.084h-.001zm4.906-2.083h2.078v1.078H12.02v.922h1.011v1.061H12.02v1.018h-1.067V7.953zm-1.92.018v.998h.973L10.02 11H9.016v1.016H6.969V7.961l2.064.01z"/><path d="M10.025.021v3.967h3.954L10.025.021zM4 9h.973v.961H4zm4 0h.969v1.983H8z"/></g>`),
		g.Group(children),
	)
}