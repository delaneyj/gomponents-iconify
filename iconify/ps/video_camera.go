package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M45 177q0 2-1 7.5t-1 7.5v149q0 18 12.5 30.5T85 384h171q17 0 30-12.5t13-30.5v-8l55 28q2 2 9 2q8 0 10-2q11-8 11-20V192q0-13-11-19q-11-7-21 0l-53 28v-9q0-2-1-7.5t-1-7.5q44-28 44-81q0-40-28-68T245 0q-45 0-74 36Q142 0 96 0Q56 0 28 28T0 96q0 52 45 81zm296 49v81l-42-21v-39zm-85-34v149H85V192h171zM245 43q23 0 38.5 15.5T299 96t-15.5 37.5T245 149t-38-15.5T192 96t15-37.5T245 43zM96 43q22 0 37.5 15.5T149 96t-15.5 37.5T96 149t-37.5-15.5T43 96t15.5-37.5T96 43zm53 213q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15zm86 0q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15zm21-160q0 11-11 11q-10 0-10-11t10-11q11 0 11 11zm-149 0q0 11-11 11T85 96t11-11t11 11z"/>`),
		g.Group(children),
	)
}