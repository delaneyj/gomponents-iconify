package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagTl4x30"><path fill-opacity=".7" d="M0 0h682.7v512H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagTl4x30)" transform="scale(.9375)"><path fill="#cb000f" d="M0 0h1031.2v512H0z"/><path fill="#f8c00c" d="M0 0c3.2 0 512 256.7 512 256.7L0 512V0z"/><path d="M0 0c2.1 0 340.6 256.7 340.6 256.7L0 512V0z"/><path fill="#fff" d="M187.7 298.2L127 284.7l-31 52.8l-5-59.7l-60.7-13.3l54.9-24.9l-3.3-59.3l40.2 43.4l55.4-25.3l-28.9 54l39.2 45.8z"/></g>`),
		g.Group(children),
	)
}