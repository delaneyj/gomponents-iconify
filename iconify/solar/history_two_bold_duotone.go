package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11.25 2a.75.75 0 0 1 .75-.75c5.937 0 10.75 4.813 10.75 10.75S17.937 22.75 12 22.75S1.25 17.937 1.25 12a.75.75 0 0 1 1.5 0A9.25 9.25 0 1 0 12 2.75a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.1 2.398a.75.75 0 0 1-.43.97a9.187 9.187 0 0 0-.419.174a.75.75 0 1 1-.608-1.371a10.4 10.4 0 0 1 .487-.203a.75.75 0 0 1 .97.43ZM5.648 4.241a.75.75 0 0 1-.026 1.06c-.11.105-.216.212-.32.321a.75.75 0 1 1-1.087-1.034c.121-.127.246-.252.373-.373a.75.75 0 0 1 1.06.026ZM3.161 7.26a.75.75 0 0 1 .38.99c-.06.138-.119.278-.173.42a.75.75 0 0 1-1.4-.54c.064-.165.131-.328.203-.488a.75.75 0 0 1 .99-.381Z" clip-rule="evenodd" opacity=".5"/><path d="M12 8.25a.75.75 0 0 1 .75.75v3.25H16a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75V9a.75.75 0 0 1 .75-.75Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}