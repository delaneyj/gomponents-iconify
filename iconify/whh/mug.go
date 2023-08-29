package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.232 704q-31 0-64-12v12q0 80-56 136t-136 56h-384q-80 0-136-56t-56-136V64q0-27 18.5-45.5T64.232 0h640q27 0 45.5 18.5t18.5 45.5v75q31-11 64-11q80 0 136 56t56 136v192q0 80-56 136t-136 56zm64-363q0-35-28-60t-68-25q-16 0-32 5v310q16 5 32 5q40 0 68-25t28-60V341z"/>`),
		g.Group(children),
	)
}