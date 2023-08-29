package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-7.31 7.225l-4.207-4.207a5.001 5.001 0 0 1-5.885-6.814l.318-.738l1.355.262L9.114 9.57l.459-.46L7.73 7.27l-.263-1.354l.739-.319a5.001 5.001 0 0 1 6.813 5.886l4.207 4.208l-3.535 3.535Zm.708-3.535l-3.71-3.71l.263-.62a3.001 3.001 0 0 0-2.46-4.157l1.91 1.909l-3.287 3.287l-1.91-1.91a3.001 3.001 0 0 0 4.157 2.462l.62-.264l3.71 3.71l.707-.707Z"/>`),
		g.Group(children),
	)
}