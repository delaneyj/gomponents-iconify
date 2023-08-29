package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagJo4x30"><path fill-opacity=".7" d="M-117.8 0h682.6v512h-682.6z"/></clipPath></defs><g clip-path="url(#flagJo4x30)" transform="translate(110.5) scale(.9375)"><g fill-rule="evenodd" stroke-width="1pt"><path d="M-117.8 0h1024v170.7h-1024z"/><path fill="#fff" d="M-117.8 170.7h1024v170.6h-1024z"/><path fill="#090" d="M-117.8 341.3h1024V512h-1024z"/><path fill="red" d="m-117.8 512l512-256l-512-256v512z"/><path fill="#fff" d="m24.5 289l5.7-24.9H4.7l23-11l-15.9-19.9l23 11l5.6-24.8l5.7 24.9L69 233.2l-16 19.9l23 11H50.6l5.7 24.9l-15.9-20z"/></g></g>`),
		g.Group(children),
	)
}