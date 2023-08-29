package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupWithStraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a.75.75 0 0 0-.75.75V7H13c-.991 0-1.25.507-1.493.987L11.5 8L11 9H8a1 1 0 0 0 0 2l1.5 18c.065.55.5 1 1 1h11c.5 0 .959-.444 1-1L24 11a1 1 0 1 0 0-2h-3l-.5-1c-.237-.491-.5-1-1.5-1h-2.25V3.5H20c.276 0 .5-.336.5-.75S20.276 2 20 2h-4Zm5.993 9l-.571 6.859H10.578L10.008 11h11.986Z"/>`),
		g.Group(children),
	)
}