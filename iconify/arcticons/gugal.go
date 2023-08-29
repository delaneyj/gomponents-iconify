package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gugal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.52 9.527c5.377-5.37 14.086-5.37 19.463 0h0c5.077 5.079 5.4 13.207.74 18.672l.52.51h1.74l10.501 10.5l-3.28 3.291L28.703 32v-1.73l-.52-.52c-5.806 4.913-14.496 4.188-19.409-1.618c-4.612-5.451-4.293-13.525.737-18.594l.01-.01Zm3.051 3.06a9.451 9.451 0 1 0 13.371 0a9.46 9.46 0 0 0-13.37 0Zm17.162 15.622l-1.55 1.54"/>`),
		g.Group(children),
	)
}