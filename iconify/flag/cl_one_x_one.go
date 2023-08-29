package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCl1x10"><path fill-opacity=".7" d="M0 0h708.7v708.7H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagCl1x10)" transform="scale(.722)"><path fill="#fff" d="M354.3 0H1063v354.3H354.3z"/><path fill="#0039a6" d="M0 0h354.3v354.3H0z"/><path fill="#fff" d="m232.3 265.3l-55-41.1l-54.5 41.5l20.3-67.5l-54.5-41.7l67.4-.6l21-67.3l21.3 67.2h67.5L211.4 198l20.8 67.4z"/><path fill="#d52b1e" d="M0 354.3h1063v354.4H0z"/></g>`),
		g.Group(children),
	)
}