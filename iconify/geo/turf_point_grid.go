package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfPointGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="11.528" cy="88.472" r="5.824" fill="currentColor"/><circle cx="88.472" cy="11.528" r="5.824" fill="currentColor"/><circle cx="69.236" cy="11.528" r="5.824" fill="currentColor"/><circle cx="50" cy="11.528" r="5.824" fill="currentColor"/><circle cx="30.764" cy="11.528" r="5.824" fill="currentColor"/><circle cx="11.528" cy="11.528" r="5.824" fill="currentColor"/><circle cx="88.472" cy="88.472" r="5.824" fill="currentColor"/><circle cx="88.472" cy="69.236" r="5.824" fill="currentColor"/><circle cx="88.472" cy="50" r="5.824" fill="currentColor"/><circle cx="88.472" cy="30.764" r="5.824" fill="currentColor"/><circle cx="69.472" cy="88.472" r="5.824" fill="currentColor"/><circle cx="69.472" cy="69.236" r="5.824" fill="currentColor"/><circle cx="69.472" cy="50" r="5.824" fill="currentColor"/><circle cx="69.472" cy="30.764" r="5.824" fill="currentColor"/><circle cx="50.472" cy="88.472" r="5.824" fill="currentColor"/><circle cx="50.472" cy="69.236" r="5.824" fill="currentColor"/><circle cx="50.472" cy="50" r="5.824" fill="currentColor"/><circle cx="50.472" cy="30.764" r="5.824" fill="currentColor"/><circle cx="30.472" cy="88.472" r="5.824" fill="currentColor"/><circle cx="30.472" cy="69.236" r="5.824" fill="currentColor"/><circle cx="30.472" cy="50" r="5.824" fill="currentColor"/><circle cx="30.472" cy="30.764" r="5.824" fill="currentColor"/><circle cx="11.472" cy="69.236" r="5.824" fill="currentColor"/><circle cx="11.472" cy="50" r="5.824" fill="currentColor"/><circle cx="11.472" cy="30.764" r="5.824" fill="currentColor"/><circle cx="88.472" cy="11.528" r="5.824" fill="currentColor"/><circle cx="69.236" cy="11.528" r="5.824" fill="currentColor"/><circle cx="50" cy="11.528" r="5.824" fill="currentColor"/><circle cx="30.764" cy="11.528" r="5.824" fill="currentColor"/><circle cx="11.528" cy="11.528" r="5.824" fill="currentColor"/>`),
		g.Group(children),
	)
}