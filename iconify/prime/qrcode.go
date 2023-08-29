package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qrcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.38 3.88v7h7v-7Zm5.5 5.5h-4v-4h4Zm-14 1.5h7v-7h-7Zm1.5-5.5h4v4h-4Zm-1.5 14h7v-7h-7Zm1.5-5.5h4v4h-4Zm7-1.5h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75zm-1.75 1.75h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75zm-5.25 1.75h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75zm-1.75 1.75h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75z"/>`),
		g.Group(children),
	)
}