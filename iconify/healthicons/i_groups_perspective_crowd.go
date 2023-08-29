package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IGroupsPerspectiveCrowd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M18 16.5a4.5 4.5 0 0 1-4.5 4.5A4.499 4.499 0 0 1 9 16.5c0-2.486 2.014-4.5 4.5-4.5s4.5 2.014 4.5 4.5ZM4 28.333C4 24.787 10.33 23 13.5 23s9.5 1.787 9.5 5.333V36H4v-7.667ZM39 16.5a4.5 4.5 0 0 1-4.5 4.5a4.499 4.499 0 0 1-4.5-4.5c0-2.486 2.014-4.5 4.5-4.5s4.5 2.014 4.5 4.5ZM27 15a3 3 0 1 1-6 0a3 3 0 1 1 6 0Zm-2 13.333C25 24.787 31.33 23 34.5 23s9.5 1.787 9.5 5.333V36H25v-7.667Z"/><path fill-rule="evenodd" d="M28.75 22.185c-.266.098-.53.202-.788.313c-1.17.5-2.353 1.176-3.272 2.08c-.246.243-.48.508-.69.797a6.48 6.48 0 0 0-.69-.797c-.919-.904-2.101-1.58-3.273-2.08a16.46 16.46 0 0 0-.788-.313C20.772 21.396 22.73 21 24 21c1.27 0 3.228.396 4.75 1.185Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}