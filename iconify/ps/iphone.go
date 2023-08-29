package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M237 512q28 0 46-18.5t18-45.5V64q0-27-18-45.5T237 0H67Q39 0 21 18.5T3 64v384q0 27 18 45.5T67 512h170zM67 43h170q22 0 22 21H45q0-21 22-21zm-22 64h214v277H45V107zm0 341v-21h214v21q0 21-22 21H67q-22 0-22-21zm128 0q0-9-6-15q-15-13-30 0q-6 6-6 15t6 15t15 6t15-6t6-15zM67 128h42v43H67v-43zm85 0h43v43h-43v-43z"/>`),
		g.Group(children),
	)
}