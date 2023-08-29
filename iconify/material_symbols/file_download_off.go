package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDownloadOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L17.15 20H6q-.825 0-1.413-.588T4 18v-3h2v3h9.15l-2.575-2.575L12 16l-5-5l.575-.575l-6.2-6.2L2.8 2.8l18.4 18.4l-1.425 1.425Zm-4.35-10.05L14 11.15l1.6-1.6L17 11l-1.575 1.575ZM13 10.15l-2-2V4h2v6.15Zm7 7l-2-2V15h2v2.15Z"/>`),
		g.Group(children),
	)
}