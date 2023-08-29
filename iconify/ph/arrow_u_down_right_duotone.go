package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUDownRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 168l-48 48v-96Z" opacity=".2"/><path d="m229.66 162.34l-48-48A8 8 0 0 0 168 120v40H88a48 48 0 0 1 0-96h88a8 8 0 0 0 0-16H88a64 64 0 0 0 0 128h80v40a8 8 0 0 0 13.66 5.66l48-48a8 8 0 0 0 0-11.32ZM184 196.69v-57.38L212.69 168Z"/></g>`),
		g.Group(children),
	)
}