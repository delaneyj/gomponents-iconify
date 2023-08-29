package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shotcut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h6.667v24H0v-.889h5.778V.889H0V0zm7.556 0v24H24v-.889H8.444V.889H24V0H7.556zm1.388 22.611H24V1.389H8.944v21.222zM5.278 1.389H0v21.222h5.278V1.389z"/>`),
		g.Group(children),
	)
}