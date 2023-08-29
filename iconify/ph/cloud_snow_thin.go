package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 196a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm32 8a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm48-16a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm-96 40a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm88 0a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm72-136a72.08 72.08 0 0 1-72 72H76a48 48 0 1 1 11.51-94.61A72.08 72.08 0 0 1 228 92Zm-8 0a64.06 64.06 0 0 0-128-3.77a4 4 0 0 1-8-.46a71.63 71.63 0 0 1 1.42-10.65A40 40 0 1 0 76 156h80a64.07 64.07 0 0 0 64-64Z"/>`),
		g.Group(children),
	)
}