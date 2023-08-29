package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1504 128q113 0 212 42t172 116t116 173t43 212q0 58-12 115t-36 110h-463l-192-200l-320 320l-448-448l-320 328H49q-24-53-36-110T1 671q0-113 42-212t116-172t173-116t212-43q109 0 208 41t177 118l95 96l95-96q77-77 176-118t209-41zm-96 896h510l-14 16q-7 8-15 17l-865 864l-865-864q-8-8-15-16t-14-17h254l192-184l448 448l320-320l64 56z"/>`),
		g.Group(children),
	)
}