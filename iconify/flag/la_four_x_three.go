package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagLa4x30"><path fill-opacity=".7" d="M0 0h640v480H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagLa4x30)"><path fill="#ce1126" d="M-40 0h720v480H-40z"/><path fill="#002868" d="M-40 119.3h720v241.4H-40z"/><path fill="#fff" d="M423.4 240a103.4 103.4 0 1 1-206.8 0a103.4 103.4 0 1 1 206.8 0z"/></g>`),
		g.Group(children),
	)
}