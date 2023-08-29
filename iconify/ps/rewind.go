package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M177 423q6 6 15 6q4 0 9-2q12-3 12-19V271l135 150q6 6 15 6q4 0 8-2q13-3 13-19V24q0-15-13-19q-14-3-23 6L213 161V24q0-15-12-19q-15-3-24 6L6 203q-11 17 0 28zM341 79v271L220 216zm-170 0v271L49 216z"/>`),
		g.Group(children),
	)
}