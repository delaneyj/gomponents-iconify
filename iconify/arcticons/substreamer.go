package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Substreamer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 9v30M9.375 15.207v17.586m9.75-17.586v17.586M4.5 20.897v6.206M33.75 9v30m-4.875-23.793v17.586m9.75-17.586v17.586M24 20.897v6.206m19.5-6.206v6.206"/>`),
		g.Group(children),
	)
}