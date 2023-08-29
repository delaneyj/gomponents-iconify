package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudHyperProtectVs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 18v6.5c0 1.9 1.1 3.7 2.9 4.5l2.1 1l2.1-1c1.7-.8 2.9-2.6 2.9-4.5V18H20zm8 6.5c0 1.2-.7 2.2-1.7 2.7l-1.3.6l-1.3-.6c-1-.5-1.7-1.6-1.7-2.7V20h6v4.5z"/><path fill="currentColor" d="M25 2H5c-1.1 0-2 .9-2 2v22c0 1.1.9 2 2 2h11V4h9v11h2V4c0-1.1-.9-2-2-2zM13.6 12.2L5 23.9V5.1l8.6 7.1zM6 26l8-10.9V26H6zm8-16.1L6.8 4H14v5.9z"/>`),
		g.Group(children),
	)
}