package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.955 2.045a.75.75 0 0 1 0 1.06l-.55.551c3.288 3.831 3.118 9.61-.51 13.238a9.712 9.712 0 0 1-6.145 2.827v1.53h1.75a.75.75 0 1 1 0 1.5h-5a.75.75 0 0 1 0-1.5h1.75v-1.53a9.709 9.709 0 0 1-5.594-2.316l-.55.55a.75.75 0 0 1-1.061-1.06l.707-.708A1.272 1.272 0 0 1 5.5 16.15a8.222 8.222 0 0 0 5.497 2.101h.006A8.25 8.25 0 0 0 17.15 4.5a1.272 1.272 0 0 1 .038-1.748l.707-.707a.75.75 0 0 1 1.061 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M4.25 10a6.75 6.75 0 1 1 13.5 0a6.75 6.75 0 0 1-13.5 0ZM11 4.75a5.25 5.25 0 0 0-5.159 6.23a2.25 2.25 0 1 1 2.56 3.583a5.252 5.252 0 0 0 7.746-3.522a3.75 3.75 0 1 1-2.89-5.782A5.23 5.23 0 0 0 11 4.75ZM10.75 9a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0ZM7.5 11.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}