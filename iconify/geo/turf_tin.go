package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfTin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="13.162" cy="54.683" r="5.824" fill="currentColor"/><circle cx="51.116" cy="15.612" r="5.824" fill="currentColor"/><circle cx="87.954" cy="83.706" r="5.824" fill="currentColor"/><circle cx="50" cy="50.218" r="5.824" fill="currentColor"/><path fill="currentColor" d="m40.7 53.359l-17.757 2.152a9.797 9.797 0 0 1-2.934 6.202l58.146 22.491c-.008-.166-.025-.33-.025-.499a9.73 9.73 0 0 1 1.299-4.854l-23.552-20.78A9.756 9.756 0 0 1 50 60.042c-4.317 0-7.985-2.803-9.3-6.683zm19.124-3.141a9.761 9.761 0 0 1-1.296 4.858l23.546 20.775a9.78 9.78 0 0 1 5.134-1.931L58.899 21.583a9.831 9.831 0 0 1-6.094 3.69l-.497 15.407c4.306 1.042 7.516 4.917 7.516 9.538zM48.311 40.55l.496-15.391a9.864 9.864 0 0 1-6.816-5.93L16.389 45.415a9.871 9.871 0 0 1 6.077 6.126l17.751-2.151c.376-4.466 3.747-8.082 8.094-8.84z"/>`),
		g.Group(children),
	)
}