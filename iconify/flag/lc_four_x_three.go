package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LcFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#65cfff" d="M0 0h640v480H0z"/><path fill="#fff" d="m318.9 42l162.7 395.3l-322.6.9L318.9 42z"/><path d="m319 96.5l140.8 340l-279 .8L319 96.5z"/><path fill="#ffce00" d="m318.9 240.1l162.7 197.6l-322.6.5l159.9-198.1z"/></g>`),
		g.Group(children),
	)
}