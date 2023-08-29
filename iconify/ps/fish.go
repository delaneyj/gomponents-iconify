package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M493 99q-20-14-41-2q-40 17-66 59q-28-54-87-77V45q0-17-15-34q-16-11-37-8q-56 12-74 66q-72 15-122.5 60T0 235q0 72 70 121.5T224 406q56 0 98.5-24t63.5-66q26 42 66 60q14 4 17 4q13 0 24-6q19-13 19-36V133q0-25-19-34zM224 65q5-12 32-20v20h-32zM43 235q0-30 21.5-57t57.5-45q27 33 27 102t-27 103q-36-18-57.5-45.5T43 235zm181 128q-24 0-62-8q30-51 30-120q0-68-30-119q24-9 62-9q62 0 100.5 35t38.5 93t-38.5 93T224 363zm245-19q-64-42-64-105q0-33 19.5-63.5T469 135v209zM128 214q0 9-6.5 15t-14.5 6q-9 0-15.5-6T85 214t6.5-15t15.5-6q8 0 14.5 6t6.5 15z"/>`),
		g.Group(children),
	)
}