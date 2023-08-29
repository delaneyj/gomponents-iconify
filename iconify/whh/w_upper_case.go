package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 72v696q0 106-75 181t-181 75H256q-106 0-181-75T0 768V64q0-26 19-45T64.5 0T110 19t19 45q0 1-.5 3t-.5 3v698q0 53 37.5 90.5T256 896h32q13 0 22.5-9.5T320 864V64q0-26 19-45t45.5-19t45 19T448 64v800q0 13 9.5 22.5T480 896h32q53 0 90.5-37.5T640 768V64q0-26 19-45t45.5-19T750 19t19 45q0 2-1 8z"/>`),
		g.Group(children),
	)
}