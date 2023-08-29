package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Honey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHoney0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="m14.613 4.843l-4.69 4.689a3.316 3.316 0 1 0 4.69 4.689l4.689-4.69a3.316 3.316 0 0 0-4.69-4.688Zm10.16 8.597l-6.252 6.252a3.316 3.316 0 1 0 4.69 4.69l6.252-6.253a3.316 3.316 0 1 0-4.69-4.69Z"/><path d="m28.68 18.91l13.77 13.77c1.028 1.028.811 2.91-.483 4.206c-1.295 1.295-3.178 1.511-4.206.484L23.991 23.6"/><path fill="#555" d="M22.428 6.406L11.487 17.347a3.316 3.316 0 1 0 4.69 4.69l10.94-10.942a3.316 3.316 0 1 0-4.689-4.69Z"/><path stroke-linecap="round" d="M13.854 23.142c1.424 4.437 1.187 7.715-.71 9.834c-2.847 3.179-2.405 11.166 4.026 11.166c6.431 0 8.32-7.987 2.872-10.79"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHoney0)"/>`),
		g.Group(children),
	)
}