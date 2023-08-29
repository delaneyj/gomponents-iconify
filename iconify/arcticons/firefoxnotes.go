package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firefoxnotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.64 4.51a2.6 2.6 0 0 0-2.78 2.78h-2.78a2.59 2.59 0 0 0-2.79 2.79v30.63a2.59 2.59 0 0 0 2.79 2.78h27.84a2.59 2.59 0 0 0 2.79-2.78V10.08a2.59 2.59 0 0 0-2.79-2.79h-2.78c0-3.71-5.57-3.71-5.57 0H18.43a2.6 2.6 0 0 0-2.79-2.78Zm-2.79 9.54h22.3m-22.3 7.49h22.3m-22.3 7.49h22.3m-22.3 7.5h22.3"/>`),
		g.Group(children),
	)
}