package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Application(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423L12 .845Zm0 2.31L4.34 7.577v8.846L12 20.845l7.66-4.422V7.577L12 3.155ZM8.723 8.613L12 10.798l3.277-2.185l1.11 1.664L13 12.535V16h-2v-3.465l-3.387-2.258l1.11-1.664Z"/>`),
		g.Group(children),
	)
}