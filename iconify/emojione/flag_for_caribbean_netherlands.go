package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCaribbeanNetherlands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#fff" d="m12.6 54.9l48-32A29.95 29.95 0 0 0 38.4 2.7l-36 24C2.2 28.4 2 30.2 2 32c0 9.2 4.1 17.4 10.6 22.9z"/><path fill="#2a5f9e" d="M62 32c0-3.2-.5-6.2-1.4-9.1l-48 32c5.2 4.4 12 7.1 19.4 7.1c16.6 0 30-13.4 30-30"/><path fill="#ffce31" d="m2.5 26.7l36-24C36.4 2.2 34.2 2 32 2C17.2 2 5 12.7 2.5 26.7z"/><path fill="#3e4347" d="m39.5 29.5l2.8-1.6l-2.8-1.6c-.7-5.8-5.3-10.4-11.1-11.1l-1.6-2.8l-1.6 2.8c-5.8.8-10.4 5.4-11.2 11.2L11.2 28l2.8 1.6c.7 5.8 5.3 10.4 11.1 11.1l1.6 2.8l1.6-2.8c5.8-.8 10.4-5.4 11.2-11.2m-12.7 9.6c-6.2 0-11.2-5-11.2-11.2s5-11.2 11.2-11.2S38 21.7 38 27.9s-5.1 11.2-11.2 11.2"/><path fill="#ed4c5c" d="m31.2 28l2.2-3.9H29l-2.2-3.8l-2.3 3.8h-4.4l2.2 3.9l-2.2 3.8h4.4l2.3 3.8l2.2-3.8h4.4z"/>`),
		g.Group(children),
	)
}