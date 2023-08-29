package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 4.5q0 .825-.588 1.413T10.5 6.5q-.325 0-.6-.088t-.55-.287q-.6.2-.963.725T8.025 8H19.85q.45 0 .738.35t.237.8l-.7 5q-.05.375-.337.613t-.663.237h-1.8l.175-1.725q.125-1.225-.688-2.125t-2.037-.9h-5.55q-1.225 0-2.038.9T6.5 13.275L6.675 15h-1.8q-.375 0-.663-.238t-.337-.612l-.7-5q-.05-.45.238-.8T4.15 8H6q0-1.225.675-2.225T8.5 4.3q.075-.775.65-1.287T10.5 2.5q.825 0 1.413.588T12.5 4.5ZM9.6 20.25h4.8q.4 0 .675-.25t.325-.65l.6-6.225q.05-.55-.3-.963t-.925-.412h-5.55q-.575 0-.925.413t-.3.962l.6 6.225q.05.4.325.65t.675.25Z"/>`),
		g.Group(children),
	)
}