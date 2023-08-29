package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundhound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.44 4.51a17.22 17.22 0 0 0-5 .56A16.69 16.69 0 0 0 14.77 33L9.1 38.59a16.91 16.91 0 0 0 28.22-7.5A16.68 16.68 0 0 0 33.06 15l5.67-5.6a17 17 0 0 0-11.29-4.89Z"/>`),
		g.Group(children),
	)
}