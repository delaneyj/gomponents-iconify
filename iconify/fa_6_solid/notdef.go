package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notdef(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 390.3L153.5 256L64 121.7v268.6zm38.5 57.7h179L192 313.7L102.5 448zm128-192L320 390.3V121.7L230.5 256zm51-192h-179L192 198.3L281.5 64zM0 48C0 21.5 21.5 0 48 0h288c26.5 0 48 21.5 48 48v416c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V48z"/>`),
		g.Group(children),
	)
}