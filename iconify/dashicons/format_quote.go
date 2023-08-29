package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatQuote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.54 12.74c0-.87-.24-1.61-.72-2.22c-.73-.92-2.14-1.03-2.96-.85c-.34-1.93 1.3-4.39 3.42-5.45L6.65 1.94C3.45 3.46.31 6.96.85 11.37C1.19 14.16 2.8 16 5.08 16c1 0 1.83-.29 2.48-.88c.66-.59.98-1.38.98-2.38zm9.43 0c0-.87-.24-1.61-.72-2.22c-.73-.92-2.14-1.03-2.96-.85c-.34-1.93 1.3-4.39 3.42-5.45l-1.63-2.28c-3.2 1.52-6.34 5.02-5.8 9.43c.34 2.79 1.95 4.63 4.23 4.63c1 0 1.83-.29 2.48-.88c.66-.59.98-1.38.98-2.38z"/>`),
		g.Group(children),
	)
}