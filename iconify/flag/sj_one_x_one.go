package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SjOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#ef2b2d" d="M0 0h512v512H0z"/><path fill="#fff" d="M128 0h128v512H128z"/><path fill="#fff" d="M0 192h512v128H0z"/><path fill="#002868" d="M160 0h64v512h-64z"/><path fill="#002868" d="M0 224h512v64H0z"/>`),
		g.Group(children),
	)
}