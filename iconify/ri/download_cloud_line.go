package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadCloudLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 14.5a6.496 6.496 0 0 1 3.064-5.519a8.001 8.001 0 0 1 15.872 0a6.5 6.5 0 0 1-2.936 12L7 21c-3.356-.274-6-3.078-6-6.5Zm15.848 4.487a4.5 4.5 0 0 0 2.03-8.309l-.807-.503l-.12-.942a6.001 6.001 0 0 0-11.903 0l-.12.942l-.805.503a4.5 4.5 0 0 0 2.029 8.309l.173.013h9.35l.173-.013ZM13 12h3l-4 5l-4-5h3V8h2v4Z"/>`),
		g.Group(children),
	)
}