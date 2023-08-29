package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagSo4x30"><path fill-opacity=".7" d="M-85.3 0h682.6v512H-85.3z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagSo4x30)" transform="translate(80) scale(.9375)"><path fill="#40a6ff" d="M-128 0h768v512h-768z"/><path fill="#fff" d="M336.5 381.2L254 327.7l-82.1 54l30.5-87.7l-82-54.2L222 239l31.4-87.5l32.1 87.3l101.4.1l-81.5 54.7l31.2 87.6z"/></g>`),
		g.Group(children),
	)
}