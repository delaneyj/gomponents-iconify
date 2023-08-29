package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdtv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640.585 703v65h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-448q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5h96v-65q-384-7-384-95V64q0-27 18.5-45.5T64.585 0h896q26 0 45 18.5t19 45.5v544q0 88-384 95zm320-639h-896v512h896V64zm-832 448V128h576z"/>`),
		g.Group(children),
	)
}