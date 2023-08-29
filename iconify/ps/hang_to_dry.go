package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangToDry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M423 21Q412 0 387 0H45Q20 0 9 21H3v448q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V21h-6zm-41 22q-13 37-59.5 61T216 128q-61 0-108-24T50 43h332zM45 469V111q65 60 171 60t171-60v358H45z"/>`),
		g.Group(children),
	)
}