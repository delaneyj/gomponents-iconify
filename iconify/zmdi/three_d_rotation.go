package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m159 458l29-28l81 81l-14 1q-100 0-173.5-68T0 277h32q6 60 40 108t87 73zm19-139q14 0 21-7t7-20q0-7-2-12t-6-8q-4-4-9.5-5.5T175 265h-16v-22h16q8 0 13-2t8-5q4-3 6-8t2-10q0-12-7-19q-6-6-19-6q-5 0-10 2q-4 1-8 4q-3 3-5 8q-2 4-2 9h-28q0-10 4-18t11-14t17-10q9-3 21-3q11 0 22 3q10 3 16 9q7 6 11 15t4 20q0 5-2 10q-1 5-4 10q-4 5-8 9q-5 4-11 7q7 3 13 7q5 4 8 9q3 4 5 11q2 5 2 12q0 11-5 20q-4 9-11.5 15.5T200 338t-22 3q-11 0-21-3q-9-3-17-9t-12-14.5t-4-20.5h27q0 6 2 10.5t6 7.5q3 3 8 5t11 2zm182.5-126.5Q371 203 377 218q5 16 5 34v8q0 19-5 34q-6 15-16 25q-10 11-25 17q-14 5-32 5h-49V171h50q18 0 31.5 5.5t24 16zM352 260v-8q0-28-12-43q-12-14-35-14h-20v123h19q12 0 21-4t15-11q6-8 9-19t3-24zM255 0q100 0 173.5 68T510 234h-32q-6-59-40.5-107T351 54l-29 28l-81-81z"/>`),
		g.Group(children),
	)
}