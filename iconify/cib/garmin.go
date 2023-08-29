package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garmin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.355 30.229H2.646c-1.025 0-1.848-.516-2.26-1.339a2.5 2.5 0 0 1 0-2.672L13.741 3.103c.511-.817 1.333-1.333 2.26-1.333c1.027 0 1.849.516 2.251 1.333l13.364 23.115a2.5 2.5 0 0 1 0 2.672c-.417.927-1.24 1.339-2.26 1.339z"/>`),
		g.Group(children),
	)
}