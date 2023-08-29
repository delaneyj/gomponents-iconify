package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 0a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8zm0 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm1.647-5.757c-.473.365-.734.566-.734 1.147v.814c0 .459-.406.834-.907.834c-.502 0-.909-.375-.909-.834V9.39c0-1.357.831-1.998 1.38-2.422c.474-.366.734-.566.734-1.146c0-.654-.541-1.188-1.205-1.188c-.665 0-1.205.533-1.205 1.188c0 .461-.408.833-.909.833c-.5 0-.909-.372-.909-.833c0-1.574 1.357-2.854 3.023-2.854c1.665 0 3.021 1.279 3.021 2.854c0 1.356-.83 1.996-1.38 2.421z"/>`),
		g.Group(children),
	)
}