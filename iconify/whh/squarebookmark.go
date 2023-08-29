package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarebookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-752q0-33-23-56.5t-55-23.5h-228q-32 0-55 23.5t-23 56.5v520q0 15 11 27.5t21 12.5l160-128l160 128q10 0 21-12.5t11-27.5V272z"/>`),
		g.Group(children),
	)
}