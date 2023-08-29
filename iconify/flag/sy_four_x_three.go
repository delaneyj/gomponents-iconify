package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M0 0h640v480H0Z"/><path fill="#fff" d="M0 0h640v320H0Z"/><path fill="#ce1126" d="M0 0h640v160H0Z"/><path fill="#007a3d" d="m161 300l39-120l39 120l-102-74.2h126M401 300l39-120l39 120l-102-74.2h126"/>`),
		g.Group(children),
	)
}