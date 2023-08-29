package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagAx4x30"><path fill-opacity=".7" d="M106.3 0h1133.3v850H106.3z"/></clipPath></defs><g clip-path="url(#flagAx4x30)" transform="matrix(.56472 0 0 .56482 -60 -.1)"><path fill="#0053a5" d="M0 0h1300v850H0z"/><g fill="#ffce00"><path d="M400 0h250v850H400z"/><path d="M0 300h1300v250H0z"/></g><g fill="#d21034"><path d="M475 0h100v850H475z"/><path d="M0 375h1300v100H0z"/></g></g>`),
		g.Group(children),
	)
}