package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCirclePlusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 56a12 12 0 0 1 12-12h12V32a12 12 0 0 1 24 0v12h12a12 12 0 0 1 0 24h-12v12a12 12 0 0 1-24 0V68h-12a12 12 0 0 1-12-12Zm70.51 54A107.88 107.88 0 1 1 146 21.49a12 12 0 0 1-4 23.67A85 85 0 0 0 128 44a83.94 83.94 0 0 0-65.95 135.94a83.48 83.48 0 0 1 29-23.42a52 52 0 1 1 74 0a83.36 83.36 0 0 1 29 23.42A83.52 83.52 0 0 0 212 128a85.2 85.2 0 0 0-1.16-14a12 12 0 0 1 23.67-4ZM128 148a28 28 0 1 0-28-28a28 28 0 0 0 28 28Zm0 64a83.53 83.53 0 0 0 48.43-15.43a60 60 0 0 0-96.86 0A83.53 83.53 0 0 0 128 212Z"/>`),
		g.Group(children),
	)
}