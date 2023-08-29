package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationAddOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-8-3v-2h2v-7q0-2.075 1.25-3.688T10.5 4.2V2h3v2.2q.45.125.875.288t.8.412q-.325.4-.563.837t-.437.913q-.475-.3-1.025-.475T12 6q-1.65 0-2.825 1.175T8 10v7h8v-3.075q.425.375.925.688t1.075.512V17h2v2H4Zm8-7.5Zm7 1.5v-3h-3V8h3V5h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}