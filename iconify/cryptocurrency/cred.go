package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cred(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-3.864-16.034l-1.253 1.271L15.618 22L26 11.539l-1.253-1.271l-9.13 9.19zm2.11-.31l1.255 1.273l5.616-5.658L19.864 10zm-2.276 4.83L7.251 15.7L6 16.97l4.734 4.762z"/>`),
		g.Group(children),
	)
}