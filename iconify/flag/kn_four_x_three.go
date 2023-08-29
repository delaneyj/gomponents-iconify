package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKn4x30"><path fill-opacity=".7" d="M-80.1 0h682.7v512H-80.1z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagKn4x30)" transform="translate(75.1) scale(.9375)"><path fill="#ffe900" d="M-107.8.2h737.6v511.3h-737.6z"/><path fill="#35a100" d="m-108.2.2l.8 368.6L466.6 0l-574.8.2z"/><path fill="#c70000" d="m630.7 511.5l-1.4-383.2l-579 383.5l580.4-.3z"/><path d="m-107.9 396.6l.5 115.4l125.3-.2l611.7-410.1L629 1.4L505.2.2l-613 396.4z"/><path fill="#fff" d="m380.4 156.6l-9.8-42.2l33.3 27l38-24.6l-17.4 41.3l33.4 27l-44.2-1.5l-17.3 41.3l-9.9-42.2l-44.1-1.5zm-275.2 179l-9.9-42.3l33.3 27l38-24.6l-17.4 41.3l33.4 27l-44.1-1.5l-17.4 41.3l-9.8-42.2l-44.1-1.5z"/></g>`),
		g.Group(children),
	)
}