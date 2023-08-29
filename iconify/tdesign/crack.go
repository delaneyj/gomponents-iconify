package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.387 3.02a9 9 0 0 0-.957 17.844l1.45-2.865A4.999 4.999 0 0 1 7.67 15.5l-.501-.866l1.731-1l.5.865c.31.536.78.967 1.345 1.227L9.882 14l2-4l-2.001-4.003l1.506-2.976Zm2.183.116l-1.45 2.867L14.117 10l-2 4l.91 1.82a3.01 3.01 0 0 0 1.571-1.32l.501-.866l1.731 1.001l-.5.866a5.018 5.018 0 0 1-2.405 2.115l.193.387l-1.506 2.977a9 9 0 0 0 .957-17.844ZM1 12C1 5.925 5.925 1 12 1c.37 0 .737.018 1.099.054C18.659 1.606 23 6.296 23 12c0 6.075-4.925 11-11 11c-.37 0-.737-.018-1.099-.054C5.341 22.394 1 17.704 1 12Zm8-4v4H7V8h2Zm8 0v4h-2V8h2Z"/>`),
		g.Group(children),
	)
}