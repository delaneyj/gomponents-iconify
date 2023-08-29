package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapAiming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm12-7.667V5.52A6.56 6.56 0 0 1 18.48 11h1.187v2H18.48A6.56 6.56 0 0 1 13 18.48v1.187h-2V18.48A6.56 6.56 0 0 1 5.52 13H4.333v-2H5.52A6.56 6.56 0 0 1 11 5.52V4.333h2Zm-1 3.111a4.556 4.556 0 1 0 0 9.112a4.556 4.556 0 0 0 0-9.112Zm0 4.445a.111.111 0 1 0 0 .222a.111.111 0 0 0 0-.222ZM9.889 12a2.111 2.111 0 1 1 4.222 0a2.111 2.111 0 0 1-4.222 0Z"/>`),
		g.Group(children),
	)
}