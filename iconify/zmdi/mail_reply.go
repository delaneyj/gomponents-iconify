package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 128q54 8 96.5 30.5T315 214t43.5 69.5T384 363q-78-109-235-109v87L0 192L149 43v85z"/>`),
		g.Group(children),
	)
}