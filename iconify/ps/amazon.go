package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amazon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M274 114q-27 2-37 4q-34 4-55 14q-54 20-54 80q0 37 21 57t53 20q18 0 35-5q23-7 47-31q15 21 28 32q6 3 12-1q5-4 19-16t19-16q6-6 0-13q-18-24-18-46V87.5l-.5-9.5l-1.5-8.5l-2-9l-2.5-7.5l-4-8l-5-7l-6.5-7Q294 6 248 6h-11q-87 6-100 76q-2 8 7 10l48 6q8-2 8-9q7-27 37-31h4q17 0 27 13q6 9 6 36v7zm0 50q0 37-9 53q-8 17-28 22l-3 1h-4q-13 0-21.5-10t-8.5-26q0-36 37-47q16-3 37-3v10zm158 241q14-12 22-31.5t8-34.5v-8l-1-3q-4-4-22.5-5.5T399 326q-9 2-20 10q-6 5 1 7q5-1 19-3q34-2 40 5q7 8-13 57q-2 4 .5 5t5.5-2zM4 339q97 87 228 87q94 0 167-45q7-4 19-12q6-5 2.5-10t-9.5-2q-2 1-6 2t-6 2q-76 31-162 31q-123 0-228-60q-4-3-6.5 0t1.5 7z"/>`),
		g.Group(children),
	)
}