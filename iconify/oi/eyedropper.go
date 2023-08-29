package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyedropper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.31 0a.5.5 0 0 0-.19.84l.63.63L.16 5.13L0 5.29v2.72h2.69l.16-.16l3.66-3.66l.63.66a.5.5 0 1 0 .72-.69l-.94-.94l.66-.66c.59-.58.59-1.54 0-2.13c-.57-.57-1.56-.57-2.13 0l-.66.66l-.94-.94a.5.5 0 0 0-.47-.16a.5.5 0 0 0-.06 0zm1.16 2.19L5.78 3.5L2.62 6.66L1.34 5.35l3.13-3.16z"/>`),
		g.Group(children),
	)
}