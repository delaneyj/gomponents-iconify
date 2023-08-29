package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M12.5 16.5v17h3v-17h-3zm6 0v17h3v-17h-3zm6 0v17h3v-17h-3zm3-12c0-2.5-.609-3-3-3h-10c-2.52 0-2.98.55-2.98 3.01L11.5 7.5h-8c-1.48 0-2 .49-2 2v1c0 1.55.52 2 2 2h1v26c0 2.49.55 3 3 3h24c2.5 0 4-.471 4-3v-26h1c1.51 0 2-.48 2-2v-1c0-1.48-.43-2-2-2h-9v-3zm-3 0v3h-10v-3h10zm-15 8h21v24h-21v-24z"/>`),
		g.Group(children),
	)
}