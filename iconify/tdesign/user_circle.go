package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9a8.96 8.96 0 0 0 1.773 5.365A5.002 5.002 0 0 1 9.5 14h5a5.002 5.002 0 0 1 4.727 3.365A8.957 8.957 0 0 0 21 12a9 9 0 0 0-9-9Zm5.5 16.125V19a3 3 0 0 0-3-3h-5a3 3 0 0 0-3 3v.125A8.957 8.957 0 0 0 12 21c2.072 0 3.979-.7 5.5-1.875ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11a10.98 10.98 0 0 1-3.85 8.36A10.96 10.96 0 0 1 12 23a10.96 10.96 0 0 1-7.15-2.64A10.977 10.977 0 0 1 1 12Zm11-6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM7.5 8.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Z"/>`),
		g.Group(children),
	)
}