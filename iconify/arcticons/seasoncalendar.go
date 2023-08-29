package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seasoncalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.15 5.44v2.92h-24.3V5.44m0 2.92h-5.4a2 2 0 0 0-1.95 2v27.3a2 2 0 0 0 2 2h35.1a2 2 0 0 0 2-2V10.31a2 2 0 0 0-2-2h-5.4M4.5 14.47h39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.1 22.93a9.57 9.57 0 0 0-3.2 1.14a9.55 9.55 0 0 0-3.18-1.14c-1.31 0-3.57 1-3.57 4.86s2.19 6.79 4.61 6.79c1.24 0 1.53-.55 2.14-.55s.88.55 2.13.55c2.43 0 4.6-2.92 4.6-6.79s-2.21-4.86-3.53-4.86Zm-3.75-.62v.36h.37a3.39 3.39 0 0 0 3.39-3.39a2.16 2.16 0 0 0 0-.36h-.36a3.4 3.4 0 0 0-3.4 3.39Zm19.85 2.15a3.74 3.74 0 0 0-5.27 0h0a2.22 2.22 0 0 0-.28.33l1.63 1.63c.19.17-1.63-1.63-1.63-1.63A93.53 93.53 0 0 0 24 33.85a.67.67 0 0 0 .89.88c1.13-.51 2.84-1.27 4.47-2.07L27.69 31l1.62 1.62a29.18 29.18 0 0 0 4.89-2.93a3.71 3.71 0 0 0 0-5.25Zm5.52-.7a2.75 2.75 0 0 0-.22-.26a3 3 0 0 0-4-.25a3 3 0 0 0-.16-4a2.43 2.43 0 0 0-.25-.2a3.58 3.58 0 0 0-.25 4.95a3.52 3.52 0 0 0 4.88-.24Z"/>`),
		g.Group(children),
	)
}