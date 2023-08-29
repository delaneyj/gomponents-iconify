package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11ZM9.525 9.526a3.5 3.5 0 0 0 4.95 4.95l.707-.708l1.414 1.415l-.707.707a5.5 5.5 0 1 1 0-7.778l.707.707l-1.414 1.414l-.707-.707a3.5 3.5 0 0 0-4.95 0Z"/>`),
		g.Group(children),
	)
}