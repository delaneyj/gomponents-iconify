package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungFinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="27.569" cy="23.856" r="7.378" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.786 29.073l1.88 1.88m-1.156 1.155l2.311-2.312l6.505 6.505l-2.312 2.312z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 31.234V12.55a3.162 3.162 0 0 0-3.162-3.163H7.662A3.162 3.162 0 0 0 4.5 12.55v18.973a3.162 3.162 0 0 0 3.162 3.162h22.195"/>`),
		g.Group(children),
	)
}