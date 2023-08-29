package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Behance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 103H277V77h107v26zM208 203.5q12 17.5 12 42.5q0 20-8 35q-7 14-21 23q-12 9-30 14q-14 4-34 4H0V56h124q12 0 34 5q13 3 26 12q11 7 18 20q6 13 6 31q0 20-9.5 33.5T172 179q24 7 36 24.5zM55 163h61q17 0 26-6q10-7 10-23q0-9-3.5-15t-9.5-9q-6-4-12-5q-9-2-15-2H55v60zm107 80q0-20-11-29q-11-8-30-8H55v73h64q7 0 17-2q8-2 13.5-5.5T159 260q3-6 3-17zm264-3H289q0 24 13 37q12 11 34 11q15 0 27-8q12-9 14-18h46q-10 35-34 50q-24 16-55 16q-22 0-40-7t-31-21q-13-13-19-32q-7-18-7-40t7-40.5t20-32.5q13-13 30-21q18-8 40-8q24 0 42 9.5t30 25.5q11 15 17 37q5 21 3 42zm-52-34q-2-18-12-30q-9-10-29-10q-13 0-21 4.5T298.5 181t-6.5 13q-3 7-3 12h85z"/>`),
		g.Group(children),
	)
}