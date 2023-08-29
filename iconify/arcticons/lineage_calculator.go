package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageCalculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="26" x="4.5" y="11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(90 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 17.322h26"/><rect width="5.551" height="5.551" x="13.336" y="19.537" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="21.224" y="19.537" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="29.112" y="19.537" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="13.336" y="27.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="21.224" y="27.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="29.112" y="27.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="13.336" y="35.734" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="21.224" y="35.734" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><rect width="5.551" height="5.551" x="29.112" y="35.734" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/>`),
		g.Group(children),
	)
}