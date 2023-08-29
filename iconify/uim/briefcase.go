package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 6V5h4v1h2V5a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v1h2z"/><path fill="currentColor" d="M9 15a1 1 0 0 1-1-1v-2a1 1 0 1 1 2 0v2a1 1 0 0 1-1 1zm6 0a1 1 0 0 1-1-1v-2a1 1 0 1 1 2 0v2a1 1 0 0 1-1 1z" opacity=".25"/><path fill="currentColor" d="M20 6H4a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h4v-1a1 1 0 1 1 2 0v1h4v-1a1 1 0 1 1 2 0v1h4a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2z" opacity=".5"/><path fill="currentColor" d="M20 13h-4v1a1 1 0 1 1-2 0v-1h-4v1a1 1 0 1 1-2 0v-1H4a2 2 0 0 1-2-2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}