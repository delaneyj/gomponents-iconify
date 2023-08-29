package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDecreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#FFF" d="M4 4h120v120H4z"/><path fill="none" stroke="#B0BEC5" stroke-miterlimit="10" stroke-width="1.993" d="M24.7 4.2v119.6M44.35 4.2v119.6M64 4.2v119.6M83.65 4.2v119.6M103.3 4.2v119.6"/><path fill="none" stroke="#B0BEC5" stroke-miterlimit="10" stroke-width="2" d="M123.8 24.66H4.15m119.66 19.67H4.16M123.83 64H4.17m119.67 19.67H4.19m119.66 19.67H4.2"/><path fill="#0B88CF" d="m5.1 5.14l-.08 6.17l26.92 48.51l23.7-21.91l63.78 85.1h3.49l.14-5.16l-66.6-88.86l-23.02 21.28L8.45 5.11z"/><path fill="#B0BEC5" d="M122.01 5.99V122H5.99V5.99h116.02M124 4H4v120h120V4z"/>`),
		g.Group(children),
	)
}