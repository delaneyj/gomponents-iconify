package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M10 402q19 27 53 27h288q32 0 53-27q21-32 5-67l-75-162l-71-132Q247 7 208 7q-40 0-55 34L82 173L5 333q-13 39 5 69zm36-47l75-160l70-133q6-12 19-12t19 12l69 133l76 162q9 14-2 23q-9 9-17 9H63q-8 0-17-9q-6-16 0-25zm162-75q21 0 21-21v-86q0-21-21-21t-21 21v86q0 21 21 21zm21 43q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5z"/>`),
		g.Group(children),
	)
}