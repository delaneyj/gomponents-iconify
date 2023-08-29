package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.787 2.042H2.02v1.021H.022v9.913h1.02l.002 1h14.902L15.967 5H7.349L5.787 2.042zm2.192 4.937h1.062V9h2.021v1.062H9.041v2.021H7.979v-2.021H5.958V9h2.021V6.979z"/><path d="M13.964 3.982v-.94h-5.94l.33.94h5.61z"/></g>`),
		g.Group(children),
	)
}