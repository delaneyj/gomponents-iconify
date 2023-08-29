package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenMustBeSupervised(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M19.25 24v-3.667c0-1.145.784-2.074 1.75-2.074v-6.222S19.687 11 17.5 11c-2.188 0-3.5 1.037-3.5 1.037v6.222c.967 0 1.75.929 1.75 2.075V24M9 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M17.373 9S16 8.125 16 7.031c0-.845.672-1.531 1.502-1.531S19 6.186 19 7.031C19 8.125 17.63 9 17.63 9h-.257ZM6.85 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C8.746 3.5 7.15 4.5 7.15 4.5h-.3Z"/>`),
		g.Group(children),
	)
}