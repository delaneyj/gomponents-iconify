package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M33.609 21.068h-7.736v8.408h7.736c2.354 0 4.268-1.887 4.268-4.203c0-2.318-1.913-4.205-4.268-4.205m0 13.455h-7.736v8.408h7.736c2.354 0 4.268-1.887 4.268-4.205s-1.913-4.203-4.268-4.203"/><path fill="currentColor" d="M52 2H12C6.477 2 2 6.476 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.524-4.477-10-10-10zm-9 36.727c0 5.102-4.211 9.25-9.391 9.25L21 48V16l12.609.023c5.18 0 9.391 4.148 9.391 9.25c0 2.648-1.137 5.04-2.956 6.727C41.863 33.688 43 36.078 43 38.727z"/>`),
		g.Group(children),
	)
}