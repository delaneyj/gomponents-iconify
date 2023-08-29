package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GmOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="red" d="M0 0h512v170.7H0z"/><path fill="#fff" d="M0 170.7h512V199H0z"/><path fill="#009" d="M0 199.1h512V313H0z"/><path fill="#fff" d="M0 312.9h512v28.4H0z"/><path fill="#090" d="M0 341.3h512V512H0z"/></g>`),
		g.Group(children),
	)
}