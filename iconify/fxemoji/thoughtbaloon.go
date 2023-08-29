package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thoughtbaloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="78.4" cy="447.9" r="29.9" fill="#B0E9FF"/><circle cx="140.5" cy="391.4" r="35.9" fill="#B0E9FF"/><circle cx="24.4" cy="488" r="18.7" fill="#B0E9FF"/><path fill="#B0E9FF" d="M481.1 193.8c3.8-13.2 5.8-27.1 5.8-41.5C486.8 68.6 419 .7 335.3.7c-57.6 0-107.7 32.2-133.4 79.5c-16.5-14.1-38-22.6-61.4-22.6c-52.3 0-94.7 42.4-94.7 94.7c0 24.1 9 46.2 23.9 62.9c-14.6 12.6-23.9 31.2-23.9 52c0 37.9 30.7 68.7 68.7 68.7c15.4 0 29.7-5.1 41.2-13.7c14.6 40.4 53.3 69.3 98.7 69.3c42.9 0 79.8-25.8 96-62.7c15.8 11.9 35.5 18.9 56.8 18.9c52.3 0 94.7-42.4 94.7-94.7c0-22.4-7.8-43-20.8-59.2z"/>`),
		g.Group(children),
	)
}