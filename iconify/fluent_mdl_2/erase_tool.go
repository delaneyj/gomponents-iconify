package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraseTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1115 1792h421v128H453L50 1516q-24-24-37-56t-13-68q0-35 13-67t38-58L1248 69l794 795l-927 928zm133-1542L538 960l614 613l709-709l-613-614zM933 1792l128-128l-613-614l-306 307q-14 14-14 35t14 35l364 365h427z"/>`),
		g.Group(children),
	)
}