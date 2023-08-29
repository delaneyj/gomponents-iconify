package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagZa4x30"><path fill-opacity=".7" d="M-71.9 0h682.7v512H-71.9z"/></clipPath></defs><g clip-path="url(#flagZa4x30)" transform="translate(67.4) scale(.93748)"><g fill-rule="evenodd" stroke-width="1pt"><path d="M-71.9 407.8V104.4L154 256.1L-72 407.8z"/><path fill="#000c8a" d="m82.2 512.1l253.6-170.6H696V512H82.2z"/><path fill="#e1392d" d="M66 0h630v170.8H335.7S69.3-1.7 66 0z"/><path fill="#ffb915" d="M-71.9 64v40.4L154 256L-72 407.8v40.3l284.5-192L-72 64z"/><path fill="#007847" d="M-71.9 64V0h95l301.2 204h371.8v104.2H324.3L23 512h-94.9v-63.9l284.4-192L-71.8 64z"/><path fill="#fff" d="M23 0h59.2l253.6 170.7H696V204H324.3L23 .1zm0 512.1h59.2l253.6-170.6H696v-33.2H324.3L23 512z"/></g></g>`),
		g.Group(children),
	)
}