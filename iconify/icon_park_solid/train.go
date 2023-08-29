package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrain0"><g fill="none"><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M9 8.84c0-.933.438-1.807 1.271-2.225C12.246 5.625 16.613 4 24 4c7.388 0 11.754 1.625 13.728 2.615C38.563 7.033 39 7.907 39 8.839V36a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2V8.84Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 38v4m-20-4v4"/><path fill="#fff" stroke="#fff" d="M20.5 32a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm10 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 44h8"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M9 12h30v14H9z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 12v14"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 12h8m-8 14h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrain0)"/>`),
		g.Group(children),
	)
}