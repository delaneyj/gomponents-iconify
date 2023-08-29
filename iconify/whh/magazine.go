package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magazine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.85 1024h-768q-27 0-45.5-19T.85 960V64q0-27 18.5-45.5T64.85 0h768q26 0 45 18.5t19 45.5v896q0 26-19 45t-45 19zm-704-960q-27 0-45.5 18.5T64.85 128v768q0 27 18.5 45.5t45.5 18.5h448q0-60-3-98t-10.5-67t-23-45.5t-36.5-27t-55-18.5q-17-4-37-4t-36 2.5t-37.5 8t-32.5 9t-29.5 9.5t-19.5 7l543-664q-15-8-31-8h-640zm704 64q0-21-13-39l-204 871h153q26 0 45-18.5t19-45.5V128z"/>`),
		g.Group(children),
	)
}