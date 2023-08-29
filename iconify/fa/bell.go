package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M848 1696q0-16-16-16q-59 0-101.5-42.5T688 1536q0-16-16-16t-16 16q0 73 51.5 124.5T832 1712q16 0 16-16zm816-288q0 52-38 90t-90 38h-448q0 106-75 181t-181 75t-181-75t-75-181H128q-52 0-90-38t-38-90q50-42 91-88t85-119.5t74.5-158.5t50-206T320 576q0-152 117-282.5T744 135q-8-19-8-39q0-40 28-68t68-28t68 28t28 68q0 20-8 39q190 28 307 158.5T1344 576q0 139 19.5 260t50 206t74.5 158.5t85 119.5t91 88z"/>`),
		g.Group(children),
	)
}