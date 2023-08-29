package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#0b7226" d="M0 0h213.3v480H0z"/><path fill="#ff0" d="M213.3 0h213.3v480H213.3z"/><path fill="#bc0000" d="M426.6 0H640v480H426.6z"/></g><path fill="#0b7226" d="M342 218.8h71.8l-56.6 43.6l20.7 69.3l-56.6-43.6l-56.6 41.6l20.7-67.3l-56.6-43.6h69.8l22.7-71.3z"/>`),
		g.Group(children),
	)
}