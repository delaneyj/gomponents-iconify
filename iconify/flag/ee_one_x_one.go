package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EeOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt" transform="scale(.482 .72)"><rect width="1063" height="708.7" rx="0" ry="0"/><rect width="1063" height="236.2" y="475.6" fill="#fff" rx="0" ry="0"/><path fill="#1791ff" d="M0 0h1063v236.2H0z"/></g>`),
		g.Group(children),
	)
}