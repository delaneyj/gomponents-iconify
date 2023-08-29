package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="red" d="M0 0h512v512H0z"/><g fill="#fff"><path d="M96 208h320v96H96z"/><path d="M208 96h96v320h-96z"/></g></g>`),
		g.Group(children),
	)
}