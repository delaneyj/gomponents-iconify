package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#e70011" d="M0 0h640v240H0Z"/><path fill="#fff" d="M0 240h640v240H0Z"/>`),
		g.Group(children),
	)
}