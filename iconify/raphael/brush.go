package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.125 29.178l1.31-1.5l1.316 1.5l1.312-1.5l1.31 1.5l1.316-1.5l1.31 1.5l1.313-1.5l1.315 1.5l1.312-1.5l1.312 1.5l1.313-1.5l1.312 1.5v-8.52H8.125v8.52zm15.25-12.022c-.354 0-5.833-.166-5.833-2.917s.75-8.835.75-8.835s.25-2.583-2.292-2.583s-2.292 2.583-2.292 2.583s.75 6.083.75 8.834s-5.48 2.916-5.833 2.916c-.5 0-.5 1.166-.5 1.166v1.27h15.75v-1.27s0-1.166-.5-1.166zM16 8.03c-.62 0-1.125-2.19-1.125-2.81S15.38 4.093 16 4.093s1.125.504 1.125 1.125S16.62 8.03 16 8.03z"/>`),
		g.Group(children),
	)
}