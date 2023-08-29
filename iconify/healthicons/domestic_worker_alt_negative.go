package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomesticWorkerAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsDomesticWorkerAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16.5 13a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm12.949 3.05A2.006 2.006 0 0 0 29 16h-9v12h-7V16h-1a6 6 0 0 0-1 11.917V40a2 2 0 1 0 4 0v-9h3v9a2 2 0 1 0 4 0V20h7c.477 0 .914-.167 1.258-.445l1.998 8.654a4 4 0 0 0-2.998 4.798l1.05 8.994l2.824-.652l-1.03-4.914l1.958-.41l1.021 4.874l6.92-1.597l-2.999-8.545a4 4 0 0 0-4.797-2.997L30.38 11.195a1 1 0 0 0-1.95.45l1.018 4.406ZM18 19v-3h-3v3h3Zm-7 1a2 2 0 1 0 0 4v-4Zm21.706 10.158l1.949-.45a2 2 0 0 1 2.398 1.5l-5.846 1.349a2 2 0 0 1 1.499-2.399Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDomesticWorkerAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}