package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Germanos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.896 45.5A21.277 21.277 0 0 1 2.854 23.997A21.277 21.277 0 0 1 23.896 2.5a20.364 20.364 0 0 1 15.14 6.867a4.956 4.956 0 0 1 1.381 4.243a4.45 4.45 0 0 1-2.904 3.35a4.354 4.354 0 0 1-4.289-.877c-3.001-2.35-5.703-4.477-9.328-4.477a12.264 12.264 0 0 0-12.132 12.392a12.264 12.264 0 0 0 12.128 12.396h.004a12.815 12.815 0 0 0 9.253-4.38a4.243 4.243 0 0 1 6.091-.09a4.42 4.42 0 0 1 1.143 1.959a4.699 4.699 0 0 1-1.028 4.338A20.48 20.48 0 0 1 23.896 45.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.59 26.105a4.305 4.305 0 0 0-1.53-5.81a4.097 4.097 0 0 0-2.08-.569H27.804a4.26 4.26 0 0 0 0 8.503H40.98a4.158 4.158 0 0 0 3.61-2.124Z"/>`),
		g.Group(children),
	)
}