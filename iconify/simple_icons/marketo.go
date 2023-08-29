package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marketo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.146 0v24l6.134-4.886V3.334zm-2.853 18.758l-4.939 2.157V2.086l4.939 1.462zm-11.572-.553l3.78-.999V5.188l-3.762-.606z"/>`),
		g.Group(children),
	)
}