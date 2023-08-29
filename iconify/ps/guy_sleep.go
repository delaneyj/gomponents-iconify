package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuySleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171zM64 213v69q-21-18-21-47q0-22 21-22zm192 256q-62 0-105.5-43.5T107 320V192q0-69 53-115q70 70 239 70q6 23 6 45v128q0 62-43.5 105.5T256 469zm192-187v-69q21 0 21 22q0 29-21 47zm-213 17q0 21 21 21t21-21v-22h-42v22zm21-86h-43q0 22-21 22q-5 0-8.5-1.5T178 231l-1-1q-6-6-6-17h-43q0 28 17 47q22 17 47 17q32 0 45-17q19-19 19-47zm64 22q-5 0-8.5-1.5T306 231l-1-1q-6-6-6-17h-43q0 28 17 47q17 17 47 17q32 0 45-17q19-19 19-47h-43q0 22-21 22zm-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15z"/>`),
		g.Group(children),
	)
}