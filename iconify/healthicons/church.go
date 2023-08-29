package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Church(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M23 13.643V11h-3V9h3V6h2v3h3v2h-3v2.643L30.222 17s.778.5.778 1.5V42h-4v-9a3 3 0 1 0-6 0v9h-4V18.5c0-1 .778-1.5.778-1.5L23 13.643ZM7.5 26.5c-.961.278-1.5 1.08-1.5 2V42h9V24l-7.5 2.5Zm34.5 2c0-.92-.539-1.722-1.5-2L33 24v18h9V28.5Z"/>`),
		g.Group(children),
	)
}