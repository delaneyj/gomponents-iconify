package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChiliOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 8.28c.6.35 1 .99 1 1.72v2.8l-4.5-4.51L12 8l1.75 1L15 8.28M12 6.5l1.75 1l1.52-.87c-.55-.97-1.36-1.69-2.3-1.98A2.996 2.996 0 0 0 10 2v2c.44 0 .8.29.94.69c-.68.23-1.28.68-1.77 1.28l1.37 1.37L12 6.5M2.39 1.73L1.11 3L8 9.9V11c0 9 8 11 8 11v-4.11l4.84 4.84l1.27-1.27L2.39 1.73Z"/>`),
		g.Group(children),
	)
}