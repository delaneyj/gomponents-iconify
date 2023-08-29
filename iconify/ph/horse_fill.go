package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M202.05 55A103.24 103.24 0 0 0 128 24h-8a8 8 0 0 0-8 8v27.53L11.81 121.19a8 8 0 0 0-2.59 11.05l13.78 22l.3.43a31.84 31.84 0 0 0 31.34 12.83c13.93-2.36 38.62-6.54 61.4 3.29l-26.6 36.57A84.71 84.71 0 0 1 69.34 194a8 8 0 1 0-10.67 12a103.32 103.32 0 0 0 69.26 26h2.17a104 104 0 0 0 72-177ZM124 112a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}