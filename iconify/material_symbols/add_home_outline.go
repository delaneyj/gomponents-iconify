package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHomeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 19h1v-2.5H21v-1h-2.5V13h-1v2.5H15v1h2.5V19Zm.5 2q-2.075 0-3.538-1.463T13 16q0-2.075 1.463-3.538T18 11q2.075 0 3.538 1.463T23 16q0 2.075-1.463 3.538T18 21ZM4 19V7l8-6l8 6v2.3q-.475-.15-.975-.225T18 9V8l-6-4.5L6 8v9h5.075q.075.525.225 1.025t.375.975H4Zm8-8.75Z"/>`),
		g.Group(children),
	)
}