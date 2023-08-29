package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatSettingsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.455 19L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455Zm1.69-6.929l-.975.563l1 1.732l.976-.563c.501.51 1.14.887 1.854 1.071V16h2v-1.126a3.996 3.996 0 0 0 1.854-1.072l.976.564l1-1.732l-.975-.563a4.004 4.004 0 0 0 0-2.142l.975-.563l-1-1.732l-.976.563A3.996 3.996 0 0 0 13 7.126V6h-2v1.126a3.996 3.996 0 0 0-1.854 1.071l-.976-.563l-1 1.732l.975.563a4.004 4.004 0 0 0 0 2.142ZM12 13a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}