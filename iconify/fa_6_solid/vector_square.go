package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M368 80h32v32h-32V80zm-16-48c-17.7 0-32 14.3-32 32H128c0-17.7-14.3-32-32-32H32C14.3 32 0 46.3 0 64v64c0 17.7 14.3 32 32 32v192c-17.7 0-32 14.3-32 32v64c0 17.7 14.3 32 32 32h64c17.7 0 32-14.3 32-32h192c0 17.7 14.3 32 32 32h64c17.7 0 32-14.3 32-32v-64c0-17.7-14.3-32-32-32V160c17.7 0 32-14.3 32-32V64c0-17.7-14.3-32-32-32h-64zM96 160c17.7 0 32-14.3 32-32h192c0 17.7 14.3 32 32 32v192c-17.7 0-32 14.3-32 32H128c0-17.7-14.3-32-32-32V160zM48 400h32v32H48v-32zm320 32v-32h32v32h-32zM48 112V80h32v32H48z"/>`),
		g.Group(children),
	)
}