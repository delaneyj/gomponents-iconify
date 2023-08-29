package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BdFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#006a4e" d="M0 0h640v480H0z"/><circle cx="280" cy="240" r="160" fill="#f42a41"/>`),
		g.Group(children),
	)
}