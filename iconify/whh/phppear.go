package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phppear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015.476 158l-119 200v603q0 27-18.5 45.5t-45.5 18.5h-640q-27 0-45.5-19t-18.5-45V358l-119-200q-14-22-7-47t30-38t48.5-6t38.5 29l95 160h596l95-160q13-22 38.5-29t48.5 6t30 38t-7 47zm-281 226q-68 17-102 78q-27-14-56-14q-35 0-64.5 18t-46.5 48l-6.5-1l-10.5-1q-80 0-136 56.5t-56 136t56 136t136 56.5t136-56.5t56-135.5q0-4-.5-10t-.5-7q29-17 47-46.5t18-64.5q0-39-22-72q11-26 34-43t52-18q-9-35-34-60zm-542-320q0-26 18.5-45t45.5-19h512q27 0 45.5 19t18.5 45v128h-640V64z"/>`),
		g.Group(children),
	)
}