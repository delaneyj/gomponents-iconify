package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="m6.198 4.256l2.19 2.513a.811.811 0 0 0 1.143 0l2.189-2.513c.314-.315.363-.876.049-1.19H9.965V1.144c0-.559-.433-1.01-.968-1.01c-.535 0-.969.451-.969 1.01v1.922H6.247c-.315.315-.364.874-.049 1.19z"/><path d="M7 10h4.031v1.031H7z"/><path fill="currentColor" d="M14.993 1.006h-3.962v.975h3.146L15.54 9.01H2.505l1.407-7.029h3.057v-.975H3.073l-2.06 8.755v4.17c0 1.334.472 2.028 1.804 2.028h12.28c1.246 0 1.885-.527 1.885-2.111V9.761l-1.989-8.755zM11.016 11H6.985V9.969h4.031V11z"/></g>`),
		g.Group(children),
	)
}