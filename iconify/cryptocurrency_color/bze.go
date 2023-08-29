package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#00AEEF"/><path fill="#FFF" d="M13.952 24.461H8.264l1.116-6.315l5.905-5.319h-5.022l4.061-2.489h6.603l-8.54 8.43l-.574 3.22s1.907.14 3.41 0c1.504-.155 3.193-.777 3.193-.777l-4.464 3.25zm8.944-13.516l-2.542 2.489s2.139 2.38.17 5.646c-1.937 3.204-8.37 2.613-8.37 2.613l.496-2.815l9.052-8.835H14.2l.76-4.293h-3.627l-1.302 7.341h4.619l-5.487 4.947L8 24.695s4.387.124 8.107 0c3.735-.125 8.09-3.329 8.726-6.953c.868-4.9-1.937-6.797-1.937-6.797z"/></g>`),
		g.Group(children),
	)
}