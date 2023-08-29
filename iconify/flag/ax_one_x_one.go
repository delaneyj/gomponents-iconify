package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagAx1x10"><path fill-opacity=".7" d="M166 0h850v850H166z"/></clipPath></defs><g clip-path="url(#flagAx1x10)" transform="translate(-100) scale(.6024)"><path fill="#0053a5" d="M0 0h1300v850H0z"/><g fill="#ffce00"><path d="M400 0h250v850H400z"/><path d="M0 300h1300v250H0z"/></g><g fill="#d21034"><path d="M475 0h100v850H475z"/><path d="M0 375h1300v100H0z"/></g></g>`),
		g.Group(children),
	)
}