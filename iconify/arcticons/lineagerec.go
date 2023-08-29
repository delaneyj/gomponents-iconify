package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lineagerec(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 9.77a2.94 2.94 0 0 0-2.95 2.94h0v22.57a2.94 2.94 0 0 0 2.94 2.95h33.11a2.94 2.94 0 0 0 3-2.94h0V12.72a2.94 2.94 0 0 0-2.94-3H7.45Zm7.46 8.86A5.6 5.6 0 0 0 9.51 24a5.61 5.61 0 0 0 5.4 5.39A5.6 5.6 0 0 0 20.3 24a5.6 5.6 0 0 0-5.39-5.37Zm18.22 0A5.59 5.59 0 0 0 27.74 24a5.6 5.6 0 0 0 5.39 5.39A5.6 5.6 0 0 0 38.52 24a5.6 5.6 0 0 0-5.39-5.37ZM14.91 29.41h18.22"/>`),
		g.Group(children),
	)
}