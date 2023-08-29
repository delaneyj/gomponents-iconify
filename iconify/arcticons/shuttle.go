package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuttle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.794 15.379v18.305"/><circle cx="11.647" cy="33.684" r="6.147" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.675 8.169H42.5v9.482H30.131l-.511-2.272"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.298 15.379H17.794V5.897h12.369l2.135 9.482zM42.5 17.651v18.305"/><circle cx="36.353" cy="35.956" r="6.147" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}