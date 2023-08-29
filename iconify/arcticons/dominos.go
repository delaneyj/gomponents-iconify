package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dominos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="17.47" height="35.91" x="15.27" y="6.04" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".9" transform="rotate(45 23.999 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.82 17.82l12.36 12.36"/><circle cx="30.55" cy="17.45" r="2.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.85" cy="30.51" r="2.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.12" cy="30.51" r="2.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}