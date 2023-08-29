package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadLockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22h-5c-.5 0-1-.5-1-1v-4c0-.5.5-1 1-1v-1.5c0-1.4 1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5V16c.5 0 1 .5 1 1v4c0 .5-.5 1-1 1M5 18h9v2H5v-2m16-2v-1.5c0-.8-.7-1.5-1.5-1.5s-1.5.7-1.5 1.5V16h3M9 16v-6H5l7-7l7 7h-4v6H9m.83-8H11v6h2V8h1.17L12 5.83L9.83 8Z"/>`),
		g.Group(children),
	)
}