package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFrameTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.25 2.031A1.002 1.002 0 0 0 11 2H7.75v4.25h3.5V2.031Zm0 5.719H2V14c0 .452 0 .867.004 1.25h9.246v-7.5Zm0 9h-3.5V20l.917.001h2.583v-3.25Zm-5-14.727c-1.61.055-2.588.239-3.274.856c-.75.675-.924 1.666-.964 3.371H6.25V2.023Zm0 14.727v3.227c-1.61-.055-2.588-.239-3.274-.856c-.59-.53-.823-1.256-.915-2.37H6.25Z"/><path d="M12.75 7.006V4h3.5v4.25h-3.5V7.006Zm0 9V9.75H22V16c0 .452 0 .867-.004 1.25H12.75v-1.244Zm0 5.962A1 1 0 0 0 13 22h3.25v-3.25h-3.5v3.218Zm5-13.718V4.023c1.61.055 2.588.239 3.274.856c.75.675.924 1.666.964 3.371H17.75Zm4.19 10.5h-4.19v3.227c1.61-.055 2.588-.239 3.274-.856c.59-.53.823-1.256.915-2.37Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}