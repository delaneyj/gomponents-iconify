package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameEmoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGameEmoji0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 30H10a6 6 0 0 0 0 12h28a6 6 0 0 0 0-12Zm-2-8a8 8 0 1 0 0-16a8 8 0 0 0 0 16ZM4 14l9-9l9 9l-9 9l-9-9Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGameEmoji0)"/>`),
		g.Group(children),
	)
}