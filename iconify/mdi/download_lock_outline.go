package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadLockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16a1.08 1.08 0 0 1 1 1v4a1.08 1.08 0 0 1-1 1h-5a1.08 1.08 0 0 1-1-1v-4a1.08 1.08 0 0 1 1-1v-1.5a2.5 2.5 0 0 1 5 0V16m-1 0v-1.5a1.5 1.5 0 0 0-3 0V16h3M13 5v6h1.17L12 13.17L9.83 11H11V5h2m2-2H9v6H5l7 7l7-7h-4V3m-1 15H5v2h9Z"/>`),
		g.Group(children),
	)
}