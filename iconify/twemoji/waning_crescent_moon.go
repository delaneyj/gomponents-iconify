package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaningCrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#FFD983"/><path fill="#66757F" d="M36 18c0 9.941-8.059 18-18 18c-.294 0-.58-.029-.87-.043C11.239 33.393 7 26.332 7 18C7 9.669 11.239 2.607 17.13.044C17.42.03 17.706 0 18 0c9.941 0 18 8.059 18 18z"/><circle cx="25.5" cy="8.5" r="3.5" fill="#5B6876"/><circle cx="16" cy="16" r="3" fill="#5B6876"/><circle cx="14.5" cy="27.5" r="3.5" fill="#5B6876"/><circle cx="15" cy="6" r="2" fill="#5B6876"/><circle cx="33" cy="18" r="1" fill="#5B6876"/><circle cx="6" cy="9" r="1" fill="#FFCC4D"/><circle cx="21" cy="31" r="1" fill="#5B6876"/><circle cx="4" cy="19" r="2" fill="#FFCC4D"/><circle cx="26" cy="23" r="2" fill="#5B6876"/>`),
		g.Group(children),
	)
}