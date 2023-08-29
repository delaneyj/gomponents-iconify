package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 160h-32v-58.66a71.89 71.89 0 0 0 29 21.42a8 8 0 0 0 6-14.83A55.78 55.78 0 0 1 200 56a8 8 0 0 0-16 0a56 56 0 0 1-112 0a8 8 0 0 0-16 0a55.78 55.78 0 0 1-35 51.93a8 8 0 0 0 6 14.83a71.89 71.89 0 0 0 29-21.42V160H24a8 8 0 0 0 0 16h32v24a8 8 0 0 0 16 0v-24h112v24a8 8 0 0 0 16 0v-24h32a8 8 0 0 0 0-16Zm-88-33.8V160h-32v-33.8a72 72 0 0 0 32 0Zm-72-25a72.47 72.47 0 0 0 24 19.27V160H72Zm88 58.8v-39.52a72.47 72.47 0 0 0 24-19.27V160Z"/>`),
		g.Group(children),
	)
}