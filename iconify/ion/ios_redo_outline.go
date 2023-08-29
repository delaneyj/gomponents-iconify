package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosRedoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M64 400h10.3l19.2-31.2c20.5-32.7 44.9-62.8 75.8-76.6 24.4-10.9 46.7-18.9 86.7-20V352l192-128L256 96v80.3c-63 2.8-108.1 20.7-143.3 56.2C60.4 285.2 64 351.5 64 368.2c.1 8.9 0 21.7 0 31.8zm208-208v-64.7l148.1 96.8L272 320.8V256c-91 0-144.6 24.6-192.2 105.4 0 0-8.8-169.4 192.2-169.4z" fill="currentColor"/>`),
		g.Group(children),
	)
}