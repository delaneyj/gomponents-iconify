package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlAngel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 192q0-56-32-107h32q21 0 21-21t-21-21h-73Q322 0 256 0T137 43H64q-21 0-21 21t21 21h32q-32 51-32 107v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222v-11zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469zm-21-170q0 21 21 21t21-21v-22h-42v22zm21-86h-43q0 22-21 22q-5 0-8.5-1.5T178 231l-1-1q-6-6-6-17h-43q0 28 17 47q22 17 47 17q32 0 45-17q19-19 19-47zm64 22q-5 0-8.5-1.5T306 231l-1-1q-6-6-6-17h-43q0 28 17 47q17 17 47 17q32 0 45-17q19-19 19-47h-43q0 22-21 22zm-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15z"/>`),
		g.Group(children),
	)
}