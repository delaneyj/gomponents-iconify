package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.83 8v.52A11.41 11.41 0 0 1 8.35 20a11.41 11.41 0 0 1-6.2-1.81h1a8.09 8.09 0 0 0 5-1.73a4 4 0 0 1-3.78-2.8a4.81 4.81 0 0 0 .77.06a4.66 4.66 0 0 0 1.06-.13A4 4 0 0 1 3 9.67a4.13 4.13 0 0 0 1.82.51A4.06 4.06 0 0 1 3 6.77a4 4 0 0 1 .54-2A11.47 11.47 0 0 0 11.85 9a4.71 4.71 0 0 1-.1-.92a4 4 0 0 1 7-2.77a7.93 7.93 0 0 0 2.56-1a4 4 0 0 1-1.78 2.22a7.94 7.94 0 0 0 2.33-.62a8.91 8.91 0 0 1-2 2.09Z"/>`),
		g.Group(children),
	)
}