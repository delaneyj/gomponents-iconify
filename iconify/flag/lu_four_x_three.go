package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00a1de" d="M0 240h640v240H0z"/><path fill="#ed2939" d="M0 0h640v240H0z"/><path fill="#fff" d="M0 160h640v160H0z"/>`),
		g.Group(children),
	)
}