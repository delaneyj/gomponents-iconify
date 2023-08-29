package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelievedFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm4.073-1.875a1 1 0 0 1 1.302-.552c.346.14.725.218 1.125.218s.779-.078 1.125-.218a1 1 0 1 1 .75 1.854a4.987 4.987 0 0 1-1.875.364a4.987 4.987 0 0 1-1.875-.364a1 1 0 0 1-.552-1.302Zm8.302-.552a1 1 0 1 0-.75 1.854a5.005 5.005 0 0 0 3.75 0a1 1 0 1 0-.75-1.854c-.346.14-.725.218-1.125.218s-.779-.078-1.125-.218Zm.126 6.793a1 1 0 0 0-1.002-1.732c-.44.255-.95.401-1.499.401a2.993 2.993 0 0 1-1.5-.4a1 1 0 0 0-1 1.73c.736.427 1.591.67 2.5.67c.91 0 1.764-.243 2.5-.67Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}