package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadSideCough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M616 304a24 24 0 1 0-24-24a24 24 0 0 0 24 24Zm-64 112a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm-64-56a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm128 104a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm0-104a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm-64-40a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm-74.78-45c-21-47.12-48.5-151.75-73.12-186.75A208.13 208.13 0 0 0 234.1 0H192C86 0 0 86 0 192c0 56.75 24.75 107.62 64 142.88V512h224v-32h64a64 64 0 0 0 64-64h-96a32 32 0 0 1 0-64h96v-32h32a32 32 0 0 0 29.22-45ZM288 224a32 32 0 1 1 32-32a32.07 32.07 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}