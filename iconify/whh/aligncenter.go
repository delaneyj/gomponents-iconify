package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aligncenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 320H64q-27 0-45.5-18.5T0 256.5T18.5 211T64 192h896q26 0 45 19t19 45.5t-19 45t-45 18.5zM832 128H192q-27 0-45.5-18.5t-18.5-45T146.5 19T192 0h640q26 0 45 19t19 45.5t-19 45t-45 18.5zM128 512q-27 0-45.5-18.5t-18.5-45T82.5 403t45.5-19h768q26 0 45 19t19 45.5t-19 45t-45 18.5H128zm-64 64h896q26 0 45 19t19 45.5t-19 45t-45 18.5H64q-27 0-45.5-18.5T0 640t18.5-45.5T64 576zm256 192h384q26 0 45 19t19 45.5t-19 45t-45 18.5H320q-27 0-45.5-18.5T256 832t18.5-45.5T320 768z"/>`),
		g.Group(children),
	)
}