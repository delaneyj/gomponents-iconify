package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swipeup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.9 1024h-384q-21 0-51.5-34.5t-53.5-81t-24-77.5q0-24-16-81.5T336.4 618t-15.5-138v-32q0-27 18.5-45.5t45.5-18.5q22 0 39 13t22 34q3-28 3-47q0-59 2-113t8-121.5T478.4 41t34.5-41q45 0 54.5 45t9.5 211q0 37 6.5 64t16 39t19 18t15.5 7h7q0-27 18.5-45.5t45.5-18.5q64 0 64 128q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v64q0-27 18.5-45.5t45.5-18.5t45.5 18.5t18.5 45.5v288q0 46-7 83t-18.5 59.5t-25 39.5t-27 24.5t-25 12t-18.5 5.5h-7zm-704-832v288q0 13-9.5 22.5t-23 9.5t-22.5-9.5t-9-22.5V192H8.9q-13-14-5-24l145-163q5-5 12-5t12 5l144 163q9 10-4 24h-120z"/>`),
		g.Group(children),
	)
}