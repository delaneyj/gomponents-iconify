package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmSecurityServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 20a2.97 2.97 0 0 0-1.855.66L22 19.423v-1.606a3 3 0 1 0-2 0v1.606l-2.145 1.239A2.97 2.97 0 0 0 16 20a3.02 3.02 0 1 0 2.925 2.353L21 21.154l2.075 1.198A2.998 2.998 0 1 0 26 20Zm-10 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm5-10a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm5 10a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M16 31A11.012 11.012 0 0 1 5 20V6.382l10.987-5.5l10.46 5.237l-.895 1.788l-9.565-4.789L7 7.618V20a9 9 0 0 0 14.4 7.201l1.2 1.599A10.908 10.908 0 0 1 16 31Z"/>`),
		g.Group(children),
	)
}