package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalNineThousand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiHal90000" fill="none" stroke-linejoin="round" d="M26 5h20v62H26z"/></defs><circle cx="36" cy="40" r="1"/><path fill="#3F3F3F" stroke="#000" stroke-linejoin="round" d="M26 5h20v62H26z"/><path fill="#9B9B9A" d="M26 53h20v14H26z"/><circle cx="36" cy="40" r="8" fill="#D22F27" stroke="#D0CFCE"/><circle r="4" fill="#EA5A47" transform="matrix(-1 0 0 1 36 40)"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41 40a4.993 4.993 0 0 0-5-5"/><circle cx="36" cy="40" r="2" fill="#F1B31C"/><circle cx="36" cy="40" r="1" fill="#FCEA2B"/><path fill="#61B2E4" d="M28 7h8v4h-8z"/><g stroke="#000"><use href="#openmojiHal90000" stroke-linejoin="round"/><circle cx="36" cy="40" r="7" fill="none"/><path stroke-linecap="round" d="M27.5 52.5h17"/><path fill="none" stroke-linejoin="round" d="M28 7h16v4H28z"/><use href="#openmojiHal90000" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}