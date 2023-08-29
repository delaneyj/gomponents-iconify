package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriverAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="22" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M30 15h-2.09A6.005 6.005 0 0 0 22 10l-.022.001A9.983 9.983 0 0 0 4.051 15H2v2h2.05a9.983 9.983 0 0 0 17.928 4.999L22 22a6.005 6.005 0 0 0 5.91-5H30ZM14 8a7.977 7.977 0 0 1 5.738 2.446A6.015 6.015 0 0 0 16.09 15H6.07A8.007 8.007 0 0 1 14 8Zm0 16a8.007 8.007 0 0 1-7.93-7h10.02a6.015 6.015 0 0 0 3.649 4.554A7.977 7.977 0 0 1 14 24Zm8-4a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}