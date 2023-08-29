package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiningRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M6 0S5 7 5 8s2.25 3 2.25 3L7 24.344C7 25.448 7.896 26 9 26s2-.552 2-1.656L10.75 11S13 8.958 13 8c0-.958-1-8-1-8h-1v6.5a.5.5 0 0 1-1 0c0-.093-.344-6.5-.344-6.5H8.344S8 6.407 8 6.5a.5.5 0 0 1-1 0V0H6zm10.719.063C16.13.204 16 .835 16 1.53v22.813c0 1.104.896 1.656 2 1.656s1.969-.553 1.969-1.656c0-5.087-1-7.799-1-10.344c0-1.148 2.031-2.626 2.031-7.969c0-3.268-2.896-5.968-4-5.968c-.104 0-.197-.02-.281 0z"/>`),
		g.Group(children),
	)
}