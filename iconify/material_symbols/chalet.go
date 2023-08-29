package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chalet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 11V9.8l-.75.75l-.7-.7L17.5 8.4v-.9h-.9l-1.45 1.45l-.7-.7l.75-.75H14v-1h1.2l-.75-.75l.7-.725l1.45 1.45h.9V5.6l-1.45-1.45l.7-.7l.75.75V3h1v1.2l.75-.75l.7.7L18.5 5.6v.9h.9l1.45-1.45l.7.7l-.75.75H22v1h-1.2l.75.75l-.7.7L19.4 7.5h-.9v.9l1.45 1.45l-.7.7l-.75-.75V11h-1ZM5 20v-4.7l-1.1 1.1L2.5 15L10 7.5l7.5 7.5l-1.4 1.425l-1.1-1.1V20h-4v-5H9v5H5Z"/>`),
		g.Group(children),
	)
}