package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalibreIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#3057F4" d="M28.423 222.25c48.024 29.86 84.58-8.211 84.58-8.211c-13.064 46.678-43.093 55.183-43.093 55.183l6.775 13.539h102.63l6.787-13.539s-30.04-8.505-43.093-55.195c0 0 36.544 38.082 84.568 8.212C275.601 192.39 275.092 76.459 128 0C-19.092 76.436-19.601 192.38 28.423 222.25Z"/>`),
		g.Group(children),
	)
}