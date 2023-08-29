package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M718 1q10 0 17 6t6 17v510q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V24Q0 14 7 7t16-6h695zM602 418l-92-162l-93 162l-139-278l-139 278h463zm-46-185q20 0 33-14t13-33t-13-33t-33-13t-33 13t-13 32.5t13 33.5t33 14z"/>`),
		g.Group(children),
	)
}