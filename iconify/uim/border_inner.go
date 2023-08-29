package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.965 13h-16a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2Z"/><path fill="currentColor" d="M11.965 21a1 1 0 0 1-1-1V4a1 1 0 0 1 2 0v16a1 1 0 0 1-1 1Z"/><circle cx="3.964" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="3.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="3.964" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="3.964" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="7.964" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="15.964" cy="4" r="1" fill="currentColor" opacity=".5"/><circle cx="7.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="15.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="16" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="20" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="8" r="1" fill="currentColor" opacity=".5"/><circle cx="19.964" cy="4" r="1" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}