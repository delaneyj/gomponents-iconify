package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libreofficeviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.91 20.12l-15-15a2 2 0 0 0-1.41-.59h-12a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-20a2 2 0 0 0-.59-1.41Zm-9.02-14.6l8.59 8.59a.6.6 0 0 0 1-.42V6.5a2 2 0 0 0-2-2h-7.17a.6.6 0 0 0-.42 1.02Z"/>`),
		g.Group(children),
	)
}