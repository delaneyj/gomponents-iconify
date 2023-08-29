package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v9h-2V4h-17v14.296L6.124 16H13v2H6.876L1.5 22.704V2ZM20 12.5v1.14a3.496 3.496 0 0 1 1.405.814l.99-.571l1 1.732l-.99.571a3.51 3.51 0 0 1 0 1.623l.99.572l-1 1.732l-.993-.573a3.496 3.496 0 0 1-1.403.81v1.145h-2V20.35a3.496 3.496 0 0 1-1.403-.81l-.992.573l-1-1.732l.99-.572a3.506 3.506 0 0 1 0-1.623l-.99-.571l1-1.732l.989.57a3.497 3.497 0 0 1 1.406-.813V12.5h2Zm-1 2.995a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}