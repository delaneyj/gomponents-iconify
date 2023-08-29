package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Registered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48a208 208 0 1 1 0 416a208 208 0 1 1 0-416zm0 464a256 256 0 1 0 0-512a256 256 0 1 0 0 512zm-96-360v208c0 13.3 10.7 24 24 24s24-10.7 24-24v-72h60.9l37.2 81.9c5.5 12.1 19.7 17.4 31.8 11.9s17.4-19.7 11.9-31.8l-34.1-75c21.8-14.3 36.3-39 36.3-67c0-44.2-35.8-80-80-80h-88c-13.3 0-24 10.7-24 24zm48 88v-64h64c17.7 0 32 14.3 32 32s-14.3 32-32 32h-64z"/>`),
		g.Group(children),
	)
}