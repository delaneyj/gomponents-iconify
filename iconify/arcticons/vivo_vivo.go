package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivoVivo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.877 12.069c-8.894 0-10.872 12.99 0 23.862c10.858-10.858 8.894-23.862 0-23.862Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.684 19.056c6.289-6.29 16.873 1.499 16.873 16.873C8.201 35.93.394 25.345 6.684 19.056Zm34.632 0c-6.289-6.29-16.873 1.499-16.873 16.873c15.356 0 23.163-10.584 16.873-16.873Z"/>`),
		g.Group(children),
	)
}