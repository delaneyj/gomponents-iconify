package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileQuestionAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.07 12h-5a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2Zm1 8h-8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V9a.14.14 0 0 0 0-.06a.86.86 0 0 0-.07-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1 1 0 0 0-.29-.19s-.05 0-.08 0a.88.88 0 0 0-.32-.11h-6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a1 1 0 0 0 0-2Zm-1-14.59L15.65 8h-1.58a1 1 0 0 1-1-1Zm5.57 14.88a1.58 1.58 0 0 0-.15-.12a1.08 1.08 0 0 0-.36-.15a1 1 0 0 0-.9.27a1 1 0 0 0 0 1.42a1 1 0 0 0 .7.29a1 1 0 0 0 .93-1.38a1.19 1.19 0 0 0-.22-.33ZM13.07 16h-5a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2Zm4.86-3a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.74 1a1 1 0 1 1 .86 1.5a1 1 0 0 0 0 2a3 3 0 0 0 0-6Zm-9.86-3h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}