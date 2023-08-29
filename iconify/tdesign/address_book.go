package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 6a4 4 0 0 1 4-4h14v20H7a4 4 0 0 1-4-4V6Zm2 8.535A3.982 3.982 0 0 1 7 14h12V4h-2v6.766l-3.5-2.1l-3.5 2.1V4H7a2 2 0 0 0-2 2v8.535ZM19 16H7a2 2 0 1 0 0 4h12v-4ZM15 4h-3v3.234l1.5-.9l1.5.9V4Z"/>`),
		g.Group(children),
	)
}