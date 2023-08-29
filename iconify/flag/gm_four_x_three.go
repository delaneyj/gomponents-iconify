package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GmFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagGm4x30"><path fill-opacity=".7" d="M0-48h640v480H0z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagGm4x30)" transform="translate(0 48)"><path fill="red" d="M0-128h640V85.3H0z"/><path fill="#fff" d="M0 85.3h640V121H0z"/><path fill="#009" d="M0 120.9h640V263H0z"/><path fill="#fff" d="M0 263.1h640v35.6H0z"/><path fill="#090" d="M0 298.7h640V512H0z"/></g>`),
		g.Group(children),
	)
}