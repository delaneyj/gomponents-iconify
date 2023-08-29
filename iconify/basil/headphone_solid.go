package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.75C8.436 3.75 5.75 6.205 5.75 9v1.512A2.55 2.55 0 0 1 6 10.5h2A1.5 1.5 0 0 1 9.5 12v5A1.5 1.5 0 0 1 8 18.5H6A2.5 2.5 0 0 1 3.5 16v-3c0-.7.287-1.332.75-1.785V9c0-3.832 3.582-6.75 7.75-6.75S19.75 5.168 19.75 9v2.215A2.49 2.49 0 0 1 20.5 13v3a2.5 2.5 0 0 1-2.5 2.5h-2a1.5 1.5 0 0 1-1.5-1.5v-5a1.5 1.5 0 0 1 1.5-1.5h2c.084 0 .168.004.25.012V9c0-2.795-2.686-5.25-6.25-5.25Z"/>`),
		g.Group(children),
	)
}