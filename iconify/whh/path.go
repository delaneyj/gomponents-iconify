package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Path(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 768q-74 0-128-35v163q0 53-37.5 90.5T384 1024h-64q-53 0-90.5-37.5T192 896q0-3 .5-7t4-15t9.5-19.5t19-15.5t31-7h19l8.5-.5l9-1l7.5-1.5l7-2.5l5-4l4.5-5.5l2.5-7.5l1-9.5V384q0-27 18.5-45.5T384 320h64q27 0 45.5 18.5T512 384v128q0 28 38.5 46t89.5 18q77 0 134.5-52T832 384q0-62-43.5-106t-115-65T512 192q-69 0-125 14t-101 43.5t-69.5 80T192 448q0 36 39 89L80 688Q0 578 0 448q0-93 29-169t77.5-127t115-86t139-50.5T512 0q78 0 151.5 15t139 46.5t115 76.5t78 108t28.5 138q0 81-24.5 150t-71 121.5t-121 82.5T640 768z"/>`),
		g.Group(children),
	)
}