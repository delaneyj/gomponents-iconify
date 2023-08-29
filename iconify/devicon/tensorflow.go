package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tensorflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#ff6f00" d="m61.55 128l-21.84-12.68V40.55L6.81 59.56l.08-28.32L61.55 0zM66.46 0v128l21.84-12.68V79.31l16.49 9.53l-.1-24.63l-16.39-9.36v-14.3l32.89 19.01l-.08-28.32z"/>`),
		g.Group(children),
	)
}