package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#c10000" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 0h218.3v175H0z"/><g fill="#c10000"><path d="M89.8 27.3h34.8v121.9H89.8z"/><path d="M168.2 70.8v34.8H46.3V70.8z"/></g></g>`),
		g.Group(children),
	)
}