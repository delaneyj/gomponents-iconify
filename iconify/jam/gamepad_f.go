package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5V4a1 1 0 1 0-2 0v1H4a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0V7h1a1 1 0 1 0 0-2H7zm2.318-4h5.364A6 6 0 0 1 24 6c0 3.314-2.686 10-6 10c-1.976 0-3.729-2.378-4.822-5h-2.356C9.73 13.622 7.976 16 6 16C2.686 16 0 9.314 0 6a6 6 0 0 1 9.318-5zM18 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-2 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm4 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-2 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}