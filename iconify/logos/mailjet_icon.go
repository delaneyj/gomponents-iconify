package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailjetIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#9585F4" d="m0 97.991l93.408 42.34l18.769-18.66l-47.795-21.715l148.187-56.744l-56.961 147.533l-21.606-47.359l-18.878 18.769l.982 2.183l41.357 90.68L256 0z"/>`),
		g.Group(children),
	)
}