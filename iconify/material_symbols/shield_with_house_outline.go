package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldWithHouseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.475-.875-5.738-3.988T4 11.1V5l8-3l8 3v6.1q0 3.8-2.263 6.913T12 22Zm0-12.475l-5.6 4.35q.475 1.575 1.4 2.863T10 18.9V15h4v3.9q1.275-.875 2.2-2.163t1.4-2.862L12 9.525Zm0-5.4l-6 2.25V11.1q0 .125.013.275t.012.275L12 7l5.975 4.65q0-.125.013-.275T18 11.1V6.375l-6-2.25Z"/>`),
		g.Group(children),
	)
}