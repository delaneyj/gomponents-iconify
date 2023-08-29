package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ajax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m758 460l-42 43q-9 9-21.5 9t-21.5-9l-33-33v362q0 80-56.5 136T448 1024H320q-80 0-136-56t-56-136h128q0 27 18.5 45.5T320 896h128q26 0 45-18.5t19-45.5V470l-33 33q-9 9-21.5 9t-21.5-9l-43-43q-9-9-9-21.5t9-21.5l152-152q9-9 31-9q21 0 31 10l152 151q8 9 8 21.5t-9 21.5zM448 128H320q-27 0-45.5 19T256 192v362l32-33q9-9 21.5-9t21.5 9l43 43q9 9 9 21.5t-9 21.5L223 759q-9 9-31 9t-32-10L9 607q-9-9-9-21.5T9 564l43-43q9-9 21.5-9t21.5 9l33 33V192q0-80 56-136T320 0h128q79 0 135.5 56T640 192H512q0-26-19-45t-45-19z"/>`),
		g.Group(children),
	)
}