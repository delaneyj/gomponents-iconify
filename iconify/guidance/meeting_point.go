package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeetingPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14 24v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1v8a2 2 0 0 1 2 2V24m-6.5 0v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1m15 17.5v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1m-6.65-2s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Zm-6.5 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C7.246 3.5 5.65 4.5 5.65 4.5h-.3Zm13 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}