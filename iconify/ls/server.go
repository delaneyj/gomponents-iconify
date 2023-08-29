package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M36 160h508c19 0 36 17 36 36v84c0 20-17 36-36 36H36c-20 0-36-16-36-36v-84c0-19 16-36 36-36zm500 86v-16c0-6-6-12-12-12h-64c-7 0-12 6-12 12v16c0 7 5 12 12 12h64c6 0 12-5 12-12zM36 344h508c19 0 36 16 36 36v84c0 20-17 36-36 36H36c-20 0-36-16-36-36v-84c0-20 16-36 36-36zm500 86v-16c0-6-6-12-12-12h-64c-7 0-12 6-12 12v16c0 7 5 12 12 12h64c6 0 12-5 12-12zM36 528h508c19 0 36 16 36 36v84c0 20-17 36-36 36H36c-20 0-36-16-36-36v-84c0-20 16-36 36-36zm500 86v-16c0-7-6-12-12-12h-64c-7 0-12 5-12 12v16c0 7 5 12 12 12h64c6 0 12-5 12-12zm8-481H36c-6 0-12 1-17 3l61-83c12-16 38-29 58-29h304c19 0 45 13 57 29l62 83c-6-2-12-3-17-3z"/>`),
		g.Group(children),
	)
}