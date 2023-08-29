package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoAccessForServiceAnimalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l14.748 14.748M23.5 23.5l-8.252-8.252M19 24v-9.357a2.5 2.5 0 0 0 2.5-2.5v-.357h-3.786c-.986 0-1.785-.8-1.785-1.786v2.782c0 .882-.242 1.732-.68 2.466M9.5 20v1.75c0 .966.784 1.75 1.75 1.75h.393v-2.71a4.81 4.81 0 0 1 2.143-4.004c.073-.049.144-.1.214-.152M14 20h1.5v.292c0 .875 0 1.458.45 2.333c.001.002.45.875 1.05.875M5.5 24v-6.5a2 2 0 0 0-2-2v-8s1.5-1 4-1s4 1 4 1v4m-4.15-7s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C9.246 3.5 7.65 4.5 7.65 4.5h-.3Z"/>`),
		g.Group(children),
	)
}