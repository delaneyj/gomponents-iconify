package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LatteMacchiato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M45 42H26m19-18H25.5m20-7H25.14"/><path fill="#fff" fill-rule="evenodd" d="M22.034 13.033a1 1 0 0 1 1-1.033h24.93a1 1 0 0 1 .998 1.036L48.856 16L47.34 58.072a2 2 0 0 1-2 1.928H25.524a2 2 0 0 1-1.999-1.934L22.132 16l-.098-2.967Z" clip-rule="evenodd"/><path fill="#A57939" d="M24.037 17.036a1 1 0 0 1 1-1.036h20.927a1 1 0 0 1 .999 1.036L45.569 56.07A2 2 0 0 1 43.571 58H26.473a1 1 0 0 1-1-.964l-1.436-40Z"/><path fill="#6A462F" fill-rule="evenodd" d="M46.108 41h-21.21l-.86-23.964A1 1 0 0 1 25.036 16h20.927a1 1 0 0 1 .999 1.036L46.108 41Z" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="M46.75 23H24.251l-.214-5.964a1 1 0 0 1 1-1.036h20.927a1 1 0 0 1 .999 1.036L46.75 23Z" clip-rule="evenodd"/><path d="m22.034 13.033l-1 .033l1-.033Zm26.928.003L47.963 13l1 .036ZM48.856 16l-1-.036L47.82 17h1.037v-1Zm0 0l1 .036l.037-1.036h-1.037v1ZM47.34 58.072l.999.036l-1-.036Zm-23.816-.006l-1 .033l1-.033ZM22.132 16l1-.033L23.1 15h-.968v1Zm0 0l-1 .033l.033.967h.967v-1Zm.901-5a2 2 0 0 0-1.998 2.066L23.033 13v-2Zm24.93 0h-24.93v2h24.93v-2Zm1.999 2.072A2 2 0 0 0 47.963 11v2l1.999.072Zm-.107 2.964l.107-2.964L47.963 13l-.107 2.964l2 .072Zm-1 .964v-2v2Zm-.998-1.036L46.34 58.036l1.999.072l1.516-42.072l-1.998-.072ZM46.34 58.036a1 1 0 0 1-1 .964v2a3 3 0 0 0 2.999-2.892l-1.999-.072Zm-1 .964H25.524v2h19.818v-2Zm-19.817 0a1 1 0 0 1-1-.967l-1.998.066A3 3 0 0 0 25.523 61v-2Zm-1-.967l-1.391-42.066l-2 .066L22.526 58.1l1.999-.066ZM22.134 17v-2v2Zm-1.098-3.934l.098 2.967l1.999-.066l-.1-2.967l-1.998.066Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M34 57h9.535a1 1 0 0 0 1-.964L46 15"/>`),
		g.Group(children),
	)
}