package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.85 22.525L1.475 4.15L2.9 2.725L21.275 21.1l-1.425 1.425ZM8 16v-2h6.175l2 2H8Zm-3-3v-2h2v2H5Zm3 0v-2h2v2H8Zm9 0v-2h2v2h-2ZM5 10V8h2v2H5Zm9 0V8h2v2h-2Zm3 0V8h2v2h-2Zm4.4 8.425L20 17V7h-9.975l-2-2H20q.825 0 1.413.588T22 7v10.025q0 .425-.163.775t-.437.625ZM4 19q-.825 0-1.413-.588T2 17V7q0-.825.588-1.413T4 5h1.175l2 2H4v10h13.175l2 2H4Zm7.025-11H13l-.025 1.975L11.025 8ZM14 11h2l-.025 1.975L14 11Zm-4.675 1ZM15 12Z"/>`),
		g.Group(children),
	)
}