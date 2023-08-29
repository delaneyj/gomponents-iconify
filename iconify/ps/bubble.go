package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M216 368q0-21-21-21h-86q-21 0-21 21t21 21h86q21 0 21-21zm-85 43q-22 0-22 21t22 21q0 22 21 22t21-22q22 0 22-21t-22-21h-42zm104-137q66-36 66-119q0-68-45-109T152 5T48 46T3 155q0 83 66 119q4 22 21.5 36.5T131 325h42q23 0 40.5-14.5T235 274zm-40-28v15q0 22-22 22h-42q-22 0-22-22v-15l-12-4q-2-1-5.5-2.5t-12.5-9T63 214t-12.5-25t-5.5-34q0-50 32-78.5T152 48t75 28.5t32 78.5q0 12-2 22.5t-5.5 18.5t-8 15t-9 11.5t-9.5 8.5t-8.5 6t-6.5 3l-3 2z"/>`),
		g.Group(children),
	)
}