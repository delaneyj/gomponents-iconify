package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthWorker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 28.013c.503-.115 2 1.987 2 1.987h10s1.497-2.102 2-1.987c5.404 1.23 11 4.782 11 8.557V42H6v-5.43c0-3.775 5.596-7.327 11-8.557ZM32 34v-2h2v2h2v2h-2v2h-2v-2h-2v-2h2Zm-8-10a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2c5.523 0 10-4.477 10-10S29.523 6 24 6s-10 4.477-10 10s4.477 10 10 10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}