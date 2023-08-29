package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHtmlLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M126 152a6 6 0 0 1-6 6h-10v50a6 6 0 0 1-12 0v-50H88a6 6 0 0 1 0-12h32a6 6 0 0 1 6 6Zm-62-6a6 6 0 0 0-6 6v22H38v-22a6 6 0 0 0-12 0v56a6 6 0 0 0 12 0v-22h20v22a6 6 0 0 0 12 0v-56a6 6 0 0 0-6-6Zm176 56h-14v-50a6 6 0 0 0-12 0v56a6 6 0 0 0 6 6h20a6 6 0 0 0 0-12Zm-46.4-55.78a6 6 0 0 0-6.75 2.69L168 180.34l-18.85-31.43A6 6 0 0 0 138 152v56a6 6 0 0 0 12 0v-34.34l12.85 21.43a6 6 0 0 0 10.3 0L186 173.66V208a6 6 0 0 0 12 0v-56a6 6 0 0 0-4.4-5.78ZM214 112a6 6 0 0 1-12 0V94h-50a6 6 0 0 1-6-6V38H56a2 2 0 0 0-2 2v72a6 6 0 0 1-12 0V40a14 14 0 0 1 14-14h96a6 6 0 0 1 4.25 1.76l56 56A6 6 0 0 1 214 88Zm-20.48-30L158 46.48V82Z"/>`),
		g.Group(children),
	)
}