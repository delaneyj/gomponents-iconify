package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Origami(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 384L704 704L256 832l-128-64L0 832l128-192l128 64l77-51l79-185L256 0l256 128l95 237l289-173l128 192H832z"/>`),
		g.Group(children),
	)
}