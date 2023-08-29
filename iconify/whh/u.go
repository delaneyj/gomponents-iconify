package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func U(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 66v702q0 106-75 181t-181 75H256q-106 0-181-75T0 768V64q0-26 19-45T64.5 0T110 19t19 45l-1 1v703q0 53 37.5 90.5T256 896h256q53 0 90.5-37.5T640 768V64q0-26 19-45t45.5-19T750 19t19 45q-1 1-1 2z"/>`),
		g.Group(children),
	)
}