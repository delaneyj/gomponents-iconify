package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DefaultRootFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#c09553" d="M27.5 5.5h-9.3l-2.1 4.2H4.4v16.8h25.2v-21Zm0 4.2h-8.2l1.1-2.1h7.1Z"/><path fill="#c09553" d="M19.735 31.25h-5.924l9.794-21.5h5.985l-9.855 21.5z"/><path fill="#ffeebe" d="M23.766 10H29.2l-9.625 21H14.2Z"/>`),
		g.Group(children),
	)
}