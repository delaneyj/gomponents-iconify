package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.861 22.317l3.041-12.48s.374-4.126 5.275-4.251c4.673-.12 8.43-.08 8.43-.08l-2.36 9.25c-.037.437-1.073 1.502-2.414 1.38c-1.497-.134-2.312-.661-2.207-1.65l.395-2.37c-.02 0 .097-.641-.518-.832c-.775-.24-1.12.716-1.12.716l-3.647 17.62l-9.75.077l-4.438-18.388l-5.668.012l-1.487-5.812h12.413l4.055 16.808ZM12.774 42.5v-9.815h3.19c1.84 0 3.313 1.473 3.313 3.313s-1.473 3.312-3.313 3.312h-3.19m22.267-6.625l-3.189 4.908l-3.19-4.908m3.19 9.815v-4.907m-5.644 1.594h-4.416M20.688 42.5L24 32.685l3.312 9.815"/>`),
		g.Group(children),
	)
}