package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func P(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 96c0-35.3 28.7-64 64-64h96c88.4 0 160 71.6 160 160s-71.6 160-160 160H64v96c0 17.7-14.3 32-32 32S0 465.7 0 448V96zm64 192h96c53 0 96-43 96-96s-43-96-96-96H64v192z"/>`),
		g.Group(children),
	)
}