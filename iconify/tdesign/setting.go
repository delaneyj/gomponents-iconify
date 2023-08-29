package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Setting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423L12 .845Zm0 2.31L4.34 7.577v8.846L12 20.845l7.66-4.422V7.577L12 3.155ZM12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/>`),
		g.Group(children),
	)
}