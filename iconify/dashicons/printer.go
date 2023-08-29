package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 11H7v1h5v-1zm1 4H7v1h6v-1zm-3-2H7v1h3v-1zm7-7h-2V2H5v4H3c-.6 0-1 .4-1 1v5c0 .6.4 1 1 1h2v5h10v-5h2c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1zm-3 11H6v-7h8v7zm0-11H6V3h8v3zm2 3h-1V8h1v1z"/>`),
		g.Group(children),
	)
}