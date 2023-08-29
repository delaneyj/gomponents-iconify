package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m149 107l-85 85l85 85v64L0 192L149 43v64zm128 21q54 8 96.5 30.5T443 214t43.5 69.5T512 363q-78-109-235-109v87L128 192L277 43v85z"/>`),
		g.Group(children),
	)
}