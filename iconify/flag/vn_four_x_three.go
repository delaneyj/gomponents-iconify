package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VnFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagVn4x30"><path fill-opacity=".7" d="M-85.3 0h682.6v512H-85.3z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagVn4x30)" transform="translate(80) scale(.9375)"><path fill="#da251d" d="M-128 0h768v512h-768z"/><path fill="#ff0" d="M349.6 381L260 314.3l-89 67.3L204 272l-89-67.7l110.1-1l34.2-109.4L294 203l110.1.1l-88.5 68.4l33.9 109.6z"/></g>`),
		g.Group(children),
	)
}