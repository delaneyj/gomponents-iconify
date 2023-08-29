package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrystalGrowth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m253.8 15.56l-79.9 84.11l2.3 58.83l50.6 36.2l31.9 182l10.8-26.9l11.8-235.4l18.7 1l-9.1 181l28.3-70.8l8.2-108l.9-17.93zm139 50.57l-46.6 50.77l-3.9 51.1l10.6-26.2l30.4-13.7c3.2-20.6 6.3-41.3 9.5-61.97zm60.3 51.17l-85.7 38.4l-102.6 255.9l14.6 83.3h7.8l147.6-293.1l16.7 8.4l-143.4 284.7h24.4l146.6-291.8zm-340.2 18.9l-54.11 99.1l69.11 259.6h93.6l-51.1-274.8l18.3-3.4l51.8 278.2h19.9l-50.7-289.4zm358.3 260.4l-65.8-5.2l-49.8 99.2l69.8-36.7zm-435.96-28l42.47 126.7h30.99L80.6 389.9z"/>`),
		g.Group(children),
	)
}