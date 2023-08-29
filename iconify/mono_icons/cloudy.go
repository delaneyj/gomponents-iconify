package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 5a3 3 0 0 1 2.557 4.57a5.96 5.96 0 0 0-1.886-.533a6.019 6.019 0 0 0-2.697-3.25A2.99 2.99 0 0 1 16 5zm-4.055.074a6 6 0 0 0-6.931 6.336A5 5 0 0 0 7 21h9a6 6 0 0 0 4.2-10.285a5 5 0 0 0-8.255-5.64zM7 11a4 4 0 0 1 7.92-.8a1 1 0 0 0 1 .8H16a4 4 0 0 1 0 8H7a3 3 0 0 1-.66-5.927a1 1 0 0 0 .757-1.194A4.017 4.017 0 0 1 7 11z"/>`),
		g.Group(children),
	)
}