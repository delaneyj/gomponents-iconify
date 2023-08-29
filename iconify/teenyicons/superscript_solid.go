package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperscriptSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.5A1.5 1.5 0 0 1 13.5 0h.293a1.207 1.207 0 0 1 .854 2.06l-.94.94H15v1h-2.5a.5.5 0 0 1-.354-.854l1.793-1.792A.207.207 0 0 0 13.793 1H13.5a.5.5 0 0 0-.5.5h-1Zm-6.625 6L1.1 1.8l.8-.6L6 6.667L10.1 1.2l.8.6l-4.275 5.7l4.275 5.7l-.8.6L6 8.333L1.9 13.8l-.8-.6l4.275-5.7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}