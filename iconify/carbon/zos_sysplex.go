package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZosSysplex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 14h8c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2h-8c-1.103 0-2 .897-2 2v1H9c-1.103 0-2 .897-2 2v3H4c-1.103 0-2 .897-2 2v8c0 1.103.897 2 2 2h3v3c0 1.103.897 2 2 2h9v1c0 1.103.897 2 2 2h8c1.103 0 2-.897 2-2v-8c0-1.103-.897-2-2-2h-8c-1.103 0-2 .897-2 2v5H9v-3h3c1.103 0 2-.897 2-2v-8c0-1.103-.897-2-2-2H9V7h9v5c0 1.103.897 2 2 2Zm1.414 14L28 21.414V28h-6.586Zm5.172-8L20 26.586V20h6.586Zm-16.034-8L4 18.552V12h6.552ZM5.38 20L12 13.38l.001 6.62h-6.62Zm16.034-8L28 5.414V12h-6.586Zm5.172-8L20 10.586V4h6.586Z"/>`),
		g.Group(children),
	)
}