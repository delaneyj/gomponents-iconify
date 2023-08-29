package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSettingsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12h-2V5H4v13.385L5.763 17H12v2H6.455L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v8Zm-7.855 7.071a4.004 4.004 0 0 1 0-2.142l-.975-.563l1-1.732l.976.563A3.996 3.996 0 0 1 17 14.127V13h2v1.126c.715.184 1.352.56 1.854 1.072l.976-.564l1 1.732l-.975.563a4.004 4.004 0 0 1 0 2.142l.975.563l-1 1.732l-.976-.564A4 4 0 0 1 19 21.874V23h-2v-1.126a3.996 3.996 0 0 1-1.854-1.072l-.976.564l-1-1.732l.975-.563ZM18 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}