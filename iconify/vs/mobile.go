package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M864 64q40 1 68 28q27 28 28 68v1344q-1 40-28 68q-28 27-68 28H96q-40-1-68-28q-27-28-28-68V160q1-40 28-68q28-27 68-28h768zm-432 96q-20 0-34 14t-14 34t14 34t34 14h96q20 0 34-14t14-34t-14-34t-34-14h-96zm96 1344q20 0 34-14t14-34t-14-34t-34-14h-96q-20 0-34 14t-14 34t14 34t34 14h96zm336-192V352H96v960h768zm458-842q150 158 150 362t-150 362q-20 18-45 18q-26 0-45-19q-20-20-20-44q0-25 20-45q112-118 112-272t-112-272q-20-20-20-45q0-24 20-44q19-19 45-19q25 0 45 18zm-180 180q74 80 74 182t-74 182q-20 18-46 18t-45-19t-19-45t18-46q38-38 38-90t-38-90q-18-19-18-46q0-26 19-45t45-19t46 18z"/>`),
		g.Group(children),
	)
}