package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKrwLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 130h-25.09l22.65-55.74a6 6 0 0 0-11.12-4.52L202 130h-44l-24.44-60.26a6 6 0 0 0-11.12 0L98 130H54L29.56 69.74a6 6 0 1 0-11.12 4.52L41.09 130H16a6 6 0 0 0 0 12h30l24.48 60.26a6 6 0 0 0 11.12 0L106 142h44l24.48 60.26a6 6 0 0 0 11.12 0L210 142h30a6 6 0 0 0 0-12ZM76 184.06L58.91 142h34.18ZM110.91 130L128 87.94L145.09 130ZM180 184.06L162.91 142h34.18Z"/>`),
		g.Group(children),
	)
}