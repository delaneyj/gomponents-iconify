package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firestrikeretro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.757 28.346c9.656 0 21.368 1.865 18.212 5.013c-6.094 6.078-35.327 4.244-37.997 1.45c-3.056-3.197 9.14-6.463 19.785-6.463Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.505 33.82c0-2.405 2.67-4.922 2.67-4.922s.993-18.31 17.306-18.31s16.136 17.404 16.136 17.404a5.33 5.33 0 0 1 2.883 4.444"/>`),
		g.Group(children),
	)
}