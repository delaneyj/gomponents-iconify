package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 31v5m-9.987-8.245l-9.986-3.244a21 21 0 0 1 39.945 0l-9.986 3.244H14.013Zm0 0a10.5 10.5 0 0 1 3.815-5.25l-3.815 5.25Zm0 0L24 31l-9.987-3.245Zm3.815-5.25A10.5 10.5 0 0 1 24 20.5l-6.172 2.005Zm0 0L24 31l-6.172-8.495ZM24 20.5a10.5 10.5 0 0 1 6.171 2.005L24 20.5Zm0 0V31V20.5Zm6.171 2.005a10.5 10.5 0 0 1 3.815 5.25l-3.815-5.25Zm0 0L24 31l6.171-8.495Zm3.815 5.25L24 31l9.986-3.245Z"/>`),
		g.Group(children),
	)
}