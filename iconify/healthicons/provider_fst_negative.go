package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProviderFstNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsProviderFstNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM7 6a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7Zm25 25a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1h-9Zm-4.793-5.207l3.184 3.184l-1.414 1.414l-3.184-3.184l1.414-1.414Zm-1.553-1.553l-3.977-3.977h4.218v-2h-7.632v7.632h2v-4.218l3.977 3.977l1.415-1.414Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsProviderFstNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}