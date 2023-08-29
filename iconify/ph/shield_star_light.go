package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldStarLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M173.57 117.77a6 6 0 0 1-3.34 7.8l-32.73 13.09l19.3 25.74a6 6 0 0 1-9.6 7.2L128 146l-19.2 25.6a6 6 0 1 1-9.6-7.2l19.3-25.74l-32.73-13.09a6 6 0 0 1 4.46-11.14L122 127.14V96a6 6 0 0 1 12 0v31.14l31.77-12.71a6 6 0 0 1 7.8 3.34ZM222 56v58.77c0 88.25-74.68 117.52-89.65 122.49a13.5 13.5 0 0 1-8.7 0C108.68 232.31 34 203 34 114.79V56a14 14 0 0 1 14-14h160a14 14 0 0 1 14 14Zm-12 0a2 2 0 0 0-2-2H48a2 2 0 0 0-2 2v58.77c0 80 67.84 106.59 81.43 111.1a1.57 1.57 0 0 0 1.13 0c13.6-4.51 81.44-31.09 81.44-111.1Z"/>`),
		g.Group(children),
	)
}