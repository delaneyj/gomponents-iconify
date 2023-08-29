package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnROne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.061 9.509v10.93a.508.508 0 0 0 .5.5a.5.5 0 0 0 .5-.5V9.509a3.5 3.5 0 0 1 3.5-3.5h9.9l.01 1.44a.5.5 0 0 0 .81.4l2.47-1.96a.5.5 0 0 0 0-.78l-2.49-1.94a.5.5 0 0 0-.81.4l.01 1.44h-9.9a4.507 4.507 0 0 0-4.5 4.5Z"/>`),
		g.Group(children),
	)
}