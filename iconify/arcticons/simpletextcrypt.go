package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simpletextcrypt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="34.27" cy="17.07" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.23" ry="9.12"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 33.6l6.52 6.45h.01l16.72-16.53m-16.2 8.58l3.76 3.72"/>`),
		g.Group(children),
	)
}