package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Framer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M0 0h256v128H128L0 0Zm0 128h128l128 128H128v128L0 256V128Z"/>`),
		g.Group(children),
	)
}