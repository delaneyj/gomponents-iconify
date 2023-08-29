package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.94 2.512a20.75 20.75 0 0 1 4.692 0l1.596.182a.75.75 0 0 1 .62 1.003l-2.203 6.016H17a.75.75 0 0 1 .702 1.014l-.19.507a35.749 35.749 0 0 1-5.535 9.745l-.391.49A.75.75 0 0 1 10.25 21v-7.152H8.429a.75.75 0 0 1-.748-.684l-.122-1.382a39.75 39.75 0 0 1 0-7.027l.122-1.382a.75.75 0 0 1 .663-.68l1.595-.18Zm4.523 1.49a19.25 19.25 0 0 0-4.354 0l-.987.113l-.069.772c-.2 2.25-.2 4.513 0 6.763l.062.698H11a.75.75 0 0 1 .75.75v5.704a34.251 34.251 0 0 0 4.163-7.589H13.57a.75.75 0 0 1-.704-1.007l2.244-6.13l-.648-.073Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}