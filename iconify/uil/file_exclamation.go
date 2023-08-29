package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.92 16.62a.76.76 0 0 0-.09-.18a1.58 1.58 0 0 0-.12-.15a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.5 1.5 0 0 0 0 .2a.84.84 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 12 18a.84.84 0 0 0 .38-.08a.9.9 0 0 0 .54-.54A.84.84 0 0 0 13 17a1.5 1.5 0 0 0 0-.2a.64.64 0 0 0-.08-.18ZM20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.94Zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-6-8a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}