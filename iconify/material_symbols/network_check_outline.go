package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.1 11.1L0 9q2.3-2.35 5.375-3.675T12 4q.6 0 1.2.037t1.2.113l-1.5 2.9Q12.6 7 12.412 7H12Q9.1 7 6.562 8.088T2.1 11.1Zm4.25 4.25l-2.1-2.15q1.4-1.4 3.213-2.212t3.937-.938L9.8 13.3q-.975.3-1.85.813t-1.6 1.237Zm4.95 4.5q-.75-.275-1.163-1.038T10.1 17.3l6-12.2q.1-.2.288-.263t.412.013q.2.075.3.25t.05.4L13.9 18.65q-.2.825-1 1.163t-1.6.037Zm6.35-4.5q-.15-.15-.325-.3t-.375-.3l.8-3.15q.525.375 1.038.763t.962.837l-2.1 2.15Zm4.25-4.25q-.75-.75-1.612-1.375T18.45 8.6l.7-3q1.35.65 2.575 1.5T24 9l-2.1 2.1Z"/>`),
		g.Group(children),
	)
}