package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardTrophyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsAwardTrophyNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM13 6a1 1 0 0 0-1 1v1H7a1 1 0 0 0-1 1v6a5 5 0 0 0 5 5h1.683A12.017 12.017 0 0 0 22 27.834V34h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6v-6.166A12.017 12.017 0 0 0 35.317 20H37a5 5 0 0 0 5-5V9a1 1 0 0 0-1-1h-5V7a1 1 0 0 0-1-1H13Zm23 4v8h1a3 3 0 0 0 3-3v-5h-4ZM8 10h4v8h-1a3 3 0 0 1-3-3v-5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAwardTrophyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}