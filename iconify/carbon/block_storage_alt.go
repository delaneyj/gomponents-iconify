package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockStorageAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20h-2v2h2v6H4v-6h2v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="M15 20H8v-7h7zm-5-2h3v-3h-3zm14 2h-7v-7h7zm-5-2h3v-3h-3zm-4-7H8V4h7zm-5-2h3V6h-3zm14 2h-7V4h7zm-5-2h3V6h-3z"/>`),
		g.Group(children),
	)
}