package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.687 4l2.607-2.607a1 1 0 0 1 1.414 0L15.315 4H19a1 1 0 0 1 1 1v3.686l2.607 2.607a1 1 0 0 1 0 1.414L20 15.314V19a1 1 0 0 1-1 1h-3.686l-2.607 2.607a1 1 0 0 1-1.414 0L8.687 20H5.001a1 1 0 0 1-1-1v-3.686l-2.607-2.607a1 1 0 0 1 0-1.414l2.607-2.607V5a1 1 0 0 1 1-1h3.686Zm3.314 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}