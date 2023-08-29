package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviantartIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#45AD47" d="M256 75.206V.004h-75.203l-7.505 7.57l-35.485 67.624l-11.158 7.528H0V186h69.626l6.201 7.5L0 338.396V413.6h75.201l7.507-7.567l35.487-67.627l11.156-7.526H256V227.605h-69.626l-6.195-7.551L256 75.206"/>`),
		g.Group(children),
	)
}