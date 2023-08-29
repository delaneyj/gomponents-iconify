package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 320q0 87-43 161T864.5 597.5T704 640h-32q-96 43-96 128q0 102 80 139q52 9 82 23t30 30q0 27-75 45.5T512 1024t-181-18.5t-75-45.5q0-16 30-30t82-23q80-37 80-139q0-85-96-128h-32q-87 0-160.5-42.5T43 481T0 320q0-96 52.5-166T193 54q-1-31-1-54h640q0 24-1 54q88 30 140.5 100t52.5 166zM205 188q-77 43-77 132q0 77 53.5 132.5T311 511q-78-121-106-323zm365 5L512 64l-58 129l-134 18l99 97l-26 140l119-68l119 68l-26-140l99-97zm249-5q-28 202-106 323q76-3 129.5-58.5T896 320q0-89-77-132z"/>`),
		g.Group(children),
	)
}