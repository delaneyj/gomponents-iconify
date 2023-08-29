package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DjFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagDj4x30"><path fill-opacity=".7" d="M-40 0h682.7v512H-40z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagDj4x30)" transform="translate(37.5) scale(.94)"><path fill="#0c0" d="M-40 0h768v512H-40z"/><path fill="#69f" d="M-40 0h768v256H-40z"/><path fill="#fffefe" d="m-40 0l382.7 255.7L-40 511V0z"/><path fill="red" d="M119.8 292L89 270l-30.7 22.4L69.7 256l-30.6-22.5l37.9-.3l11.7-36.3l12 36.2h37.9l-30.5 22.7l11.7 36.4z"/></g>`),
		g.Group(children),
	)
}