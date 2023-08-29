package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlearrowup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm237-623L559 211q-22-21-54-17h-3q-3 0-6 1l-1.5.5l-2.5.5l-2 1l-3 1q-15 5-23 13L274 401q-19 19-19 46t19 46t46 19t46-19l82-81v356q0 27 18.5 45.5t45 18.5t45-18.5T575 768V412l82 81q19 19 46 19t46-19t19-46t-19-46z"/>`),
		g.Group(children),
	)
}