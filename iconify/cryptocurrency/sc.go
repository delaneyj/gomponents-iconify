package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm0-24.5a8.5 8.5 0 1 0 0 17h8.5V16A8.5 8.5 0 0 0 16 7.5zm5.1 13.6h-5.023c-2.82 0-5.163-2.255-5.177-5.074a5.106 5.106 0 0 1 5.126-5.126c2.819.014 5.074 2.358 5.074 5.177V21.1z"/>`),
		g.Group(children),
	)
}