package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWorker0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M32 16a8 8 0 1 1-16 0"/><path fill="#fff" stroke="#fff" d="M24 8a8 8 0 0 0-8 8h16a8 8 0 0 0-8-8Z"/><path stroke="#fff" d="M12 16h24M24 4v4"/><path fill="#fff" stroke="#fff" d="M24 27c-9.389 0-17 7.163-17 16h34c0-8.837-7.611-16-17-16Z"/><path stroke="#000" d="M18 34v4m12-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWorker0)"/>`),
		g.Group(children),
	)
}