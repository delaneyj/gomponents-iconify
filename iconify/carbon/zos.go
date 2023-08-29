package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 22h-5v-2h5v-3h-3c-1.103 0-2-.897-2-2v-3c0-1.103.897-2 2-2h5v2h-5v3h3c1.103 0 2 .897 2 2v3c0 1.103-.897 2-2 2zm-9 0h-3c-1.103 0-2-.897-2-2v-8c0-1.103.897-2 2-2h3c1.103 0 2 .897 2 2v8c0 1.103-.897 2-2 2zm-3-10v8h3v-8h-3zm-6.054-2l-4 12h2.108l4-12h-2.108zM6 22H0v-2.303L3.798 14H0v-2h6v2.303L2.202 20H6v2z"/>`),
		g.Group(children),
	)
}