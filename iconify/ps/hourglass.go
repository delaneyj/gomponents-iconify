package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 363v106H27q-22 0-22 22q0 9 6 15t16 6h298q10 0 16-6t6-15q0-22-22-22h-21V363q0-69-58-107q58-38 58-107V43h21q22 0 22-22q0-9-6-15t-16-6H27Q17 0 11 6T5 21q0 22 22 22h21v106q0 69 58 107q-58 38-58 107zm43-214V43h170v106q0 35-25 60.5T176 235t-60-25.5T91 149zm0 214q0-35 25-60.5t60-25.5t60 25.5t25 60.5v106H91V363zm149-256l-128 42v22q1 5 4.5 12t19.5 18.5t40 11.5t40-10.5t20-20.5l4-11v-64zM112 448h128v-85l-128 21v64z"/>`),
		g.Group(children),
	)
}