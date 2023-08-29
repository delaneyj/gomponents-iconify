package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 416L448 960q-1 1-4 7t-5 9.5l-6 10.5l-7 11l-8.5 9.5l-10 8.5l-11 5.5l-12.5 2.5q-27 0-45.5-15T320 960q0-20 32-64q139-262 261-511q12-24 18.5-56.5T639 276l1-20q0-53-37.5-90.5T512 128H64q-27 0-45.5-18.5T0 64.5T18.5 19T64 0h448q106 0 181 75t75 181q0 40-11 90t-21 70z"/>`),
		g.Group(children),
	)
}