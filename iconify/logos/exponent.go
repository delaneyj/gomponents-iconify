package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M94.104-.001L0 217.125l42.193 22.519L128 55.703l85.57 183.941L256 217.125L161.896-.001H94.104z" fill="#023C69"/>`),
		g.Group(children),
	)
}