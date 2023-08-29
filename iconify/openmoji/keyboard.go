package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiKeyboard0" d="M41 44H9c-.6 0-1-.4-1-1V29c0-.6.4-1 1-1h32c.6 0 1 .4 1 1v14c0 .6-.4 1-1 1zm22 0h-8c-.6 0-1-.4-1-1V29c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v14c0 .6-.4 1-1 1zm-13.7-3.7v-.7c0-.6-.4-1-1-1h-.7c-.6 0-1 .4-1 1v.7c0 .6-.4 1-1 1H45c-.6 0-1 .4-1 1v.7c0 .6.4 1 1 1h6c.6 0 1-.4 1-1v-.7c0-.6-.4-1-1-1h-.7c-.5 0-1-.4-1-1zM51 34h-6c-.6 0-1-.4-1-1v-4c0-.6.4-1 1-1h6c.6 0 1 .4 1 1v4c0 .6-.4 1-1 1z"/></defs><path fill="#FFF" d="M4 24h64v24H4z"/><path fill="#D0CFCE" d="M59 24h9v24H32z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M4 24h64v24H4z"/><use href="#openmojiKeyboard0"/><use href="#openmojiKeyboard0" fill="#3F3F3F"/>`),
		g.Group(children),
	)
}