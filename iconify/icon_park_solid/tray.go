package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTray0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="28" height="30" x="24.762" y="3.243" fill="#fff" stroke="#fff" rx="2" transform="rotate(45 24.762 3.243)"/><path stroke="#fff" d="m38.197 16.677l4.242-4.242l-7.07-7.071l-4.244 4.242"/><path stroke="#000" d="M18 21h12m-12 6h4m6 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTray0)"/>`),
		g.Group(children),
	)
}