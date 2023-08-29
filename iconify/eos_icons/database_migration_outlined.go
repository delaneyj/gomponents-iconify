package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseMigrationOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21h-2v-2.59L15.41 21L14 19.59L16.59 17H14v-2h6v6zM5.5 12C2 12 0 12.7 0 14v8c0 1.3 2.83 2 5.5 2s5.5-.7 5.5-2v-8c0-1.3-2-2-5.5-2Zm0 1c2.485 0 4.5.672 4.5 1.5S7.985 16 5.5 16S1 15.328 1 14.5S3.015 13 5.5 13ZM1 20.238v-1.056A10.956 10.956 0 0 0 5.5 20a10.957 10.957 0 0 0 4.5-.818v1.089A11.468 11.468 0 0 1 5.5 21a11.64 11.64 0 0 1-4.5-.762ZM18.5 0C15 0 13 .7 13 2v8c0 1.3 2.83 2 5.5 2s5.5-.7 5.5-2V2c0-1.3-2-2-5.5-2Zm0 1c2.485 0 4.5.672 4.5 1.5S20.985 4 18.5 4S14 3.328 14 2.5S16.015 1 18.5 1ZM14 8.238V7.182A10.956 10.956 0 0 0 18.5 8a10.957 10.957 0 0 0 4.5-.818v1.089A11.468 11.468 0 0 1 18.5 9a11.64 11.64 0 0 1-4.5-.762Z"/>`),
		g.Group(children),
	)
}