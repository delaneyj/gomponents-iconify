package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Portdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M19.205 45.035C-2.857 34.876 2.955 12.425 4.461 7.71a1.995 1.995 0 0 1 1.174-1.244c4.991-1.997 11.302-3.682 13.829-3.682c2.047 0 10.003 1.689 14.999 3.688a1.98 1.98 0 0 1 1.16 1.242c1.508 4.721 7.312 27.165-14.745 37.321a2.018 2.018 0 0 1-1.673 0Z"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M23.11 18.494h6.145"/><circle cx="31.432" cy="18.494" r="1.766" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M22.605 26.737h5.182"/><circle cx="29.98" cy="26.737" r="1.766" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M21.975 22.756h19.541"/><circle cx="43.591" cy="22.756" r="1.766" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="m21.633 22.911l1.028 6.56a.658.658 0 0 1-.651.761h-4.006a.66.66 0 0 1-.651-.762l1.041-6.595a.688.688 0 0 0-.272-.639a3.072 3.072 0 1 1 3.794.028a.689.689 0 0 0-.283.647Z"/>`),
		g.Group(children),
	)
}