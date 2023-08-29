package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudHyperProtectDbaas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24 30l-2.1-1c-1.7-.8-2.9-2.6-2.9-4.5V18h10v6.5c0 1.9-1.1 3.7-2.9 4.5L24 30zm-3-10v4.5c0 1.2.7 2.2 1.7 2.7l1.3.6l1.3-.6c1-.5 1.7-1.6 1.7-2.7V20h-6z"/><circle cx="10" cy="23" r="1" fill="currentColor"/><circle cx="10" cy="15" r="1" fill="currentColor"/><circle cx="10" cy="7" r="1" fill="currentColor"/><path fill="currentColor" d="M23 2H7c-1.1 0-2 .9-2 2v22c0 1.1.9 2 2 2h9v-2H7v-6h9v-2H7v-6h16v3h2V4c0-1.1-.9-2-2-2zM7 10V4h16v6H7z"/>`),
		g.Group(children),
	)
}