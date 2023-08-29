package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ditmm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="32.949" cy="37.12" r=".75" fill="currentColor" transform="rotate(-50.126 32.949 37.12)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.885 32.886l-.055 1.31m-9.283-16.588a9.721 9.721 0 0 1 3.018-6.87a9.78 9.78 0 0 1 7.365-2.31a9.895 9.895 0 0 1 9.562 10.038a9.72 9.72 0 0 1-3.018 6.87c-2.34 1.65-7.439 4.056-7.607 7.987M4.77 15.77a5.612 5.612 0 1 1 10.684-3.435m6.368 19.896a5.612 5.612 0 1 1-10.685 3.436L4.77 15.77m10.684-3.435l6.368 19.896m-3.184-9.948L7.954 25.719"/>`),
		g.Group(children),
	)
}