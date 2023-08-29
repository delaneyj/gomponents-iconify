package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmoticonTongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2m-4.73 9c-.17-.29-.27-.64-.27-1c0-1.11.89-2 2-2a2 2 0 0 1 2 2c0 .36-.1.71-.27 1c-.34-.6-.99-1-1.73-1s-1.39.4-1.73 1M16 15h-1c0 2-.9 3-2 3s-2-1-2-3H8v-2h8v2m.73-4c-.34-.6-.99-1-1.73-1s-1.39.4-1.73 1c-.17-.29-.27-.64-.27-1c0-1.11.89-2 2-2a2 2 0 0 1 2 2c0 .36-.1.71-.27 1Z"/>`),
		g.Group(children),
	)
}