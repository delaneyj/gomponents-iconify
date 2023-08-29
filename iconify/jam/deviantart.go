package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.985 0H8.846L7.141 3.211a1.204 1.204 0 0 1-1.066.631H.033v5.1h3.27c.362 0 .596.373.43.686l-3.7 6.967v3.375h4.131l1.706-3.21a1.204 1.204 0 0 1 1.066-.632h6.049v-5.1H9.707a.471.471 0 0 1-.43-.686l3.708-6.98V0z"/>`),
		g.Group(children),
	)
}