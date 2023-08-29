package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTrayStackedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M48 336v96a48.14 48.14 0 0 0 48 48h320a48.14 48.14 0 0 0 48-48v-96"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M48 336h144m128 0h144m-272 0a64 64 0 0 0 128 0"/><path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M384 32H128c-26 0-43 14-48 40L48 192v96a48.14 48.14 0 0 0 48 48h320a48.14 48.14 0 0 0 48-48v-96L432 72c-5-27-23-40-48-40Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M48 192h144m128 0h144m-272 0a64 64 0 0 0 128 0"/>`),
		g.Group(children),
	)
}