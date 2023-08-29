package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeAutoit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="12.551" fill="#5d83ac"/><path fill="#f0f0f0" d="M2 16a14 14 0 1 1 14 14A14 14 0 0 1 2 16ZM16 4.789A11.211 11.211 0 1 0 27.211 16A11.211 11.211 0 0 0 16 4.789Z"/><path fill="#f0f0f0" d="m24.576 20.156l-6.4-9.264a3.131 3.131 0 0 0-.819-.819a2.36 2.36 0 0 0-2.442.023a3.543 3.543 0 0 0-.812.8l-6.57 9.26h3.752l4.808-6.8l1.838 2.71q.26.368.544.789t.5.7q-.368-.031-.865-.031h-3.53l-1.914 2.634Z"/>`),
		g.Group(children),
	)
}