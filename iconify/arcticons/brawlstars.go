package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brawlstars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.47 33.96c-9.61-7.854-6.365-19.144.928-25.212s17.395-3.24 17.395-3.24s10.163 2.6 13.444 11.503s.447 20.302-11.803 22.299M10.47 33.96l-.712 2.307s-.818 2.092 2.632 2.946l6.877 1.842m11.167-1.745l-.537 2.352s-.337 2.222-3.752 1.236l-6.877-1.843"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.597 33.934l-2.204-.591l-2.204-.59l1.613-1.614l1.614-1.614l.59 2.204Z"/><circle cx="14.45" cy="24.573" r="5.112" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.9 17.86l2.321 4.02"/><circle cx="31.668" cy="29.186" r="5.112" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.305 20.647l-4.02 2.32"/>`),
		g.Group(children),
	)
}