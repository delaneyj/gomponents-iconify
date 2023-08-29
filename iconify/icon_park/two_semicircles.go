package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoSemicircles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"><path d="M44 25C44 13.9543 35.0457 5 24 5C12.9543 5 4 13.9543 4 25H44Z"/><path d="M14 32C14 37.5228 18.4772 42 24 42C29.5228 42 34 37.5228 34 32H14Z"/></g>`),
		g.Group(children),
	)
}