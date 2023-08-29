package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkedin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 274v170h-98V285q0-67-50-67q-37 0-51 36q-3 7-3 24v166h-99q1-269 0-297h99v42l-1 1h1v-1q32-49 89-49q51 0 82 34t31 100zM58 4Q33 4 17.5 18.5T2 55t15 37t39 15h1q25 0 40.5-15T113 55q-1-22-16-36.5T58 4zM7 444h99V147H7v297z"/>`),
		g.Group(children),
	)
}