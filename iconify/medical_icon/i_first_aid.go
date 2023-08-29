package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IFirstAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M62.321 20.418H42.908V1.005h-22.34v19.413H1.155v22.339h19.413V62.17h22.34V42.757h19.413z"/>`),
		g.Group(children),
	)
}