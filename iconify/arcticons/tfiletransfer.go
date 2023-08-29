package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tfiletransfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.533 5.5l10 10.867H22v11h-8.933v-11H7.533Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.867 19.9v11.4h5.6l-9.934 11.2l-10.266-11.2h5.466V19.9Z"/>`),
		g.Group(children),
	)
}