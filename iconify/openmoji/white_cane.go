package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteCane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="m54.26 17.472l6.503-6.503a1.234 1.234 0 0 0 0-1.745l-.719-.72a1.226 1.226 0 0 0-1.734 0l-6.509 6.51Z"/><path fill="#d0cfce" d="m19.956 46.859l31.845-31.845l2.459 2.459l-31.845 31.844z"/><path fill="#ea5a47" d="m9.842 56.973l10.114-10.115l2.459 2.459l-9.909 9.909l-2.664-2.253z"/><path fill="#d0cfce" d="m9.636 57.056l.206-.083l2.664 2.253l-.043.166c-.176 1.121-2.846 1.948-3.324 1.47l-.723-.724c-.48-.48-.01-2.467 1.22-3.082Z"/><g fill="none" stroke="#000"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M54.26 17.472L12.506 59.226l-.043.166c-.176 1.121-2.846 1.948-3.324 1.47m45.121-43.39l6.503-6.503a1.234 1.234 0 0 0 0-1.745l-.719-.72"/><path stroke-miterlimit="10" d="M61.635 8.867s1.23-2.459 4.918-2.459s4.917 7.376 3.688 8.605s-3.688-2.458-4.918-3.687a17.684 17.684 0 0 0-3.688-2.459Z"/></g>`),
		g.Group(children),
	)
}