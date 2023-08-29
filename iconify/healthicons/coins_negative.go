package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinsNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCoinsNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm37.972 27.975A12.939 12.939 0 0 0 40 21c0-7.18-5.82-13-13-13c-2.567 0-4.96.744-6.975 2.027a16.953 16.953 0 0 0-3.954.698A14.958 14.958 0 0 1 27 6c8.284 0 15 6.716 15 15a14.96 14.96 0 0 1-4.725 10.929c.381-1.263.62-2.587.697-3.954ZM21 42c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15c-8.284 0-15 6.716-15 15c0 8.284 6.716 15 15 15Zm-3-18a2 2 0 0 1 2-2v4a2 2 0 0 1-2-2Zm4-5v1a4.001 4.001 0 0 1 3.772 2.667a1 1 0 1 1-1.885.666A2.001 2.001 0 0 0 22 22v4a4 4 0 0 1 0 8v1h-2v-1a4.001 4.001 0 0 1-3.772-2.667a1 1 0 1 1 1.885-.666A2.001 2.001 0 0 0 20 32v-4a4 4 0 0 1 0-8v-1h2Zm2 11a2 2 0 0 1-2 2v-4a2 2 0 0 1 2 2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCoinsNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}