package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Surprised(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm5-2a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm8 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-2 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}