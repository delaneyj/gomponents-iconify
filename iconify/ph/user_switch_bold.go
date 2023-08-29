package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSwitchBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m252.49 140.49l-20 20a12 12 0 0 1-17 0l-20-20a12 12 0 0 1 16.32-17.57A84 84 0 0 0 62.42 75.5a12 12 0 0 1-18.73-15a108 108 0 0 1 192.2 62.69a12 12 0 0 1 16.6 17.31Zm-40.18 55a108 108 0 0 1-192.2-62.69a12 12 0 0 1-16.6-17.31l20-20a12 12 0 0 1 17 0l20 20a12 12 0 0 1-16.32 17.57a83.55 83.55 0 0 0 17.88 46.86a83.48 83.48 0 0 1 29-23.42a52 52 0 1 1 74 0A83.39 83.39 0 0 1 194 180a12 12 0 0 1 18.3 15.49ZM128 148a28 28 0 1 0-28-28a28 28 0 0 0 28 28Zm0 64a83.6 83.6 0 0 0 48.48-15.37a60 60 0 0 0-96.91-.06A83.53 83.53 0 0 0 128 212Z"/>`),
		g.Group(children),
	)
}