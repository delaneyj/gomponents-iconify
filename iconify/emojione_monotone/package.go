package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2L2 19.887V44.08L32 62l30-17.92V19.92L32 2zm0 2.109l27.298 16.307l-11.073 6.613l-27.198-16.378L32 4.109zm-.909 55.239L3.818 43.057V21.998L31.09 38.287v21.061zM4.677 20.401l10.935-6.52l27.457 16.228L32 36.721L4.677 20.401zm55.505 22.656L49.12 49.664v-7.293l-5.153 3.078v7.293l-11.058 6.605v-21.06l11.058-6.593v7.277l5.153-3.077v-7.272l11.062-6.595v21.03z"/><path fill="currentColor" d="m26.803 8.47l6.429 3.927l5.424-3.235l-6.428-3.927zm32.885 14.575l-5.153 3.072v8.333l5.153-3.077zM33.71 57.932l7.031-4.2V42.371l-7.031 4.192z"/>`),
		g.Group(children),
	)
}