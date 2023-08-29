package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSignalCellularOneBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M4.41 22H20c1.1 0 2-.9 2-2V4.41c0-.89-1.08-1.34-1.71-.71L3.71 20.29c-.63.63-.19 1.71.7 1.71z"/><path fill="currentColor" d="m12 12l-8.29 8.29c-.63.63-.19 1.71.7 1.71H12V12z"/>`),
		g.Group(children),
	)
}