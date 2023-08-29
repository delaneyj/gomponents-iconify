package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m6.75 6.75l3.75 3.75M6.75 6.75L8.5 5v-.25l-1.44-.6a12 12 0 0 1-2.88-1.706L3 1.5h-.25L1.5 2.75V3l.944 1.18A12 12 0 0 1 4.15 7.06l.6 1.44H5l1.75-1.75ZM13.734 17a102.1 102.1 0 0 1 4.516 5.5h.25l4-4v-.25s-2.47-1.852-5.5-4.516m.5-2.234a5 5 0 0 0 4.584-7H21.5l-3 3h-2v-2l3-3v-.584a5 5 0 0 0-6.703 6.287L11 10c-4.5 4.5-9.5 8.25-9.5 8.25v.25l4 4h.25S9.5 17.5 14 13l1.797-1.797a4.989 4.989 0 0 0 1.703.297Z"/>`),
		g.Group(children),
	)
}