package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSettingsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14v8H4a8 8 0 0 1 8-8Zm0-1c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Zm2.595 5.811a3.505 3.505 0 0 1 0-1.622l-.992-.573l1-1.732l.992.573A3.498 3.498 0 0 1 17 14.645V13.5h2v1.145c.532.158 1.012.44 1.405.812l.992-.573l1 1.732l-.991.573a3.512 3.512 0 0 1 0 1.622l.991.573l-1 1.732l-.992-.573a3.495 3.495 0 0 1-1.405.812V22.5h-2v-1.145a3.495 3.495 0 0 1-1.405-.812l-.992.573l-1-1.732l.992-.573ZM18 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}