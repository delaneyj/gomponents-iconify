package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Defibrilator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M39.843 9.972c2.339 2.575 2.745 6.191 1.524 9.297c-2.087 5.31-6.784 10.12-11.24 14.684c-2.179 2.231-4.3 4.404-6.03 6.547c-.4-.39-.813-.79-1.237-1.2c-5.689-5.52-13.324-12.926-16.18-20.031a8.888 8.888 0 0 1 .192-7.073a8.96 8.96 0 0 1 2.164-2.905c3.633-3.244 9.32-2.739 12.777.69l2.283 2.264l2.645-2.622a9.036 9.036 0 0 1 3.041-1.987a9.096 9.096 0 0 1 7.132.19a9.023 9.023 0 0 1 2.93 2.146ZM18.97 24.802l7.94-10.266l-1.527 7.75h4.526a.5.5 0 0 1 .396.806l-7.94 10.266l1.527-7.75h-4.527a.5.5 0 0 1-.395-.806Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}