package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.65 0 0 268.65 0 600s268.65 600 600 600s600-268.65 600-600S931.35 0 600 0zm0 139.16c254.499 0 460.84 206.341 460.84 460.84S854.499 1060.84 600 1060.84S139.16 854.499 139.16 600S345.501 139.16 600 139.16zM450 300.439V899.56L900 600L450 300.439z"/>`),
		g.Group(children),
	)
}