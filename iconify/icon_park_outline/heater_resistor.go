package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeaterResistor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="28" height="12" x="9.858" y="29.657" rx="2" transform="rotate(-45 9.858 29.657)"/><path stroke-linecap="round" d="m7.03 40.97l7.07-7.07m19.801-19.8l7.071-7.07M14.808 24.707l8.485 8.485m-3.535-13.435l8.485 8.486m-3.536-13.435l8.485 8.485m-20.505 3.536l14.142-14.143m-5.657 22.628l14.142-14.142"/></g>`),
		g.Group(children),
	)
}