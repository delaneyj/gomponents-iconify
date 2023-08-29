package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25a.75.75 0 0 1 .75.75v1.282a8.752 8.752 0 0 1 7.968 7.968H22a.75.75 0 0 1 0 1.5h-1.282a8.752 8.752 0 0 1-7.968 7.968V22a.75.75 0 0 1-1.5 0v-1.282a8.752 8.752 0 0 1-7.968-7.968H2a.75.75 0 0 1 0-1.5h1.282a8.752 8.752 0 0 1 7.968-7.968V2a.75.75 0 0 1 .75-.75Zm0 3.5a7.25 7.25 0 1 0 0 14.5a7.25 7.25 0 0 0 0-14.5Zm0 5a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM8.25 12a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}