package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Broom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1005 110L697 418q19 51-9 79t-79 9l-1 1q29 47 30.5 95T621 689t-54 73q-89 89-128 262q-46-23-101-60q-11-50 4-99t52-83q-46 15-79 50t-45 81q-49-39-85-74q-2-2-5.5-6t-4.5-6q34-34 48-77q-35 31-78 45q-31-35-55-67q44-14 78-46.5t48-77.5q-32 36-78 51.5T44 663q-27-43-44-78q173-40 262-128q34-35 73-54t87-17.5t95 30.5l1-1q-19-51 9-79t79-9L914 19q19-19 45.5-19t45.5 19t19 45.5t-19 45.5z"/>`),
		g.Group(children),
	)
}