package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M481 31c-35.9-35.8-94.1-35.8-130 0l-15 15l-13.1-13c-28.1-28.1-73.7-28.1-101.8 0L135 119c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l86.1-86c9.4-9.4 24.6-9.4 33.9 0L302.1 80L186.3 195.7l130 130L481 161c35.9-35.9 35.9-94.1 0-129.9zM293.7 348.3l-130-130l-64.2 64.2c-48 48-80.8 109.2-94.1 175.8l-5 25c-1.6 7.9.9 16 6.6 21.7s13.8 8.1 21.7 6.6l25-5c66.6-13.3 127.8-46.1 175.8-94.1l64.2-64.2z"/>`),
		g.Group(children),
	)
}