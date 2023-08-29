package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.75 12a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M11.46 1.838a.75.75 0 0 1 1.08 0L15.111 4.5h3.638a.75.75 0 0 1 .75.75v3.638l2.662 2.573a.75.75 0 0 1 0 1.078L19.5 15.111v3.639a.75.75 0 0 1-.75.75h-3.638l-2.573 2.661a.75.75 0 0 1-1.078 0L8.888 19.5H5.25a.75.75 0 0 1-.75-.75v-3.64l-2.661-2.57a.75.75 0 0 1 0-1.078L4.5 8.888V5.25a.75.75 0 0 1 .75-.75h3.638l2.573-2.662ZM12 7.25a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}