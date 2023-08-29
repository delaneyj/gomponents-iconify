package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#F44336" d="m69.8 88.1l1.8-80c.1-2.2-1.8-4.1-4-4.1H55.9c-2.3 0-4.1 1.9-4 4.1l1.8 80c.1 2.2 1.8 3.9 4 3.9h8c2.2 0 4-1.7 4.1-3.9z"/><g fill="#FFF"><path d="M64 50c.3-1.9-.5-26.2-.5-26.2s.1-3.5-2.8-3.8c-2.8-.3-3.8 2.3-3.8 4.1c0 1.8.4 21.5.6 23.3c.1 1.8 1.9 3.7 3.6 4.2s2.7-.2 2.9-1.6z" opacity=".2"/><circle cx="59.6" cy="13.4" r="3.3" opacity=".2"/></g><circle cx="64.1" cy="112" r="12" fill="#C33"/><circle cx="62.4" cy="112" r="10.3" fill="#F44336"/><path fill="#FFF" d="M56.5 108.4c1.2-1.8 3.8-3.3 6.5-3.7c.7-.1 1.3-.1 1.9.1c.4.2.8.6.5 1c-.2.4-.7.5-1.1.6c-2.5.7-4.8 2.4-6.2 4.4c-.5.8-1.4 2.9-2.4 2.4c-1-.7-.7-2.6.8-4.8z" opacity=".2"/><path fill="#C33" d="M71.9 4h-4.3c2.3 0 4.1 1.9 4 4.1l-1.8 80c-.1 2.2-1.8 3.9-4 3.9H70c2.2 0 4-1.7 4-3.9l1.8-80C76 5.9 74.1 4 71.9 4z"/>`),
		g.Group(children),
	)
}