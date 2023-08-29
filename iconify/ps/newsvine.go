package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newsvine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m144 299l55-40q8 5 17 5q14 0 23.5-9.5T249 232q0-14-9.5-23.5T216 199q-13 0-22.5 9.5T184 232v4l-40 29v-91l55-40q7 4 17 4q14 0 23.5-9.5T249 106q0-14-9.5-23.5T216 73q-13 0-22.5 9.5T184 106v4l-40 29V2h-32v72L72 45v-6q0-14-9.5-23.5T40 6q-14 0-23.5 9.5T7 39t9.5 23.5T40 72q7 0 15-4l57 41v91l-40-30v-5q0-14-9.5-23.5T40 132q-14 0-23.5 9.5T7 165t9.5 23t23.5 9q9 0 16-4l56 41v97l-40-29v-6q0-14-9-23t-23-9t-23.5 9T7 296t9.5 23.5T40 329q6 0 16-4l56 40v97h32v-32l55-39q8 5 17 5q14 0 23.5-10t9.5-23t-9.5-23t-23.5-10q-13 0-22.5 9.5T184 363v4l-40 29v-97z"/>`),
		g.Group(children),
	)
}