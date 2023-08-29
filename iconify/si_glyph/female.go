package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.982 9.965c2.243-.47 3.934-2.48 3.934-4.883c0-2.751-2.215-4.988-4.938-4.988c-2.721 0-4.936 2.237-4.936 4.988c0 2.412 1.704 4.43 3.959 4.888v2.073H5.062v1.925h2.939v2.001h1.98v-2.001h2.893v-1.925H9.981V9.965h.001zm-4.048-5.01c0-1.768 1.367-3.205 3.045-3.205c1.68 0 3.045 1.438 3.045 3.205c0 1.767-1.365 3.203-3.045 3.203c-1.678 0-3.045-1.436-3.045-3.203z"/>`),
		g.Group(children),
	)
}