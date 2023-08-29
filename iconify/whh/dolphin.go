package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolphin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M805.193 888q17 33 24 63.5t.5 51.5t-29.5 21q-61-82-119-180q-28-20-85-50.5t-102-55.5t-103-74t-103-110q0 12-8 22.5t-16 25t-8 38.5v8q0 16-.5 23t-3 16.5t-9.5 13t-19 3.5q-25 0-60.5-50t-35.5-78q0-54 19-121.5t43-70.5q-32-58-42-78q-24-39-65-100t-62-95t-21-47q0-26 18.5-45t45.5-19t45.5 19t18.5 45q0 7 5 32q46-32 123-32q48 0 132.5 69t123.5 123q0 1 1 2l1 1q55 9 113 36.5t99.5 69.5t41.5 83v1q-15 1-59.5-.5t-67.5-1.5q20 36 31 65q13 36 22 86q8 4 10 10q18 44 57.5 88t83 78t85 66t68 63t26.5 57l-34-10l-88.5-28l-96.5-34z"/>`),
		g.Group(children),
	)
}