package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartLife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 21.431L24 8.743L5.5 21.431v20.856h37V21.431zm-5.035-8.718a4.399 4.399 0 0 0-4.804-2.917l-.008.013m8.694 1.685a8.65 8.65 0 0 0-9.425-5.687"/>`),
		g.Group(children),
	)
}