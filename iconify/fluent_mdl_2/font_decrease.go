package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1027 768l342 1024h-135l-85-256H771l-85 256H551L893 768h134zm79 640L960 971l-146 437h292zM2048 0l-320 320L1408 0h640z"/>`),
		g.Group(children),
	)
}