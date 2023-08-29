package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTextToSpeechOneSpeechAutomatedTranslateLanguageVoiceTechnology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3.5 13.5l3.14-7.33a.39.39 0 0 1 .72 0l3.14 7.33m-5.59-3.29h4.18M.5 4A3.5 3.5 0 0 1 4 .5"/><path d="M.5 4A3.5 3.5 0 0 1 4 .5M3 4a1 1 0 0 1 1-1"/><path d="M3 4a1 1 0 0 1 1-1m6-2.5A3.5 3.5 0 0 1 13.5 4"/><path d="M10 .5A3.5 3.5 0 0 1 13.5 4M10 3a1 1 0 0 1 1 1"/><path d="M10 3a1 1 0 0 1 1 1"/></g>`),
		g.Group(children),
	)
}