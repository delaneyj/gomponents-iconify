package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClassicalBuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#acb8bf" d="M2 60.1h60V62H2zm3.5-3.7h53v1.9h-53z"/><path fill="#dae3ea" d="M3.8 58.2h56.5v1.9H3.8zm3.4-3.7h49.5v1.9H7.2z"/><path fill="#94989b" d="M4.8 27.8h54.4v1.9H4.8z"/><path fill="#dae3ea" d="M2 29.6h60v1.9H2z"/><path fill="#94989b" d="M32 4.8L2 21.2L4.5 25L32 6.7L59.5 25l2.5-3.8z"/><g fill="#dae3ea"><path d="M32 2L2 21.2v2.9L32 4.8l30 19.3v-2.9z"/><path d="M32 6.7L2 25.9v1.9h60v-1.9z"/></g><path fill="#acb8bf" d="M32 9.5L7.1 25.9h49.8z"/><g fill="#94989b"><path d="m32.8 22.6l-.4.2c0-.1-.1-.2-.2-.2h-.3c-.1 0-.2.1-.2.2l-.4-.2l-1-.6v2.2l1-.6l.4-.2c0 .1.1.2.2.2h.3c.1 0 .2-.1.2-.2l.4.2l1 .6V22l-1 .6m-2.3-3.1c.2.3.4.4.6.6c.3.1.6.2.9.2c.3 0 .6-.1.9-.2c.3-.1.5-.3.6-.6c0 .3-.1.6-.4.8c-.3.2-.7.4-1.1.4c-.4 0-.8-.1-1.1-.4c-.3-.2-.4-.5-.4-.8"/><path d="M36.3 16.6v-.2c-.1 0-.1-.1-.2-.1h-.2c-.6-1.5-2.1-2.6-3.9-2.6c-1.7 0-3.2 1.1-3.9 2.6h-.2c-.1 0-.1 0-.2.1v.2c0 .2 0 .1.1.2c0 0 .1 0 .1.1c-.1.4-.2.7-.2 1.1c0 2.3 1.9 4.2 4.2 4.2s4.2-1.9 4.2-4.2c0-.4-.1-.8-.2-1.1l.1-.1c.3-.1.3-.1.3-.2M32 14.2c1.4 0 2.7.8 3.3 2h-1.4c-.5 0-1 .2-1.3.3c-.3.2-.8.2-1.2 0c-.4-.2-.8-.3-1.3-.3h-1.4c.6-1.2 1.9-2 3.3-2m0 7.5c-2.1 0-3.7-1.7-3.7-3.7v-.1c.1.5.4.8.9 1c.5.2 1 .2 1.5 0c.3-.1.5-.3.7-.5c.3-.5.2-.8.4-1.2c.1-.3.5-.3.7 0c.2.4.1.7.4 1.2c.2.3.4.4.7.5c.5.2 1 .2 1.5 0c.4-.2.7-.5.9-1v.1c-.3 2-1.9 3.7-4 3.7"/></g><path fill="#dae3ea" d="M8.5 33.2h6v19.6h-6z"/><path fill="#94989b" d="M7.8 52.8h7.5v1.8H7.8zm0-21.3h7.5v1.7H7.8z"/><path fill="#c8d0d6" d="M9.1 34h.7v18h-.7zm4.1 0h.7v18h-.7zm-1.4 0h.7v18h-.7zm-1.3 0h.7v18h-.7z"/><path fill="#dae3ea" d="M22.2 33.2h6v19.6h-6z"/><path fill="#94989b" d="M21.4 52.8h7.5v1.7h-7.5zm0-21.3h7.5v1.7h-7.5z"/><path fill="#c8d0d6" d="M22.8 34h.7v18h-.7zm4.1 0h.7v18h-.7zm-1.4 0h.7v18h-.7zm-1.4 0h.7v18h-.7z"/><path fill="#dae3ea" d="M35.8 33.2h6v19.6h-6z"/><path fill="#94989b" d="M35.1 52.8h7.5v1.7h-7.5zm0-21.3h7.5v1.7h-7.5z"/><path fill="#c8d0d6" d="M36.4 34h.7v18h-.7zm4.1 0h.7v18h-.7zm-1.3 0h.7v18h-.7zm-1.4 0h.7v18h-.7z"/><path fill="#dae3ea" d="M49.5 33.2h6v19.6h-6z"/><path fill="#94989b" d="M48.8 52.8h7.5v1.8h-7.5zm0-21.3h7.5v1.7h-7.5z"/><path fill="#c8d0d6" d="M50.1 34h.7v18h-.7zm4.1 0h.7v18h-.7zm-1.4 0h.7v18h-.7zm-1.3 0h.7v18h-.7z"/>`),
		g.Group(children),
	)
}