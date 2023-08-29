package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSnapcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#82bfa1" fill-rule="evenodd" d="m18.105 8.876l6.173 2.667l-6.173 6ZM6.318 29L17.21 18.413L13.887 15.2ZM2 3l15.582 15.052V8.474Z"/><path fill="#fa6340" fill-rule="evenodd" d="M27.436 8.473h-8.972L30 13.457Z"/>`),
		g.Group(children),
	)
}