package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0z"/><path fill="#002f6c" d="M0 186.2h512v139.6H0z"/><path fill="#002f6c" d="M123.2 0h139.6v512H123.1z"/>`),
		g.Group(children),
	)
}