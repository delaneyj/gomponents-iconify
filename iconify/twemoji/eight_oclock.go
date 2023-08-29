package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#99AAB5"/><circle cx="18" cy="18" r="14" fill="#E1E8ED"/><path fill="#66757F" d="M17 18a1 1 0 1 0 2 0V7a1 1 0 1 0-2 0v11z"/><path fill="#66757F" d="M9.34 23a.998.998 0 0 0 1.365.367l7.795-4.5a1 1 0 1 0-1-1.732l-7.795 4.5A.998.998 0 0 0 9.34 23z"/>`),
		g.Group(children),
	)
}