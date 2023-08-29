package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Island(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 896q0 31-27 54t-71.5 37t-112 22.5t-139 11.5t-162.5 3q-110 0-200-8t-161.5-24.5t-111-45T0 880q0-42 51-87.5t139.5-83T384 654v-78q0-107-40-206q-18 19-46 58t-49 67.5t-53 53t-65 27.5q-3-14-3-29.5t1.5-36T131 479q0-65 30.5-123T243 256h-19q-69 0-118.5 53T32 448q-9-9-20.5-31.5T0 372q0-112 64.5-178T224 128q105 0 154 54q18-79 81.5-130.5T605 0q96 0 161.5 50T832 192q-73-64-192-64q-49 0-75.5 19.5T523 205q41-13 83-13q94 0 160 51t66 123q0 45-27.5 83.5T730 512q-17-85-75-133.5T513 322q10 20 21 50.5T560.5 481T576 640v1q105 6 211 48.5T958.5 789t65.5 107z"/>`),
		g.Group(children),
	)
}