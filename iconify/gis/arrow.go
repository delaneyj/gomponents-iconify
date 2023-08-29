package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M17.09 1.853a4.999 4.999 0 0 0-5.276 5.596l7.557 81.087c.483 3.938 5.137 5.773 8.176 3.223l15.947-12.932l7.15 12.385c4.112 7.122 10.636 8.872 17.758 4.76s8.87-10.638 4.758-17.76l-7.125-12.34l18.896-7.244c3.728-1.357 4.467-6.306 1.3-8.693L19.784 2.847a4.995 4.995 0 0 0-2.695-.994Z" color="currentColor"/>`),
		g.Group(children),
	)
}