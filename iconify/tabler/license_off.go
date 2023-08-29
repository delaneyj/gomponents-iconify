package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LicenseOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 21H6a3 3 0 0 1-3-3v-1h10v2a2 2 0 1 0 4 0v-2m0-4V5a2 2 0 1 1 2 2h-2m2-4H8a3 3 0 0 0-.864.126M5.122 5.151A3 3 0 0 0 5 6v11m6-10h2m-4 4h2M3 3l18 18"/>`),
		g.Group(children),
	)
}