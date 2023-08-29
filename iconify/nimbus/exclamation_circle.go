package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.328 8c0 3.377-2.76 6.25-6.328 6.25c-3.567 0-6.328-2.873-6.328-6.25S4.432 1.75 8 1.75c3.567 0 6.328 2.873 6.328 6.25ZM15.5 8a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0ZM8.59 3.101l-.013 7.3l-1.172-.002l.013-7.3l1.172.002Zm.015 9.249a.673.673 0 0 0-.179-.46a.59.59 0 0 0-.43-.19a.59.59 0 0 0-.432.19a.673.673 0 0 0-.178.46c0 .172.064.338.178.46a.59.59 0 0 0 .431.19a.59.59 0 0 0 .431-.19a.673.673 0 0 0 .179-.46Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}