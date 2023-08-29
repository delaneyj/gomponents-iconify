package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneDisabledSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.1 14.35l-1.45-1.45q.225-.275.738-1t.637-1l-2.85-2.875L15.1 3H21v1.05q0 2.725-.975 5.3t-2.925 5ZM4.05 21H3v-5.875L8 14.1l2.9 2.9q.725-.45 1.137-.738t.763-.562L1.4 4.3l1.4-1.4l18.4 18.4l-1.4 1.4l-5.55-5.55q-2.3 1.875-4.925 2.863T4.05 21Z"/>`),
		g.Group(children),
	)
}