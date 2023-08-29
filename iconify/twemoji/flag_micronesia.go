package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMicronesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#75B2DD" d="M32 5H4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4z"/><path fill="#FFF" d="m18.3 7.6l-.584 1.797h-1.889l1.528 1.11l-.583 1.796l1.528-1.11l1.528 1.11l-.583-1.796l1.528-1.11h-1.889zm8.603 9.816v-1.889l-1.11 1.528l-1.796-.583L25.107 18l-1.11 1.528l1.796-.583l1.11 1.528v-1.889L28.7 18zm-7.658 8.077l.583-1.796l-1.528 1.11l-1.528-1.11l.583 1.796l-1.528 1.11h1.889L18.3 28.4l.584-1.797h1.889zm-8.438-8.438l-1.11-1.528v1.889L7.9 18l1.797.584v1.889l1.11-1.528l1.796.583L11.493 18l1.11-1.528z"/>`),
		g.Group(children),
	)
}