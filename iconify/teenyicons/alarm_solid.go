package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.854.854l-3 3l-.708-.708l3-3l.708.708Zm10.292 3l-3-3l.708-.708l3 3l-.707.708Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 8.5a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0ZM8 8V5H7v4h3V8H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}