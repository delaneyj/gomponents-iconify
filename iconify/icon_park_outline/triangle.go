package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22.27 6.99c.77-1.33 2.69-1.33 3.46 0l18.532 32.008c.772 1.333-.19 3.002-1.73 3.002H5.468c-1.54 0-2.503-1.669-1.731-3.002L22.269 6.99Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}