package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.085 43.296V4.5h8.78c9.189 0 17.05 7.76 17.05 17.05v4.9c0 9.189-7.76 17.05-17.05 17.05h-8.78v-.204Zm0-19.398h10.84"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.192 17.772c4.697.92 8.27 2.961 10.618 6.126c-2.348 3.165-5.82 5.207-10.618 6.126c1.123-4.084 1.123-8.168 0-12.252Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.192 17.772c.817-2.246.817-4.288 0-6.33c7.045 1.532 9.495 6.432 10.618 12.456c-1.02 6.024-3.471 10.924-10.618 12.456c.817-2.042.817-4.084 0-6.33"/>`),
		g.Group(children),
	)
}