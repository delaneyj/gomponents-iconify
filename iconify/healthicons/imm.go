package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M26 25.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path fill-rule="evenodd" d="M33 24a9 9 0 1 1-18 0a9 9 0 0 1 18 0Zm-2 0a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z" clip-rule="evenodd"/><path d="M14.5 11a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7ZM39 31.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path fill-rule="evenodd" d="M31 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM16 32a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M44 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-2 0c0 9.941-8.059 18-18 18S6 33.941 6 24S14.059 6 24 6s18 8.059 18 18Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}