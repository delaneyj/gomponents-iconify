package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M491 219h-90q2-22 2-54l77-57q7-5 8-14.5T484 78q-4-7-13.5-8T454 74l-55 42q-11-48-51.5-79.5T256 5t-91.5 31.5T111 116L55 74q-6-5-15.5-4T26 78q-5 6-4 15.5t8 14.5l77 57q0 7 1 25t1 29H21q-21 0-21 21t21 21h86q4 0 8-2q1 5 11 60l-77 77q-14 14 0 30q6 6 15 6t15-6l60-60q38 109 117 109q77 0 115-109l60 60q6 6 15 6t15-6q14-16 0-30l-77-77q9-44 11-60q1 0 3.5 1t4.5 1h86q21 0 21-21t-19-21zM256 432q-56 0-87-115q-6-20-18-92q38-28 105-28q71 0 102 28q-9 85-34 146t-68 61zm107-256q-46-21-107-21t-107 21v-21q0-46 34-77l11 23q6 11 19 11q7 0 9-2q8-3 10.5-11.5T230 82l-12-25q15-9 38-9q21 0 36 6l-13 26q-7 19 9 28q4 4 11 4q13 0 19-13l11-23q34 34 34 79v21z"/>`),
		g.Group(children),
	)
}