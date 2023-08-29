package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handbag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiHandbag0" d="M61 60.958H11L13 28.5h46z"/></defs><g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#6A462F" d="M61 60.958H11L13 28.5h46z"/><path fill="#D0CFCE" d="M21.954 31.943h9.21v9.21h-9.21zm20.93 0h9.21v9.21h-9.21z"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiHandbag0"/><path d="M21.954 31.943h9.21v9.21h-9.21zm20.93 0h9.21v9.21h-9.21zM24.52 25.187c0-8.566 1.31-15.511 11.484-15.511S47.49 16.62 47.49 25.187"/><use href="#openmojiHandbag0"/></g>`),
		g.Group(children),
	)
}