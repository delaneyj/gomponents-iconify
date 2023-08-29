package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobilePhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10 2v-.5a.5.5 0 0 0-1 0V2H5a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1ZM6 13H5v-1h1Zm0-2H5v-1h1Zm0-2H5V8h1Zm2 4H7v-1h1Zm0-2H7v-1h1Zm0-2H7V8h1Zm2 4H9v-1h1Zm0-2H9v-1h1Zm0-2H9V8h1Zm0-2.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5Z"/>`),
		g.Group(children),
	)
}