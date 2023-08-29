package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JmOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path d="m0 0l256 256L0 512zm512 0L256 256l256 256z"/><path fill="#090" d="m0 0l256 256L512 0zm0 512l256-256l256 256z"/><path fill="#fc0" d="M512 0h-47.7L0 464.3V512h47.7L512 47.7z"/><path fill="#fc0" d="M0 0v47.7L464.3 512H512v-47.7L47.7 0z"/></g>`),
		g.Group(children),
	)
}