package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveShortcut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 22q-3.425-.7-5.712-3.425Q8 15.85 8 12.2q0-2.575 1.175-4.713Q10.35 5.35 12.3 4H8V2h8v8h-2V5.3q-1.8 1.05-2.9 2.862Q10 9.975 10 12.2q0 2.8 1.7 4.937q1.7 2.138 4.3 2.813Z"/>`),
		g.Group(children),
	)
}