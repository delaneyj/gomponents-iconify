package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m189 395l263-290q13-13 13-30t-13-30l-34-34Q407 0 388 0q-20 0-30 11L189 183l-76-78Q97 85 83 85q-13 0-30 17l-34 37Q6 152 6 169.5T19 198zM83 139l106 106L386 41l34 36l-231 254L49 171z"/>`),
		g.Group(children),
	)
}