package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloverSpiked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c-24 72-72 72-72 96s48 48 72 48s72-24 72-48s-48-24-72-96zM112 184c-24 0-24 48-96 72c72 24 72 72 96 72s48-48 48-72s-24-72-48-72zm288 0c-24 0-48 48-48 72s24 72 48 72s24-48 96-72c-72-24-72-72-96-72zm-141.906.03A72 72 0 0 0 184 256a72 72 0 0 0 144 0a72 72 0 0 0-69.906-71.97zM256 352c-24 0-72 24-72 48s48 24 72 96c24-72 72-72 72-96s-48-48-72-48z"/>`),
		g.Group(children),
	)
}