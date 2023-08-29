package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrescriptionDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M27 15V4H10a1 1 0 0 0-1 1v38a1 1 0 0 0 1 1h28a1 1 0 0 0 1-1V16H28a1 1 0 0 1-1-1Zm2-1V4.586L38.414 14H29Zm-12 5a1 1 0 0 1 1-1h5a4 4 0 0 1 .395 7.98L27 29.587l3.293-3.293l1.414 1.414L28.414 31l3.293 3.293l-1.414 1.414L27 32.414l-3.293 3.293l-1.414-1.414L25.586 31l-5-5H19v7h-2V19Zm6 5h-4v-4h4a2 2 0 1 1 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}