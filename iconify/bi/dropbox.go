package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.01 4.555L4.005 7.11L8.01 9.665L4.005 12.22L0 9.651l4.005-2.555L0 4.555L4.005 2L8.01 4.555Zm-4.026 8.487l4.006-2.555l4.005 2.555l-4.005 2.555l-4.006-2.555Zm4.026-3.39l4.005-2.556L8.01 4.555L11.995 2L16 4.555L11.995 7.11L16 9.665l-4.005 2.555L8.01 9.651Z"/>`),
		g.Group(children),
	)
}