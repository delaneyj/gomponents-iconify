package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restaurantmenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768.338 896h-64v19q0 53-29.5 84.5t-71.5 22.5l-502-99q-42-8-71.5-51t-29.5-96V128q0-53 37.5-90.5t90.5-37.5h640q53 0 90.5 37.5t37.5 90.5v640q0 53-37.5 90.5t-90.5 37.5zm-640-479q0 46 30 84q4 6 21.5 19t31 27.5t13.5 28.5v128q0 15 31.5 27.5t96.5 28.5q39 10 67.5 7.5t44.5-11.5t16-20V608q0-14 13.5-26.5t32-22.5t21.5-13q5-4 14.5-9t12-10.5t2.5-19.5q0-45-29-83t-69-46q-25-5-47 5q-30-51-79.5-61t-78.5 29q-23-18-47-23q-41-8-69.5 18t-28.5 71zm704-243q0-46-32.5-78t-77.5-32h-465l346 68q42 9 71.5 52t29.5 95v553h18q45 0 77.5-32.5t32.5-77.5V174z"/>`),
		g.Group(children),
	)
}