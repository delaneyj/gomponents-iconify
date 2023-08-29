package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.5 24v-7.845c-1.306-.344-2.253-1.514-2.253-2.905s.947-2.562 2.232-2.901l.021-.005V0H24v10.344a3 3 0 0 1 .021 5.807l-.021.005v7.845zm-9.75 0V11.155c-1.306-.344-2.253-1.514-2.253-2.905s.947-2.562 2.232-2.901l.021-.005V0h1.5v5.344c1.306.344 2.253 1.514 2.253 2.905s-.947 2.562-2.232 2.901l-.021.005V24zm-10.5 0v-3.845C.944 19.811-.003 18.641-.003 17.25s.947-2.562 2.232-2.901l.021-.005V0h1.5v14.344c1.306.344 2.253 1.514 2.253 2.905s-.947 2.562-2.232 2.901l-.021.005V24z"/>`),
		g.Group(children),
	)
}