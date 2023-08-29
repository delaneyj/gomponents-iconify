package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EeFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><rect width="640" height="477.9" rx="0" ry="0"/><rect width="640" height="159.3" y="320.7" fill="#fff" rx="0" ry="0"/><path fill="#1791ff" d="M0 0h640v159.3H0z"/></g>`),
		g.Group(children),
	)
}