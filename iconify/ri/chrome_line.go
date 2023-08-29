package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.364 19.833l1.93-3.342a4.499 4.499 0 0 1-4.234-2.315L4.795 8.52a8.003 8.003 0 0 0 5.57 11.313Zm2.226.146A8 8 0 0 0 19.602 9.5h-3.86c.479.715.758 1.575.758 2.5c0 .848-.234 1.64-.642 2.318l-3.268 5.66Zm1.553-6.691l.022-.038a2.5 2.5 0 1 0-4.354-.042l.024.042a2.499 2.499 0 0 0 4.308.037Zm-8.108-6.62l1.928 3.34A4.5 4.5 0 0 1 12 7.5h6.615A7.992 7.992 0 0 0 12 4a7.98 7.98 0 0 0-5.965 2.669ZM12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/>`),
		g.Group(children),
	)
}