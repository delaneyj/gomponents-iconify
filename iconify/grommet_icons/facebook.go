package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23 0H1a1 1 0 0 0-1 1v22a1 1 0 0 0 1 1h11.75v-9h-3v-3.75h3v-3c0-3.1 1.963-4.625 4.728-4.625c1.324 0 2.463.099 2.794.142v3.24l-1.917.001c-1.504 0-1.855.715-1.855 1.763v2.479h3.75L19.5 15h-3l.06 9H23a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1"/>`),
		g.Group(children),
	)
}