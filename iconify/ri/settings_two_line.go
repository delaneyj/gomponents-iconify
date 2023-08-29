package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.687 4l2.607-2.607a1 1 0 0 1 1.414 0L15.315 4H19a1 1 0 0 1 1 1v3.686l2.607 2.607a1 1 0 0 1 0 1.414L20 15.314V19a1 1 0 0 1-1 1h-3.686l-2.607 2.607a1 1 0 0 1-1.414 0L8.687 20H5.001a1 1 0 0 1-1-1v-3.686l-2.607-2.607a1 1 0 0 1 0-1.414l2.607-2.607V5a1 1 0 0 1 1-1h3.686ZM6.001 6v3.515L3.516 12L6 14.485V18h3.515L12 20.485L14.486 18h3.515v-3.515L20.486 12l-2.485-2.485V6h-3.515l-2.485-2.485L9.516 6H6Zm6 10a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}