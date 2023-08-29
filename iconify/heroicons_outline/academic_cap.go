package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcademicCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path d="m12 14l9-5l-9-5l-9 5l9 5Z"/><path d="m12 14l6.16-3.422a12.083 12.083 0 0 1 .665 6.479A11.952 11.952 0 0 0 12 20.055a11.952 11.952 0 0 0-6.824-2.998a12.078 12.078 0 0 1 .665-6.479L12 14Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 14l9-5l-9-5l-9 5l9 5Zm0 0l6.16-3.422a12.083 12.083 0 0 1 .665 6.479A11.952 11.952 0 0 0 12 20.055a11.952 11.952 0 0 0-6.824-2.998a12.078 12.078 0 0 1 .665-6.479L12 14Zm-4 6v-7.5l4-2.222"/></g>`),
		g.Group(children),
	)
}