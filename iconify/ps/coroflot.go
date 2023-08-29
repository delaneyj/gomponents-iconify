package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coroflot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M164 303q0 8-6.5 14.5T143 324q-9 0-15-6.5t-6-14.5q0-9 6-15t15-6q8 0 14.5 6t6.5 15zm233-23q0 46-32.5 78.5T286 391h-3q-15 32-45 51.5T171 462q-50 0-86-35.5T49 340q0-40 23.5-72t59.5-44q1-61-18-99T67 89q-16 1-26.5 1.5t-20-1t-14-5.5T3 74q2-9 18.5-9.5T40 52q1-9 0-16t-2-12t2-10T52 4q17-8 23.5 7.5T80 53q30-8 60.5 3T195 91t38.5 42.5T260 171q24-3 26-3q46 0 78.5 33t32.5 79zm-237-62q4-1 11-1t21 2q17-28 51-42Q169 62 96 71q59 45 64 147zm108 122q0-32-18.5-57T202 249q-13-4-22-5h-9q-40 0-68 28t-28 68t28 68t68 28q57 0 84-49l9-21q4-16 4-26zm107-60q0-37-26-63.5T286 190q-44 0-71 35q35 14 57 45t22 70q0 17-4 29q36-1 60.5-27t24.5-62z"/>`),
		g.Group(children),
	)
}