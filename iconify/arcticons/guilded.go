package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guilded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.107 14.664a23.9 23.9 0 0 0 2.408 9.41c2.44 4.605 5.804 7.563 8.58 8.741a17.14 17.14 0 0 0 7.15-6.175h-7.873a11.888 11.888 0 0 1-4.008-7.726H42.46a34.477 34.477 0 0 1-5.858 12.953a27.8 27.8 0 0 1-12.49 9.196h-.073A28.282 28.282 0 0 1 8.36 26.698C6.472 22.938 4.5 16.114 4.5 6.938h39a55.626 55.626 0 0 1-.52 7.726Z"/>`),
		g.Group(children),
	)
}