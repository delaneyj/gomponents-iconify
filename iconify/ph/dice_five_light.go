package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFiveLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 34H64a30 30 0 0 0-30 30v128a30 30 0 0 0 30 30h128a30 30 0 0 0 30-30V64a30 30 0 0 0-30-30Zm18 158a18 18 0 0 1-18 18H64a18 18 0 0 1-18-18V64a18 18 0 0 1 18-18h128a18 18 0 0 1 18 18ZM102 92a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm36 36a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm36-36a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-72 72a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm72 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Z"/>`),
		g.Group(children),
	)
}