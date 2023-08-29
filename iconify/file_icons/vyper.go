package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vyper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m313 134l-25-45h-64l-25.65 45L256 234l57-100zm71 122l-64-111l-57 100l64 110l57-99zM192 145l-64 111l57.65 99.85L249 245l-57-100zm64 111l-64 111l64 111l64-111l-64-111zM128 34H0l64 111l64-111zm320 111.15L512 34H384l64 111.15zM70 156l52 90l64-112l-52-90l-64 112zM378 44l-52 90l63.77 112l52.45-90.85L378 44z"/>`),
		g.Group(children),
	)
}