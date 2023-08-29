package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ceur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M128 384V128h64V64H64v384h192v-64ZM384 64H256v64h128v256h-64v64h128V64Z"/>`),
		g.Group(children),
	)
}