package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumericAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1314 1313q0-63-44-116t-103-53q-52 0-83 37t-31 94t36.5 95t104.5 38q50 0 85-27t35-68zm-610 127q0 12-10 24l-319 319q-10 9-23 9q-12 0-23-9L9 1463q-15-16-7-35q8-20 30-20h192V32q0-14 9-23t23-9h192q14 0 23 9t9 23v1376h192q14 0 23 9t9 23zm750-69q0 62-13 121.5t-41 114t-68 95.5t-98.5 65.5T1106 1792q-62 0-108-16q-24-8-42-15l39-113q15 7 31 11q37 13 75 13q84 0 134.5-58.5T1302 1468h-2q-21 23-61.5 37t-84.5 14q-106 0-173-71.5T914 1275q0-105 72-178t181-73q123 0 205 94.5t82 252.5zm-30-717v114H955V654h167V222q0-7 .5-19t.5-17v-16h-2l-7 12q-8 13-26 31l-62 58l-82-86L1136 0h123v654h165z"/>`),
		g.Group(children),
	)
}