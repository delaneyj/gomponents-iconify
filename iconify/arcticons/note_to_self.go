package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteToSelf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.2 17.065c-1.679-3.274-6.328-3.264-8.042-.008A40.408 40.408 0 0 0 6.32 27.782a40.401 40.401 0 0 0-.777 9.979c.131 2.825 3.735 3.928 5.4 1.642A289.97 289.97 0 0 1 37.148 8.3c1.852-1.91 5.08-.702 5.263 1.952c.275 4.006 0 9.222-1.991 14.891a37.103 37.103 0 0 1-6.385 11.297c-1.59 1.93-4.62 1.678-5.824-.514a735.848 735.848 0 0 1-10.01-18.86Z"/>`),
		g.Group(children),
	)
}