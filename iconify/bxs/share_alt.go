package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12c0 1.654 1.346 3 3 3c.794 0 1.512-.315 2.049-.82l5.991 3.424c-.018.13-.04.26-.04.396c0 1.654 1.346 3 3 3s3-1.346 3-3s-1.346-3-3-3c-.794 0-1.512.315-2.049.82L8.96 12.397c.018-.131.04-.261.04-.397s-.022-.266-.04-.397l5.991-3.423c.537.505 1.255.82 2.049.82c1.654 0 3-1.346 3-3s-1.346-3-3-3s-3 1.346-3 3c0 .136.022.266.04.397L8.049 9.82A2.982 2.982 0 0 0 6 9c-1.654 0-3 1.346-3 3z"/>`),
		g.Group(children),
	)
}