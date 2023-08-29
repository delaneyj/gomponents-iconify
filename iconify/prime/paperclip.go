package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.94 20.74a5.85 5.85 0 0 1-4-1.56a5.23 5.23 0 0 1 0-7.68l7.56-7.14a4.22 4.22 0 0 1 5.69 0A4.1 4.1 0 0 1 19.5 7.3a3.46 3.46 0 0 1-1.1 2.55L10.83 17a2.47 2.47 0 0 1-3.36 0a2.23 2.23 0 0 1 0-3.28l7-6.59a.75.75 0 0 1 1.06 0a.75.75 0 0 1 0 1.06l-7 6.59a.73.73 0 0 0 0 1.1a1 1 0 0 0 1.3 0l7.57-7.13A2 2 0 0 0 18 7.3a2.57 2.57 0 0 0-.84-1.84a2.67 2.67 0 0 0-3.63 0L6 12.59a3.73 3.73 0 0 0 0 5.5a4.4 4.4 0 0 0 6 0L19.49 11a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06L13 19.18a5.84 5.84 0 0 1-4.06 1.56Z"/>`),
		g.Group(children),
	)
}