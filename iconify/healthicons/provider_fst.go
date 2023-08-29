package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProviderFst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 7a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7Zm25 25a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1v-9Zm-.609-3.023l-3.184-3.184l-1.414 1.414l3.184 3.184l1.414-1.414Zm-8.714-8.714l3.977 3.977l-1.414 1.415l-3.977-3.978v4.218h-2v-7.632h7.632v2h-4.218Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}