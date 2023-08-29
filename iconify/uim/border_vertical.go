package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="7" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="15" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="7" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="15" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="7" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="3" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="15" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19" cy="20" r="1" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}