package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Livejournal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M365 452h-1l-156-65q-2-2-2-3q-2-2 1-4q27-26 27-65q0-19-5-31v-1L115 130q-45 39-42 78q0 4-4 4q-4 2-4-3q-4-47 51-90q52-41 90-28q3 0 3 5q-2 3-6 2q-34-10-82 27l117 156q22 14 49.5 13t49.5-16q2-2 4 0q2 0 2 3l26 167q0 3-1 3q-1 1-3 1zm-148-70l143 60l-24-154q-48 28-96 3q3 12 3 24q0 36-26 67zM38 169q-4 0-4-4q-4-47 51-90q52-41 90-28q3 1 3 5q-2 3-6 3q-34-12-82 26q-51 41-48 83q0 4-4 4v1zm274 257q10-34 44-40l9 62zM130 20q-43 0-84.5 45T25 147l175 232l165 69l-28-181L162 36q-10-16-32-16zm0-16q30 0 46 22l174 231q3 3 3 7l28 182q1 9-6 15q-5 3-10 3q-4 0-6-1l-166-69q-5-2-6-5L12 157q0-1-.5-1t-.5-1q-20-34 5-77q18-31 51.5-52.5T130 4z"/>`),
		g.Group(children),
	)
}