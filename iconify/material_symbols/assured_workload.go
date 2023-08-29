package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssuredWorkload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17v-7h2v7H5Zm6 0v-7h2v7h-2ZM2 8V6l10-5l10 5v2H2Zm0 13v-2h12.05q.05.525.125 1.012T14.4 21H2Zm15-7.75V10h2v2.25l-2 1ZM20 24q-1.725-.425-2.863-1.988T16 18.55V16l4-2l4 2v2.55q0 1.9-1.137 3.463T20 24Zm-.725-3l3.475-3.45l-1.05-1.05l-2.425 2.375l-.975-.975l-1.05 1.075L19.275 21Z"/>`),
		g.Group(children),
	)
}