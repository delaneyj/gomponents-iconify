package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chapel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32.196 20.641s-2.046 1.08-2.046 4.375v3.286h4.093v-3.286c0-3.295-2.047-4.375-2.047-4.375z"/><path fill="currentColor" d="M55.175.7H8.391C3.807.7.077 4.431.077 9.016v46.781c0 4.586 3.729 8.316 8.314 8.316h46.784c4.584 0 8.315-3.73 8.315-8.316V9.016C63.49 4.431 59.759.7 55.175.7zm-4.88 55.891H35.697v-7.762c0-3.98-3.602-5.284-3.602-5.284s-3.602 1.303-3.602 5.284v7.762H14.124V40.495l12.113-5.85V19.116l5.966-12.468l5.973 12.468v15.529l12.119 5.85v16.096z"/>`),
		g.Group(children),
	)
}