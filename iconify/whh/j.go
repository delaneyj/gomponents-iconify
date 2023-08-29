package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func J(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 768q0 106-75 181t-181 75H256q-106 0-181-75T0 768q0-26 19-45t45.5-19t45 18.5T128 768q0 53 37.5 90.5T256 896h256q53 0 90.5-37.5T640 768V64q0-26 19-45t45.5-19t45 19T768 64v704z"/>`),
		g.Group(children),
	)
}