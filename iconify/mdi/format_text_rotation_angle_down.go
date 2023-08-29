package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextRotationAngleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.25 21h-4.22l1.41-1.41l-8.86-8.86l1.45-1.4l8.81 8.86l1.41-1.41M12.61 8l2.62 2.64l2.2-4.87m1.98-.85l-4.46 11.11l-1.45-1.45l.89-2.2l-3.51-3.57l-2.2.94l-1.46-1.5l11.11-4.41Z"/>`),
		g.Group(children),
	)
}