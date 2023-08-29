package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4d5357" d="M32 2L16 18h10v15h12V18h10zm0 60c-3.3 0-6-2.8-6-6.3v-9.3c0-3.5 2.7-6.3 6-6.3s6 2.8 6 6.3v9.3c0 3.5-2.7 6.3-6 6.3m0-18.6c-1.7 0-3 1.3-3 3v9.3c0 1.6 1.3 3 3 3s3-1.3 3-3v-9.3c0-1.7-1.3-3-3-3M22 40H10v3h4.5v19h3V43H22zm26 0h-6v22h3v-9.3h3c3.3 0 6-2.8 6-6.3c0-3.6-2.7-6.4-6-6.4m0 9.3h-3v-5.9h3c1.7 0 3 1.3 3 3s-1.3 2.9-3 2.9"/>`),
		g.Group(children),
	)
}