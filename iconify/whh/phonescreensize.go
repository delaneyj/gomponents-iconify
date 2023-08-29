package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonescreensize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024H64q-27 0-45.5-18.5T0 960V64q0-26 18.5-45T64 0h640q27 0 45.5 19T768 64v896q0 27-18.5 45.5T704 1024zm0-896H64v768h640V128zM256 256l320 448l64-64v192H448l64-64l-320-448l-64 64V192h192z"/>`),
		g.Group(children),
	)
}