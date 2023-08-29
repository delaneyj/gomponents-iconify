package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymptomsVirusHeadacheTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.326 23.25v-3.99h2.763a1.33 1.33 0 0 0 1.33-1.33v-2.867h1.536a1.022 1.022 0 0 0 .933-1.44c-1.005-2.237-2.473-4.987-2.473-4.987V6.457a2.04 2.04 0 0 0-.182-.849a7.3 7.3 0 0 0-4.976-4.527C8.279-.487 2.067 3.7 2.022 9.876a9.185 9.185 0 0 0 3.07 6.934v6.44"/><path d="M10.984 12.475a2.776 2.776 0 1 0 0-5.552a2.776 2.776 0 0 0 0 5.552Zm-.462-7.634h.925m-.463 0v2.082m3.108-.986l.654.654m-.327-.327l-1.472 1.472m2.895 1.5v.926m0-.463H13.76m.986 3.108l-.654.654m.327-.327l-1.472-1.472m-1.5 2.895h-.925m.462 0v-2.082m-3.108.986l-.654-.654m.327.327l1.472-1.472m-2.895-1.5v-.926m0 .463h2.082m-.986-3.108l.654-.654m-.327.327l1.472 1.472"/></g>`),
		g.Group(children),
	)
}