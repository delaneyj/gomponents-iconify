package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 928q-3 36-32.5 66t-63.5 30H160q-34 0-63.5-30T64 928L0 512q0-49 28-88.5T96 384h57q-25-47-25-64q0-14 37-33t91-36q2-40 18.5-97.5t47-105.5T384 0t62.5 48t47 105.5T512 251q54 17 91 36t37 33q0 17-25 64h57q40 0 68 39.5t28 88.5z"/>`),
		g.Group(children),
	)
}