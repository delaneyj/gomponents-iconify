package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymptomsVirusLossSmellOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.886 10.392a3.506 3.506 0 1 0 0-7.012a3.506 3.506 0 0 0 0 7.012ZM6.302.75h1.169m-.585 0v2.63m3.926-1.246l.827.827m-.414-.414l-1.859 1.86m3.657 1.895v1.169m0-.585h-2.63m1.246 3.926l-.827.827m.413-.414L9.366 9.366m-1.895 3.657H6.302m.584 0v-2.63m-3.925 1.246l-.827-.827m.413.413l1.86-1.859M.75 7.471V6.302m0 .584h2.63M2.134 2.961l.827-.827m-.414.413l1.86 1.86M22.468.75s-2.82 5.83-5.475 7.821c-3.128 2.346-1.564 5.474.782 5.474h5.475"/><path d="M19.578 10.977c.767-1.022 1.534-1.022 3.068-1.022m-8.166 7.67a6.498 6.498 0 0 0 1.007-1.534m-5.114 4.091a8.386 8.386 0 0 0 2.046-.842m7.765 1.865a3.98 3.98 0 0 0 .416 2.045m1.023-6.136c-.341.681-.682 1.363-.947 2.045m-5.189 3.068c-.549.56-1.269.92-2.046 1.023m4.091-4.09a8.37 8.37 0 0 1-.758 1.533"/></g>`),
		g.Group(children),
	)
}