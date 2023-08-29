package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SbFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagSb4x30"><path fill-opacity=".7" d="M0 0h682.7v512H0z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagSb4x30)" transform="scale(.9375)"><path fill="#0000d6" d="M0 507.2L987.4 0H0v507.2z"/><path fill="#006000" d="M1024 0L27.2 512H1024V0z"/><path fill="#fc0" d="M1024 0h-54.9L0 485.4V512h54.9L1024 27.6V0z"/><path fill="#fff" d="m71.4 9.1l11.8 34.5h38.5L90.6 64.7l11.9 34.4L71.4 78L40.3 99.2l11.9-34.4l-31.1-21.3h38.4zm191.1 0l11.9 34.5h38.5l-31.2 21.2l12 34.4L262.4 78l-31 21.3l11.9-34.4l-31.2-21.3h38.5zm0 144.5l11.9 34.5h38.5l-31.2 21.2l12 34.4l-31.2-21.3l-31 21.3l11.9-34.4l-31.2-21.3h38.5zm-95-71.4l11.9 34.4h38.4l-31 21.3l11.8 34.4l-31-21.3l-31.2 21.3l12-34.4l-31.2-21.3h38.5zm-96.1 71.4l11.8 34.5h38.5l-31.1 21.2l11.9 34.4l-31.1-21.3l-31.1 21.3l12-34.4L21 188h38.4z"/></g>`),
		g.Group(children),
	)
}