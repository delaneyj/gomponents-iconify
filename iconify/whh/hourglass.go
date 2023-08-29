package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.988 128h-64v64q0 99-46.5 183t-125.5 137q79 53 125.5 137.5t46.5 182.5v64h64q26 0 45 18.5t19 45.5t-19 45.5t-45 18.5h-896q-26 0-45-18.5t-19-45t19-45.5t45-19h64v-64q0-98 46.5-182.5t125.5-137.5q-79-53-125.5-137t-46.5-183v-64h-64q-26 0-45-18.5t-19-45t19-45.5t45-19h896q26 0 45 19t19 45.5t-19 45t-45 18.5zm-704 704v64h512q2-9 2-19.5t-1-24t-1-20.5q0-106-75-181t-181-75t-181 75t-75 181zm512-704h-512v64q0 106 75 181t181 75t181-75t75-181v-64z"/>`),
		g.Group(children),
	)
}