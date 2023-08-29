package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrerenderIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#F0DA50" d="M64 0v64h32c17.673 0 32-14.327 32-32V0H64Z"/><path fill="#4BC69A" d="M32 0C14.327 0 0 14.327 0 32v128c0 17.673 14.327 32 32 32h32V0H32Z"/><path fill="#F0DA50" d="M160 128c-17.673 0-32 14.327-32 32v32h64v-64h-32Z"/><path fill="#5D7B8C" d="M224 0h-32v192h32c17.673 0 32-14.327 32-32V32c0-17.673-14.327-32-32-32Z"/>`),
		g.Group(children),
	)
}