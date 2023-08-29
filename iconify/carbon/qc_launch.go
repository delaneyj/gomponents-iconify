package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QcLaunch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M25 22h-6a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h6v2h-6v10h6z" fill="currentColor"/><path d="M13 8H9a2.002 2.002 0 0 0-2 2v10a2.002 2.002 0 0 0 2 2h1v2a2.002 2.002 0 0 0 2 2h2v-2h-2v-2h1a2.002 2.002 0 0 0 2-2V10a2.002 2.002 0 0 0-2-2zM9 20V10h4v10z" fill="currentColor"/><path d="M28 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v24a2.003 2.003 0 0 1-2 2zM4 4v24h24V4z" fill="currentColor"/>`),
		g.Group(children),
	)
}