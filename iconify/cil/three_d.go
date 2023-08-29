package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m68.983 382.642l171.35 98.928a32.082 32.082 0 0 0 32 0l171.352-98.929a32.093 32.093 0 0 0 16-27.713V157.071a32.092 32.092 0 0 0-16-27.713L272.334 30.429a32.086 32.086 0 0 0-32 0L68.983 129.358a32.09 32.09 0 0 0-16 27.713v197.858a32.09 32.09 0 0 0 16 27.713ZM272.333 67.38l155.351 89.691v177.378l-155.351-87.807Zm-16.051 206.947l157.155 88.828l-157.1 90.7l-157.158-90.73ZM84.983 157.071l155.35-89.691v179.2l-155.35 87.81Z"/>`),
		g.Group(children),
	)
}