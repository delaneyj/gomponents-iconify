package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(3)"><path d="M5.031 13.225v1.434c-2.154.083-3.936.613-3.936 1.26a.05.05 0 0 0 .004.017l8.806.009s.005-.018.005-.025c0-.646-1.818-1.177-3.973-1.26v-1.48"/><ellipse cx="5.012" cy="7.031" rx="5.012" ry="5.031"/><path d="M5.04 13.971a6.882 6.882 0 0 1-4.341-1.549a.499.499 0 0 1-.071-.697a.49.49 0 0 1 .69-.072a5.811 5.811 0 0 0 3.722 1.33c3.263 0 5.916-2.678 5.916-5.967c0-2.918-2.151-5.453-5.005-5.896a5.918 5.918 0 0 0-.911-.07a.493.493 0 0 1-.491-.494c0-.273.22-.495.491-.495c.352 0 .708.027 1.061.081c3.328.518 5.837 3.475 5.837 6.875c0 3.834-3.094 6.954-6.898 6.954z"/></g>`),
		g.Group(children),
	)
}