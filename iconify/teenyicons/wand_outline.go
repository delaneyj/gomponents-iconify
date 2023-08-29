package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WandOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.853 8.848l.354-.353l-.707-.707l-.353.353l.706.707ZM9 10.493v.5h1v-.5H9Zm1-.999v-.5H9v.5h1ZM9 1.5V2h1v-.5H9Zm1-1V0H9v.5h1ZM4.5 4.997H4v1h.5v-1Zm1 1H6v-1h-.5v1Zm8-1H13v1h.5v-1Zm1 1h.5v-1h-.5v1Zm-2.353-3.852l-.354.354l.707.707l.353-.354l-.706-.707Zm1.706-.292l.354-.353L13.5.792l-.353.354l.706.707Zm-8-.707L5.5.792l-.707.708l.354.353l.706-.707Zm.294 1.706l.353.354l.707-.707l-.354-.354l-.706.707Zm6.706 5.29l-.353-.354l-.707.707l.354.353l.706-.707Zm.294 1.706l.353.353l.707-.707l-.354-.354l-.706.708ZM.853 14.844l6-5.996l-.706-.707l-6 5.996l.706.707ZM10 10.495v-1H9v1h1ZM10 1.5v-1H9v1h1ZM4.5 5.997h1v-1h-1v1Zm9 0h1v-1h-1v1Zm-.647-3.145l1-.999l-.706-.707l-1 1l.706.706Zm-7.706-.999l1 1l.706-.708l-1-1l-.706.708Zm7 6.995l1 1l.706-.708l-1-.999l-.706.707Z"/>`),
		g.Group(children),
	)
}