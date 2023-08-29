package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.25 9a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0ZM12 6.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5Z"/><path d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12ZM12 2.75a9.25 9.25 0 0 0-6.558 15.773c.18-.973.535-1.89 1.246-2.628C7.753 14.791 9.454 14.25 12 14.25c2.546 0 4.247.541 5.311 1.645c.712.738 1.066 1.655 1.247 2.629A9.25 9.25 0 0 0 12 2.75Zm5.194 16.905c-.102-1.212-.365-2.1-.962-2.719c-.65-.673-1.853-1.186-4.232-1.186c-2.379 0-3.582.513-4.232 1.186c-.597.62-.86 1.507-.962 2.72A9.207 9.207 0 0 0 12 21.25a9.208 9.208 0 0 0 5.194-1.595Z"/></g>`),
		g.Group(children),
	)
}