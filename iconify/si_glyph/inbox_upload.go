package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.768 2.757L9.579.244a.811.811 0 0 0-1.143 0L6.247 2.757c-.315.315-.363.876-.049 1.19H8v1.922c0 .559.434 1.01.969 1.01c.535 0 .969-.451.969-1.01V3.947h1.781c.315-.315.364-.875.049-1.19z"/><path d="M14.993 1.006h-3.29l.891.975h1.583L15.54 9.01H2.505l1.407-7.029h1.682l.766-.975H3.073l-2.06 8.755v4.17c0 1.334.472 2.028 1.804 2.028h12.28c1.246 0 1.885-.527 1.885-2.111V9.761l-1.989-8.755zM11.016 11H6.985V9.969h4.031V11z"/></g>`),
		g.Group(children),
	)
}