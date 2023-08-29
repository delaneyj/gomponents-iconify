package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.919 24.65l-.476 2.698a2.488 2.488 0 0 1-2.351 1.999h0a1.66 1.66 0 0 1-1.31-.586"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.5 21.352l-.581 3.298a2.488 2.488 0 0 1-2.351 1.999h0a1.631 1.631 0 0 1-1.647-1.999l.582-3.298M8.5 25.772a1.858 1.858 0 0 0 1.806.876h1.183a2.485 2.485 0 0 0 2.347-1.998h0a1.63 1.63 0 0 0-1.642-1.999h-1.308a1.63 1.63 0 0 1-1.641-1.998h0a2.485 2.485 0 0 1 2.346-1.999h1.184a1.858 1.858 0 0 1 1.806.876m5.533 5.12a2.488 2.488 0 0 1-2.351 1.999h0a1.63 1.63 0 0 1-1.646-1.999l.581-3.298"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.561 26.648a.652.652 0 0 1-.658-.8l.792-4.496m2.079 1.998a2.488 2.488 0 0 1 2.351-1.999h0a1.631 1.631 0 0 1 1.646 1.999l-.229 1.3a2.488 2.488 0 0 1-2.35 1.999h0a1.631 1.631 0 0 1-1.647-1.999m-.352 1.998l1.409-7.994m5.499 4.696a2.488 2.488 0 0 1 2.352-1.999h0a1.631 1.631 0 0 1 1.646 1.999l-.229 1.3a2.488 2.488 0 0 1-2.351 1.999h0a1.63 1.63 0 0 1-1.647-1.999m-.352 1.998l1.41-7.994"/>`),
		g.Group(children),
	)
}