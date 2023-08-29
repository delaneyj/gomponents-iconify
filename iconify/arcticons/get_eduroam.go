package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GetEduroam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.277 38.951a7.428 7.428 0 1 1-12.41-8.165"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 34.869a7.428 7.428 0 1 0-14.856 0H42.5m-25.758 2.956c-1.475-10.627 5.224-22.85 20.485-21.376"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.798 37.825C3.441 20.572 15.154 3.902 37.227 5.773"/>`),
		g.Group(children),
	)
}