package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1519 646q18 84-4 204q-87 444-565 444h-44q-25 0-44 16.5t-24 42.5l-4 19l-55 346l-2 15q-5 26-24.5 42.5T708 1792H457q-21 0-33-15t-9-36q9-56 26.5-168t26.5-168t27-167.5t27-167.5q5-37 43-37h131q133 2 236-21q175-39 287-144q102-95 155-246q24-70 35-133q1-6 2.5-7.5t3.5-1t6 3.5q79 59 98 162zm-172-282q0 107-46 236q-80 233-302 315q-113 40-252 42q0 1-90 1l-90-1q-100 0-118 96q-2 8-85 530q-1 10-12 10H57q-22 0-36.5-16.5T9 1538L241 67q5-29 27.5-48T320 0h598q34 0 97.5 13T1127 45q107 41 163.5 123t56.5 196z"/>`),
		g.Group(children),
	)
}