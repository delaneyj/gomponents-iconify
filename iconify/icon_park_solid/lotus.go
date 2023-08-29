package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lotus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLotus0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 16c.196-2.143 1.235-7.143 5-10c1.372 1.667 4.53 6 5 10"/><path fill="#fff" d="M23.752 42C15.282 40.545-.3 31.31 5.12 6c7.152 1.636 20.892 11.127 18.633 36Z"/><path fill="#fff" d="M24.248 42C32.718 40.545 48.3 31.31 42.88 6c-7.152 1.636-20.892 11.127-18.633 36Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLotus0)"/>`),
		g.Group(children),
	)
}