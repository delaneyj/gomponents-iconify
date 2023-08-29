package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fonoma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.706 21.919a4.295 4.295 0 0 0 4.293-4.296a4.295 4.295 0 1 0-4.293 4.296zM4.296 10.672a4.295 4.295 0 0 0 4.293-4.295a4.295 4.295 0 1 0-4.294 4.295zm10.412 0h4.975a4.277 4.277 0 0 0 4.293-4.295a4.277 4.277 0 0 0-4.293-4.296h-4.975a4.277 4.277 0 0 0-4.294 4.296a4.277 4.277 0 0 0 4.294 4.295zM4.295 21.92h4.976a4.277 4.277 0 0 0 4.293-4.296a4.277 4.277 0 0 0-4.293-4.295H4.295a4.277 4.277 0 0 0-4.293 4.295c.068 2.318 1.976 4.296 4.293 4.296z"/>`),
		g.Group(children),
	)
}