package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 640H384q-53 0-90.5-37.5T256 512V128q0-53 37.5-90.5T384 0h512q53 0 90.5 37.5T1024 128v384q0 53-37.5 90.5T896 640zm0-448q0-27-18.5-45.5T832 128H448q-26 0-45 18.5T384 192v256q0 26 19 45t45 19h384q27 0 45.5-19t18.5-45V192zM704 960q0 27-18.5 45.5T640 1024H256q-26 0-45-18.5T192 960t19-45.5t45-18.5h384q27 0 45.5 18.5T704 960zm-639.5 64q-26.5 0-45.5-18.5T0 960t19-45.5T64.5 896t45 18.5T128 960t-18.5 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}