package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M392 195q-74 0-145-152q-2-3-3.5-6t-3-7t-2.5-6Q225 3 200 3H93Q76 3 63.5 15.5T51 45q0 86-17 116q-46 56-15 162H8v64q0 17 12.5 29.5T51 429h426q18 0 30.5-12.5T520 387v-64q0-19-6.5-39.5t-20-41t-40-34T392 195zm21 44q27 6 43 26.5t18.5 34T477 323h-64v-84zM68 186q25-31 25-141h107q1 3 4.5 9t4.5 8q0 3 4 7q-1 0-7 4l-21 21q-14 16 0 30q6 7 15 7t15-7l17-17q2 5 8.5 15.5T249 137l-21 21q-16 16 0 30q6 7 15 7q8 0 15-7l17-17q15 19 34 34l-17 17q-16 16 0 30q6 7 15 7q8 0 15-7l21-21q2-2 2-4q10 4 26 8v88H63v-7q-5-15-8.5-36T53 230.5T68 186zm409 201H51v-22h426v22zm-320-86q28 0 46-18t18-46q0-27-18-45.5T157 173q-27 0-45.5 18.5T93 237q0 28 18.5 46t45.5 18zm0-85q10 0 16 6t6 15q0 22-22 22q-9 0-15-6t-6-16q0-9 6-15t15-6z"/>`),
		g.Group(children),
	)
}