package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBlackFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M56 25.902c-2.864-7.742-5.73-15.484-8.598-23.226c-15.84-4.854-21.363 18.166-37.205 13.312L8 16.776L24.744 62h2.488l-8.437-22.786l.001.002C34.639 44.067 40.161 21.049 56 25.902"/>`),
		g.Group(children),
	)
}