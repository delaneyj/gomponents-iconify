package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KwFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKw4x30"><path fill-opacity=".7" d="M0 0h682.7v512H0z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagKw4x30)" transform="scale(.9375)"><path fill="#fff" d="M0 170.6h1024v170.7H0z"/><path fill="#f31830" d="M0 341.3h1024V512H0z"/><path fill="#00d941" d="M0 0h1024v170.7H0z"/><path d="M0 0v512l255.4-170.7l.6-170.8L0 0z"/></g>`),
		g.Group(children),
	)
}