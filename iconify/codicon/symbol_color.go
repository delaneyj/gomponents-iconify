package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1.003a7 7 0 0 0-7 7v.43c.09 1.51 1.91 1.79 3 .7a1.87 1.87 0 0 1 2.64 2.64c-1.1 1.16-.79 3.07.8 3.2h.6a7 7 0 1 0 0-14l-.04.03zm0 13h-.52a.58.58 0 0 1-.36-.14a.56.56 0 0 1-.15-.3a1.24 1.24 0 0 1 .35-1.08a2.87 2.87 0 0 0 0-4a2.87 2.87 0 0 0-4.06 0a1 1 0 0 1-.9.34a.41.41 0 0 1-.22-.12a.42.42 0 0 1-.1-.29v-.37a6 6 0 1 1 6 6l-.04-.04zM9 3.997a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm3 7.007a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm-7-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0zM13 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}