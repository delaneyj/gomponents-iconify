package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 20.5a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1h16a1 1 0 0 1 0 2h-15v15a1 1 0 0 1-1 1Z"/><circle cx="19.5" cy="11.5" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="7.5" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="15.5" r="1" fill="currentColor" opacity=".5"/><circle cx="7.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="11.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="15.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/><circle cx="19.5" cy="19.5" r="1" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}