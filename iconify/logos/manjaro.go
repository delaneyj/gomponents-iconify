package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manjaro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#35BF5C" d="M256 0v256h-74.925V0H256Zm-90.54 90.536V256H90.535V90.536h74.925Zm0-90.536v74.925H74.67V256H0V0h165.46Z"/>`),
		g.Group(children),
	)
}