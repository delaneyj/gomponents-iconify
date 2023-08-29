package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPr4x30"><path fill-opacity=".7" d="M-37.3 0h682.7v512H-37.3z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagPr4x30)" transform="translate(35) scale(.9375)"><path fill="#ed0000" d="M-37.3 0h768v512h-768z"/><path fill="#fff" d="M-37.3 102.4h768v102.4h-768zm0 204.8h768v102.4h-768z"/><path fill="#0050f0" d="m-37.3 0l440.7 255.7L-37.3 511V0z"/><path fill="#fff" d="M156.4 325.5L109 290l-47.2 35.8l17.6-58.1l-47.2-36l58.3-.4l18.1-58l18.5 57.8l58.3.1l-46.9 36.3l18 58z"/></g>`),
		g.Group(children),
	)
}