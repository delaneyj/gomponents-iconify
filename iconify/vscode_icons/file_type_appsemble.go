package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeAppsemble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#a6cfff" d="M8 3h1.607a5 5 0 0 1 5 5v6.607H3V8a5 5 0 0 1 5-5Z"/><path fill="#4a90e2" d="M22.393 3H29v11.607h-6.607a5 5 0 0 1-5-5V8a5 5 0 0 1 5-5z"/><path fill="#a6cfff" d="M14.607 29H3V17.393h11.607z"/><path fill="#4a90e2" d="M24 29h-6.607V17.393H24a5 5 0 0 1 5 5V24a5 5 0 0 1-5 5z"/>`),
		g.Group(children),
	)
}