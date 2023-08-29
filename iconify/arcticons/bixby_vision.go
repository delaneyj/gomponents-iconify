package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BixbyVision(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="1.501" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="8.392" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37.934c8.9 0 19.5-8.18 19.5-13.934S32.9 10.066 24 10.066m0 27.868c-8.9 0-19.5-8.18-19.5-13.934S15.1 10.066 24 10.066"/>`),
		g.Group(children),
	)
}