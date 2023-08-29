package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toilet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSToilet0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" stroke-linecap="round" stroke-linejoin="round" d="M19.999 10h8v12h-8V10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 35l-5 9h24l-5-9m2.999-13V5a1 1 0 0 0-1-1h-18a1 1 0 0 0-1 1v17"/><path d="M6.08 22.364A.3.3 0 0 1 6.372 22h35.254a.3.3 0 0 1 .292.364c-1.224 5.508-4.635 10.452-10 12.2C29.436 35.374 26.656 36 24 36c-2.655 0-5.436-.627-7.92-1.436c-5.365-1.749-8.776-6.692-10-12.2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSToilet0)"/>`),
		g.Group(children),
	)
}