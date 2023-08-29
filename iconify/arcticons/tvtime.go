package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tvtime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h12.333v12.333H5.5zm12.334 0h12.333v12.333H17.834zm12.333 0H42.5v12.333H30.167zM17.834 17.834h12.333v12.333H17.834zm0 12.333h12.333V42.5H17.834z"/>`),
		g.Group(children),
	)
}