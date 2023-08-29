package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminPlugins(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.11 4.36L9.87 7.6L8 5.73l3.24-3.24c.35-.34 1.05-.2 1.56.32c.52.51.66 1.21.31 1.55zm-8 1.77l.91-1.12l9.01 9.01l-1.19.84c-.71.71-2.63 1.16-3.82 1.16H6.14L4.9 17.26c-.59.59-1.54.59-2.12 0a1.49 1.49 0 0 1 0-2.12l1.24-1.24v-3.88c0-1.13.4-3.19 1.09-3.89zm7.26 3.97l3.24-3.24c.34-.35 1.04-.21 1.55.31c.52.51.66 1.21.31 1.55l-3.24 3.25z"/>`),
		g.Group(children),
	)
}