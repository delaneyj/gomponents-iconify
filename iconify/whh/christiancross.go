package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Christiancross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 512H512v448q0 26-18.5 45t-45.5 19H320q-26 0-45-19t-19-45V512H64q-26 0-45-19T0 448V320q0-27 18.5-45.5T64 256h192V64q0-27 19-45.5T320 0h128q27 0 45.5 18.5T512 64v192h192q27 0 45.5 18.5T768 320v128q0 26-18.5 45T704 512z"/>`),
		g.Group(children),
	)
}