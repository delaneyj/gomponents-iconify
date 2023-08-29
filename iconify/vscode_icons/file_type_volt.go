package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeVolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="gray" d="M12.068 2h6.866l-2.73 8.421l7.322-1.331L13.115 30l3.037-14.505L8.474 17l3.594-15z"/>`),
		g.Group(children),
	)
}