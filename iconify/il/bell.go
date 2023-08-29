package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M579 372q0 24 9 45t25 37t37 25t45 9v93H0v-93q24 0 45-9t37-25t25-37t9-45V233q0-48 18-90t50-74t73-50t90-18t90 18t74 50t50 74t18 90v139zM347 696q-32 0-56-20t-33-49h179q-8 30-32 49t-58 20z"/>`),
		g.Group(children),
	)
}