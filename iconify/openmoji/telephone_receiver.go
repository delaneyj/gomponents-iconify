package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelephoneReceiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiTelephoneReceiver0" d="M13.67 21.31c1.415 9.9 26.87 35.355 36.77 36.77c0 0 5.657 0 8.486-5.658l-9.193-9.192l-3.535 3.536c-7.071-1.415-19.8-14.143-21.214-21.214l3.536-3.535l-9.192-9.193c-5.657 2.829-5.657 8.486-5.657 8.486z"/></defs><path fill="#d0cfce" d="M13.67 21.31c1.415 9.9 26.87 35.355 36.77 36.77c0 0 5.657 0 8.486-5.658l-9.193-9.192l-3.535 3.536c-7.071-1.415-19.8-14.143-21.214-21.214l3.536-3.535l-9.192-9.193c-5.657 2.829-5.657 8.486-5.657 8.486z"/><path fill="#9b9b9a" d="m25 26l-9.875-9.75L19 13l8.75 8.875zm31 31l-9.875-9.75L50 44l8.75 8.875z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiTelephoneReceiver0"/><use href="#openmojiTelephoneReceiver0"/></g>`),
		g.Group(children),
	)
}