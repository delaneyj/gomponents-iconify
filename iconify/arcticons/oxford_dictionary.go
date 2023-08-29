package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OxfordDictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a1.95 1.95 0 0 0 1.95 1.95h2.376v-39H10.35A1.95 1.95 0 0 0 8.4 6.45Zm4.326-1.95v39H37.65a1.95 1.95 0 0 0 1.95-1.95V6.45a1.95 1.95 0 0 0-1.95-1.95Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.663 18h17v12h-17zm8.5 12V18m-8.5 6h17m-17-6l17 12"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.663 18l-6.634 4.683L17.663 30"/>`),
		g.Group(children),
	)
}