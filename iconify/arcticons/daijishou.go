package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Daijishou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.37 17.875L24 11.466l-6.37 6.409h12.74zm0 12.189L24 36.8l-6.37-6.736h12.74z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.413 25.913L24 44.5l18.587-18.587H5.413zM42.75 22.25L24 3.5L5.25 22.25h37.5z"/>`),
		g.Group(children),
	)
}