package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RootFolderOpened(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1 6.257V2.5l.5-.5h5l.35.15l.86.85h5.79l.5.5V6h1.13l.48.63l-2.63 7l-.48.37H8.743a5.48 5.48 0 0 0 .657-1h2.73l2.37-6H8.743a5.534 5.534 0 0 0-.72-.724l.127-.126L8.5 6H13V4H7.5l-.35-.15L6.29 3H2l.01 2.594c-.361.184-.7.407-1.01.663z" clip-rule="evenodd"/><path d="M6 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z"/><path fill-rule="evenodd" d="M8 10.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0zM4.5 13a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}