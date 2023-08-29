package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CgFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCg4x30"><path fill-opacity=".7" d="M-79.5 32h640v480h-640z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagCg4x30)" transform="translate(79.5 -32)"><path fill="#ff0" d="M-119.5 32h720v480h-720z"/><path fill="#00ca00" d="M-119.5 32v480l480-480h-480z"/><path fill="red" d="M120.5 512h480V32l-480 480z"/></g>`),
		g.Group(children),
	)
}