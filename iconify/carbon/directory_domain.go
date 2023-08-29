package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectoryDomain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 26h-9.184A2.996 2.996 0 0 0 17 24.184V19h7a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2H8a2.002 2.002 0 0 0-2 2v13a2.002 2.002 0 0 0 2 2h7v5.184A2.996 2.996 0 0 0 13.184 26H4v2h9.184a2.982 2.982 0 0 0 5.632 0H28Zm-4-14H8V9h16Zm0-8v3H8V4ZM8 14h16v3H8Zm8 14a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}