package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Domino(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="35.91" height="17.47" x="6.052" y="15.266" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".9" ry=".9" transform="rotate(-45 24.007 24.001)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.82 17.82l12.36 12.36"/><circle cx="30.55" cy="17.45" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.55" cy="17.45" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.55" cy="17.45" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.82" cy="30.18" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.82" cy="30.18" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.82" cy="30.18" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}