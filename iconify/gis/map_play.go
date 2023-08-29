package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M65.809 5a2.5 2.5 0 0 0-1.03.232L34.166 19.453L3.553 5.233A2.5 2.5 0 0 0 0 7.5v70.29a2.5 2.5 0 0 0 1.447 2.267l31.666 14.71A2.5 2.5 0 0 0 34.19 95a2.5 2.5 0 0 0 1.032-.232l30.613-14.221l30.613 14.22A2.5 2.5 0 0 0 100 92.5V22.21a2.5 2.5 0 0 0-1.447-2.267L66.887 5.233A2.5 2.5 0 0 0 65.809 5zm1.142 5.775L95 23.805v64.777L67.322 75.725l-.37-64.95zm-2.998.354l.37 64.605l-28.677 13.323l-.369-64.606L63.953 11.13zM5 11.418l27.275 12.67l.371 64.95L5 76.192V11.418z" color="currentColor"/><path fill="currentColor" d="M50 24a26 26 0 0 0-26 26a26 26 0 0 0 26 26a26 26 0 0 0 26-26a26 26 0 0 0-26-26zm-6.979 9.502a1.526 1.5 0 0 1 1 .299l20.348 15a1.526 1.5 0 0 1 0 2.398l-20.348 15A1.526 1.5 0 0 1 41.581 65V35a1.526 1.5 0 0 1 1.44-1.498z"/>`),
		g.Group(children),
	)
}