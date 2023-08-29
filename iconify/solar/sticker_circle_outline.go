package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickerCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75a9.25 9.25 0 0 0-.75 18.47c.002-2.257.034-3.633.435-4.78a7.75 7.75 0 0 1 4.755-4.755c1.147-.401 2.523-.433 4.78-.435A9.25 9.25 0 0 0 12 2.75Zm9.04 10.001c-2.188.006-3.249.05-4.104.35a6.25 6.25 0 0 0-3.835 3.835c-.3.855-.344 1.916-.35 4.104c.15-.085.294-.195.427-.328l7.534-7.534c.133-.133.243-.277.328-.427ZM1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12c0 .884-.408 1.669-.977 2.239l-7.534 7.534c-.57.57-1.355.977-2.239.977c-5.937 0-10.75-4.813-10.75-10.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}