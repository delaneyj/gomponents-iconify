package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><g opacity=".8"><path d="M5.25 3.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-7Z"/><path fill-rule="evenodd" d="M7.25 5v4h8V5h-8Zm-1.5-2a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-11Z" clip-rule="evenodd"/><path d="M5.25 10.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-7Z"/><path fill-rule="evenodd" d="M7.25 12v4h8v-4h-8Zm-1.5-2a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-11Zm-2.5-7A.75.75 0 0 1 4 2.25h14.5a.75.75 0 0 1 0 1.5H4A.75.75 0 0 1 3.25 3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m5.576 16.5l1.636-5.725l-1.924-.55l-1.908 6.682a1.25 1.25 0 0 0 1.2 1.593h13.34c.83 0 1.43-.795 1.201-1.593l-1.909-6.682l-1.922.55l1.635 5.725H5.576Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M3.5 3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V3Zm12 0h-11v7h11V3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.25 2.5a.5.5 0 0 1 .5-.5h14.5a.5.5 0 0 1 0 1H2.75a.5.5 0 0 1-.5-.5Zm2.24 7.598L3.11 17h13.78l-1.38-6.902l.98-.196l1.38 6.902A1 1 0 0 1 16.89 18H3.11a1 1 0 0 1-.98-1.196l1.38-6.902l.98.196Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}