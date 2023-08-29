package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YamlAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 63.5h114.744L255.89 205.912L396.854 63.5H512l-383.215 385H14.58l184.334-185.014L0 63.5z"/>`),
		g.Group(children),
	)
}