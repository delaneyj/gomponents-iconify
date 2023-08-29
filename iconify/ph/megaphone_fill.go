package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 72h-39.8c-2.91-.17-53.62-3.74-101.91-44.24A16 16 0 0 0 24 40v160a16 16 0 0 0 26.29 12.25c37.77-31.68 77-40.76 93.71-43.3v31.72a16 16 0 0 0 7.12 13.33l11 7.33A16 16 0 0 0 186.5 212l11.77-44.36A48 48 0 0 0 192 72Zm-21 135.89v.11l-11-7.33V168h21.6ZM192 152h-32V88h32a32 32 0 1 1 0 64Z"/>`),
		g.Group(children),
	)
}