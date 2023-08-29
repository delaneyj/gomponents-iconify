package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M139.6 0H46.5C20.9 0 0 20.9 0 46.5v418.9C0 491.1 20.9 512 46.5 512h93.1c25.7 0 46.5-20.9 46.5-46.5v-419C186.2 20.9 165.3 0 139.6 0zm325.9 0h-93.1c-25.7 0-46.5 20.9-46.5 46.5v418.9c0 25.7 20.9 46.5 46.5 46.5h93.1c25.7 0 46.5-20.9 46.5-46.5V46.5C512 20.9 491.1 0 465.5 0z"/>`),
		g.Group(children),
	)
}