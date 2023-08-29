package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.49 13.94l-.34.2a1 1 0 0 0-.35 1.37a1 1 0 0 0 .86.49a1 1 0 0 0 .51-.14l.34-.2a1 1 0 0 0-1-1.72Zm-8.84-7.58l.35-.21l7 4l-1.76 1a1 1 0 0 0 .5 1.87a1 1 0 0 0 .5-.13L21.5 11a1 1 0 0 0 0-1.74l-9-5.19a1 1 0 0 0-1 0l-.85.49a1 1 0 0 0 1 1.74ZM3.71 2.29a1 1 0 0 0-1.42 1.42l3.64 3.63l-3.43 2a1 1 0 0 0 0 1.74l9 5.2a1.09 1.09 0 0 0 .5.13a1.13 1.13 0 0 0 .5-.13l1.5-.88l1.45 1.46l-3.44 2l-8.51-4.93a1 1 0 0 0-1 1.74l9 5.2a1 1 0 0 0 1 0l4.41-2.55l3.38 3.39a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm8.29 12l-7-4.1l2.4-1.38l5.12 5.13Z"/>`),
		g.Group(children),
	)
}