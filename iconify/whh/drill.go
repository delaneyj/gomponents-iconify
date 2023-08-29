package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 192h-96v32q0 13-9.5 22.5T864 256H750q-17 29-46 46.5T640 320H512v78q0 21-13.5 35.5T466 448h-44l-38 192h128q27 0 45.5 18.5T576 704v256q0 26-19 45t-45 19H64q-27 0-45.5-19T0 960V704q0-27 18.5-45.5T64 640l64-512q0-53 37.5-90.5T256 0h384q35 0 64 17.5T750 64h114q13 0 22.5 9.5T896 96v32h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 192zM160 704q-13 0-22.5 9.5T128 736t9.5 22.5T160 768h256q13 0 22.5-9.5T448 736t-9.5-22.5T416 704H160zM608 64H288q-13 0-22.5 9.5T256 96t9.5 22.5T288 128h320q13 0 22.5-9.5T640 96t-9.5-22.5T608 64z"/>`),
		g.Group(children),
	)
}