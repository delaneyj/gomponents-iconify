package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ICursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 5v2h2c1.668 0 3 1.332 3 3v4h-3v2h3v6c0 1.668-1.332 3-3 3h-2v2h2c1.633 0 3.086-.813 4-2.031c.914 1.218 2.367 2.031 4 2.031h2v-2h-2c-1.668 0-3-1.332-3-3v-6h3v-2h-3v-4c0-1.668 1.332-3 3-3h2V5h-2c-1.633 0-3.086.813-4 2.031C15.086 5.813 13.633 5 12 5z"/>`),
		g.Group(children),
	)
}