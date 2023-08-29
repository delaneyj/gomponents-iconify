package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffeebean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 909q-9-45-111-171T564.5 461.5T288 209T117 98q84-78 209.5-94T587 38.5t242.5 158t158 242.5t34.5 260.5T928 909zM461.5 564.5Q612 715 738 817t171 111q-84 78-209.5 94T439 987.5t-242.5-158T38.5 587T4 326.5T98 117q9 45 111 171t252.5 276.5z"/>`),
		g.Group(children),
	)
}