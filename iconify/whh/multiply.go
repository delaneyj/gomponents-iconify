package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m701.393 511l303 304q20 20 20 47.5t-20 46.5l-94 94q-19 20-47 20t-47-20l-304-303l-304 303q-19 20-47 20t-47-20l-94-94q-20-19-20-46.5t20-47.5l303-304l-303-303q-20-20-20-47.5t20-47.5l94-94q19-19 47-19t47 19l304 304l304-304q19-19 47-19t47 19l94 94q20 20 20 47.5t-20 47.5z"/>`),
		g.Group(children),
	)
}