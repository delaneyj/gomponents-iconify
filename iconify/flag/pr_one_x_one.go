package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPr1x10"><path fill-opacity=".7" d="M51.6 0h708.7v708.7H51.6z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagPr1x10)" transform="translate(-37.3) scale(.72249)"><path fill="#ed0000" d="M0 0h1063v708.7H0z"/><path fill="#fff" d="M0 141.7h1063v141.8H0zm0 283.5h1063v141.7H0z"/><path fill="#0050f0" d="m0 0l610 353.9L0 707.3V0z"/><path fill="#fff" d="m268.2 450.5l-65.7-49l-65.3 49.5l24.3-80.5l-65.3-49.7l80.7-.7l25-80.2l25.6 80l80.7.1l-64.9 50.2l24.9 80.3z"/></g>`),
		g.Group(children),
	)
}