package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MvOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#d21034" d="M0 0h512v512H0z"/><path fill="#007e3a" d="M128 128h256v256H128z"/><circle cx="288" cy="256" r="85.3" fill="#fff"/><ellipse cx="308.6" cy="256" fill="#007e3a" rx="73.9" ry="85.3"/>`),
		g.Group(children),
	)
}