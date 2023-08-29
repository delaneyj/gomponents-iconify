package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1005.488 915l-91 90q-19 19-46 19t-46-19l-566-563v325q0 27-19 45.5t-45 18.5h-128q-26 0-45-18.5t-19-45.5V64q0-27 18.5-45.5T64.488 0h704q26 0 45 18.5t19 45.5v128q0 26-19 45t-45 19h-334l571 568q19 19 19 45.5t-19 45.5z"/>`),
		g.Group(children),
	)
}