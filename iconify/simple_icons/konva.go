package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Konva(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0zm1.391 18.541l-.239-3.76l-2.391-1.608l.152 5.129l-4.325.152l-.173-13.409L10.5 4.98l.087 5.346l2.217-1.608l.109-3.781l4.412.283l-.348 4.586l-2.608 1.608l2.673 1.174l.913 5.694l-4.564.259z"/>`),
		g.Group(children),
	)
}