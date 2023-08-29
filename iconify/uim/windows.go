package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.03 4.832l8.147-1.11l.004 7.86l-8.144.046l-.008-6.796Zm8.144 7.655l.006 7.867l-8.144-1.12l-.001-6.8l8.138.053Zm.987-8.91L21.965 2v9.482l-10.804.085v-7.99Zm10.807 8.984L21.965 22l-10.804-1.525l-.015-7.932Z"/>`),
		g.Group(children),
	)
}