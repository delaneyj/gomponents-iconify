package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Technorati(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M706 564q-105 0-193.5-34.5t-140-93.5T321 307.5t51.5-129t140-93.5T706 51q28 0 61 3q108 62 176.5 166.5T1024 452q-53 51-137 81.5T706 564zM192 307q0 105 69 193t185 139l-61 182l321-129q170 0 307-77q-40 173-179.5 285.5T513 1013q-104 0-199-40.5T150 863T40.5 699T0 500q0-181 112.5-320.5T397 0q-95 54-150 134.5T192 307z"/>`),
		g.Group(children),
	)
}