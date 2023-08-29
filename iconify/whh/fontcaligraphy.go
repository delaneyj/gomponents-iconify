package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontcaligraphy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 128q-24 0-37 .5t-34 3.5t-34.5 9.5T794 161t-26 31q-9 26-27 71t-55 127t-78 162.5t-96.5 168t-110 153t-118.5 108t-123 42.5q-66 0-113-47T0 864q0-42 16-69t48-27q28 0 46 16.5t18 47.5q0 28-16 55.5t-16 41t22 22.5t42 9t43-13t48-40t49.5-57t52.5-75t51-83t51.5-92t47.5-91t45-91t38.5-80.5T620 267q14-31 20-43q-26 21-69.5 58.5t-74 62t-70.5 51t-82.5 39.5t-87.5 13q-93 0-142.5-49.5T64 256Q64 0 416 0q48 0 110.5 18.5t125.5 43T740 92Q835 0 928 0q40 0 68 19t28 45.5t-28 45t-68 18.5zM416 64q-96 0-160 53t-64 130q0 59 24 98t72 39q39 0 187-101t186-134q-14-6-68.5-31.5T495 78t-79-14z"/>`),
		g.Group(children),
	)
}