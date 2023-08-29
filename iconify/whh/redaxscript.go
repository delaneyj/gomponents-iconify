package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redaxscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640.81 1024v-64h32q32-33 32-64q0-19-25.5-54t-69-86.5t-65.5-83.5q-28-40-67-76q-34 36-61 76q-5 7-24.5 35t-33 48t-26 43.5t-12.5 33.5q0 16 8 32t16 24l8 8h32v64H.81v-64h32q9-7 25-20t56.5-56t72.5-91q24-36 49.5-90.5t39.5-90.5l13-36q-5-19-14.5-49t-36.5-97t-51-103q-20-30-43.5-55t-42-39t-35-24t-24.5-13.5t-9-3.5h-32V64h384v64h-32q-17 38 0 64q7 11 45 59.5t51 68.5q5 7 10 15t11 17t10 14q19-25 33-46q21-30 57-75t53.5-71.5t17.5-45.5q0-14-8-30t-16-25l-8-9h-32V0h384v64h-32q-9 7-25 20t-56.5 56t-72.5 91q-24 36-57.5 106.5T661.81 460l-21 52q144 200 198 281q67 101 151 165l3 2h32v64h-384z"/>`),
		g.Group(children),
	)
}