package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CfFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCf4x30"><path fill-opacity=".7" d="M-12.4 32h640v480h-640z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagCf4x30)" transform="translate(12.4 -32)"><path fill="#00f" d="M-52 32h719.3v119H-52z"/><path fill="#ff0" d="M-52 391.6h719.3V512H-52z"/><path fill="#009a00" d="M-52 271.3h719.3v120.3H-52z"/><path fill="#fff" d="M-52 151h719.3v120.3H-52z"/><path fill="red" d="M247.7 32.5h119.9V512H247.7z"/><path fill="#ff0" d="m99.3 137.7l-31.5-21.8l-31.3 22L47.4 101L16.9 78l38.2-1l12.5-36.3L80.3 77l38.1.7L88.2 101"/></g>`),
		g.Group(children),
	)
}