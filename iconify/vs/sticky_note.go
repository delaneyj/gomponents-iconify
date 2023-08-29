package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1152 1216h412q-13 60-94 148.5t-173.5 156T1152 1600v-384zm-128-64v448H0q79-242 103.5-447.5T128 576V64h1536v512q0 65-3.5 125t-6.5 97t-12 90t-12.5 69t-16 70.5t-13.5 60.5h-512q-40 0-52 12t-12 52z"/>`),
		g.Group(children),
	)
}