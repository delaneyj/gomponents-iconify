package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriodosBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.582 21.679c5.448 4.35 6.781 10.899 3.962 13.718c-4.387 4.387-17.842 5.456-25.067-12.46"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.275 36.793a14.326 14.326 0 0 1-5.125 1.11c-6.561 0-9.805-1.547-9.805-6.34c0-4.865 5.308-14.08 19.146-14.08a21.86 21.86 0 0 1 8.795 1.665"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.076 14.865c0-5.566 4.307-11.888 9.416-11.888s8.795 4.847 8.795 16.275c0 6.8-4.1 12.112-8.801 15.299"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.492 2.977c-12.07 0-22.76 10.966-21.949 23.13m39.001 9.29c5.898-5.898 6.653-21.472-7.963-31.185M5.904 34.467c5.413 7.27 13.08 14.864 30.626 7.64"/>`),
		g.Group(children),
	)
}