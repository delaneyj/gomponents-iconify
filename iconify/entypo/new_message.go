package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.174 1.826c-1.102-1.102-2.082-.777-2.082-.777L7.453 9.681L6 14l4.317-1.454l8.634-8.638s.324-.98-.777-2.082zm-7.569 9.779l-.471.47l-1.473.5a2.216 2.216 0 0 0-.498-.74a2.226 2.226 0 0 0-.74-.498l.5-1.473l.471-.47s.776-.089 1.537.673c.762.761.674 1.538.674 1.538zM16 17H3V4h5l2-2H3c-1.1 0-2 .9-2 2v13c0 1.1.9 2 2 2h13c1.1 0 2-.9 2-2v-7l-2 2v5z"/>`),
		g.Group(children),
	)
}