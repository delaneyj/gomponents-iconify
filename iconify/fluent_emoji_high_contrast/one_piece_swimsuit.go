package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnePieceSwimsuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g clip-path="url(#fluentEmojiHighContrastOnePieceSwimsuit0)"><path fill="currentColor" d="M5.205 10.696a5.39 5.39 0 0 1-.197-1.828V2.88a2.9 2.9 0 0 1 2.9-2.9h.19a2.9 2.9 0 0 1 2.9 2.9v1.21h9.022V2.9a2.9 2.9 0 0 1 2.9-2.9h.19a2.9 2.9 0 0 1 2.9 2.9v6.025a5.16 5.16 0 0 1-.2 1.804l-1.716 6.002l2.512 5.433a3.05 3.05 0 0 1-.61 3.427l-.137.138l-7.768 5.242a4.88 4.88 0 0 1-5.69-.018l-7.308-5.232l-.136-.145a3.05 3.05 0 0 1-.5-3.445l2.449-5.398l-1.701-6.037Zm18.685-.526c.112-.38.153-.776.12-1.17V2.9a.9.9 0 0 0-.9-.9h-.19a.9.9 0 0 0-.9.9v2.25a.943.943 0 0 1-.276.665a.94.94 0 0 1-.665.275H9.94A.94.94 0 0 1 9 5.15V2.88a.9.9 0 0 0-.9-.9h-.19a.9.9 0 0 0-.9.9v6.06a3.39 3.39 0 0 0 .12 1.21L7.933 13H23.08l-.001.003l.81-2.833Zm-1.35 4.718a1.35 1.35 0 0 0-.035.112H8.495l.282 1h13.446l.317-1.112ZM22.477 18H8.527l-2.269 5a1.05 1.05 0 0 0 .16 1.21l7.15 5.12a2.88 2.88 0 0 0 3.38 0l7.631-5.15a1.05 1.05 0 0 0 .21-1.18l-2.312-5Z"/></g><defs><clipPath id="fluentEmojiHighContrastOnePieceSwimsuit0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}