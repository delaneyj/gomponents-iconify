package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseAcceptableButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M8 9a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2h-.875a.125.125 0 0 0-.125.125V23a1 1 0 0 1-1 1h-4a1 1 0 1 1 0-2h2.875a.125.125 0 0 0 .125-.125v-11.75a.125.125 0 0 0-.125-.125H9a1 1 0 0 1-1-1Z"/><path d="M10 13a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-4.875c-.069 0-.124.056-.133.125A1 1 0 0 1 10 20v-7Zm5.875 1h-3.75a.125.125 0 0 0-.125.125v3.75c0 .069.056.125.125.125h3.75a.125.125 0 0 0 .125-.125v-3.75a.125.125 0 0 0-.125-.125Z"/><path d="M16 1C7.716 1 1 7.716 1 16c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15ZM3 16C3 8.82 8.82 3 16 3s13 5.82 13 13s-5.82 13-13 13S3 23.18 3 16Z"/></g>`),
		g.Group(children),
	)
}