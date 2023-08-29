package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Raphael(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M962 662L662 962q-62 62-150 62t-150-62L62 662Q0 600 0 512t62-150L362 62Q424 0 512 0t150 62l300 300q62 62 62 150t-62 150zm-187 37q-121 20-232-37T354 470q-51-87-41-193q-57 110-57 235t55 223.5T454 885q42 17 89 7.5t82-43.5zM512 128q-53 0-90.5 37.5T384 256t37.5 90.5T512 384t90.5-37.5T640 256t-37.5-90.5T512 128zM175 625l36 35q-19-70-19-148q0-74 19-149l-36 37q-47 46-47 112t47 113zm674-225L704 254v2q0 80-56 136t-136 56q-65 0-118-41q8 18 15 31q43 74 100.5 121.5t118 64.5t122 11T869 601q31-46 26-103.5T849 400z"/>`),
		g.Group(children),
	)
}