package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseAcceptableButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#eda454" d="M62 32c0 16.5-13.5 30-30 30S2 48.5 2 32S15.5 2 32 2s30 13.5 30 30z"/><g fill="#fff"><path d="M48 18v4.9h-3.6v21.3c0 2.7-.5 4.1-2.4 5c-1.8.8-4.5.9-8.3.9c-.2-1.5-1-3.7-1.7-5.1c2.6.2 5.4.2 6.2.1c.8 0 1.1-.3 1.1-1V22.9H16V18h32"/><path d="M19 26.5v17h4.6v-2.7H35V26.5H19m11.1 9.9h-6.4v-5.5h6.4v5.5"/></g>`),
		g.Group(children),
	)
}