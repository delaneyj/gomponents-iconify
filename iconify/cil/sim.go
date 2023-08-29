package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 40H230.627A31.791 31.791 0 0 0 208 49.373L97.373 160A31.791 31.791 0 0 0 88 182.627V448a32.036 32.036 0 0 0 32 32h264a32.036 32.036 0 0 0 32-32V72a32.036 32.036 0 0 0-32-32Zm0 408H120V182.627L230.627 72H384Z"/><path fill="currentColor" d="M208 416h144V216H208Zm32-168h80v136h-80Z"/>`),
		g.Group(children),
	)
}