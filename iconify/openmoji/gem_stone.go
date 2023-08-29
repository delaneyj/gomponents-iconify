package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GemStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiGemStone0" d="M56.377 11.98h-40L4 23.348l32 40.736l32-40.736z"/></defs><path fill="#61B2E4" d="M56.377 11.98h-40L4 23.348l32 40.736l32-40.736z"/><path fill="#92D3F5" d="m37.37 62.336l18.379-38.988L36 11.98h20.377L68 23.348L37.37 62.336"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiGemStone0"/><path d="M55.749 23.86L36 64.084V11.98L16.256 23.348m0 .512l19.748 40.224M4 23.348h64M36 11.98l19.749 11.368"/><use href="#openmojiGemStone0"/></g>`),
		g.Group(children),
	)
}