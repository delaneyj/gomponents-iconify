package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpYard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20V2zM8 8.22a1.562 1.562 0 0 1 2.45-1.28l-.01-.12c0-.86.7-1.56 1.56-1.56s1.56.7 1.56 1.56l-.01.12A1.562 1.562 0 0 1 16 8.22c0 .62-.37 1.16-.89 1.4c.52.25.89.79.89 1.41c0 .86-.7 1.56-1.56 1.56c-.33 0-.64-.11-.89-.28l.01.12c0 .86-.7 1.56-1.56 1.56s-1.56-.7-1.56-1.56l.01-.12A1.562 1.562 0 0 1 8 11.03c0-.62.37-1.16.89-1.4C8.37 9.38 8 8.84 8 8.22zM12 19c-3.31 0-6-2.69-6-6c3.31 0 6 2.69 6 6c0-3.31 2.69-6 6-6c0 3.31-2.69 6-6 6z"/><circle cx="12" cy="9.62" r="1.56" fill="currentColor"/>`),
		g.Group(children),
	)
}