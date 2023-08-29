package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymptomsVirusLossSmellTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.409 17.5a6.498 6.498 0 0 0 1.007-1.534M11.3 20.058a8.379 8.379 0 0 0 2.046-.842m7.767 1.864a3.98 3.98 0 0 0 .417 2.046m1.022-6.137c-.341.682-.682 1.364-.947 2.046M16.416 22.1c-.549.56-1.269.92-2.046 1.023m4.091-4.088a8.37 8.37 0 0 1-.758 1.534M22.468.874S19.648 6.7 16.993 8.7c-3.128 2.346-1.564 5.475.782 5.475h5.475"/><path d="M19.578 11.1c.767-1.022 1.534-1.022 3.068-1.022M6.375 12.124a5.625 5.625 0 1 0 0-11.25a5.625 5.625 0 0 0 0 11.25Zm-3.977-1.648l7.954-7.954"/></g>`),
		g.Group(children),
	)
}