package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stitcher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.59 8.516H24v6.928h-4.41zM0 8.854h4.41v7.803H0zm4.914-1.328h4.388v8.572H4.914zm4.892.725h4.388v8.81H9.806zm4.892-1.312h4.388v9.158h-4.388Z"/>`),
		g.Group(children),
	)
}