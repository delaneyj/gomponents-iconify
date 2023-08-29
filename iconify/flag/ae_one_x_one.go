package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AeOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00732f" d="M0 0h512v170.7H0z"/><path fill="#fff" d="M0 170.7h512v170.6H0z"/><path d="M0 341.3h512V512H0z"/><path fill="red" d="M0 0h180v512H0z"/>`),
		g.Group(children),
	)
}