package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketMinusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.383 5L6.93.757L6.07.243L3.217 5H.703A.703.703 0 0 0 0 5.703v.439c0 2.944.685 5.847 2.002 8.48a.685.685 0 0 0 .612.378h9.772c.26 0 .496-.146.612-.379A18.96 18.96 0 0 0 15 6.141v-.438A.703.703 0 0 0 14.297 5h-2.514L8.93.243l-.86.514L10.617 5H4.383ZM5 10h5V9H5v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}