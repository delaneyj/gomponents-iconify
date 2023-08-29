package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePdfLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 152a6 6 0 0 1-6 6h-26v20h18a6 6 0 0 1 0 12h-18v18a6 6 0 0 1-12 0v-56a6 6 0 0 1 6-6h32a6 6 0 0 1 6 6ZM90 172a26 26 0 0 1-26 26H54v10a6 6 0 0 1-12 0v-56a6 6 0 0 1 6-6h16a26 26 0 0 1 26 26Zm-12 0a14 14 0 0 0-14-14H54v28h10a14 14 0 0 0 14-14Zm84 8a34 34 0 0 1-34 34h-16a6 6 0 0 1-6-6v-56a6 6 0 0 1 6-6h16a34 34 0 0 1 34 34Zm-12 0a22 22 0 0 0-22-22h-10v44h10a22 22 0 0 0 22-22ZM42 112V40a14 14 0 0 1 14-14h96a6 6 0 0 1 4.25 1.76l56 56A6 6 0 0 1 214 88v24a6 6 0 0 1-12 0V94h-50a6 6 0 0 1-6-6V38H56a2 2 0 0 0-2 2v72a6 6 0 0 1-12 0Zm116-30h35.52L158 46.48Z"/>`),
		g.Group(children),
	)
}