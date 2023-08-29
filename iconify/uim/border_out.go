package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM5 19h14V5H5Z"/><circle cx="12" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="12" r="1" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}