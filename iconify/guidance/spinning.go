package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6 17.52a7.558 7.558 0 0 0-.5-.02a3 3 0 1 0 0 6c2 0 9.913-3.152 12-3.75v-.5c-1.017-.17-5.255-.914-8.5-1.376m8.5-2.374V11m0 4.5a4 4 0 1 1-3.71 5.5m3.71-5.5a4.002 4.002 0 0 0-3.874 3M19 8l3-3m-3 4.5h-2.969a3 3 0 0 1-2.785-1.886L12 4.5h-.914a3 3 0 0 0-1.92.695L5.5 8.25v.25l3 4v.25S7.5 15 7.5 19M10 9.5l3 2v.25s-2 1.25-2.5 3M14.805 4.5s1.81-.557 2.135-1.776A1.768 1.768 0 0 0 15.698.561a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.963 2.61.963 2.61l.29.079Z"/>`),
		g.Group(children),
	)
}