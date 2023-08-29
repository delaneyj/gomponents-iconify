package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#89664c" d="M3 29v32h58V29C61-10 3-9.4 3 29z"/><path fill="#594640" d="M7.2 31.2V61h49.5V31.2C56.8-5 7.2-4.5 7.2 31.2z"/><path fill="#a37d64" d="M8 31.7V61h48V31.7c0-35.8-48-35.2-48 0z"/><path fill="#94989b" d="M0 61h64v3H0z"/><path fill="#594640" d="M35.2 30.5h16.6c-.6-12.1-8.3-18.6-16.6-19.8v19.8m-5.5 0V10.8c-8.3 1.2-16 7.9-16.6 19.7h16.6m5.5 5.6h16.7V56H35.2zm-22.2 0h16.7V56H13z"/><path fill="#89664c" d="M34.8 30.1h16.6c-.6-12.1-8.3-18.6-16.6-19.8v19.8m-5.6 0V10.4c-8.3 1.2-16 7.9-16.6 19.7h16.6m5.6 5.6h16.7v19.9H34.8zm-22.3 0h16.7v19.9H12.5z"/><path fill="#ffce31" d="M47.5 28.9h7v14h-7z"/><path fill="#f2b200" d="M48 29.4h6v13h-6z"/><path fill="#3e4347" d="m52 41.2l-.4-2.7c.2-.2.4-.5.4-.8c0-.6-.4-1-1-1s-1 .4-1 1c0 .3.2.6.4.8l-.4 2.7h2"/><circle cx="51" cy="33" r="2.9" fill="#b78300"/><path fill="#ffd664" d="M53.8 32.5c0 1.6-1.3 2.9-2.9 2.9c-1.6 0-2.9-1.3-2.9-2.9c0-1.6 1.3-2.9 2.9-2.9c1.6 0 2.9 1.3 2.9 2.9"/>`),
		g.Group(children),
	)
}