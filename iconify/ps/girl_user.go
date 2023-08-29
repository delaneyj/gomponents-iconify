package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 512h320v-64q0-72-52-124q18-4 26-10q26-18 26-58q0-68-43-134v-5q0-48-34.5-82.5T163 0T80 34.5T45 117v5Q3 193 3 256q0 36 25 58q10 6 26 10Q3 375 3 448v64zm85-395q0-9 4-21q58 46 145 53v11q0 31-22 53t-52 22q-31 0-53-22t-22-53v-43zm64 192v-32h21v32q0 11-10 11q-11 0-11-11zm119-30q-11 6-25 5t-24-5l-4 11q-2 0-4-2v-23q40-21 55-62q9 27 9 53q2 17-7 23zm-36-172q-77-6-121-45q22-19 49-19t48 18.5t24 45.5zM45 256q0-15 9-53q19 43 55 62v23q-2 0-4 2l-4-11q-30 13-49 0v-1q-1-1-2.5-3t-2.5-4t-1.5-6t-.5-9zm0 192q0-36 19.5-67.5T116 333q14 30 47 30q31 0 47-30q30 16 50 48t20 67v21H45v-21z"/>`),
		g.Group(children),
	)
}