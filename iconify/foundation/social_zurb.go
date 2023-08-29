package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialZurb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M79 14H21a7 7 0 0 0-7 7v58a7 7 0 0 0 7 7h58a7 7 0 0 0 7-7V21a7 7 0 0 0-7-7zm-8.647 43.23v5.339a2.257 2.257 0 0 1-1.973 2.019l-30.145.002c-2.232 0-4.135-.72-5.707-2.157c-1.571-1.438-2.357-3.177-2.357-5.221c0-1.626.531-3.071 1.594-4.345c1.064-1.27 2.411-2.154 4.041-2.656l26.341-7.527H30.172v-5.096a2.263 2.263 0 0 1 1.779-2.177h30.338c2.232 0 4.133.714 5.705 2.139c1.572 1.427 2.359 3.13 2.359 5.112c0 1.655-.536 3.111-1.607 4.371c-1.071 1.262-2.449 2.141-4.134 2.635L38.235 57.23h32.118z"/>`),
		g.Group(children),
	)
}