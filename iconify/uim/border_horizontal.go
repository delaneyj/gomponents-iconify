package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 13H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/><circle cx="12" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="4" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="20" cy="4" r="1" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}