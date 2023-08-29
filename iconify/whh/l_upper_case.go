package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024H64q-26 0-45-18.5T0 960V64q0-26 19-45T64.5 0t45 18.5T128 64v800q0 13 9.5 22.5T160 896h544q27 0 45.5 19t18.5 45.5t-18.5 45T704 1024z"/>`),
		g.Group(children),
	)
}