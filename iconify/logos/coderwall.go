package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coderwall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="#3E8DCC"><circle cx="28.823" cy="28.823" r="28.051"/><circle cx="128" cy="28.823" r="28.051"/><circle cx="227.177" cy="28.823" r="28.051"/><circle cx="128" cy="128" r="28.051"/><circle cx="227.177" cy="128" r="28.051"/><circle cx="227.177" cy="227.177" r="28.051"/></g>`),
		g.Group(children),
	)
}