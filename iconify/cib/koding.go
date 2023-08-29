package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Koding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1.068 0h29.599v8H1.334V0zm0 12H24v8H.932v-8zm0 12h29.599v8H1.334v-8z"/>`),
		g.Group(children),
	)
}