package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiotechnicaconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.43 41.93L25.22 5.28a1.33 1.33 0 0 0-1.65-.71a1.25 1.25 0 0 0-.79.71L8.57 41.93a1.17 1.17 0 0 0 .82 1.51a1.29 1.29 0 0 0 .4.06h28.43a1.25 1.25 0 0 0 1.28-1.21a.92.92 0 0 0-.07-.36Zm-8.09-20.85l-8.7 22.42"/>`),
		g.Group(children),
	)
}