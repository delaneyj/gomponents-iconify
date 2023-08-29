package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindReplaceLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.033 16.618l4.28 4.282l-1.413 1.414l-4.282-4.281A8.963 8.963 0 0 1 11 20a8.998 8.998 0 0 1-8.065-5H9l-1.304 2.173A6.972 6.972 0 0 0 11 18a6.977 6.977 0 0 0 4.875-1.975l.15-.15A6.977 6.977 0 0 0 18 11c0-.695-.101-1.366-.29-2h2.067c.146.643.223 1.313.223 2a8.963 8.963 0 0 1-1.967 5.618ZM19.065 7H13l1.304-2.173A6.972 6.972 0 0 0 11 4a6.999 6.999 0 0 0-6.71 9H2.223A9.038 9.038 0 0 1 2 11c0-4.973 4.027-9 9-9a8.997 8.997 0 0 1 8.065 5Z"/>`),
		g.Group(children),
	)
}