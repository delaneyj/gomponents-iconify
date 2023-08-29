package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M488.401 47.197H23.6V0h464.802v47.197zm0 45.763H23.6v47.198h464.802V92.96zM209.52 185.92H23.6v47.198h185.92v-47.197zm278.881 0h-185.92v47.198H488.4v-47.197zM209.52 278.88H23.6v47.198h185.92v-47.197zm278.881 0h-185.92v47.198H488.4v-47.197zm0 92.962H23.6v47.197h464.802v-47.197zm-278.881 92.96H23.6V512h185.92v-47.197zm278.881 0h-185.92V512H488.4v-47.197z"/>`),
		g.Group(children),
	)
}