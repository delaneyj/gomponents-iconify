package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forrst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 462h162v-95l-49-34l12-17l37 25v-74h50v47l36-19l11 19l-47 25v26l71-35l9 19l-80 40v73h172L192 2z"/>`),
		g.Group(children),
	)
}