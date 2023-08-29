package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudPal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m30.5 17.7l-6-3.5c-.1-.2-.3-.2-.5-.2s-.4 0-.5.1l-6 3.5c-.3.2-.5.5-.5.9v6.9c0 .4.2.7.5.9l6 3.5c.1.2.3.2.5.2s.4 0 .5-.1l6-3.5c.3-.2.5-.5.5-.9v-6.9c0-.4-.2-.7-.5-.9zM29 24.9l-5 3l-5-3v-5.8l5-3l5 3v5.8z"/><path fill="currentColor" d="M14 26H9c-4.4 0-8-3.6-8-8c0-3.7 2.6-6.9 6.2-7.8C8 5.5 12.1 2 17 2c5.5 0 10 4.5 10 10h-2c0-4.4-3.6-8-8-8c-4.1 0-7.5 3.1-8 7.2v.8l-.8.1C5.2 12.5 3 15 3 18c0 3.3 2.7 6 6 6h5v2z"/>`),
		g.Group(children),
	)
}