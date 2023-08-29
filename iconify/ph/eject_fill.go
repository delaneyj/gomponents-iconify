package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 176v24a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16v-24a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16ZM48.24 144h159.52a16.18 16.18 0 0 0 14.93-9.76a15.59 15.59 0 0 0-3.1-17.12l-73.73-77.51a24.76 24.76 0 0 0-35.72 0l-73.73 77.51a15.59 15.59 0 0 0-3.1 17.12A16.18 16.18 0 0 0 48.24 144Z"/>`),
		g.Group(children),
	)
}