package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageEmoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMessageEmoji0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 7H4v30h7v5l10-5h23V7Z"/><path d="M31 16v1m-14-1v1m14 8s-2 4-7 4s-7-4-7-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMessageEmoji0)"/>`),
		g.Group(children),
	)
}