package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func One(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 1024H64q-26 0-45-18.5t-19-45T19 915t45-19h128V256q0-53-37-90t-89-38h-2q-27 0-45.5-18.5T0 64.5T18.5 19T64 0v1v-1q106 0 181 75t75 181v640h128q27 0 45.5 19t18.5 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}