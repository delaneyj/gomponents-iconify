package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0Z"/><path fill="#d92223" d="M0 512V0h512v170.7z"/><path fill="#fcd955" d="M0 512V0h341.3z"/><path fill="#003d88" d="M0 512V0h170.7z"/><path fill="#007a39" d="m0 512l512-170.7V512Z"/>`),
		g.Group(children),
	)
}