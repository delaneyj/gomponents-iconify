package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PluginsChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m13.6 5.1l-3.1 3.1l1.8 1.8l3.1-3.1c.3-.3.2-1-.3-1.5s-1.1-.6-1.5-.3zm.3-4.8c-.7-.4-9.8 7.3-9.8 7.3S.6 5.5.1 5.9c-.5.4 4 5 4 5S14.6.6 13.9.3zm5.5 9.3c-.5-.5-1.2-.6-1.5-.3l-3.1 3.1l1.8 1.8l3.1-3.2c.3-.2.2-.9-.3-1.4zm-11.7-1c-.7.7-1.1 2.7-1.1 3.8v3.8l-1.2 1.2c-.6.6-.6 1.5 0 2.1s1.5.6 2.1 0l1.2-1.2h3.8c1.2 0 3-.4 3.7-1.1l1.2-.8l-8.9-8.9l-.8 1.1z"/>`),
		g.Group(children),
	)
}