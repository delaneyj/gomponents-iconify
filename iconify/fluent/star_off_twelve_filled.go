package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOffTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.854 1.146a.5.5 0 1 0-.708.708L3.23 3.936l-1.543.224a.8.8 0 0 0-.444 1.364l1.8 1.754l-.425 2.476a.8.8 0 0 0 1.16.843l2.224-1.169l2.223 1.169a.802.802 0 0 0 1.15-.515l.772.772a.5.5 0 0 0 .708-.708l-9-9Zm8.905 4.378L9.182 7.061L4.75 2.628l.534-1.082a.8.8 0 0 1 1.435 0L7.83 3.799l2.486.36a.8.8 0 0 1 .443 1.365Z"/>`),
		g.Group(children),
	)
}