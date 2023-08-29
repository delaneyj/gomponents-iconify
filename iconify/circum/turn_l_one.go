package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnLOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.939 9.509v10.93a.508.508 0 0 1-.5.5a.5.5 0 0 1-.5-.5V9.509a3.5 3.5 0 0 0-3.5-3.5h-9.9l-.01 1.44a.5.5 0 0 1-.81.4l-2.47-1.96a.5.5 0 0 1 0-.78l2.49-1.94a.5.5 0 0 1 .81.4l-.01 1.44h9.9a4.507 4.507 0 0 1 4.5 4.5Z"/>`),
		g.Group(children),
	)
}