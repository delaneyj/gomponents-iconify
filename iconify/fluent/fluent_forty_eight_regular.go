package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluentFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.83 4.158a1.5 1.5 0 0 1 1.34 0l12 6.001a1.5 1.5 0 0 1 0 2.684l-9.316 4.659l9.317 4.659a1.5 1.5 0 0 1 0 2.683L26 30.43V42.5a1.5 1.5 0 0 1-2.298 1.27l-12-7.54A1.5 1.5 0 0 1 11 34.96V11.5a1.5 1.5 0 0 1 .83-1.34l12-6.002ZM14 12.428v21.703l9 5.655V29.503a1.5 1.5 0 0 1 .83-1.342l9.316-4.659l-9.317-4.659a1.5 1.5 0 0 1 0-2.683l9.317-4.66L24.5 7.178L14 12.428Z"/>`),
		g.Group(children),
	)
}