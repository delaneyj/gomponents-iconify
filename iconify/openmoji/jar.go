package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiJar0" d="M53 29.5v30.2a3.135 3.135 0 0 1-2.9 3.3H21.9a3.135 3.135 0 0 1-2.9-3.3V29.5a15.038 15.038 0 0 1 2.2-7.8l1.8-3v-4.2h-.9a1.58 1.58 0 0 1-1.6-1.6v-2.4a1.58 1.58 0 0 1 1.6-1.6h27.8a1.58 1.58 0 0 1 1.6 1.6v2.4a1.58 1.58 0 0 1-1.6 1.6H49v4.2l1.8 3a14.774 14.774 0 0 1 2.2 7.8Zm-3.4-15H23"/></defs><defs><path id="openmojiJar1" d="M26.3 33.1h19.3a1.685 1.685 0 0 1 1.7 1.7v14.1a1.685 1.685 0 0 1-1.7 1.7H26.3a1.685 1.685 0 0 1-1.7-1.7V34.8a1.828 1.828 0 0 1 1.7-1.7Z"/></defs><g fill="#fff"><use href="#openmojiJar1"/><use href="#openmojiJar0"/></g><path fill="#9b9b9a" d="M49.6 14.7H22.4l-1.5-1.3l-.7-2.3l1-2h29.1l1.5.9l-.3 3.7l-1.9 1z"/><path fill="#d0cfce" d="M53 28.1v33.2L50.5 63h-3.2V36c0-11.3-11.9-17-24.9-17l.3-4.5H49v2.6l.6 2.5l2.2 3.9Z"/><path fill="#3f3f3f" d="M51.8 10v3.4l-1.1.9v.2h-4.5V9.1h4.6v.1l1 .8z"/><g fill="none" stroke="#000" stroke-width="2"><use href="#openmojiJar1" stroke-miterlimit="10"/><use href="#openmojiJar0" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}