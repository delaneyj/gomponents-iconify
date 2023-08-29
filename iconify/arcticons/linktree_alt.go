package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinktreeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.126 17.707h9.598L10.9 11.202l3.776-3.882l6.505 6.687V4.5h5.64v9.508l6.506-6.672l3.773 3.866l-6.82 6.49h9.595v5.368h-9.65l6.866 6.672l-3.764 3.79l-9.325-9.37l-9.326 9.37l-3.775-3.775l6.868-6.672H8.126v-5.367Zm13.04 13.056h5.64V43.5h-5.64V30.763Z"/>`),
		g.Group(children),
	)
}