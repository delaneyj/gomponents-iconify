package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DidiGreaterChina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.805 8.3H36.36v7.306h6.55s4.255 17.342-10.068 22.19c-9.568 3.24-16.858 2.367-22.516-2.76C2.918 28.325 4.602 8.224 4.805 8.299Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.22 15.877c-.448 7.107-.36 13.794 6.82 15.534"/>`),
		g.Group(children),
	)
}