package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.969 8.969V7.011h-2.07a5.943 5.943 0 0 0-1.019-2.465l1.463-1.463l-1.414-1.414l-1.463 1.463a5.954 5.954 0 0 0-2.465-1.021V.042h-2v2.069a5.94 5.94 0 0 0-2.465 1.021L3.073 1.669L1.688 3.053l1.459 1.459a5.95 5.95 0 0 0-1.046 2.499H.032v1.958h2.063a5.943 5.943 0 0 0 1.026 2.507l-1.463 1.463l1.414 1.414l1.463-1.463c.72.513 1.557.867 2.465 1.021v2.069h2v-2.069a5.955 5.955 0 0 0 2.499-1.046l1.458 1.458l1.385-1.384l-1.463-1.463a5.95 5.95 0 0 0 1.026-2.507h2.064zm-7.948 2.093a3.084 3.084 0 0 1 0-6.166a3.084 3.084 0 0 1 0 6.166z"/>`),
		g.Group(children),
	)
}