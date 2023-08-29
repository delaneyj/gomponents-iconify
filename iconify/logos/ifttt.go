package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ifttt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M0 0h43.18v135.71H0V0Zm246.747 0h-74.024v43.18h30.843v92.53h43.181V43.18h30.843V0h-30.843Zm117.205 0h-74.024v43.18h30.843v92.53h43.18V43.18h30.844V0h-30.843Zm117.205 0h-74.024v43.18h30.843v92.53h43.18V43.18H512V0h-30.843ZM160.386 43.18V0h-98.7v135.71h43.181V98.7h37.013V55.52h-37.013V43.18h55.519Z"/>`),
		g.Group(children),
	)
}