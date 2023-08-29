package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BsOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagBs1x10"><path fill-opacity=".7" d="M56.6 26.4H537v480.3H56.6z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagBs1x10)" transform="matrix(1.066 0 0 1.067 -60.4 -28.1)"><path fill="#fff" d="M990 506.2H9.4V27.6H990z"/><path fill="#ffe900" d="M990 370.6H9.4V169.2H990z"/><path fill="#08ced6" d="M990 506.2H9.4V346.7H990zm0-319H9.4V27.9H990z"/><path d="M9 25.9c2.1 0 392.3 237 392.3 237L7.8 505.3L9 25.9z"/></g>`),
		g.Group(children),
	)
}