package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 640H64q-27 0-45.5-19T0 576V448q0-27 19-45.5T64 384h896q27 0 45.5 18.5T1024 448v128q0 26-18.5 45T960 640zM576 256H448q-27 0-45.5-19T384 192V64q0-27 19-45.5T448 0h128q27 0 45.5 18.5T640 64v128q0 26-18.5 45T576 256zM448 768h128q27 0 45.5 18.5T640 832v128q0 27-19 45.5t-45 18.5H448q-27 0-45.5-19T384 960V832q0-27 19-45.5t45-18.5z"/>`),
		g.Group(children),
	)
}