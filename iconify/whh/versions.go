package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Versions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H448q-27 0-45.5-18.5T384 960V64q0-27 18.5-45.5T448 0h512q27 0 45.5 18.5T1024 64v896q0 27-18.5 45.5T960 1024zm0-960H448v896h512V64zM192 832V192q0-27 18.5-45.5T256 128h64v64h-64v640h64v64h-64q-27 0-45.5-18.5T192 832zM0 704V320q0-26 18.5-45T64 256h64v64H64v384h64v64H64q-27 0-45.5-18.5T0 704z"/>`),
		g.Group(children),
	)
}