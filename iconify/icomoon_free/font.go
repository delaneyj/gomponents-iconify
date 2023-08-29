package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.494.253C11.08.253 10.172 0 8.715 0C4.007 0 1.812 2.681 1.812 5.404c0 1.604.76 2.132 2.259 2.132c-.106-.232-.296-.486-.296-1.626c0-3.188 1.203-4.117 2.744-4.18c0 0-1.264 12.396-4.934 13.883v.385h4.947l1.688-8h3.091l.689-2H8.642l.812-3.847c.929.19 1.837.38 2.618.38c.971 0 1.858-.296 2.343-2.533c-.591.19-1.224.253-1.921.253z"/>`),
		g.Group(children),
	)
}