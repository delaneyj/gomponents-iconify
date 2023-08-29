package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#99AAB5"/><circle cx="18" cy="18" r="14" fill="#E1E8ED"/><path fill="#66757F" d="M19 18a1 1 0 1 1-2 0V7a1 1 0 0 1 2 0v11z"/><path fill="#66757F" d="M26.661 13a1 1 0 0 1-.366 1.366l-7.795 4.5a1 1 0 0 1-1-1.732l7.795-4.5a1 1 0 0 1 1.366.366z"/>`),
		g.Group(children),
	)
}