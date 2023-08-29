package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pluralsight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.959 2.339C13.438-.401 5.083 3.5 2.36 11c-2.761 7.599 1.14 15.943 8.661 18.683c7.541 2.739 15.943-1.161 18.676-8.683C32.442 13.437 28.541 5.083 21 2.339zM16 32C7.197 32 0 24.803 0 16S7.197 0 16 0s16 7.197 16 16s-7.197 16-16 16zM11.901 7.74v16.52L26.24 16zm1.402 2.359L23.5 16l-10.197 5.901V10.098zm-4.704-.558v12.917l11.204-6.459zM10 11.937L17.083 16L10 20.083v-8.165z"/>`),
		g.Group(children),
	)
}