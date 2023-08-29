package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M238 297q88 45 118 116q-51 41-60 49q-43-82-115-101q-60 77-149 101q-15-31-32-71q43-10 65-27.5t44-59.5q-2-53 12.5-96.5T167 145q13 4 40 18t29 15q-2 13 22 18q43 10 76-20l34 57q-7 7-18.5 16T304 270t-65 4q-1 9-1 23zm8-164q-27 0-46.5-19.5T180 67t19.5-46T246 2t46.5 19T312 67t-19.5 46.5T246 133z"/>`),
		g.Group(children),
	)
}