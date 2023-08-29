package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.57 17.29a1 1 0 0 0-1.41 0a1.06 1.06 0 0 0-.22.33a1.07 1.07 0 0 0 0 .76a1.19 1.19 0 0 0 .22.33a1 1 0 0 0 .32.21a1 1 0 0 0 .39.08a1 1 0 0 0 .92-1.38a.91.91 0 0 0-.22-.33ZM20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.94Zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-6.13-9a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.73 1a1 1 0 0 1 .87-.5a1 1 0 0 1 0 2a1 1 0 1 0 0 2a3 3 0 0 0 0-6Z"/>`),
		g.Group(children),
	)
}