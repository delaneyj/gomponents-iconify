package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 1C8.268 1 2 7.268 2 15v14.06A1.94 1.94 0 0 0 3.94 31h6.12A1.94 1.94 0 0 0 12 29.06v-14c0-2.22 1.776-4.044 3.957-4.06h.004A4.006 4.006 0 0 1 20 15v14.06A1.94 1.94 0 0 0 21.94 31h6.12A1.94 1.94 0 0 0 30 29.06V15c0-7.732-6.268-14-14-14ZM4 23h6v6H4v-6Zm18 0h6v6h-6v-6Z"/>`),
		g.Group(children),
	)
}