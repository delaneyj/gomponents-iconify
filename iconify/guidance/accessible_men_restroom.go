package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibleMenRestroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24M4.35 4.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C6.246 3.5 4.65 4.5 4.65 4.5h-.3ZM19 14V8.5h-1.158a3 3 0 0 0-2.91 2.272L14 14.527M24 23.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4v-1.5h-3.258m-3.742 7a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm4.35-17s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}