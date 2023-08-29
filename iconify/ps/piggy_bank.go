package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiggyBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M477 131h-32q-7 0-10-7q-11-16-20-21l-2-2V3h-21q-27 0-62 27q-2 0-2 3q-17-9-64-9q-35 0-85 17t-75 43L61 71L44 52q-14-16-30 0q-14 14 0 30l26 25l36 13q-25 34-25 85v32l19 175q1 16 13.5 27.5T113 451h29q32 0 41-30l15-43q33 10 72 7l11 36q11 30 41 30h29q17 0 29.5-11.5T394 412l9-85q29-29 40-60q3-8 13-8h4q24 0 42-17.5t18-42.5v-26q0-17-12.5-29.5T477 131zm0 68q0 6-6.5 11.5T456 216q-36 0-53 36q-16 34-37 49l-6 5l-9 102h-29l-24-68l-17 2q-11 2-28 2q-29 0-64-11l-21-8l-26 83h-29L93 235v-30q0-42 26-74l17-20q49-44 128-44q39 0 53 8q18 10 41-13q2-1 6-4.5t5-3.5v64l6 6q5 6 11 11q9 9 10 11q17 27 47 27h34v26zm-128-68q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5z"/>`),
		g.Group(children),
	)
}