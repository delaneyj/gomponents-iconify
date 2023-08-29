package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shadowweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M31.222 15.887A10.918 10.918 0 0 0 11.1 23.474a6.998 6.998 0 0 0 .399 13.985h21.166a10.835 10.835 0 1 0-7.65-18.508"/>`),
		g.Group(children),
	)
}