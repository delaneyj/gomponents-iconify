package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FmFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagFm4x30"><path fill-opacity=".7" d="M-81.3 0h682.6v512H-81.3z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagFm4x30)" transform="translate(76.3) scale(.94)"><path fill="#6797d6" d="M-252 0H772v512H-252z"/><path fill="#fff" d="m259.8 123l-32.4 22.2l12.4-35.9l-32.5-22.2h40.1l12.4-35.9l12.4 36h40l-32.4 22.1l12.4 35.9M259.8 390l-32.4-22.2l12.4 36l-32.5 22.1h40.1l12.4 35.9l12.4-36l40 .1l-32.4-22.2l12.4-35.9m-188.4-92.4L79.3 306l1.4-38l-37.5-11.7l38.4-11.7l1.3-38l22.3 30.8l38.4-11.8l-24.6 30.7l22.4 30.7m274.2-11.7l24.6 30.7l-1.4-38l37.5-11.7l-38.4-11.7l-1.3-38l-22.3 30.8l-38.4-11.8l24.6 30.7l-22.4 30.7"/></g>`),
		g.Group(children),
	)
}