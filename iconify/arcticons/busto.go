package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Busto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.135 12.856H12.641v20.663h20.76m6.1 0h2.947"/><circle cx="36.373" cy="31.007" r="3.89" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.373" cy="31.007" r="1.778" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 33.52h0Zm-19.723 0a3.882 3.882 0 0 0 5.932 0M12.64 12.856h-2.468v5.772s.089 1.653 2.284 1.653m15.545-7.425v20.592M12.874 29.65h14.921m.459-5.473h14.06m-3.694-.099V13.19m-10.619-.334h-2.468v5.772s.089 1.653 2.284 1.653m-15.193 7.41h3.597s1.92.121 1.92 1.92m9.853-1.92h-3.597s-1.92.121-1.92 1.92"/>`),
		g.Group(children),
	)
}