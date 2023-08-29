package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monstersinc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m704 768l-43-192l-53 192H416l-53-192l-43 192H0L128 0h320l64 192q-93 0-170.5 52.5T224 384q40 87 117.5 139.5T512 576t170.5-52.5T800 384q-40-87-117.5-139.5T512 192L576 0h320l128 768H704zM512 512q-53 0-90.5-37.5T384 384t37.5-90.5T512 256t90.5 37.5T640 384t-37.5 90.5T512 512z"/>`),
		g.Group(children),
	)
}