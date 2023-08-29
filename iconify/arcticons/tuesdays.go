package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuesdays(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 14.051h5.592m-2.796 8.441v-8.441m17.047 11.457l-2.796 4.22l-2.796-4.22m2.796 8.441v-4.221m4.528-8.161c.517.674 1.167.925 2.07.925h1.25a2.106 2.106 0 0 0 2.105-2.105v-.01a2.106 2.106 0 0 0-2.106-2.105h-1.378a2.108 2.108 0 0 1-2.108-2.108h0c0-1.167.946-2.113 2.112-2.113h1.243c.903 0 1.552.251 2.07.925m-5.258 18.048c.517.674 1.167.925 2.07.925h1.25a2.106 2.106 0 0 0 2.105-2.106v-.01a2.106 2.106 0 0 0-2.106-2.105h-1.378a2.108 2.108 0 0 1-2.108-2.108h0c0-1.166.946-2.112 2.112-2.112h1.243c.903 0 1.552.25 2.07.925M18.255 14.051v5.645a2.796 2.796 0 1 0 5.592 0v-5.645m2.275 4.221h2.752m1.469 4.22h-4.221v-8.441h4.221M10.796 33.949v-8.441h1.9a3.693 3.693 0 0 1 3.692 3.693v1.055a3.693 3.693 0 0 1-3.693 3.693h-1.9Zm12.125-2.797h-3.74m-.926 2.797l2.796-8.441l2.797 8.441"/>`),
		g.Group(children),
	)
}