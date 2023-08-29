package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropmenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-640q-53 0-90.5-37.5t-37.5-90.5V384q-53 0-90.5-37.5T.428 256V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-12-928h-296q-8 0-11 14.5t5 23.5l135 177q8 9 19 9t19-9l135-177q8-9 5-23.5t-11-14.5zm12 288h-640v128h640V384zm0 192h-640v128h640V576zm0 192h-640v96q0 13 9.5 22.5t22.5 9.5h576q13 0 22.5-9.5t9.5-22.5v-96z"/>`),
		g.Group(children),
	)
}