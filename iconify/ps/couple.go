package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Couple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M5 341v43h214v-43q0-38-24.5-67T133 237v-4q28-7 46-30.5t18-53.5v-42q0-35-25-60.5T112 21T52 46.5T27 107v42q0 30 18 53t46 31v4q-37 8-61.5 37T5 341zm171 0H48q0-27 18.5-45.5T112 277t45.5 18.5T176 341zM69 149v-42q0-18 13-30.5T112 64t30 12.5t13 30.5v42q0 18-13 30.5T112 192t-30-12.5T69 149zm192 192v43h214v-43q0-38-24.5-67T389 237v-4q28-7 46-30.5t18-53.5v-42q0-35-25-60.5T368 21t-60 25.5t-25 60.5v42q0 30 18 53t46 31v4q-37 8-61.5 37T261 341zm171 0H304q0-27 18.5-45.5T368 277t45.5 18.5T432 341zM325 149v-42q0-18 13-30.5T368 64t30 12.5t13 30.5v42q0 18-13 30.5T368 192t-30-12.5t-13-30.5z"/>`),
		g.Group(children),
	)
}