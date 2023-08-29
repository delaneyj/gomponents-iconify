package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-4 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm8-14H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Zm-1 16H5V5h14ZM8 13a1 1 0 1 0-1-1a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}