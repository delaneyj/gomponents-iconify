package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BgFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#d62612" d="M0 320h640v160H0z"/><path fill="#fff" d="M0 0h640v160H0z"/><path fill="#00966e" d="M0 160h640v160H0z"/></g>`),
		g.Group(children),
	)
}