package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 320q0 56-23 106t-64 86q41 36 64 86t23 106v256q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5V704q0-53-37.5-90.5T512 576H128v384q0 27-18.5 45.5t-45 18.5t-45.5-18.5T0 960V64q0-26 19-45T64.5 0t45 18.5T128 64v384h384q53 0 90.5-37.5T640 320V64q0-26 19-45t45.5-19t45 19T768 64v256z"/>`),
		g.Group(children),
	)
}