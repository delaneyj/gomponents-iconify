package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#0768a9" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 0h170.7v512H0z"/><path fill="#fc0" d="M341.3 0H512v512H341.3z"/></g>`),
		g.Group(children),
	)
}