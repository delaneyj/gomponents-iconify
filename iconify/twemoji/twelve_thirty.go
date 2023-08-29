package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwelveThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#99AAB5"/><circle cx="18" cy="18" r="14" fill="#E1E8ED"/><path fill="#66757F" d="M17 18a1 1 0 1 1 2 0v11a1 1 0 0 1-2 0V18z"/><path fill="#66757F" d="M19 18a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0v9z"/>`),
		g.Group(children),
	)
}