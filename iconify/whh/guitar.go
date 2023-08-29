package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guitar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M932.408 932q-85 85-199 91.5t-189-68.5q-26-27-41-46.5t-23-44.5q0-25-13-57.5t-37-57.5q-30-30-62.5-44.5t-55.5-10.5q-49-57-53.5-144.5t38.5-156.5l-141-141q-40 12-69-16l-67-67q-20-21-20-50t20-49l50-49q20-21 49-21t50 21l67 66q28 29 16 69l141 141q69-43 156.5-38.5t144.5 53.5q-4 23 10.5 55.5t44.5 62.5q25 24 57.5 37t57.5 13q25 8 44.5 23t46.5 41q75 75 68.5 189t-91.5 199zm-388-484q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm314 256l-154 154l38 38l154-154z"/>`),
		g.Group(children),
	)
}