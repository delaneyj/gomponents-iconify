package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Endocrine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M23.147 44C31.293 44 40 43.01 40 30.18c0-6.216-1.492-11.276-4.474-15.18c-2.787 4.667-4.852 7-6.196 7c-2.016 0 0-6.09-2.331-10.709c-1.554-3.08-4.311-5.51-8.27-7.291c.55 4.437.431 7.655-.357 9.655c-1.184 3-11.22 8.255-10.314 17.28C8.965 39.96 15 44 23.147 44Z" clip-rule="evenodd"/><path d="M21.708 24.008c-3.805 3.459-5.427 6.66-4.865 9.603c.483 2.534 2.406 3.765 3.854 4.163c1.123.307 3.067.48 4.803-1.067c1.84-1.639 1.107-3.968-1.56-6.85c-1.126-1.215-1.87-3.165-2.232-5.849Z"/></g>`),
		g.Group(children),
	)
}