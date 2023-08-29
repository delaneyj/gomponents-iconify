package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="2" height="2" x="9" y="6" fill="currentColor" rx=".5"/><rect width="2" height="2" x="13" y="6" fill="currentColor" rx=".5"/><rect width="2" height="2" x="9" y="9.5" fill="currentColor" rx=".5"/><rect width="2" height="2" x="13" y="9.5" fill="currentColor" rx=".5"/><rect width="2" height="2" x="9" y="13" fill="currentColor" rx=".5"/><rect width="2" height="2" x="13" y="13" fill="currentColor" rx=".5"/><path fill="currentColor" d="M18.25 19.25h-.5V4a.76.76 0 0 0-.75-.75H7a.76.76 0 0 0-.75.75v15.25h-.5a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5Zm-2 0H11V17a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v2.25H7.75V4.75h8.5Z"/>`),
		g.Group(children),
	)
}