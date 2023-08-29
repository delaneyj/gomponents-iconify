package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileJpgFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M48 120h160a8 8 0 0 0 8-8V88a8 8 0 0 0-2.34-5.66l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v72a8 8 0 0 0 8 8Zm104-76l44 44h-44Zm-32 100h-16a8 8 0 0 0-8 8v56a8 8 0 0 0 16 0v-8h8a28 28 0 0 0 0-56Zm0 40h-8v-24h8a12 12 0 0 1 0 24Zm96 16.87a8 8 0 0 1-2.22 5.53A30.06 30.06 0 0 1 192 216c-17.64 0-32-16.15-32-36s14.36-36 32-36a29.38 29.38 0 0 1 16.48 5.12a8 8 0 0 1-9 13.26A13.21 13.21 0 0 0 192 160c-8.82 0-16 9-16 20s7.18 20 16 20a13.63 13.63 0 0 0 8-2.71V192a8 8 0 0 1 0-16h8a8 8 0 0 1 8 8ZM80 152v38a26 26 0 0 1-52 0a8 8 0 0 1 16 0a10 10 0 0 0 20 0v-38a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}