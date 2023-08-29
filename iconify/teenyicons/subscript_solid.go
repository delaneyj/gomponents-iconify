package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 8.333L1.9 13.8l-.8-.6l4.275-5.7L1.1 1.8l.8-.6L6 6.667L10.1 1.2l.8.6l-4.275 5.7l4.275 5.7l-.8.6L6 8.333Zm6 4.167a1.5 1.5 0 0 1 1.5-1.5h.293a1.207 1.207 0 0 1 .854 2.06l-.94.94H15v1h-2.5a.5.5 0 0 1-.354-.854l1.793-1.792a.207.207 0 0 0-.146-.354H13.5a.5.5 0 0 0-.5.5h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}