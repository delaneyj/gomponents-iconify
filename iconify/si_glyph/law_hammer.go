package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LawHammer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.44 7.389L9.662 2.61l.57-.569l-.917-.919l-4.247 4.245l.92.918l.53-.53l1.955 1.954l-8.535 8.535l.83.829l8.533-8.534l1.994 1.993l-.529.531l.918.917l4.254-4.253l-.917-.918l-.581.58z"/>`),
		g.Group(children),
	)
}