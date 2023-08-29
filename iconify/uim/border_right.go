package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="16.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="8.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="12" r="1" fill="currentColor" opacity=".5"/><circle cx="16.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="8.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="16.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="12.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8.5" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="4.5" cy="4" r="1" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}