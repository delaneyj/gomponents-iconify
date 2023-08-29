package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Datocms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0H.076v24H12c5.964 0 11.924-5.373 11.924-11.998C23.924 5.376 17.963 0 12 0zm0 17.453a5.453 5.453 0 1 1 5.455-5.451A5.45 5.45 0 0 1 12 17.452z"/>`),
		g.Group(children),
	)
}