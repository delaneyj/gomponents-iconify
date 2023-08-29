package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="560.87" height="176.25" x="-24.43" y="167.88" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="88.12" ry="88.12" transform="rotate(-45 256 256.002)"/><rect width="176" height="196" x="169.41" y="156.59" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="32" ry="32" transform="rotate(45 257.409 254.582)"/><circle cx="256" cy="208" r="16" fill="currentColor"/><circle cx="304" cy="256" r="16" fill="currentColor"/><circle cx="208" cy="256" r="16" fill="currentColor"/><circle cx="256" cy="304" r="16" fill="currentColor"/>`),
		g.Group(children),
	)
}