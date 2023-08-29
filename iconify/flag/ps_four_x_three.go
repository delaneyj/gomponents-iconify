package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPs4x30"><path fill-opacity=".7" d="M-118 0h682.7v512H-118z"/></clipPath></defs><g clip-path="url(#flagPs4x30)" transform="translate(110.6) scale(.9375)"><g fill-rule="evenodd" stroke-width="1pt"><path d="M-246 0H778v170.7H-246z"/><path fill="#fff" d="M-246 170.7H778v170.6H-246z"/><path fill="#090" d="M-246 341.3H778V512H-246z"/><path fill="red" d="m-246 512l512-256L-246 0v512z"/></g></g>`),
		g.Group(children),
	)
}