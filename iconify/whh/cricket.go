package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cricket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.556 256q-53 0-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5zm108 653l-95 95q-19 20-47 20t-48-20l-474-474q-20-20-20-48t20-47l9-9l-150-150l-177-96q-22-21-22-52t22-53l53-53q22-22 53-22t52 22l96 177l150 150l9-9q19-20 47-20t48 20l474 474q20 20 20 48t-20 47z"/>`),
		g.Group(children),
	)
}