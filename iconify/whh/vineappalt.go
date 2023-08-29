package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vineappalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-128-512q-25 0-49.5-4t-51-15t-46.5-28.5t-32.5-46.5t-12.5-66q0-29 21-62.5t43-33.5q23 0 43.5 19.5t20.5 44.5q0 86-14 128h63q15-46 15-136q0-51-34.5-85.5t-93.5-34.5q-45 0-86.5 41t-41.5 99q0 50 13 94t32 70t38 45.5t32 26.5l13 8q0 17-28 56.5t-59.5 71.5t-40.5 32q-4 0-24-18t-48.5-55.5t-55-85t-45.5-116.5t-19-141h-64q0 101 24.5 190.5t59.5 146t73 98t64.5 59.5t34.5 18q7 0 36.5-22.5t64.5-56.5t63-83.5t28-93.5q40 3 64 0v-64z"/>`),
		g.Group(children),
	)
}