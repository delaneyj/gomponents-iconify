package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xsplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M32 26c-.932-.136-3.333-.401-3.599-.401c-.136 0-3.735 3.068-5.333 4.532l-.136.135c.136-1.735.401-3.468.532-5.197c-7.865-.932-15.6-1.871-23.464-2.803V1.735h.401C1.197 1.864 30.53 5.198 32 5.334zm-3.469-3.469V8.266c-8.395-.932-16.932-2-25.333-2.932V19.47c8.401 1.061 16.939 2 25.333 3.061z"/>`),
		g.Group(children),
	)
}