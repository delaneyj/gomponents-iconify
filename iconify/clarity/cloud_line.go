package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M27.14 33H10.62C5.67 33 1 28.19 1 23.1a10 10 0 0 1 8-9.75a10.19 10.19 0 0 1 20.33 1.06a10.07 10.07 0 0 1-.33 2.25a8.29 8.29 0 0 1 6 8C35 29.1 31.33 33 27.14 33ZM19.09 6.23a8.24 8.24 0 0 0-8.19 8v.87l-.86.1A7.94 7.94 0 0 0 3 23.1c0 4 3.77 7.9 7.62 7.9h16.52c3.07 0 5.86-3 5.86-6.35a6.31 6.31 0 0 0-5.37-6.26l-1.18-.18l.39-1.13a8.18 8.18 0 0 0-7.75-10.85Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}