package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.452 11.16l3.694 3.694a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708l1.946 1.945a3.25 3.25 0 0 0-.134 4.732l4.707 4.708a.5.5 0 0 0 .707 0l2.08-2.08Zm-.707-.708L8.02 12.178L3.665 7.824A2.25 2.25 0 0 1 3.8 4.508l5.944 5.944Zm2.601-2.599L11.16 9.04l.707.707l1.188-1.188a3.25 3.25 0 0 0-.012-4.593a3.252 3.252 0 0 0-4.601-.012l-.447.448l-.454-.453a3.257 3.257 0 0 0-2.428-.956L6.5 4.378c.118.08.23.172.335.277l.81.809a.5.5 0 0 0 .715-.009l.79-.795a2.252 2.252 0 0 1 3.186.012a2.25 2.25 0 0 1 .011 3.181Z"/>`),
		g.Group(children),
	)
}