package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyframes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.225 18.412A1.595 1.595 0 0 1 8 19c-.468 0-.914-.214-1.225-.588l-4.361-5.248a1.844 1.844 0 0 1 0-2.328l4.361-5.248A1.595 1.595 0 0 1 8 5c.468 0 .914.214 1.225.588l4.361 5.248a1.844 1.844 0 0 1 0 2.328l-4.361 5.248zM17 5l4.586 5.836a1.844 1.844 0 0 1 0 2.328L17 19"/><path d="m13 5l4.586 5.836a1.844 1.844 0 0 1 0 2.328L13 19"/></g>`),
		g.Group(children),
	)
}