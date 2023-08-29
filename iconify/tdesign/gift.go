package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 22V6h-4.535A4 4 0 0 0 12 1.354A4 4 0 0 0 5.535 6H1v16h22ZM13 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm1.414 4H21v6H3V8h6.586l-3 3L8 12.414l4-4l4 4L17.414 11l-3-3ZM3 16h18v4H3v-4Zm8-12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}