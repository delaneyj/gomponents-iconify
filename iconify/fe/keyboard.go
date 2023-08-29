package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feKeyboard0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feKeyboard1" fill="currentColor"><path id="feKeyboard2" d="M2 8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8Zm15 6v2h3v-2h-3Zm-3 0v2h2v-2h-2Zm-7 0v2h6v-2H7Zm-3 0v2h2v-2H4Zm14-3v2h2v-2h-2Zm-3 0v2h2v-2h-2Zm-3 0v2h2v-2h-2Zm-3 0v2h2v-2H9Zm-5 0v2h4v-2H4Zm12-3v2h4V8h-4Zm-3 0v2h2V8h-2Zm-3 0v2h2V8h-2ZM7 8v2h2V8H7ZM4 8v2h2V8H4Z"/></g></g>`),
		g.Group(children),
	)
}