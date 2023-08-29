package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modernizr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#D91B77" d="M0 122.88V81.92h40.96V40.96h40.96V0h40.96v122.88H0ZM133.12 0C200.986 0 256 55.016 256 122.879H133.12V0Z"/>`),
		g.Group(children),
	)
}