package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#E23237" d="M.1 45.522L125.908.697l129.196 44.028l-20.919 166.45l-108.277 59.966l-106.583-59.169L.1 45.522Z"/><path fill="#B52E31" d="M255.104 44.725L125.908.697v270.444l108.277-59.866l20.919-166.55Z"/><path fill="#FFF" d="M126.107 32.274L47.714 206.693l29.285-.498l15.739-39.347h70.325l17.233 39.845l27.99.498l-82.179-174.917Zm.2 55.882l26.496 55.383h-49.806l23.31-55.383Z"/>`),
		g.Group(children),
	)
}