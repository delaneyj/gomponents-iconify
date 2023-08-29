package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisherOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFireExtinguisherOne0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M16 18a8 8 0 1 1 16 0v24a2 2 0 0 1-2 2H18a2 2 0 0 1-2-2V18Z"/><path stroke="#000" stroke-linecap="round" d="M24 24v10"/><path fill="#fff" stroke="#fff" d="M20 5h9v4h-9zm9 0l9-1v6l-9-1V5Z"/><path stroke="#fff" stroke-linecap="round" d="M20 7c-3 0-7.5-.5-10 2c-2.417 2.416-2 5-2 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFireExtinguisherOne0)"/>`),
		g.Group(children),
	)
}