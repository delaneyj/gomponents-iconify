package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Html(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 15V9h1.5v2h2V9H5v6H3.5v-2.5h-2V15H0Zm7.75 0v-4.5H6V9h5v1.5H9.25V15h-1.5ZM12 15v-5q0-.425.288-.713T13 9h4.5q.425 0 .713.288T18.5 10v5H17v-4.5h-1V14h-1.5v-3.5h-1V15H12Zm8 0V9h1.5v4.5H24V15h-4Z"/>`),
		g.Group(children),
	)
}