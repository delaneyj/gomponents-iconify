package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pagoda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPagoda0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M18 16h12s2.424 2.672 4.5 3.429c1.301.474 3.5.571 3.5.571s-1.242.775-2.1 1.143C34.815 21.608 33 22 33 22H15s-1.815-.392-2.9-.857C11.242 20.775 10 20 10 20s2.199-.098 3.5-.571C15.576 18.673 18 16 18 16Zm-1 12h14s2.627 2.672 5 3.429c1.487.474 4 .571 4 .571s-1.42.775-2.4 1.143C36.36 33.608 34 34 34 34H14s-2.36-.392-3.6-.857C9.42 32.775 8 32 8 32s2.513-.098 4-.571C14.373 30.673 17 28 17 28Zm1-20.571C20.45 6.299 24 4 24 4s3.55 2.298 6 3.429c.772.356 2 .857 2 .857s-.726.556-1.2.857C30.183 9.535 29 10 29 10H19s-1.183-.465-1.8-.857c-.474-.301-1.2-.857-1.2-.857s1.228-.501 2-.857Z"/><path stroke-linecap="round" d="M21 10v6m6-6v6m-8 6v6m10-6v6m-13 6v10h16V34"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPagoda0)"/>`),
		g.Group(children),
	)
}