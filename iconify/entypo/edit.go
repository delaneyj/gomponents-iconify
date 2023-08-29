package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.561 2.439c-1.442-1.443-2.525-1.227-2.525-1.227L8.984 7.264L2.21 14.037L1.2 18.799l4.763-1.01l6.774-6.771l6.052-6.052c-.001 0 .216-1.083-1.228-2.527zM5.68 17.217l-1.624.35a3.71 3.71 0 0 0-.69-.932a3.742 3.742 0 0 0-.932-.691l.35-1.623l.47-.469s.883.018 1.881 1.016c.997.996 1.016 1.881 1.016 1.881l-.471.468z"/>`),
		g.Group(children),
	)
}