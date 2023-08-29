package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindmillEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.375 3.938L5.937 3.5l.438-.438h.438L9 .876L8.125 0L5.937 2.188v.437l-.437.438l-.438-.438v-.438L2.876 0L2 .875l2.188 2.188h.437l.438.437l-.438.438h-.438L2 6.124L2.875 7l2.188-2.188v-.437l.437-.438l.438.438v.438L8.124 7L9 6.125L6.812 3.937h-.437zM8.5 10H8L7 7.5L5.5 6L4 7.5L3 10h-.5a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1zM6 9.972a.028.028 0 0 1-.028.028h-.944A.028.028 0 0 1 5 9.972V9a.5.5 0 1 1 1 0v.972z" fill="currentColor"/>`),
		g.Group(children),
	)
}