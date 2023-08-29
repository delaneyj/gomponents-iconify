package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 625q0 39-17 55l-326 326q-18 18-41 18t-41-18L17 680Q0 663 0 630V64q0-26 19-45T64.5 0T110 19t19 45h-1v553l256 256l256-256V64q0-26 19-45t45.5-19t45 19T768 64v561z"/>`),
		g.Group(children),
	)
}