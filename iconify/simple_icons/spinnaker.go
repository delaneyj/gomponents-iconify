package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinnaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.343 0C17.785 8.741 11.317 21.987.815 23.882c10.806 1.064 19.481-5.327 21.646-8.066C24.627 13.076 21.343 0 21.343 0zM.815 23.882L.8 23.88v.004l.015-.003zM17.182 5.8C15.409 10.988 10.477 18.547 5.4 20.39c.885.033 1.74-.019 2.561-.132c3.989-3.221 7.14-8.037 9.577-12.771c-.193-.981-.356-1.687-.356-1.687z"/>`),
		g.Group(children),
	)
}