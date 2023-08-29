package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TokyoTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#dae3ea" d="M38.7 47.5H25.3L24 44h16z"/><path fill="#ed4c5c" d="M38.7 48.2H25.3C23.6 54.4 21.2 59 16 64h7.7s2.7-9.5 8.3-9.5s8.3 9.5 8.3 9.5H48c-5.2-5-7.6-9.6-9.3-15.8"/><g fill="#dae3ea"><path d="M18 25.7c0 .9-.7 1.7-1.7 1.7H1.7c-.9 0-1.7-.7-1.7-1.7S.7 24 1.7 24h14.7c.9 0 1.6.7 1.6 1.7m6.5-3.4c0 .9-.7 1.7-1.7 1.7H8.2c-.9 0-1.7-.7-1.7-1.7c0-.9.7-1.7 1.7-1.7h14.7c.9.1 1.6.8 1.6 1.7M64 40.7c0 .9-.7 1.7-1.7 1.7H47.7c-.9 0-1.7-.7-1.7-1.7c0-.9.7-1.7 1.7-1.7h14.7c.9 0 1.6.8 1.6 1.7"/><path d="M60.2 37.4c0 .9-.7 1.7-1.7 1.7H43.8c-.9 0-1.7-.7-1.7-1.7c0-.9.7-1.7 1.7-1.7h14.7c.9 0 1.7.7 1.7 1.7"/></g><path fill="#ed4c5c" d="M37.5 43.3h-11l1.6-5.6h7.8z"/><path fill="#dae3ea" d="M35.9 37h-7.8l.8-7h6z"/><path fill="#ed4c5c" d="M35 30h-6l.8-6.8h4.4z"/><path fill="#dae3ea" d="M28.4 19.6h7.2v2.9h-7.2zm0-4.6h7.2v1.8h-7.2z"/><path fill="#ed4c5c" d="M28.4 13.6h7.2V15h-7.2z"/><path fill="#c94747" d="M30.7 8.2h2.6v5.4h-2.6z"/><path fill="#dae3ea" d="M30.7 1.5h2.6v6.8h-2.6z"/><path fill="#ed4c5c" d="M30.7 0h2.6v1.5h-2.6z"/><path fill="#c5d0d8" d="M30 16.8h4v2.9h-4zM28.3 30h7.5v.7h-7.5z"/><path fill="#b2c1c0" d="M30.7 7.5h2.6v.7h-2.6zm-2.3 14.3h7.2v.7h-7.2zm-.1 8.9h7.5v.7h-7.5zm-.2 2.1H36v.7h-7.9zm-.4 2.1h8.6v.7h-8.6z"/><path fill="#c5d0d8" d="M28.1 32.1H36v.7h-7.9zm-.4 2.1h8.6v.7h-8.6zm-.3 2.1h9.3v.7h-9.3z"/><path fill="#c94747" d="M25.3 47.5h13.3v.7H25.3zm4.5-25h4.3v.7h-4.3zm-.2 21h-1.3v-3.3c0-.4.3-.6.6-.6c.4 0 .6.3.6.6l.1 3.3m2 0h-1.3v-3.3c0-.4.3-.6.6-.6c.4 0 .6.3.6.6l.1 3.3m2.1 0h-1.3v-3.3c0-.4.3-.6.6-.6c.4 0 .6.3.6.6l.1 3.3m2 0h-1.3v-3.3c0-.4.3-.6.6-.6c.4 0 .6.3.6.6l.1 3.3M31.6 30h-1.3v-3.3c0-.4.3-.6.6-.6c.4 0 .6.3.6.6l.1 3.3m2.1 0h-1.3v-3.3c0-.4.3-.6.6-.6c.4 0 .6.3.6.6l.1 3.3"/><path fill="#b2c1c0" d="M24 43.3h16v.7H24z"/>`),
		g.Group(children),
	)
}