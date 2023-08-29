package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Questionnaire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704V2Zm2 2v14.296L6.124 16H20.5V4h-17ZM12 7.5a1 1 0 0 0-1 1v1H9v-1a3 3 0 1 1 6 0c0 .676-.172 1.246-.474 1.71a2.957 2.957 0 0 1-1.029.95a4.136 4.136 0 0 1-.494.238v.352h-2v-1c0-.424.245-.687.361-.79c.12-.105.24-.165.296-.192c.107-.05.233-.094.309-.12l.018-.007c.19-.066.36-.127.52-.218a.964.964 0 0 0 .343-.305c.072-.11.15-.294.15-.618a1 1 0 0 0-1-1Zm-1 5h2.004v2.004H11V12.5Z"/>`),
		g.Group(children),
	)
}