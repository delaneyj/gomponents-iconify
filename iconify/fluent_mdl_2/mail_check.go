package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1939 1363l90 90l-557 558l-269-270l90-90l179 178l467-466zm109-979v907q-32-32-65-63t-63-66V648l-896 447l-896-447v888h1099l-128 128H0V384h2048zM1024 953l881-441H143l881 441z"/>`),
		g.Group(children),
	)
}