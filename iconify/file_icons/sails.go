package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M47.593 512S-148.257 200.549 254.635 0v512H47.593M317 512V194l195 318H317"/>`),
		g.Group(children),
	)
}