package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m306 106l-26-40C100 187 0 334 0 455c0 117 86 172 159 172c92 0 157-78 157-160c0-69-44-128-103-150c-17-6-33-11-33-40c0-37 27-92 126-171zm397 0l-26-40C499 187 397 334 397 455c0 117 88 172 161 172c93 0 159-78 159-160c0-69-45-128-106-150c-17-6-32-11-32-40c0-37 28-92 124-171z"/>`),
		g.Group(children),
	)
}