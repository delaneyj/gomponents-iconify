package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rstudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M71.4 38.8c-1.5-.6-3.9-1-6.9-1.1c-4.2-.1-9 .4-9.2.5v20c13.3.6 15.5-1.7 15.5-1.7c11.6-5.9 4.3-16.2.6-17.7z"/><path fill="currentColor" d="M64 0C28.6 0 0 28.6 0 64s28.6 64 64 64s64-28.6 64-64S99.3 0 64 0zm28.6 89.8H82L64.4 63.5h-9V84h9v5.8H41.5v-5.7l7.6-.1l-.1-45.9c-.8-.2-7.5-.8-7.5-.8V32c1 1 7.9 1.2 7.9 1.2c1.6.1 3.9.2 5.2-.1c9.3-1.7 16.4-.4 16.4-.4c14 3.2 14.2 15.8 10.3 22.6c-3.5 5.8-10.3 7.2-10.3 7.2l14.4 21.8l7.2-.1v5.6z"/>`),
		g.Group(children),
	)
}