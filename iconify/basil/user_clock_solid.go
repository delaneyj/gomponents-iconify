package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserClockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3.75a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5Zm-4 9.5A3.75 3.75 0 0 0 2.25 17v1.188c0 .754.546 1.396 1.29 1.517c2.21.361 4.443.536 6.676.524c.156-.001.25-.174.174-.31A6.968 6.968 0 0 1 9.5 16.5c0-.787.13-1.544.37-2.25a.213.213 0 0 0-.192-.28a7.247 7.247 0 0 1-1.928-.35l-.866-.284a1.752 1.752 0 0 0-.543-.086H6ZM17.25 15a.75.75 0 0 0-1.5 0v1.773c0 .24.115.465.309.606l1 .728a.75.75 0 0 0 .882-1.214l-.691-.502V15Z"/><path fill="currentColor" fill-rule="evenodd" d="M16.5 22a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}