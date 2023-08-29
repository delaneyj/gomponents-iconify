package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M69 512h342q26 0 44.5-17.5T475 450v-2q0-24-13-45L304 192V43h21q8 0 15-7t7-15t-7-14.5T325 0H155q-8 0-15 6.5T133 21t7 15t15 7h21v147L20 401Q7 422 7 444l-2 4q1 27 20 45.5T69 512zm107-250l11-12q24-34 32-77V43h42v128q9 47 32 76l11 13l43 58H133zM54 427l58-79v15h267l47 64q2 2 6 19q0 9-6.5 16t-14.5 7H69q-8 0-14.5-7T48 446q0-13 6-19z"/>`),
		g.Group(children),
	)
}