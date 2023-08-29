package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chromecast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 6c-1.103 0-2 .897-2 2v4h2V8h20v16h-8v2h8c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2H6zm-2 8v2c5.17 0 9.436 3.942 9.95 8.979c.033.335.05.676.05 1.021h2c0-6.617-5.383-12-12-12zm0 4v2c3.309 0 6 2.691 6 6h2c0-4.411-3.589-8-8-8zm0 4v4h4a4 4 0 0 0-4-4z"/>`),
		g.Group(children),
	)
}